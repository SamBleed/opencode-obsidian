---
name: drawio-diagrams
description: >
  Use this skill whenever the user needs to create, improve, or export technical diagrams
  as production-quality SVG files using Draw.io XML format. Triggers include: any mention
  of "diagrama", "diagram", "draw.io", "drawio", "arquitectura", "architecture diagram",
  "pipeline diagram", "flujo", "flowchart", "SVG diagram", "diagrama profesional", or
  requests to replace/improve Mermaid diagrams with something that renders correctly on
  GitHub, Notion, or any markdown viewer. Also triggers when the user complains that a
  Mermaid diagram looks bad, has rendering issues, or wants a diagram they can commit to
  a repo. Use this skill to produce .drawio XML files and their exported .svg counterparts
  — both committed to the repo so any dev can edit them. Do NOT use for simple inline
  Mermaid that the user explicitly wants; only when quality, stability, and visual control
  matter.
---

# Draw.io Diagrams — Production SVG for Repos

This skill produces professional technical diagrams as **Draw.io XML** (`.drawio`) paired
with **exported SVG** (`.svg`). Both files are committed to the repository:

- The `.svg` renders natively on GitHub, Notion, and any Markdown viewer — no renderer dependency.
- The `.drawio` (plain XML) can be opened and edited by any dev using VS Code + Draw.io extension or draw.io desktop.

Read `references/diagram-types.md` for type-specific layout rules before starting any diagram.
Read `references/visual-system.md` for the color palette, typography, and spacing system.
Read `references/svg-export.md` for the exact SVG output format and embedding instructions.

---

## Decision: Draw.io vs Mermaid

Use **Draw.io SVG** when:
- The diagram has more than ~12 nodes OR nested subgroups
- Visual quality matters (docs site, README of a public repo, client-facing material)
- The diagram will be edited by multiple people over time
- Rendering consistency across dark/light mode is required

Use **Mermaid** when:
- The diagram is simple and linear (< 10 nodes, no nesting)
- It lives in internal docs that won't be shared
- The user explicitly prefers Mermaid for editability

---

## Workflow

### Step 1 — Understand the diagram

Before touching any XML, answer these questions (from context or by asking):

1. **Type**: flowchart, architecture, sequence, entity-relationship, state machine, network?
2. **Nodes**: list all entities, services, or steps — their names and roles
3. **Edges**: list all connections — direction, label, style (solid/dashed/bidirectional)
4. **Grouping**: are there logical clusters? (e.g. "Worker", "Infraestructura", "Auth")
5. **Tone**: technical/engineering doc, executive overview, public README?

Extract this from existing Mermaid or description. If ambiguous, list your assumptions and proceed — don't block on clarification.

### Step 2 — Choose layout

Read `references/diagram-types.md` for the layout rules for the identified type.

General layout principles:
- **Flow diagrams**: top-to-bottom (TD) for processes, left-to-right (LR) for data pipelines
- **Architecture**: LR with clients on left, core services center, infra/storage right
- **Keep edge crossings minimal** — arrange nodes so most arrows flow in one direction
- **Group related nodes** visually with swimlanes or colored containers

### Step 3 — Apply the visual system

Read `references/visual-system.md` and apply:
- Use the semantic color palette (not arbitrary colors)
- Consistent node shapes per entity type (rectangle = service, cylinder = database, etc.)
- Font size hierarchy: group labels 14px bold, node labels 12px, edge labels 10px
- Adequate padding: minimum 20px inside containers, 40px between unrelated groups

### Step 4 — Write the Draw.io XML

Generate the complete `.drawio` XML. See format spec in `references/svg-export.md`.

Key rules:
- Every node gets a unique numeric `id` starting from 2 (0 and 1 are reserved for root/default layer)
- Use `mxGeometry` with absolute `x`, `y`, `width`, `height` — no relative positioning
- Group containers use `swimlane` style with `startSize` for the header height
- Edges reference source and target by `id`, not by label
- Never leave `x` or `y` as 0,0 for all nodes — lay them out spatially

### Step 5 — Export as SVG

After generating the `.drawio` XML, produce the corresponding inline SVG.

The SVG must:
- Use `viewBox` instead of fixed `width`/`height` so it scales in any context
- Embed all fonts as base64 or use system-safe fallbacks (`Helvetica, Arial, sans-serif`)
- Work in both light and dark mode (use explicit dark-safe colors — never rely on white background)
- Be self-contained — no external `<image>` src references that could break

### Step 6 — Deliver both files

Output structure:
```
docs/diagrams/
├── pipeline-flow.drawio     ← XML source, editable in Draw.io
└── pipeline-flow.svg        ← exported, committed, rendered in README
```

Embed in README as:
```markdown
![Pipeline Flow](docs/diagrams/pipeline-flow.svg)
```

Not as a Mermaid block. Not as a code block. A direct image reference.

---

## Node Shape Vocabulary

Use these shapes consistently across all diagrams so readers build a mental model:

| Entity type | Draw.io style | When to use |
|-------------|--------------|-------------|
| Service / API | `rounded=1` rectangle | Any running process or application |
| Database / Storage | Cylinder (`shape=mxgraph.flowchart.stored_data` or ellipse) | PostgreSQL, Redis, S3 |
| Queue / Bus | Parallelogram or hexagon | NATS, Kafka, RabbitMQ |
| Decision | Diamond `rhombus` | Yes/no branch, condition check |
| User / Actor | Rounded stadium shape | Developer, end user, client |
| Group / Zone | Swimlane container | Logical cluster (Worker, Auth, Infra) |
| External system | Dashed-border rectangle | Third party, outside system boundary |
| Terminal / Result | Bold-border rounded rectangle | APPROVED, BLOCKED, final state |

---

## Edge Style Vocabulary

| Relationship | Style |
|--------------|-------|
| Primary data flow | Solid arrow, medium weight (`strokeWidth=2`) |
| Async / queue message | Dashed arrow (`dashed=1`) |
| Bidirectional sync | Arrow on both ends (`startArrow=block;endArrow=block`) |
| Optional / fallback | Dotted thin arrow (`dashed=1;strokeWidth=1`) |
| Error / alert path | Red solid arrow (`strokeColor=#d32f2f`) |
| Internal call | Thin solid, no label needed |

---

## Semantic Color Palette

These are the only fill colors to use. Never pick arbitrary colors.

| Role | Fill | Stroke | Text | Use for |
|------|------|--------|------|---------|
| **Neutral** | `#F5F5F5` | `#9E9E9E` | `#212121` | Default nodes, services |
| **Primary** | `#E3F2FD` | `#1565C0` | `#0D47A1` | Core services, main flow |
| **Success** | `#E8F5E9` | `#2E7D32` | `#1B5E20` | Approved, healthy, output |
| **Warning** | `#FFF8E1` | `#F57F17` | `#E65100` | Manual review, caution |
| **Danger** | `#FFEBEE` | `#C62828` | `#B71C1C` | Blocked, error, critical |
| **Purple** | `#F3E5F5` | `#6A1B9A` | `#4A148C` | Auth, security, identity |
| **Teal** | `#E0F2F1` | `#00695C` | `#004D40` | Infrastructure, data layer |
| **Amber** | `#FFF3E0` | `#E65100` | `#BF360C` | Cache, optimization, perf |
| **Group BG** | `#FAFAFA` | `#BDBDBD` | `#424242` | Swimlane/container fill |

---

## Common Mistakes to Avoid

- **Don't center all nodes at 0,0** — XML needs real coordinates or everything stacks on top
- **Don't produce only Mermaid** and call it Draw.io — they are entirely different formats
- **Don't embed raster images** inside SVG for icons — use SVG path icons or unicode symbols
- **Don't use `width="100%"` on root `<svg>`** — use `viewBox` and let the container control size
- **Don't forget the `.drawio` file** — SVG alone is not editable; always deliver both
- **Don't make every node identical size** — use dimensions to convey relative importance
- **Don't cross more than 2 edges** without restructuring the layout first
- **Don't use `white` as background** — use `none` or `#FAFAFA` so it works on dark backgrounds

---

## Quality Checklist

Before delivering, verify:

- [ ] Every node has a meaningful label (no "Node 1", "Box A")
- [ ] Edge labels describe the relationship, not just an arrow
- [ ] Groups/swimlanes have clear, short titles
- [ ] Color usage is semantic (danger = red, success = green, neutral = gray/blue)
- [ ] SVG renders cleanly at 800px wide — text is legible, no overlap
- [ ] `.drawio` XML is well-formed (no unclosed tags, all IDs unique)
- [ ] Both files are named consistently (`pipeline-flow.drawio` + `pipeline-flow.svg`)
- [ ] README embedding uses `![alt text](path/to/file.svg)` — not a code block
- [ ] No hardcoded `white` backgrounds that break in dark mode

---

## Reference Files

- `references/diagram-types.md` — Layout rules per diagram type (flowchart, architecture, sequence, ER, state machine, network)
- `references/visual-system.md` — Color palette deep dive, typography scale, spacing system, dark mode rules, icon usage
- `references/svg-export.md` — Draw.io XML format spec, complete node/edge XML examples, SVG export format, README embedding patterns
