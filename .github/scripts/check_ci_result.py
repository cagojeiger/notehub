#!/usr/bin/env python3
"""Public 작업 계획의 완료를 확인하고 워크플로 자체의 성공·실패로 전달한다."""
import json
import os
from pathlib import Path
import re
import sys


SELECTIONS = {"rust_build", "rust", "web", "browser", "postgres", "image"}


def load_plan(path):
    plans = json.loads(path.read_text())
    if type(plans) is not dict or set(plans) != {"ci", "dependency-security", "security-dast", "private-image"}:
        raise ValueError("Invalid execution plans")
    for plan in plans.values():
        if type(plan) is not dict or not plan:
            raise ValueError("Empty execution plan")
        for name, item in plan.items():
            if (not re.fullmatch(r"[a-z][a-z0-9-]*", name) or name in {"gate", "result"}
                    or type(item) is not dict or not {"name", "workflow"} <= set(item)
                    or set(item) - {"name", "workflow", "task", "selection"}
                    or not isinstance(item["name"], str) or not item["name"].isprintable()
                    or not item["name"] or not re.fullmatch(r"[a-z][a-z0-9-]*\.yml", item["workflow"])
                    or "selection" in item and item["selection"] not in SELECTIONS
                    or (item["workflow"] == "rust.yml") != ("task" in item)
                    or "task" in item and not re.fullmatch(r"[a-z][a-z0-9-]*", item["task"])):
                raise ValueError("Invalid execution plan entry")
    return plans


def verify(kind, needs, plans):
    plan = plans[kind]
    if type(needs) is not dict or set(needs) != {"gate", *plan}:
        raise ValueError("Incomplete execution results")
    if needs["gate"]["result"] != "success":
        raise ValueError("Request gate did not succeed")
    selection = needs["gate"]["outputs"]
    if kind == "ci" and any(selection.get(key) not in {"true", "false"} for key in SELECTIONS):
        raise ValueError("Incomplete job selection")
    for name, item in plan.items():
        enabled = "selection" not in item or selection[item["selection"]] == "true"
        if needs[name]["result"] != ("success" if enabled else "skipped"):
            raise ValueError("Job result does not match the execution plan")


def main():
    try:
        plans = load_plan(Path(__file__).resolve().parents[1] / "ci-plan.json")
        verify(os.environ["CI_KIND"], json.loads(os.environ["NEEDS_JSON"]), plans)
    except (ValueError, TypeError, KeyError, OSError):
        print("Public CI execution plan did not complete successfully.", file=sys.stderr)
        return 1
    print("Public CI execution plan completed successfully.")
    return 0


if __name__ == "__main__":
    sys.exit(main())
