#!/usr/bin/env bash
set -euo pipefail

image=${1:?image required}
platform=${2:?platform required}
container=$(docker run --detach --platform "$platform" --publish 127.0.0.1::8080 "$image")
work=$(mktemp -d)
cleanup() {
  local status=$?
  if (( status != 0 )); then
    docker logs "$container" >&2 || true
  fi
  docker rm -f "$container" >/dev/null || true
  rm -rf "$work"
}
trap cleanup EXIT
address=$(docker port "$container" 8080/tcp)

curl --fail-with-body --silent --show-error --max-time 10 \
  --retry 8 --retry-all-errors --retry-delay 1 \
  --header 'Content-Type: text/plain; charset=utf-8' \
  --data-binary @testdata/architecture-1024.d2 \
  --dump-header "$work/headers" --output "$work/diagram.svg" \
  "http://$address/render/d2"
grep -qi '^Content-Type: image/svg+xml' "$work/headers"
python3 - "$work/diagram.svg" <<'PY'
import re
import sys
import xml.etree.ElementTree as ET
from pathlib import Path

source = Path("testdata/architecture-1024.d2").read_text()
labels = re.findall(r"^\s+n\d+: (.+)$", source, re.MULTILINE)
diagram = ET.parse(sys.argv[1]).getroot()
namespace = "{http://www.w3.org/2000/svg}"
actual = {"".join(element.itertext()) for element in diagram.iter(namespace + "text")}
assert diagram.tag == namespace + "svg"
assert len(set(labels)) == 1024 and set(labels) <= actual
print("Rendered and checked all 1,024 D2 node labels.")
PY
