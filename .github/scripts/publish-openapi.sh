#!/usr/bin/env bash
set -euo pipefail

version=$1
schema=$2
[[ "$version" =~ ^(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)$ ]]
jq -e --arg version "$version" '
  (.info.version == $version) and
  (.openapi | startswith("3.")) and
  (.paths | length > 0 and all(keys[]; startswith("/api/v2/"))) and
  ([.. | objects | .["$ref"]? | strings] | all(startswith("#/")))
' "$schema" >/dev/null

publish_tree=$(mktemp -d)
trap 'rm -rf "$publish_tree"' EXIT
git init --quiet "$publish_tree"
git -C "$publish_tree" remote add origin "$(git remote get-url origin)"

for attempt in 1 2 3; do
  git -C "$publish_tree" fetch --no-tags --depth=1 origin refs/heads/openapi
  git -C "$publish_tree" checkout --quiet --detach FETCH_HEAD
  version_file="$publish_tree/openapi/v2/v$version.json"

  if [[ -f "$version_file" ]]; then
    if ! cmp -s "$schema" "$version_file"; then
      echo "The OpenAPI contract for v$version already exists with different bytes." >&2
      exit 1
    fi
  else
    cp "$schema" "$version_file"
  fi

  latest_file="$publish_tree/openapi/v2/latest.txt"
  latest=$(cat "$latest_file")
  [[ "$latest" =~ ^v(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)$ ]]
  if [[ "$(printf '%s\n' "$latest" "v$version" | sort -V | tail -1)" == "v$version" ]]; then
    printf 'v%s\n' "$version" > "$latest_file"
    cat > "$publish_tree/openapi/README.md" <<README
# NoteHub API V2

[Latest deployed OpenAPI (v$version)](v2/v$version.json) · [Version history](v2)

Send an agent API key in the \`Authorization: Bearer\` header.
README
  fi

  git -C "$publish_tree" add -- openapi
  if git -C "$publish_tree" diff --cached --quiet; then
    exit 0
  fi
  git -C "$publish_tree" -c user.name='github-actions[bot]' \
    -c user.email='41898282+github-actions[bot]@users.noreply.github.com' \
    commit -qm "docs(api): publish OpenAPI for v$version"
  if git -C "$publish_tree" -c credential.helper= \
      -c 'credential.helper=!gh auth git-credential' push origin HEAD:refs/heads/openapi; then
    exit 0
  fi
  echo "OpenAPI push attempt $attempt failed. Rechecking the publication branch." >&2
done

echo 'Could not publish the OpenAPI contract after three push attempts.' >&2
exit 1
