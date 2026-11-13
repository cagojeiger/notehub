"""DAST 설정을 컴파일하는 소스에서 읽어 구·신 키 이름을 함께 지원한다."""
import os
from pathlib import Path
import subprocess
import tempfile
import textwrap
import unittest

WORKFLOW = Path(__file__).resolve().parents[1] / "workflows/security-dast.yml"


class DastCryptoTests(unittest.TestCase):
    def invoke(self, example, layout="deploy/.env.example"):
        block = WORKFLOW.read_text().split("      - name: Load source-owned crypto fixtures\n", 1)[1]
        block = block.split("\n      - name:", 1)[0]
        script = textwrap.dedent(block.split("        run: |\n", 1)[1])
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            fixture = root / layout
            fixture.parent.mkdir(parents=True, exist_ok=True)
            fixture.write_text(example)
            output = root / "env"
            output.write_text("existing=true\n")
            result = subprocess.run(["bash", "-e", "-o", "pipefail", "-c", script], cwd=root,
                                    env={**os.environ, "GITHUB_ENV": str(output)},
                                    text=True, capture_output=True)
            self.assertFalse((root / "injected").exists())
            return result, output.read_text()

    def example(self, role):
        return (f"NOTEGATE_ENC_ROOT_{role}KEY_ID=fixture-enc\n"
                f"NOTEGATE_ENC_ROOT_{role}SECRET=fixture-enc-secret\n"
                "NOTEGATE_LOOKUP_ROOT_KEY_ID=fixture-lookup\n"
                "NOTEGATE_LOOKUP_ROOT_SECRET=fixture-lookup-secret\n")

    def test_old_and_new_source_names_and_layouts(self):
        for role in ("", "PRIMARY_"):
            for layout in ("deploy/.env.example", ".env.example"):
                with self.subTest(role=role, layout=layout):
                    example = self.example(role)
                    result, output = self.invoke(example, layout)
                    self.assertEqual(result.returncode, 0, result.stderr)
                    self.assertEqual(result.stdout, "")
                    self.assertEqual(output, "existing=true\n" + example)

    def test_comments_unrelated_settings_and_shell_text_are_not_executed(self):
        example = self.example("PRIMARY_").replace("fixture-enc-secret", "literal-$(touch injected)")
        result, output = self.invoke(example + "GITHUB_TOKEN=unrelated\n# NOTEGATE_ENC_ROOT_SECONDARY_SECRET=ignored\n")
        self.assertEqual(result.returncode, 0, result.stderr)
        self.assertEqual(output, "existing=true\n" + example)
        self.assertEqual(result.stdout, "")

    def test_incomplete_or_duplicate_fixtures_do_not_publish_partial_settings(self):
        example = self.example("PRIMARY_")
        for invalid in (example.replace("NOTEGATE_LOOKUP_ROOT_SECRET=fixture-lookup-secret\n", ""),
                        example + "NOTEGATE_LOOKUP_ROOT_SECRET=duplicate\n",
                        example.replace("fixture-enc-secret", "")):
            result, output = self.invoke(invalid)
            self.assertNotEqual(result.returncode, 0)
            self.assertEqual(output, "existing=true\n")
            self.assertEqual(result.stdout, "")


if __name__ == "__main__":
    unittest.main()
