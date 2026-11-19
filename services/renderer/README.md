# Renderer

D2 원문을 SVG로 변환하는 Go HTTP 서비스입니다.

```sh
cd services/renderer
go run .
```

기본 주소는 `127.0.0.1:8080`입니다. `LISTEN_ADDR`로 변경할 수 있습니다. 인증이 없는 서비스이므로 운영에서는 내부 네트워크에 둡니다.

```sh
curl --fail-with-body -H 'Content-Type: text/plain' --data-binary 'a -> b' \
  http://127.0.0.1:8080/render/d2 -o diagram.svg
```

요청·응답은 [OpenAPI 명세](openapi.yaml)를 참고하세요. 릴리스는 [VERSION](VERSION)을 수정하고 main의 Renderer CI가 통과한 뒤 [Renderer Release](../../.github/workflows/renderer-release.yml)를 실행합니다.

[Apache-2.0](LICENSE) · [제3자 고지](THIRD_PARTY_NOTICES.md)
