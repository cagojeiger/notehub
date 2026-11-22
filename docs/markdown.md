# Authoring documents in NoteHub

Save notes and folder landing pages as `.md` text documents. Write ordinary Markdown by default. Use GitHub Flavored Markdown for tables, task lists, and footnotes. When a diagram helps, choose its format by what it needs to communicate.

| Use case | What to save |
| --- | --- |
| Prose, lists, tables, and links | A `.md` text document |
| A flowchart or sequence inside a note | A `mermaid` fenced block in the `.md` document |
| A chart or infographic inside a note | An `infographic` fenced block in the `.md` document |
| A standalone or reusable architecture diagram | A `.d2` text document, optionally embedded in Markdown |

## Markdown links and folder landing pages

Resolve relative links from the folder containing the Markdown document. A path starting with `/` resolves from the root of the same Space. Use the exact names of existing files, including their extensions. Use ordinary Markdown links for in-Space navigation and keep the target path within that Space.

```markdown
[Project folder](./project/)
[Project README](./project/README.md)
[Space index](/README.md)
[See the notes](#notes)
![System overview](./assets/overview.png)
```

Place `README.md` directly inside a folder to give it a landing page. The web folder view opens on the README tab, including at the Space root. Use `[Project folder](./project/)` when readers should land in that folder and see its file list and README. Use `[Project README](./project/README.md)` to open the document directly. Incoming links follow the exact linked target, even when the two views show the same README body.

For a bundle of related materials, use a folder with a README that links to its notes and attachments. Use a single path on its own line in an `audio` block to add a player for an audio file already stored in the same Space.

`````markdown
```audio
./recording.webm
```
`````

## Mermaid

Put a diagram directly inside the note using a `mermaid` fence.

```mermaid
flowchart LR
    Inbox --> Review
    Review --> Archive
```

See the [Mermaid syntax reference](https://mermaid.js.org/intro/syntax-reference.html) for other diagram types.

## AntV Infographic

Use an `infographic` fence. Start with `infographic <template>` and indent the data under `data`. Choose a template that matches the data shape. For a bar chart, use `values` with a label and a numeric value per entry.

```infographic
infographic chart-bar-plain-text
data
  title Notes this week
  desc Documents written by category
  values
    - label Meeting notes
      value 12
    - label Work notes
      value 18
    - label Research
      value 9
    - label Retrospectives
      value 6
```

See the [AntV data reference](https://infographic.antv.vision/learn/data) for additional data shapes.

## D2

Save the D2 source as a separate text document named `architecture.d2`. Put the source directly in the file without Markdown fences.

```d2
client: Client
api: NoteHub API
db: Database

client -> api: HTTPS
api -> db: Read and write
```

To embed the rendered SVG, reference the original D2 document with Markdown image syntax.

```markdown
![NoteHub architecture](./architecture.d2)
```

Use `[Open the diagram](./architecture.d2)` to open the standalone diagram. NoteHub uses its configured D2 renderer for the SVG preview. See the [D2 language guide](https://d2lang.com/tour/intro/) for more syntax.

## Preview and verification

- Use the web preview to verify that the Markdown, links, and diagrams display as intended.
- MCP Apps can display standalone and embedded D2 when rendering is available. Mermaid and Infographic require the host to support the isolated diagram renderer. The viewer can show the source if that renderer is unavailable.
- The CLI and MCP text tools work with the saved source. Use the existing text write operation for both `.md` and `.d2` files. Store image and audio attachments as files in the same Space and link to those paths.
- MCP and CLI command targets use `<space>:/absolute/path`, such as `daily:/project/README.md`. Inside a Markdown document, use `./architecture.d2` or a Space-root path such as `/project/architecture.d2`.
