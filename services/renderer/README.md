# Renderer

D2 원문을 SVG로 변환하는 Go HTTP 서비스입니다. 요청을 받은 Go 프로세스가 [D2 v0.9.0](https://d2lang.com/releases/0.9.0/)의 라이브러리를 직접 호출합니다.

## 실행

`services/renderer/`에서 실행합니다. Go 1.27.4 이상이 필요합니다.

```sh
go run .
```

기본 주소는 `127.0.0.1:8080`입니다. 컨테이너 이미지의 기본 주소는 `:8080`이며 `LISTEN_ADDR`로 바꿀 수 있습니다. API 인증 없이 동작하므로 운영에서는 허용된 서비스가 접근하는 내부 네트워크에 둡니다.

```sh
docker run --rm -p 127.0.0.1:8080:8080 ghcr.io/cagojeiger/d2-diagram-renderer:v0.1.0
```

## API

[OpenAPI 3.1.1 명세](openapi.yaml)를 제공합니다.

```sh
curl --fail-with-body \
  -H 'Content-Type: text/plain; charset=utf-8' \
  --data-binary @testdata/architecture-1024.d2 \
  http://127.0.0.1:8080/render/d2 \
  -o architecture-1024.svg
```

성공은 `200 OK`와 `image/svg+xml` 본문을 반환합니다. 오류는 `{"code":"invalid_diagram","message":"…"}` 형태의 JSON입니다. 잘못된 D2는 `422`, 진행 중인 렌더 때문에 실행할 수 없으면 `503`, 응답 기한을 넘기면 `504`를 반환합니다.

한 번에 한 요청을 변환합니다. 입력은 128 KiB, 반환할 SVG는 2 MiB로 제한합니다. 응답 기한은 변환을 시작한 시점부터 3초입니다. D2 내부 계산이 취소 신호를 즉시 처리하지 않는 구간에서는 HTTP 응답이 끝나도 계산 종료까지 변환 자리를 점유합니다.

D2 파일 import는 비활성화합니다. 반환한 SVG는 브라우저에서 `<img>`로 표시할 수 있습니다. HTTP 응답의 콘텐츠 보안 정책은 그림을 직접 열었을 때 스크립트와 외부 리소스 요청을 제한합니다.

`testdata/architecture-1024.d2`는 각 영역에 128개씩, 총 1,024개 서비스 노드를 담은 합성 샘플입니다.

```sh
go test ./...
```

## 릴리즈

새 버전은 `VERSION`과 `openapi.yaml`의 `info.version`에 적습니다. 동일한
main 커밋의 **Renderer CI**가 통과하면 **Renderer Release**에서 main과 버전을
선택해 실행합니다. 렌더러 관련 변경이 없는 main 커밋은 **Renderer CI**를
수동으로 실행할 수 있습니다.

Linux amd64·arm64 이미지에서 각각 1,024개 노드 샘플을 렌더링합니다.
테스트한 digest를 기존 공개 패키지 `ghcr.io/cagojeiger/d2-diagram-renderer`에
`vX.Y.Z` 태그로 발행합니다. GitHub 릴리즈는 `renderer/vX.Y.Z` 태그를
사용합니다. 이미지에는 SBOM과 빌드 provenance를 첨부합니다.

공개 렌더러 패키지의 [Manage Actions access](https://docs.github.com/en/packages/learn-github-packages/configuring-a-packages-access-control-and-visibility#github-actions-access-for-packages-scoped-to-personal-accounts)에서
`cagojeiger/notehub`에 Write 권한을 부여해야 발행할 수 있습니다.

## 라이선스

[Apache-2.0](LICENSE). D2의 라이선스와 소스는 [제3자 고지](THIRD_PARTY_NOTICES.md)에 있습니다.
