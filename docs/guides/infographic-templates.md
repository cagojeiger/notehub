# Infographic templates · 인포그래픽 템플릿

[English authoring guide](document-authoring.md) · [한국어 작성 가이드](document-authoring.ko.md)

NoteHub uses the built-in registry from `@antv/infographic@0.2.20`. Copy an exact ID below. The six categories contain 273 template IDs. Three additional IDs are deprecated aliases.

NoteHub가 사용하는 버전은 `@antv/infographic@0.2.20`입니다. 아래 ID를 그대로 복사해 사용하세요. 6개 분류에 템플릿 ID 273개가 있고, 이전 문서와의 호환을 위한 별칭이 3개 더 있습니다.

## Choose a template · 템플릿 선택하기

| What you need · 찾는 정보 | Where to look · 확인할 곳 |
| --- | --- |
| Preview the layout · 그림 모양 확인 | [AntV gallery · 공식 갤러리](https://infographic.antv.vision/gallery). Search by template ID. 템플릿 ID로 검색하세요. |
| Pick an ID supported by NoteHub · 지원 ID 선택 | [NoteHub template index · 분류별 목록](#template-index). Copy the exact ID for the installed version. 현재 버전에 등록된 ID를 복사하세요. |
| Write the fenced block · 코드 블록 작성 | Follow the example under your category. 각 분류의 예제를 사용하세요. [AntV syntax · 공식 문법](https://infographic.antv.vision/learn/infographic-syntax) |
| Fill in labels, values, or relationships · 데이터 채우기 | [AntV data reference · 데이터 필드](https://infographic.antv.vision/learn/data) |

## Template index

| Category · 분류 | Data key | Templates |
| --- | --- | ---: |
| [Numeric charts · 수치 차트](#chart) | `values` | 11 |
| [Lists · 목록](#list) | `lists` | 29 |
| [Sequences · 순서·과정](#sequence) | `sequences` | 83 |
| [Hierarchies · 계층](#hierarchy) | `root` | 112 |
| [Comparisons and quadrants · 비교·사분면](#compare) | `compares` | 20 |
| [Relationships · 관계](#relation) | `nodes + relations` | 18 |

Each example is a complete Markdown block. IDs below it are grouped by their registered layout. · 예제는 Markdown에 넣을 전체 코드 블록입니다. 각 예제 아래에는 등록된 레이아웃별로 템플릿 ID를 모았습니다.

## chart

Numeric charts · 수치 차트

Bar · 가로 막대, column · 세로 막대, line · 꺾은선, pie · 원형, donut · 도넛, word cloud · 워드 클라우드.

````markdown
```infographic
infographic chart-bar-plain-text
data
  values
    - label Notes
      value 12
    - label Tasks
      value 8
```
````

### chart-bar

- `chart-bar-plain-text`

### chart-column

- `chart-column-simple`

### chart-line

- `chart-line-plain-text`

### chart-pie

- `chart-pie-compact-card`
- `chart-pie-donut-compact-card`
- `chart-pie-donut-pill-badge`
- `chart-pie-donut-plain-text`
- `chart-pie-pill-badge`
- `chart-pie-plain-text`

### chart-wordcloud

- `chart-wordcloud`
- `chart-wordcloud-rotate`

## list

Lists · 목록

````markdown
```infographic
infographic list-row-simple-horizontal-arrow
data
  lists
    - label Notes
    - label Tasks
```
````

### list-column

- `list-column-done-list`
- `list-column-simple-vertical-arrow`
- `list-column-vertical-icon-arrow`

### list-grid

- `list-grid-badge-card`
- `list-grid-candy-card-lite`
- `list-grid-circular-progress`
- `list-grid-compact-card`
- `list-grid-done-list`
- `list-grid-horizontal-icon-arrow`
- `list-grid-progress-card`
- `list-grid-ribbon-card`
- `list-grid-simple`

### list-pyramid

- `list-pyramid-badge-card`
- `list-pyramid-compact-card`
- `list-pyramid-rounded-rect-node`

### list-row

- `list-row-circular-progress`
- `list-row-horizontal-icon-arrow`
- `list-row-horizontal-icon-line`
- `list-row-simple-horizontal-arrow`
- `list-row-simple-illus`

### list-sector

- `list-sector-half-plain-text`
- `list-sector-plain-text`
- `list-sector-simple`

### list-waterfall

- `list-waterfall-badge-card`
- `list-waterfall-compact-card`

### list-zigzag-down

- `list-zigzag-down-compact-card`
- `list-zigzag-down-simple`

### list-zigzag-up

- `list-zigzag-up-compact-card`
- `list-zigzag-up-simple`

## sequence

Sequences · 순서·과정

````markdown
```infographic
infographic sequence-timeline-simple
data
  sequences
    - label Draft
    - label Publish
```
````

### sequence-ascending-stairs-3d

- `sequence-ascending-stairs-3d-simple`
- `sequence-ascending-stairs-3d-underline-text`

### sequence-ascending-steps

- `sequence-ascending-steps`

### sequence-circle-arrows

- `sequence-circle-arrows-indexed-card`

### sequence-circular

- `sequence-circular-simple`
- `sequence-circular-underline-text`

### sequence-color-snake-steps

- `sequence-color-snake-steps-horizontal-icon-line`
- `sequence-color-snake-steps-simple-illus`

### sequence-cylinders-3d

- `sequence-cylinders-3d-simple`

### sequence-filter-mesh

- `sequence-filter-mesh-simple`
- `sequence-filter-mesh-underline-text`

### sequence-funnel

- `sequence-funnel-simple`

### sequence-horizontal-zigzag

- `sequence-horizontal-zigzag-horizontal-icon-line`
- `sequence-horizontal-zigzag-plain-text`
- `sequence-horizontal-zigzag-simple`
- `sequence-horizontal-zigzag-simple-horizontal-arrow`
- `sequence-horizontal-zigzag-simple-illus`
- `sequence-horizontal-zigzag-underline-text`

### sequence-interaction

Each entry in `sequences` is a lane. Give its steps IDs under `children`, then connect two step IDs with a labeled message in `relations`.
· `sequences`의 각 항목이 레인입니다. 각 레인의 단계를 `children`에 ID와 함께 작성하고, `relations`에서 두 단계의 ID를 메시지로 연결하세요.

````markdown
```infographic
infographic sequence-interaction-compact-badge-card
data
  sequences
    - label Browser
      children
        - id submit
          label Submit
    - label Server
      children
        - id store
          label Store
  relations
    - from submit
      to store
      label Save note
```
````

- `sequence-interaction-compact-animated-badge-card`
- `sequence-interaction-compact-animated-capsule-item`
- `sequence-interaction-compact-animated-compact-card`
- `sequence-interaction-compact-animated-rounded-rect-node`
- `sequence-interaction-compact-badge-card`
- `sequence-interaction-compact-capsule-item`
- `sequence-interaction-compact-compact-card`
- `sequence-interaction-compact-dashed-badge-card`
- `sequence-interaction-compact-dashed-capsule-item`
- `sequence-interaction-compact-dashed-compact-card`
- `sequence-interaction-compact-dashed-rounded-rect-node`
- `sequence-interaction-compact-rounded-rect-node`
- `sequence-interaction-default-animated-badge-card`
- `sequence-interaction-default-animated-capsule-item`
- `sequence-interaction-default-animated-compact-card`
- `sequence-interaction-default-animated-rounded-rect-node`
- `sequence-interaction-default-badge-card`
- `sequence-interaction-default-capsule-item`
- `sequence-interaction-default-compact-card`
- `sequence-interaction-default-dashed-badge-card`
- `sequence-interaction-default-dashed-capsule-item`
- `sequence-interaction-default-dashed-compact-card`
- `sequence-interaction-default-dashed-rounded-rect-node`
- `sequence-interaction-default-rounded-rect-node`
- `sequence-interaction-wide-animated-badge-card`
- `sequence-interaction-wide-animated-capsule-item`
- `sequence-interaction-wide-animated-compact-card`
- `sequence-interaction-wide-animated-rounded-rect-node`
- `sequence-interaction-wide-badge-card`
- `sequence-interaction-wide-capsule-item`
- `sequence-interaction-wide-compact-card`
- `sequence-interaction-wide-dashed-badge-card`
- `sequence-interaction-wide-dashed-capsule-item`
- `sequence-interaction-wide-dashed-compact-card`
- `sequence-interaction-wide-dashed-rounded-rect-node`
- `sequence-interaction-wide-rounded-rect-node`

### sequence-mountain

- `sequence-mountain-underline-text`

### sequence-pyramid

- `sequence-pyramid-simple`

### sequence-roadmap-vertical

- `sequence-roadmap-vertical-badge-card`
- `sequence-roadmap-vertical-pill-badge`
- `sequence-roadmap-vertical-plain-text`
- `sequence-roadmap-vertical-quarter-circular`
- `sequence-roadmap-vertical-quarter-simple-card`
- `sequence-roadmap-vertical-simple`
- `sequence-roadmap-vertical-underline-text`

### sequence-snake-steps

- `sequence-snake-steps-compact-card`
- `sequence-snake-steps-pill-badge`
- `sequence-snake-steps-simple`
- `sequence-snake-steps-simple-illus`
- `sequence-snake-steps-underline-text`

### sequence-stairs-front

- `sequence-stairs-front-compact-card`
- `sequence-stairs-front-pill-badge`
- `sequence-stairs-front-simple`

### sequence-steps

- `sequence-steps-badge-card`
- `sequence-steps-simple`
- `sequence-steps-simple-illus`

### sequence-timeline

- `sequence-timeline-done-list`
- `sequence-timeline-plain-text`
- `sequence-timeline-rounded-rect-node`
- `sequence-timeline-simple`
- `sequence-timeline-simple-illus`

### sequence-zigzag-pucks-3d

- `sequence-zigzag-pucks-3d-indexed-card`
- `sequence-zigzag-pucks-3d-simple`
- `sequence-zigzag-pucks-3d-underline-text`

### sequence-zigzag-steps

- `sequence-zigzag-steps-underline-text`

## hierarchy

Hierarchies · 계층

````markdown
```infographic
infographic hierarchy-tree-curved-line-compact-card
data
  root
    label Workspace
    children
      - label Notes
      - label Tasks
```
````

### hierarchy-mindmap

- `hierarchy-mindmap-branch-gradient-capsule-item`
- `hierarchy-mindmap-branch-gradient-circle-progress`
- `hierarchy-mindmap-branch-gradient-compact-card`
- `hierarchy-mindmap-branch-gradient-lined-palette`
- `hierarchy-mindmap-branch-gradient-rounded-rect`
- `hierarchy-mindmap-level-gradient-capsule-item`
- `hierarchy-mindmap-level-gradient-circle-progress`
- `hierarchy-mindmap-level-gradient-compact-card`
- `hierarchy-mindmap-level-gradient-lined-palette`
- `hierarchy-mindmap-level-gradient-rounded-rect`

### hierarchy-structure

- `hierarchy-structure`
- `hierarchy-structure-mirror`

### hierarchy-tree

- `hierarchy-tree-bt-curved-line-badge-card`
- `hierarchy-tree-bt-curved-line-capsule-item`
- `hierarchy-tree-bt-curved-line-compact-card`
- `hierarchy-tree-bt-curved-line-ribbon-card`
- `hierarchy-tree-bt-curved-line-rounded-rect-node`
- `hierarchy-tree-bt-dashed-arrow-badge-card`
- `hierarchy-tree-bt-dashed-arrow-capsule-item`
- `hierarchy-tree-bt-dashed-arrow-compact-card`
- `hierarchy-tree-bt-dashed-arrow-ribbon-card`
- `hierarchy-tree-bt-dashed-arrow-rounded-rect-node`
- `hierarchy-tree-bt-dashed-line-badge-card`
- `hierarchy-tree-bt-dashed-line-capsule-item`
- `hierarchy-tree-bt-dashed-line-compact-card`
- `hierarchy-tree-bt-dashed-line-ribbon-card`
- `hierarchy-tree-bt-dashed-line-rounded-rect-node`
- `hierarchy-tree-bt-distributed-origin-badge-card`
- `hierarchy-tree-bt-distributed-origin-capsule-item`
- `hierarchy-tree-bt-distributed-origin-compact-card`
- `hierarchy-tree-bt-distributed-origin-ribbon-card`
- `hierarchy-tree-bt-distributed-origin-rounded-rect-node`
- `hierarchy-tree-bt-tech-style-badge-card`
- `hierarchy-tree-bt-tech-style-capsule-item`
- `hierarchy-tree-bt-tech-style-compact-card`
- `hierarchy-tree-bt-tech-style-ribbon-card`
- `hierarchy-tree-bt-tech-style-rounded-rect-node`
- `hierarchy-tree-curved-line-badge-card`
- `hierarchy-tree-curved-line-capsule-item`
- `hierarchy-tree-curved-line-compact-card`
- `hierarchy-tree-curved-line-ribbon-card`
- `hierarchy-tree-curved-line-rounded-rect-node`
- `hierarchy-tree-dashed-arrow-badge-card`
- `hierarchy-tree-dashed-arrow-capsule-item`
- `hierarchy-tree-dashed-arrow-compact-card`
- `hierarchy-tree-dashed-arrow-ribbon-card`
- `hierarchy-tree-dashed-arrow-rounded-rect-node`
- `hierarchy-tree-dashed-line-badge-card`
- `hierarchy-tree-dashed-line-capsule-item`
- `hierarchy-tree-dashed-line-compact-card`
- `hierarchy-tree-dashed-line-ribbon-card`
- `hierarchy-tree-dashed-line-rounded-rect-node`
- `hierarchy-tree-distributed-origin-badge-card`
- `hierarchy-tree-distributed-origin-capsule-item`
- `hierarchy-tree-distributed-origin-compact-card`
- `hierarchy-tree-distributed-origin-ribbon-card`
- `hierarchy-tree-distributed-origin-rounded-rect-node`
- `hierarchy-tree-lr-curved-line-badge-card`
- `hierarchy-tree-lr-curved-line-capsule-item`
- `hierarchy-tree-lr-curved-line-compact-card`
- `hierarchy-tree-lr-curved-line-ribbon-card`
- `hierarchy-tree-lr-curved-line-rounded-rect-node`
- `hierarchy-tree-lr-dashed-arrow-badge-card`
- `hierarchy-tree-lr-dashed-arrow-capsule-item`
- `hierarchy-tree-lr-dashed-arrow-compact-card`
- `hierarchy-tree-lr-dashed-arrow-ribbon-card`
- `hierarchy-tree-lr-dashed-arrow-rounded-rect-node`
- `hierarchy-tree-lr-dashed-line-badge-card`
- `hierarchy-tree-lr-dashed-line-capsule-item`
- `hierarchy-tree-lr-dashed-line-compact-card`
- `hierarchy-tree-lr-dashed-line-ribbon-card`
- `hierarchy-tree-lr-dashed-line-rounded-rect-node`
- `hierarchy-tree-lr-distributed-origin-badge-card`
- `hierarchy-tree-lr-distributed-origin-capsule-item`
- `hierarchy-tree-lr-distributed-origin-compact-card`
- `hierarchy-tree-lr-distributed-origin-ribbon-card`
- `hierarchy-tree-lr-distributed-origin-rounded-rect-node`
- `hierarchy-tree-lr-tech-style-badge-card`
- `hierarchy-tree-lr-tech-style-capsule-item`
- `hierarchy-tree-lr-tech-style-compact-card`
- `hierarchy-tree-lr-tech-style-ribbon-card`
- `hierarchy-tree-lr-tech-style-rounded-rect-node`
- `hierarchy-tree-rl-curved-line-badge-card`
- `hierarchy-tree-rl-curved-line-capsule-item`
- `hierarchy-tree-rl-curved-line-compact-card`
- `hierarchy-tree-rl-curved-line-ribbon-card`
- `hierarchy-tree-rl-curved-line-rounded-rect-node`
- `hierarchy-tree-rl-dashed-arrow-badge-card`
- `hierarchy-tree-rl-dashed-arrow-capsule-item`
- `hierarchy-tree-rl-dashed-arrow-compact-card`
- `hierarchy-tree-rl-dashed-arrow-ribbon-card`
- `hierarchy-tree-rl-dashed-arrow-rounded-rect-node`
- `hierarchy-tree-rl-dashed-line-badge-card`
- `hierarchy-tree-rl-dashed-line-capsule-item`
- `hierarchy-tree-rl-dashed-line-compact-card`
- `hierarchy-tree-rl-dashed-line-ribbon-card`
- `hierarchy-tree-rl-dashed-line-rounded-rect-node`
- `hierarchy-tree-rl-distributed-origin-badge-card`
- `hierarchy-tree-rl-distributed-origin-capsule-item`
- `hierarchy-tree-rl-distributed-origin-compact-card`
- `hierarchy-tree-rl-distributed-origin-ribbon-card`
- `hierarchy-tree-rl-distributed-origin-rounded-rect-node`
- `hierarchy-tree-rl-tech-style-badge-card`
- `hierarchy-tree-rl-tech-style-capsule-item`
- `hierarchy-tree-rl-tech-style-compact-card`
- `hierarchy-tree-rl-tech-style-ribbon-card`
- `hierarchy-tree-rl-tech-style-rounded-rect-node`
- `hierarchy-tree-tech-style-badge-card`
- `hierarchy-tree-tech-style-capsule-item`
- `hierarchy-tree-tech-style-compact-card`
- `hierarchy-tree-tech-style-ribbon-card`
- `hierarchy-tree-tech-style-rounded-rect-node`

## compare

Comparisons and quadrants · 비교·사분면

````markdown
```infographic
infographic compare-binary-horizontal-compact-card-vs
data
  compares
    - label Before
    - label After
```
````

### compare-binary-horizontal

- `compare-binary-horizontal-badge-card-arrow`
- `compare-binary-horizontal-badge-card-fold`
- `compare-binary-horizontal-badge-card-vs`
- `compare-binary-horizontal-compact-card-arrow`
- `compare-binary-horizontal-compact-card-fold`
- `compare-binary-horizontal-compact-card-vs`
- `compare-binary-horizontal-simple-arrow`
- `compare-binary-horizontal-simple-fold`
- `compare-binary-horizontal-simple-vs`
- `compare-binary-horizontal-underline-text-arrow`
- `compare-binary-horizontal-underline-text-fold`
- `compare-binary-horizontal-underline-text-vs`

### compare-hierarchy-left-right

- `compare-hierarchy-left-right-circle-node-pill-badge`
- `compare-hierarchy-left-right-circle-node-plain-text`

### compare-hierarchy-row

- `compare-hierarchy-row-letter-card-compact-card`
- `compare-hierarchy-row-letter-card-rounded-rect-node`
- `compare-swot`

### quadrant

- `compare-quadrant-quarter-circular`
- `compare-quadrant-quarter-simple-card`
- `compare-quadrant-simple-illus`

## relation

Relationships · 관계

````markdown
```infographic
infographic relation-dagre-flow-lr-compact-card
data
  nodes
    - id client
      label Client
    - id api
      label API
  relations
    - from client
      to api
```
````

### relation-circle

- `relation-circle-circular-progress`
- `relation-circle-icon-badge`

### relation-dagre-flow

- `relation-dagre-flow-lr-animated-badge-card`
- `relation-dagre-flow-lr-animated-capsule`
- `relation-dagre-flow-lr-animated-compact-card`
- `relation-dagre-flow-lr-animated-simple-circle-node`
- `relation-dagre-flow-lr-badge-card`
- `relation-dagre-flow-lr-compact-card`
- `relation-dagre-flow-lr-simple-circle-node`
- `relation-dagre-flow-tb-animated-badge-card`
- `relation-dagre-flow-tb-animated-capsule`
- `relation-dagre-flow-tb-animated-compact-card`
- `relation-dagre-flow-tb-animated-simple-circle-node`
- `relation-dagre-flow-tb-badge-card`
- `relation-dagre-flow-tb-compact-card`
- `relation-dagre-flow-tb-simple-circle-node`

### relation-network

- `relation-network-icon-badge`
- `relation-network-simple-circle-node`

## Deprecated aliases · 이전 별칭

Use the corresponding `compare-quadrant-*` ID in new documents. · 새 문서에는 오른쪽 ID를 사용하세요.

| Alias | Use |
| --- | --- |
| `quadrant-quarter-circular` | `compare-quadrant-quarter-circular` |
| `quadrant-quarter-simple-card` | `compare-quadrant-quarter-simple-card` |
| `quadrant-simple-illus` | `compare-quadrant-simple-illus` |
