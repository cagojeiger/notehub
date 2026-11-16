# NoteHub

NoteHub is a workspace for organizing, editing, and searching notes and documents.

## NoteHub CLI

macOS·Linux의 arm64·x86_64에서 사용할 수 있다. 공식 설치 프로그램이 운영체제와 CPU에 맞는 실행 파일을 다운로드하고 체크섬과 버전을 검증한다.

```sh
curl -fsSL https://github.com/cagojeiger/notehub/releases/latest/download/notehub-cli-installer.sh | sh
export PATH="$HOME/.local/bin:$PATH"
notehub-cli auth login
```

기본 서비스 주소는 `https://note.project-jelly.io`다. 로그인 명령이 안내하는 링크를 브라우저에서 열고 본인 계정의 접근을 승인한다. 로그인 정보는 OS 키체인에 저장한다.

실행 파일은 `~/.local/bin/notehub-cli`에 설치된다. 위 PATH 설정은 현재 터미널에 적용되므로, 이후에도 사용하려면 사용 중인 셸 설정에 같은 PATH 설정을 추가한다.

```sh
notehub-cli update --check
notehub-cli update
notehub-cli rollback
```

공식 설치 프로그램으로 설치한 경우 자체 업데이트를 사용할 수 있다. `update`는 새 파일 검증과 직전 버전 백업 후 실행 파일을 교체한다. `rollback`은 저장된 직전 버전으로 복구하고, 성공하면 백업을 소비한다. 백업은 업데이트 시 생성되며 자동 업데이트·자동 롤백은 하지 않는다.

[모든 플랫폼의 다운로드 파일](https://github.com/cagojeiger/notehub/releases/latest)
