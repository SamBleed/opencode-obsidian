# Mermaid vs Draw.io — Decision Guide

Use this reference to decide which tool to use for a given diagram request.

---

## Use Mermaid when

All of these are true:
- The flow is **linear or has a single primary dimension** (left-to-right or top-to-bottom with
  no significant branching back)
- Fewer than **15 nodes**
- **No swimlanes** or at most 1 level of grouping
- The diagram lives in a context where Mermaid is reliably rendered (GitHub markdown, Notion,
  GitLab — but NOT in plain HTML files, PDFs, or offline viewers)
- The user will **update the diagram frequently** and wants to edit text, not a visual tool
- No specific styling requirements beyond Mermaid's default theme

**Good Mermaid candidates:**
- Simple dedup flow (linear: input → hash check → embedding → decision)
- Git branch model (gitGraph)
- Simple class or ER diagram for a small schema
- Short sequence diagram (< 6 participants, < 10 messages)

---

## Use Draw.io SVG when

Any of these is true:
- **More than 15 nodes** or the diagram has complex branching/loops
- **Multiple layers or swimlane groups** that need to be visually separated
- The diagram needs to **render correctly in all contexts** (GitHub Pages, PDFs, offline,
  presentations, Confluence, Notion exports) — SVG always works
- The user needs **specific visual hierarchy** (entry points visually distinct, severity-based
  color coding, icon-like shapes for different service types)
- Mermaid's auto-layout produces **crossing connectors or illegible output**
- The diagram will be embedded in a **presentation or design document** where aesthetics matter
- The user wants **long-term editability** with a proper graphical tool (VS Code + Draw.io
  extension, diagrams.net)
- The existing Mermaid diagram **looks bad or is hard to read**

**Good Draw.io SVG candidates:**
- Full system architecture (multiple services, queues, databases, auth, observability)
- CI/CD or DevSecOps pipeline with parallel stages and decision points
- Network topology
- Data flow through a multi-stage processing pipeline
- Any diagram that will be shown to non-technical stakeholders

---

## The rendering problem with Mermaid

Mermaid diagrams look different depending on the renderer:

| Context | Mermaid behavior |
|---|---|
| GitHub light mode | Renders, default theme, acceptable |
| GitHub dark mode | Renders, but `style fill:#xxx` blocks look broken |
| claude.ai previewer | Renders with its own theme, colors may clash |
| Notion | Renders only with plugin, not native |
| Confluence | Requires plugin |
| PDF export | Does not render — shows raw code |
| Offline HTML | Does not render without CDN script |
| VS Code preview | Requires extension |

Draw.io SVG committed to the repo renders correctly in **all of these** because it is a
standard image format. No plugin, no renderer, no CDN.

---

## Migration: Mermaid → Draw.io SVG

When the user has existing Mermaid and wants to migrate:

1. Parse the Mermaid source to extract: nodes, node labels, subgraphs, edges, edge labels
2. Map subgraphs → swimlane groups
3. Map node shapes: default rect → rectangle, `([...])` → stadium, `[(...)]` → cylinder,
   `{...}` → diamond, `((...))` → circle
4. Apply the visual system from `references/visual-system.md`
5. Choose a layout strategy from `references/layout-strategies.md`
6. Generate SVG following `references/svg-generation.md`
7. Remove the Mermaid block from the markdown and replace with `![...](path/to/diagram.svg)`

Keep the Mermaid source as a comment in the `.drawio` file's metadata if the user wants
to preserve it as a record.
