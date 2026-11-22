# NoteHub 문서 작성 가이드

[English](document-authoring.md)

## 목차

- [Markdown 파일 작성 가이드](#markdown-파일-작성-가이드)
  - [기본 노트 작성하기](#기본-노트-작성하기)
  - [문서와 폴더 연결하기](#문서와-폴더-연결하기)
  - [이미지와 음성 넣기](#이미지와-음성-넣기)
  - [Mermaid와 Infographic 작성하기](#mermaid와-infographic-작성하기)
  - [전체 Mermaid 종류와 공식 예제](mermaid-diagrams.md)
  - [전체 Infographic 템플릿 목록](infographic-templates.md)
- [D2 파일 작성 가이드](#d2-파일-작성-가이드)
  - [지원 도형·연결선·레이아웃과 공식 예제](d2-reference.md)

## Markdown 파일 작성 가이드

### 기본 노트 작성하기

다음 Markdown을 `notes.md`라는 텍스트 문서로 저장하면 제목과 작업 목록이 표시됩니다. 노트에는 `.md` 또는 `.markdown`을 사용합니다. 표와 각주는 GitHub Flavored Markdown 문법을 따릅니다.

```markdown
# 회의 메모

- [ ] 제안서 검토
```

문서 속성은 파일 첫 줄부터 두 `---` 구분선 사이에 YAML 키와 값을 작성합니다. 접을 수 있는 본문은 다음과 같이 작성합니다.

```markdown
<details>
<summary>회의 메모</summary>

- 제안서 검토

</details>
```

### 문서와 폴더 연결하기

MCP와 CLI 명령에서는 `daily:/project/README.md`처럼 Space 이름을 포함한 경로를 사용합니다. Markdown 본문에서는 같은 Space 안의 경로를 씁니다. `./`와 `../`는 현재 문서의 폴더 기준이고, `/`는 Space 루트 기준입니다. 파일명과 확장자를 정확히 적습니다.

Space 루트의 `README.md`에서 아래 두 링크는 각각 프로젝트 폴더와 README 문서를 엽니다. 웹의 폴더 화면에서는 파일 목록과 해당 폴더의 README가 보입니다. Incoming 링크는 본문에서 가리킨 대상별로 집계됩니다.

```markdown
[Project folder](./project/)
[Project README](./project/README.md)
```

### 이미지와 음성 넣기

먼저 첨부 파일을 같은 Space에 업로드합니다. 노트가 있는 폴더 아래에 `assets/overview.png`와 `recordings/meeting.webm`가 있다면 아래 원문을 노트에 넣어 이미지와 재생기를 표시합니다. `audio` 블록에는 파일 경로 한 줄을 씁니다.

````markdown
![System overview](./assets/overview.png)

```audio
./recordings/meeting.webm
```
````

### Mermaid와 Infographic 작성하기

여닫는 코드 펜스까지 Markdown 본문에 넣습니다. 흐름도나 시퀀스 다이어그램에는 `mermaid`를 사용합니다.

````markdown
```mermaid
flowchart LR
    Inbox --> Review
    Review --> Archive
```
````

[현재 버전의 전체 Mermaid 종류·선언·공식 예제](mermaid-diagrams.md)

코드 펜스의 언어는 `infographic`이고 블록 안의 첫 줄은 `infographic <템플릿 ID>`입니다. 수치 차트, 목록, 순서, 계층, 비교, 관계 중에서 분류를 고르고 해당 템플릿의 데이터 키를 사용합니다. 아래 막대 차트는 `values`를 사용합니다. [전체 ID 및 분류별 작성 예제](infographic-templates.md)

````markdown
```infographic
infographic chart-bar-plain-text
data
  title 이번 주 기록
  desc 종류별로 작성한 문서 수
  values
    - label 회의록
      value 12
    - label 작업 노트
      value 18
```
````

템플릿 모양은 [AntV 공식 갤러리](https://infographic.antv.vision/gallery)에서 확인하세요. 문법 참고 자료는 [Mermaid](https://mermaid.js.org/intro/syntax-reference.html) · [AntV Infographic](https://infographic.antv.vision/learn/data)

## D2 파일 작성 가이드

[flow.d2](examples/request-flow/flow.d2)를 텍스트 문서로 저장합니다. 파일 본문에는 Markdown 펜스 없이 D2 원문을 그대로 넣습니다.

```d2
client: Browser
api: NoteHub
client -> api: Send request
```

같은 폴더에 [README.md](examples/request-flow/README.md)를 다음 내용으로 저장합니다.

```markdown
# Request flow

![Request flow](./flow.d2)

[Open the diagram](./flow.d2)
```

이미지 문법은 노트 안에 `flow.d2`를 SVG로 표시합니다. 일반 링크를 클릭하면 같은 파일을 독립 다이어그램 뷰어에서 엽니다. Markdown 안의 `d2` 코드 블록은 소스 코드를 표시합니다.

[지원 도형·연결선·레이아웃과 공식 예제](d2-reference.md)
