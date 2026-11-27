# NoteHub Markdown Guide

Write notes in Markdown and embed files from your Space. Use Mermaid and D2 for diagrams, or OpenUI to combine cards and charts in one block.

## Contents

- [Write a Markdown note](#write-a-markdown-note)
- [Write LaTeX equations](#write-latex-equations)
- [Link to files and folders](#link-to-files-and-folders)
- [Images, diagrams, and audio](#images-diagrams-and-audio)
- [Draw a Mermaid diagram](#draw-a-mermaid-diagram)
- [Compose an OpenUI block](#compose-an-openui-block)
- References — [Mermaid diagrams](#mermaid-diagram-reference) · [D2](#d2-reference)

## Write a Markdown note

Save notes with the extension `.md` or `.markdown`. The preview supports task lists, tables, and footnotes.

```markdown
# Meeting notes

- [ ] Review the proposal.
```

To show document properties, start the file with a YAML mapping between two lines containing `---`. To add a collapsible section, use the following markup.

```markdown
<details>
<summary>Meeting notes</summary>

- Review the proposal.

</details>
```

## Write LaTeX equations

Use `$...$` for an inline equation. For a displayed equation, put the opening and closing `$$` on their own lines. Escape a literal dollar sign as `\$` when writing prices alongside equations.

```markdown
The price is \$15 and the quantity is $n$.

$$
\frac{a}{b} + c
$$
```

## Link to files and folders

Use a target such as `daily:/project/README.md` in MCP and CLI commands. Inside Markdown, use paths within the same Space. Resolve `./` and `../` from the document’s folder and `/` from the Space root. Include the exact filename and extension.

In a Space-root `README.md`, these links open the project folder and its README respectively. The web folder view displays the file list and previews its README. Incoming links are counted for the target you link to.

```markdown
[Project folder](./project/)
[Project README](./project/README.md)
```

## Images, diagrams, and audio

### Embed or open a D2 diagram

With `flow.d2` beside your note, use image syntax to render the diagram as SVG or a regular link to open its viewer.

```markdown
# Request flow

![Request flow](./flow.d2)

[Open the diagram](./flow.d2)
```

A fenced `d2` block displays the source code. To write the diagram itself, see the [D2 reference](#d2-reference).

### Embed an image or audio recording

Upload the attachments to the same Space first. If `assets/overview.png` and `recordings/meeting.webm` are under the folder containing your note, put the following source in the note to show the image and audio player. The `audio` block takes one file path on its own line.

````markdown
![System overview](./assets/overview.png)

```audio
./recordings/meeting.webm
```
````

## Draw a Mermaid diagram

Start a fenced block with `mermaid`. The first line inside the block selects the diagram type. For example, `flowchart LR` draws a flowchart from left to right.

````markdown
```mermaid
flowchart LR
    Inbox --> Review
    Review --> Archive
```
````

[All Mermaid diagram types, declarations and official examples for this version](#mermaid-diagram-reference)

## Compose an OpenUI block

Copy the whole example, including the opening and closing fences, into a Markdown note. Start the block with `root = Col([...])`, then define the cards, labels, or charts it should contain. Open the note in Preview to see the result.

### Show a chart

````markdown
```openui
root = Col([Card([Title("Notes per week"), chart])])
chart = Chart("bar", data, "week", ["notes"])
data = [{week: "Week 1", notes: 12}, {week: "Week 2", notes: 18}, {week: "Week 3", notes: 15}]
```
````

### Bind a slider to a value

````markdown
```openui
root = Col([Card([Title("Budget split"), company, personal, slider])])
company = Caption("Company " + $share + "%")
personal = Caption("Personal " + (100 - $share) + "%")
slider = Slider($share, 0, 100, 1, "Company share", "%")
$share = 70
```
````

Moving the slider updates both percentages in the preview. To change the starting value in the saved note and its PDF, edit `$share = 70` in the source.

### Components available in NoteHub

| Purpose | Components |
| --- | --- |
| Layout | `Col`, `Row`, `Grid`, `Box`, `Card`, `Divider`, `Spacer` |
| Text | `Title`, `Caption`, `Label`, `Badge` |
| Charts and drawings | `Chart`, `Svg`, `G`, `Rect`, `Circle`, `Line`, `Path`, `SvgText` |
| Interaction | `Slider` |

When asking an AI to write a block, give it one of the examples and ask it to use the components listed here. NoteHub runs the block in an isolated frame. Forms, tool calls, and external media are outside the document catalog. If the block cannot render, NoteHub shows its source so you can copy or edit it. Existing `infographic` blocks also display as copyable source code. [OpenUI Lang specification](https://www.openui.com/docs/openui-lang/specification)

## Mermaid diagram reference

NoteHub uses Mermaid 12.3.0. In a Markdown document, use a `mermaid` fence and start the diagram source with the declaration for the type you choose. Copy the declaration exactly, including capitalization and any `-beta` suffix shown. Each row links to the official syntax and an example in the 12.3.0 source.

````markdown
```mermaid
flowchart LR
    Draft --> Review
```
````

[Official syntax overview](https://mermaid.js.org/intro/syntax-reference.html) · [Live Editor](https://mermaid.live/)

### Flows and interactions

| Diagram | Declaration | Official syntax and examples |
| --- | --- | --- |
| Flowchart | `flowchart` | [Syntax](https://mermaid.js.org/syntax/flowchart.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/flowchart.md#L688-L689) |
| Swimlanes | `swimlane-beta` | [Syntax](https://mermaid.js.org/syntax/swimlanes.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/swimlanes.md#L136-L140) |
| Sequence | `sequenceDiagram` | [Syntax](https://mermaid.js.org/syntax/sequenceDiagram.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/sequenceDiagram.md#L435-L437) |
| Agentflow | `agentflow-beta` | [Syntax](https://mermaid.js.org/syntax/agentflow.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/agentflow.md#L176-L187) |
| Event modeling | `eventmodeling` | [Syntax](https://mermaid.js.org/syntax/eventmodeling.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/eventmodeling.md#L42-L46) |
| State | `stateDiagram-v2` | [Syntax](https://mermaid.js.org/syntax/stateDiagram.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/stateDiagram.md#L104-L105) |
| User journey | `journey` | [Syntax](https://mermaid.js.org/syntax/userJourney.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/userJourney.md#L12-L20) |

### Software and systems

| Diagram | Declaration | Official syntax and examples |
| --- | --- | --- |
| Class | `classDiagram` | [Syntax](https://mermaid.js.org/syntax/classDiagram.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/classDiagram.md#L367-L368) |
| Entity relationship | `erDiagram` | [Syntax](https://mermaid.js.org/syntax/entityRelationshipDiagram.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/entityRelationshipDiagram.md#L154-L155) |
| Requirements | `requirementDiagram` | [Syntax](https://mermaid.js.org/syntax/requirementDiagram.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/requirementDiagram.md#L145-L152) |
| Use case | `usecase-beta` | [Syntax](https://mermaid.js.org/syntax/usecase.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/usecase.md#L259-L262) |
| Architecture | `architecture-beta` | [Syntax](https://mermaid.js.org/syntax/architecture.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/architecture.md#L170-L178) |
| C4 | `C4Dynamic`, `C4Component`, `C4Container`, `C4Context`, `C4Deployment` | [Syntax](https://mermaid.js.org/syntax/c4.html) · [C4Dynamic](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/c4.md#L312-L327) · [C4Component](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/c4.md#L270-L305) · [C4Container](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/c4.md#L228-L263) · [C4Context](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/c4.md#L10-L51) · [C4Deployment](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/c4.md#L334-L381) |
| Block | `block` | [Syntax](https://mermaid.js.org/syntax/block.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/block.md#L66-L67) |
| Git graph | `gitGraph` | [Syntax](https://mermaid.js.org/syntax/gitgraph.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/gitgraph.md#L55-L58) |
| Packet | `packet` | [Syntax](https://mermaid.js.org/syntax/packet.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/packet.md#L63-L69) |
| Railroad | `railroad-ebnf-beta`, `railroad-abnf-beta`, `railroad-peg-beta`, `railroad-beta` | [Syntax](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/railroad.md) · [railroad-ebnf-beta](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/railroad.md#L76-L77) · [railroad-abnf-beta](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/railroad.md#L208-L213) · [railroad-peg-beta](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/railroad.md#L241-L246) · [railroad-beta](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/railroad.md#L265-L287) |

### Planning and analysis

| Diagram | Declaration | Official syntax and examples |
| --- | --- | --- |
| Gantt | `gantt` | [Syntax](https://mermaid.js.org/syntax/gantt.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/gantt.md?plain=1#L24-L32) |
| Timeline | `timeline` | [Syntax](https://mermaid.js.org/syntax/timeline.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/timeline.md#L52-L57) |
| Kanban | `kanban` | [Syntax](https://mermaid.js.org/syntax/kanban.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/kanban.md#L14-L16) |
| Mindmap | `mindmap` | [Syntax](https://mermaid.js.org/syntax/mindmap.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/mindmap.md#L95-L96) |
| TreeView | `treeView-beta` | [Syntax](https://mermaid.js.org/syntax/treeView.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/treeView.md#L219-L223) |
| Ishikawa | `ishikawa-beta` | [Syntax](https://mermaid.js.org/syntax/ishikawa.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/ishikawa.md#L14-L33) |
| Wardley map | `wardley-beta` | [Syntax](https://mermaid.js.org/syntax/wardley.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/wardley.md?plain=1#L30-L49) |
| Cynefin | `cynefin-beta` | [Syntax](https://mermaid.js.org/syntax/cynefin.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/cynefin.md#L153-L159) |

### Charts and sets

| Diagram | Declaration | Official syntax and examples |
| --- | --- | --- |
| Pie | `pie` | [Syntax](https://mermaid.js.org/syntax/pie.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/pie.md#L13-L16) |
| XY chart | `xychart` | [Syntax](https://mermaid.js.org/syntax/xyChart.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/xyChart.md#L254-L258) |
| Quadrant | `quadrantChart` | [Syntax](https://mermaid.js.org/syntax/quadrantChart.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/quadrantChart.md#L12-L25) |
| Sankey | `sankey` | [Syntax](https://mermaid.js.org/syntax/sankey.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/sankey.md#L120-L126) |
| Radar | `radar-beta` | [Syntax](https://mermaid.js.org/syntax/radar.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/radar.md#L50-L61) |
| Treemap | `treemap-beta` | [Syntax](https://mermaid.js.org/syntax/treemap.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/treemap.md#L49-L55) |
| Venn | `venn-beta` | [Syntax](https://mermaid.js.org/syntax/venn.html) · [Example](https://github.com/mermaid-js/mermaid/blob/mermaid%4012.3.0/packages/mermaid/src/docs/syntax/venn.md#L82-L85) |

### Alternate declarations

The following declarations also rendered in NoteHub’s frame.

| Preferred | Also accepted |
| --- | --- |
| `flowchart` | `graph`, `flowchart-elk` |
| `classDiagram` | `classDiagram-v2` |
| `stateDiagram-v2` | `stateDiagram` |
| `block` | `block-beta` |
| `packet` | `packet-beta` |
| `sankey` | `sankey-beta` |
| `xychart` | `xychart-beta` |
| `ishikawa-beta` | `ishikawa` |
| `treemap-beta` | `treemap` |

### Renderer scope

- `info` displays Mermaid version information.

- ZenUML requires a separate plugin. NoteHub 12.3.0-based previews support the built-in `sequenceDiagram` syntax for message exchanges. [ZenUML](https://mermaid.js.org/syntax/zenuml.html)

## D2 reference

Save this source as `flow.d2` in the same folder as your Markdown note. Omit the Markdown fences.

```d2
client: Browser
api: NoteHub
client -> api: Send request
```

This reference covers NoteHub renderer **0.1.1**, D2 **0.9.0**, and the **Dagre** layout engine. File imports are disabled. [Renderer configuration](../../services/renderer/render.go) · [D2 dependency](../../services/renderer/go.mod)

- [Shapes](#shapes)
- [Connections and arrowheads](#connections)
- [Layout](#layout)

### Shapes

Set `shape` on a node. A node with no explicit shape uses `rectangle`.

| Use | Values | Official syntax and examples |
| --- | --- | --- |
| Geometry | `rectangle`, `square`, `oval`, `circle`, `diamond`, `hexagon`, `parallelogram` | [Shape catalog](https://d2lang.com/tour/shapes/) |
| Documents and annotations | `page`, `document`, `callout` | [Shape examples](https://d2lang.com/tour/shapes/#example) |
| Storage and infrastructure | `cylinder`, `stored_data`, `queue`, `package`, `cloud` | [Shape catalog](https://d2lang.com/tour/shapes/) |
| People and steps | `person`, `c4-person`, `step` | [Shape catalog](https://d2lang.com/tour/shapes/) |
| Text and code | `text`, `code` | [Text and code blocks](https://d2lang.com/tour/text/) |
| UML class and database table | `class`, `sql_table` | [UML classes](https://d2lang.com/tour/uml-classes/) · [SQL tables](https://d2lang.com/tour/sql-tables/) |
| Image | `image` (set `icon` to the image URL) | [Icons and images](https://d2lang.com/tour/icons/) |
| Sequence | `sequence_diagram` | [Sequence diagram examples](https://d2lang.com/tour/sequence-diagrams/) |
| Hierarchy container | `hierarchy` | [v0.9.0 shape registry](https://github.com/d2lang/d2/blob/v0.9.0/d2target/d2target.go#L1244-L1298) · [Containers](https://d2lang.com/tour/containers/) |

```d2
reader: Reader {shape: person}
store: Archive {shape: cylinder}
reader -> store: Save note
```

An `image` shape can compile to SVG containing an image reference. Displaying that resource depends on the viewer's image policy. The `hierarchy` shape is accepted by the compiler; NoteHub uses Dagre for graph placement.

### Connections

Connect **node keys**. Add a label after `:`. [Connection syntax and examples](https://d2lang.com/tour/connections/)

| D2 | Result |
| --- | --- |
| `a -- b` | Line without arrowheads |
| `a -> b` | Arrow toward `b` |
| `a <- b` | Arrow toward `a` |
| `a <-> b` | Arrows at both ends |

Choose an arrowhead on an end that has an arrow. Set `source-arrowhead.shape` on the source or `target-arrowhead.shape` on the target. [Arrowhead examples](https://d2lang.com/tour/connections/#arrowheads)

| Use | Values |
| --- | --- |
| Standard ends | `triangle` (default), `arrow`, `diamond`, `circle`, `box`, `cross`, `none` |
| Entity cardinality | `cf-one`, `cf-one-required`, `cf-many`, `cf-many-required` |

```d2
author <-> note: owns {
  source-arrowhead.shape: cf-one-required
  target-arrowhead.shape: cf-many
  style.stroke-dash: 4
}
```

Use `style.filled: false` on a triangle arrowhead for an outline, or `style.filled: true` on a diamond, circle, or box for a solid end. Edge styles such as `style.stroke-dash` belong on the connection. [Styles](https://d2lang.com/tour/style/) · [Chains, repeated edges and indexed references](https://d2lang.com/tour/connections/)

### Layout

NoteHub resolves graph layout to **Dagre**. The source can choose flow direction and organize nodes into containers, grids, or sequence diagrams. ELK and TALA are upstream engine options; this service always selects Dagre. [Dagre examples](https://d2lang.com/tour/dagre/) · [Upstream engine comparison](https://d2lang.com/tour/layouts/)

| Intent | Source | Official examples |
| --- | --- | --- |
| Flow direction | Top-level `direction: down` (default), `up`, `left`, `right` | [Layouts](https://d2lang.com/tour/layouts/) |
| Group nodes | Nest children in `group: { ... }`; connect by path, e.g. `group.api` | [Containers](https://d2lang.com/tour/containers/) |
| Grid | Set `grid-columns` or `grid-rows` on the grid | [Grid diagrams](https://d2lang.com/tour/grid-diagrams/) |
| Ordered messages | Set `shape: sequence_diagram`; write the messages in order | [Sequence diagrams](https://d2lang.com/tour/sequence-diagrams/) |
| Anchor a title or legend | Use a constant such as `near: top-center` on a separate node | [Positioning examples](https://d2lang.com/tour/positions/) |

```d2
direction: right
client -> services.api: Request
services: {
  api -> db
  db.shape: cylinder
}
```
