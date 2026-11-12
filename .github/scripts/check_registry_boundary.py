"""공개 리포의 Actions 토큰이 비공개 이미지를 읽을 수 없는지 확인한다."""
import base64
import json
import os
from urllib.error import HTTPError
from urllib.request import Request, urlopen

image = os.environ["REGISTRY_IMAGE"]
if image not in {"project-jelly/cloud", "cagojeiger/notehub"}:
    raise SystemExit("Unexpected private image")
expected_access = os.environ.get("EXPECT_REGISTRY_ACCESS") == "true"
actor = "cagojeiger" if expected_access else os.environ["GITHUB_ACTOR"]
credential = os.environ.get("REGISTRY_TOKEN") if expected_access else os.environ.get("GH_TOKEN")
if not credential:
    raise SystemExit("Missing registry verification credential")
credentials = base64.b64encode(f"{actor}:{credential}".encode()).decode()
request = Request(f"https://ghcr.io/token?service=ghcr.io&scope=repository:{image}:pull",
                  headers={"Authorization": "Basic " + credentials})
try:
    with urlopen(request, timeout=30) as response:
        token = json.load(response)["token"]
    request = Request(f"https://ghcr.io/v2/{image}/manifests/latest", headers={
        "Authorization": "Bearer " + token,
        "Accept": "application/vnd.oci.image.index.v1+json, application/vnd.docker.distribution.manifest.list.v2+json"})
    with urlopen(request, timeout=30):
        pass
except HTTPError as error:
    status = error.code
    error.close()
    if status not in {401, 403}:
        raise SystemExit(f"Private registry boundary could not be verified: HTTP {status}") from None
    if expected_access:
        raise SystemExit("Protected registry token could not read the private image.") from None
    print("Public Actions token cannot read the private image.")
else:
    if not expected_access:
        raise SystemExit("Public repository still has private image access; remove its package Actions grant.")
    print("Protected main environment can read the private image.")

