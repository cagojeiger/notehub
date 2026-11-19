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
nodes=1024
python3 - "$nodes" > "$work/diagram.d2" <<'PYTHON'
import sys

print("grid-columns: 32")
for index in range(1, int(sys.argv[1]) + 1):
    print(f"n{index}: Node {index:04d}")
PYTHON

curl --fail-with-body --silent --show-error --max-time 10 \
  --retry 8 --retry-all-errors --retry-delay 1 \
  --header 'Content-Type: text/plain; charset=utf-8' \
  --data-binary "@$work/diagram.d2" \
  --dump-header "$work/headers" --output "$work/diagram.svg" \
  "http://$address/render/d2"
grep -qi '^Content-Type: image/svg+xml' "$work/headers"
python3 - "$work/diagram.svg" "$nodes" <<'PY'
import sys
import xml.etree.ElementTree as ET

labels = {f"Node {index:04d}" for index in range(1, int(sys.argv[2]) + 1)}
diagram = ET.parse(sys.argv[1]).getroot()
namespace = "{http://www.w3.org/2000/svg}"
actual = {"".join(element.itertext()) for element in diagram.iter(namespace + "text")}
assert diagram.tag == namespace + "svg"
assert labels <= actual
print(f"Rendered and checked all {len(labels):,} D2 node labels.")
PY
