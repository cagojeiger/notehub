"""Public 계획·실행 결과의 누락과 거짓 성공을 검증한다."""
import copy
import json
import os
from pathlib import Path
import re
import subprocess
import sys
import tempfile
import unittest

import check_ci_result as result


ROOT = Path(__file__).resolve().parents[2]
PLAN = ROOT / ".github/ci-plan.json"


class ResultTests(unittest.TestCase):
    def setUp(self):
        self.plans = result.load_plan(PLAN)

    def needs(self, kind="ci", enabled=True):
        flags = {key: str(enabled).lower() for key in result.SELECTIONS}
        return {"gate": {"result": "success", "outputs": flags},
                **{name: {"result": "success" if "selection" not in entry or enabled else "skipped"}
                   for name, entry in self.plans[kind].items()}}

    def test_full_and_scoped_plans_and_all_request_kinds(self):
        for kind in self.plans:
            for enabled in (True, False):
                with self.subTest(kind=kind, enabled=enabled):
                    result.verify(kind, self.needs(kind, enabled), self.plans)

    def test_any_selected_job_failure_cancel_or_skip_is_rejected(self):
        for kind, plan in self.plans.items():
            for name in ("gate", *plan):
                for status in ("failure", "cancelled", "skipped", None):
                    with self.subTest(kind=kind, name=name, status=status):
                        needs = self.needs(kind)
                        needs[name]["result"] = status
                        with self.assertRaises(ValueError):
                            result.verify(kind, needs, self.plans)

    def test_missing_duplicate_shape_or_unknown_results_are_rejected(self):
        needs = self.needs()
        for changed in ({key: value for key, value in needs.items() if key != "rust-lint"},
                        {key: value for key, value in needs.items() if key != "browser-ui"},
                        {**needs, "unplanned": {"result": "success"}}, [], None):
            with self.subTest(changed=changed), self.assertRaises(ValueError):
                result.verify("ci", changed, self.plans)

    def test_scoped_checks_cannot_run_instead_of_expected_skip(self):
        for name, entry in self.plans["ci"].items():
            if "selection" in entry:
                needs = self.needs(enabled=False)
                needs[name]["result"] = "success"
                with self.subTest(name=name), self.assertRaises(ValueError):
                    result.verify("ci", needs, self.plans)

    def test_selection_must_be_complete_and_boolean_strings(self):
        for key in result.SELECTIONS:
            for value in (None, "", "yes", True, "${{ true }}"):
                needs = self.needs()
                needs["gate"]["outputs"][key] = value
                with self.subTest(key=key, value=value), self.assertRaises(ValueError):
                    result.verify("ci", needs, self.plans)

    def test_plan_is_data_and_never_an_arbitrary_command_or_workflow_path(self):
        with tempfile.TemporaryDirectory() as temporary:
            path = Path(temporary) / "plan.json"
            for field, value in (("workflow", "../../private.yml"), ("command", "echo secret"),
                                 ("selection", "untrusted"), ("task", "build; echo secret")):
                plan = copy.deepcopy(self.plans)
                plan["ci"]["rust-build"][field] = value
                path.write_text(json.dumps(plan))
                with self.subTest(field=field), self.assertRaises(ValueError):
                    result.load_plan(path)

    def test_cli_failure_prints_no_input_or_diagnostic_secrets(self):
        env = {**os.environ, "CI_KIND": "ci", "NEEDS_JSON": "private-secret"}
        run = subprocess.run([sys.executable, "-I", "-B", str(Path(result.__file__))], env=env,
                             capture_output=True, text=True, timeout=5)
        self.assertEqual(run.returncode, 1)
        self.assertNotIn("private-secret", run.stdout + run.stderr)
        self.assertEqual(run.stderr, "Public CI execution plan did not complete successfully.\n")

    def test_root_graph_cannot_omit_result_dependencies_or_ignore_failures(self):
        for kind, plan in self.plans.items():
            text = (ROOT / f".github/workflows/{kind}.yml").read_text()
            names = re.findall(r"(?m)^  ([a-z][a-z0-9-]*):$", text.split("\njobs:\n")[1])
            self.assertEqual(names, ["gate", *plan, "result"])
            tail = text.split("\n  result:\n")[1]
            self.assertIn("    if: always()\n", tail)
            self.assertIn("    needs: [gate, " + ", ".join(plan) + "]\n", tail)
            self.assertNotIn("continue-on-error", text)
            self.assertIn("NEEDS_JSON: ${{ toJSON(needs) }}", tail)
            for entry in plan.values():
                self.assertIn("uses: ./.github/workflows/" + entry["workflow"], text)

    def test_security_audits_are_unconditional_and_every_selection_is_consumed(self):
        for name in ("source", "rust-audit", "npm-audit"):
            self.assertNotIn("selection", self.plans["ci"][name])
        self.assertEqual({entry["selection"] for entry in self.plans["ci"].values() if "selection" in entry},
                         result.SELECTIONS)
        for entry in self.plans["dependency-security"].values():
            self.assertNotIn("selection", entry)

    def test_browser_matrix_requires_all_shards_and_rejects_empty_suites(self):
        text = (ROOT / ".github/workflows/browser.yml").read_text()
        self.assertIn("shard: [1, 2, 3]", text)
        self.assertIn("fail-fast: false", text)
        self.assertNotIn("continue-on-error", text)
        self.assertIn("ci_browser_ui_timings.sh", text)
        self.assertIn('BROWSER_UI_SHARD: ${{ matrix.shard }}/3', text)

    def test_private_scan_requires_protected_token_and_never_publishes_reports(self):
        text = (ROOT / ".github/workflows/private-scan.yml").read_text()
        root = (ROOT / ".github/workflows/private-image.yml").read_text()
        self.assertIn("      GHCR_TOKEN: ${{ secrets.GHCR_TOKEN }}", root)
        self.assertIn("      GHCR_TOKEN:\n        required: true", text)
        self.assertIn("environment: core-read", text)
        self.assertIn("REGISTRY_USER: cagojeiger", text)
        self.assertIn("REGISTRY_TOKEN: ${{ secrets.GHCR_TOKEN }}", text)
        self.assertIn("NATIVE_REQUEST_DEADLINE: ${{ steps.worker.outputs.deadline }}", text)
        self.assertIn("bash .control/.github/scripts/private_run.sh .control/.github/scripts/ci_private_image_check.sh", text)
        self.assertNotIn("upload-artifact", text)
        self.assertNotIn("upload-sarif", text)


if __name__ == "__main__":
    unittest.main()
