#!/bin/bash
set -euo pipefail
SOURCE_SHA=$(git -C core rev-parse HEAD)
version=$(cat core/apps/desktop/VERSION)
tag=quick-v$version
dmg=NoteHub-Quick-$version.dmg
feed="$RUNNER_TEMP/feed"
repo="repos/$GITHUB_REPOSITORY"
notes="NoteHub Quick $version (core $SOURCE_SHA)"
branch=desktop-updates

# 버전별 릴리즈에 원본 피드도 보관해 게시 도중 실패하면 같은 파일로 복구한다.
if gh release view "$tag" --json body > "$RUNNER_TEMP/release.json"; then
    test "$(jq -r .body "$RUNNER_TEMP/release.json")" = "$notes"
else
    gh release create "$tag" --draft --latest=false --target "$GITHUB_SHA" \
        --title "NoteHub Quick $version" --notes "$notes"
fi
if [[ $(gh release view "$tag" --json isDraft --jq .isDraft) = true ]]; then
    gh release upload "$tag" "$feed/$dmg" "$feed/appcast.xml" --clobber
    gh release edit "$tag" --draft=false --latest=false
fi
mkdir "$RUNNER_TEMP/published"
gh release download "$tag" -p appcast.xml -p "$dmg" -D "$RUNNER_TEMP/published"
cd "$RUNNER_TEMP/published"
# 공개 다운로드에서 받은 파일의 길이와 Ed25519 서명을 피드와 대조한다.
python3 - "$version" "$dmg" <<'PY'
from pathlib import Path
import sys
import xml.etree.ElementTree as ET
version, name = sys.argv[1:]
ns = "{http://www.andymatuschak.org/xml-namespaces/sparkle}"
item = ET.parse("appcast.xml").find("./channel/item")
assert item.findtext(ns + "version") == version
entry = item.find("enclosure")
assert entry.get("url") == f"https://github.com/cagojeiger/notehub/releases/download/quick-v{version}/{name}"
assert int(entry.attrib["length"]) == Path(name).stat().st_size
Path("signature").write_text(entry.attrib[ns + "edSignature"])
PY
# 키는 빌드에 고정된 공개 키로 검증한다.
swift "$GITHUB_WORKSPACE/.github/scripts/verify-desktop.swift" \
    "$GITHUB_WORKSPACE/core/apps/desktop/sparkle-public-key" signature "$dmg"

if ! gh api "$repo/git/ref/heads/$branch" --silent; then
    gh api "$repo/git/refs" -f ref="refs/heads/$branch" -f sha="$GITHUB_SHA" --silent
fi
args=()
if gh api "$repo/contents/appcast.xml?ref=$branch" > current.json; then
    jq -r .content current.json | base64 -D > current.xml
    cmp -s current.xml appcast.xml && exit 0
    python3 - "$version" <<'PY'
import sys
import xml.etree.ElementTree as ET
old = ET.parse("current.xml").findtext(".//{http://www.andymatuschak.org/xml-namespaces/sparkle}version")
assert tuple(map(int, old.split("."))) < tuple(map(int, sys.argv[1].split("."))), "feed version must increase"
PY
    args=(-f "sha=$(jq -r .sha current.json)")
fi
gh api --method PUT "$repo/contents/appcast.xml" -f branch="$branch" \
    -f message="chore(desktop): publish $version update feed" \
    -f "content=$(base64 < appcast.xml | tr -d '\n')" ${args[@]+"${args[@]}"} --silent
