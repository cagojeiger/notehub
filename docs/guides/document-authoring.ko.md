# NoteHub 문서 작성 가이드

[English](document-authoring.md)

노트는 `.md` 또는 `.markdown` 텍스트 문서로 저장합니다. 표, 작업 목록, 각주는 GitHub Flavored Markdown 문법을 사용합니다.

| 작성할 내용 | 작성 방법 |
| --- | --- |
| 본문, 목록, 표, 링크 | Markdown 텍스트 문서 |
| 노트 안의 흐름도나 시퀀스 다이어그램 | `mermaid` 코드 블록 |
| 노트 안의 차트나 인포그래픽 | `infographic` 코드 블록 |
| 독립적으로 열거나 여러 문서에서 재사용할 그림 | `.d2` 텍스트 문서 |
| Markdown 안에 표시할 D2 그림 | `![설명](./flow.d2)` |

문서 속성은 파일 첫 줄부터 두 `---` 구분선 사이에 YAML 키와 값을 작성합니다. 접을 수 있는 본문은 다음과 같이 작성합니다.

```markdown
<details>
<summary>회의 메모</summary>

- 제안서 검토

</details>
```

## 링크, 첨부, 폴더 첫 화면

상대 경로는 현재 문서의 폴더를 기준으로 합니다. `../`는 상위 폴더, `/`는 현재 Space의 루트를 가리킵니다. 확장자를 포함한 실제 이름을 사용합니다.

```markdown
[프로젝트 폴더](./project/)
[프로젝트 소개 문서](./project/README.md)
[Space 목차](/README.md)
[요약으로 이동](#요약)
![시스템 개요](./assets/overview.png)
```

폴더 바로 아래에 `README.md`를 두면 웹에서 해당 폴더나 Space 루트를 열 때 기본 탭에 본문이 표시됩니다. 폴더 링크는 파일 목록과 README가 있는 폴더 화면을 열고, README 링크는 문서를 직접 엽니다. Incoming은 링크가 가리킨 대상별로 집계됩니다. 관련 자료를 묶을 때는 한 폴더에 두고 README에서 각 자료로 연결합니다.

같은 Space에 올려 둔 음성을 본문에서 재생하려면 `audio` 블록에 파일 경로 한 줄을 넣습니다.

````markdown
```audio
./recording.webm
```
````

## Mermaid

노트에 흐름도를 넣을 때는 `mermaid` 블록에 소스를 작성합니다.

```mermaid
flowchart LR
    Inbox --> Review
    Review --> Archive
```

다른 유형은 [Mermaid 문법](https://mermaid.js.org/intro/syntax-reference.html)을 참고하세요.

## AntV Infographic

`infographic` 블록 첫 줄에 `infographic <템플릿>`을 쓰고 `data` 아래를 들여씁니다. 막대 차트는 `values`에 각 항목의 이름과 숫자를 넣습니다.

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

템플릿별 입력은 [AntV 데이터 문법](https://infographic.antv.vision/learn/data)을 참고하세요.

## 독립 D2 문서 작성과 본문 삽입

[flow.d2](examples/request-flow/flow.d2)의 D2 원문을 코드 펜스 없이 텍스트 문서로 저장합니다. [예제 README.md](examples/request-flow/README.md)를 같은 폴더에 둡니다. NoteHub에서 README를 열면 이미지는 SVG로 표시되고 일반 링크는 독립 다이어그램 뷰어를 엽니다.

```markdown
![Request flow](./flow.d2)

[Open the diagram](./flow.d2)
```

독자에게 소스 코드를 보여줄 때는 `d2` 코드 블록을 사용합니다. 그림을 보여줄 때는 위의 이미지 문법을 사용합니다. CLI·MCP 명령에서 파일을 지정할 때는 `<space>:/folder/flow.d2`, `<space>:/folder/README.md`를 쓰고 README 본문에서는 상대 경로를 씁니다.
