# D2 reference · D2 작성 레퍼런스

[English authoring guide](document-authoring.md) · [한국어 작성 가이드](document-authoring.ko.md)

Write raw D2 in a `.d2` text document. This reference covers NoteHub renderer **0.1.1**, D2 **0.9.0**, and the **Dagre** layout engine. File imports are disabled. · `.d2` 텍스트 문서에 원문을 작성하세요. 아래 목록은 NoteHub 렌더러 **0.1.1**에 내장된 D2 **0.9.0**과 **Dagre** 기준입니다. 파일 import는 차단됩니다. [Renderer configuration · 렌더러 설정](../../services/renderer/render.go) · [D2 dependency · 버전](../../services/renderer/go.mod)

- [Shapes · 도형](#shapes)
- [Connections and arrowheads · 연결선과 화살촉](#connections)
- [Layout · 배치](#layout)

## Shapes

Set `shape` on a node. A node with no explicit shape uses `rectangle`. · 노드의 `shape`에 아래 값을 사용합니다. 생략하면 직사각형입니다.

| Use · 용도 | Values · 작성할 값 | Official syntax and examples · 공식 문법과 예제 |
| --- | --- | --- |
| Geometry · 기본 도형 | `rectangle`, `square`, `oval`, `circle`, `diamond`, `hexagon`, `parallelogram` | [Shape catalog](https://d2lang.com/tour/shapes/) |
| Documents and annotations · 문서·주석 | `page`, `document`, `callout` | [Shape examples](https://d2lang.com/tour/shapes/#example) |
| Storage and infrastructure · 저장소·인프라 | `cylinder`, `stored_data`, `queue`, `package`, `cloud` | [Shape catalog](https://d2lang.com/tour/shapes/) |
| People and steps · 사용자·단계 | `person`, `c4-person`, `step` | [Shape catalog](https://d2lang.com/tour/shapes/) |
| Text and code · 텍스트·코드 | `text`, `code` | [Text and code blocks](https://d2lang.com/tour/text/) |
| UML class and database table · 클래스·테이블 | `class`, `sql_table` | [UML classes](https://d2lang.com/tour/uml-classes/) · [SQL tables](https://d2lang.com/tour/sql-tables/) |
| Image · 이미지 | `image` (set `icon` to the image URL · `icon`에 이미지 URL 지정) | [Icons and images](https://d2lang.com/tour/icons/) |
| Sequence · 메시지 순서 | `sequence_diagram` | [Sequence diagram examples](https://d2lang.com/tour/sequence-diagrams/) |
| Hierarchy container · 계층 컨테이너 | `hierarchy` | [v0.9.0 shape registry](https://github.com/d2lang/d2/blob/v0.9.0/d2target/d2target.go#L1244-L1298) · [Containers](https://d2lang.com/tour/containers/) |

```d2
reader: Reader {shape: person}
store: Archive {shape: cylinder}
reader -> store: Save note
```

An `image` shape can compile to SVG containing an image reference. Displaying that resource depends on the viewer's image policy. The `hierarchy` shape is accepted by the compiler; NoteHub uses Dagre for graph placement. · 이미지 도형은 이미지 참조를 담은 SVG를 생성할 수 있으며 실제 표시는 뷰어 정책을 따릅니다. `hierarchy`는 컴파일러가 인식하는 도형이며 NoteHub의 그래프 배치는 Dagre가 담당합니다.

## Connections

Connect **node keys**. Add a label after `:`. · 양끝에는 **노드 키**를 쓰고 `:` 뒤에 선의 라벨을 붙입니다. [Connection syntax and examples · 연결선 문법과 예제](https://d2lang.com/tour/connections/)

| D2 | Result · 표시 |
| --- | --- |
| `a -- b` | Line without arrowheads · 화살촉 없는 선 |
| `a -> b` | Arrow toward b · b 방향 |
| `a <- b` | Arrow toward a · a 방향 |
| `a <-> b` | Arrows at both ends · 양방향 |

Choose an arrowhead on an end that has an arrow. Set `source-arrowhead.shape` on the source or `target-arrowhead.shape` on the target. · 화살촉이 있는 끝에 `source-arrowhead.shape` 또는 `target-arrowhead.shape`를 지정합니다. [Arrowhead examples · 화살촉 예제](https://d2lang.com/tour/connections/#arrowheads)

| Use · 용도 | Values · 작성할 값 |
| --- | --- |
| Standard ends · 일반 끝 모양 | `triangle` (default · 기본), `arrow`, `diamond`, `circle`, `box`, `cross`, `none` |
| Entity cardinality · 관계의 개수 | `cf-one`, `cf-one-required`, `cf-many`, `cf-many-required` |

```d2
author <-> note: owns {
  source-arrowhead.shape: cf-one-required
  target-arrowhead.shape: cf-many
  style.stroke-dash: 4
}
```

Use `style.filled: false` on a triangle arrowhead for an outline, or `style.filled: true` on a diamond, circle, or box for a solid end. Edge styles such as `style.stroke-dash` belong on the connection. · 화살촉 내부의 `style.filled`로 채움을 설정하고, 점선 등의 스타일은 연결선에 설정합니다. [Styles](https://d2lang.com/tour/style/) · [Chains, repeated edges and indexed references](https://d2lang.com/tour/connections/)

## Layout

NoteHub resolves graph layout to **Dagre**. The source can choose flow direction and organize nodes into containers, grids, or sequence diagrams. ELK and TALA are upstream engine options; this service always selects Dagre. · NoteHub에서 선택할 수 있는 그래프 배치 엔진은 **Dagre 하나**입니다. D2 원문으로 진행 방향과 컨테이너·격자·시퀀스 배치를 지정할 수 있습니다. [Dagre examples](https://d2lang.com/tour/dagre/) · [Upstream engine comparison](https://d2lang.com/tour/layouts/)

| Intent · 원하는 배치 | Source · 작성 규칙 | Official examples · 공식 예제 |
| --- | --- | --- |
| Flow direction · 진행 방향 | Top-level `direction: down` (default · 기본), `up`, `left`, `right` | [Layouts](https://d2lang.com/tour/layouts/) |
| Group nodes · 노드 묶기 | Nest children in `group: { ... }`; connect by path, e.g. `group.api` | [Containers](https://d2lang.com/tour/containers/) |
| Grid · 격자 | Set `grid-columns` or `grid-rows` on the grid | [Grid diagrams](https://d2lang.com/tour/grid-diagrams/) |
| Ordered messages · 순서가 있는 메시지 | Set `shape: sequence_diagram`; write the messages in order | [Sequence diagrams](https://d2lang.com/tour/sequence-diagrams/) |
| Anchor a title or legend · 제목·범례 위치 | Use a constant such as `near: top-center` on a separate node | [Positioning examples](https://d2lang.com/tour/positions/) |

```d2
direction: right
client -> services.api: Request
services: {
  api -> db
  db.shape: cylinder
}
```

For a complete Markdown document embedding a diagram, see the [example README](examples/request-flow/README.md) and its [D2 source](examples/request-flow/flow.d2). · 본문에 삽입하는 방법은 두 예제 파일을 함께 보세요.
