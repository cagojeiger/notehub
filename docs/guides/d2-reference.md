# D2 reference

[Document authoring guide](document-authoring.md)

Write raw D2 in a `.d2` text document. This reference covers NoteHub renderer **0.1.1**, D2 **0.9.0**, and the **Dagre** layout engine. File imports are disabled. [Renderer configuration](../../services/renderer/render.go) · [D2 dependency](../../services/renderer/go.mod)

- [Shapes](#shapes)
- [Connections and arrowheads](#connections)
- [Layout](#layout)

## Shapes

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

## Connections

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

## Layout

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

For a complete Markdown document embedding a diagram, see the [example README](examples/request-flow/README.md) and its [D2 source](examples/request-flow/flow.d2).
