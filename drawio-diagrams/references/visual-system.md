# Visual System — Colors, Typography, Spacing, Dark Mode

This is the complete visual specification for all Draw.io diagrams produced by this skill.
Apply it consistently across every diagram so they form a cohesive visual language.

---

## Color Palette

### Semantic Colors (use these, nothing else)

Every color has a role. Don't pick colors for aesthetics — pick them for meaning.

| Token | Fill | Stroke | Text | Semantic meaning |
|-------|------|--------|------|-----------------|
| `neutral` | `#F5F5F5` | `#9E9E9E` | `#212121` | Default services, generic nodes |
| `primary` | `#E3F2FD` | `#1565C0` | `#0D47A1` | Core services, main data flow |
| `success` | `#E8F5E9` | `#2E7D32` | `#1B5E20` | APPROVED, healthy, final success |
| `warning` | `#FFF8E1` | `#F57F17` | `#E65100` | MANUAL REVIEW, caution, cache |
| `danger` | `#FFEBEE` | `#C62828` | `#B71C1C` | BLOCKED, error, critical finding |
| `purple` | `#F3E5F5` | `#6A1B9A` | `#4A148C` | Auth, security, identity, RBAC |
| `teal` | `#E0F2F1` | `#00695C` | `#004D40` | Infrastructure, DB, storage, queue |
| `amber` | `#FFF3E0` | `#E65100` | `#BF360C` | Cache, performance optimization |
| `blue-dark` | `#1A237E` | `#0D1B6B` | `#FFFFFF` | Header bars, prominent labels |
| `group-bg` | `#FAFAFA` | `#BDBDBD` | `#424242` | Swimlane/container background |

### Color Assignment Rules

- **One color per logical group** — all nodes in "Auth" use `purple`, all in "Infra" use `teal`
- **Never use the same color for two unrelated groups** in the same diagram
- **Terminal states always use their semantic color**: success node = `success`, blocked = `danger`
- **Edge colors match their semantic role**:
  - Normal flow: `#1565C0` (primary stroke)
  - Error/alert: `#C62828` (danger stroke)
  - Async/dashed: `#757575` (neutral, dashed)
  - Retry: `#E65100` (amber stroke, dashed)

### Dark Mode Safety

SVGs committed to repos must survive dark mode (GitHub, VS Code, Notion dark).

Rules:
- Never use `fill="white"` or `background-color: white` on the SVG root — use `fill="none"` or omit
- All text must be `#212121` or darker on light fills, `#FFFFFF` on dark fills
- Stroke colors must have at least 4.5:1 contrast ratio against both white and `#1e1e1e`
- Test mentally: would the diagram be readable on a `#1e1e1e` background? If containers have
  light fills (e.g. `#F5F5F5`) they will show as light boxes on dark — that is fine as long as
  borders and text are visible

---

## Typography

### Font Stack

```
Helvetica Neue, Helvetica, Arial, sans-serif
```

Never specify a Google Font or external font. SVGs committed to repos must render without
network access.

### Size Hierarchy

| Element | Size | Weight | Color |
|---------|------|--------|-------|
| Diagram title | 18px | Bold | `#1A237E` |
| Group / swimlane header | 13px | Bold | `#424242` |
| Node primary label | 12px | Regular | `#212121` |
| Node secondary label (subtitle, type) | 10px | Regular | `#757575` |
| Edge label | 10px | Regular | `#424242` |
| Legend text | 11px | Regular | `#616161` |

### Label Writing Rules

- **Be specific**: "API Gateway · Go · :8080" not just "API"
- **Use · (middle dot) as separator** within a label, not "/" or "|"
- **Break at natural points**: if a label needs 2 lines, break after the primary name
- **Edge labels**: describe the action or protocol, not the direction
  - Good: "PUB scans.new", "REST :8080", "SUB results.*"
  - Bad: "sends to", "→", "calls"
- **Max label length**: 35 characters per line. If longer, abbreviate.

---

## Spacing System

### Canvas Padding
- Minimum 40px between outermost elements and SVG edge
- If the diagram has a title, add 60px at the top

### Between Groups
- 50px minimum between swimlane containers
- 60px when an edge label crosses the gap

### Inside Groups
- 20px padding between group border and first node
- 30px between nodes within a group
- Horizontal row: 25px between nodes
- Vertical column: 20px between nodes

### Node Sizing Standards

| Node type | Default size | Notes |
|-----------|-------------|-------|
| Service rectangle | 140 × 50px | Increase width for longer labels |
| Database cylinder | 100 × 65px | Fixed aspect ratio |
| Queue parallelogram | 130 × 45px | |
| Decision diamond | 80 × 65px | |
| Terminal/result | 130 × 40px | Bold border |
| Actor/user | 100 × 40px | Stadium shape |
| Swimlane container | Content + 30px all sides | startSize=30 for header |

---

## Icons and Symbols

Use Unicode symbols sparingly in node labels for quick visual scanning.
Always include a text label — never rely solely on an icon.

| Symbol | Meaning |
|--------|---------|
| ⚙️ | Service / process |
| 🗄 | Database |
| 📨 | Message queue |
| 🔑 | Authentication / key |
| 🛡 | Security / RBAC |
| 📋 | Policy / config |
| 🔧 | Worker / processor |
| 📣 | Notifier / alert |
| 🪣 | Object storage / S3 |
| 🤖 | AI / ML service |
| 📊 | Metrics / monitoring |
| 🔭 | Observability / tracing |
| ✅ | Approved / success |
| ⚠️ | Warning / review |
| 🚫 | Blocked / error |

Icon placement: start of first line in the label, followed by the service name.

---

## Diagram Composition

### Visual Hierarchy (most to least prominent)
1. Terminal states (APPROVED/BLOCKED) — largest, most distinct color
2. Core processing groups — largest swimlanes
3. Primary service nodes — standard size, primary color
4. Supporting infrastructure — neutral color
5. Edge labels — smallest, never compete with node labels

### What to Omit
- Implementation details (variable names, internal-only config)
- Every possible edge — show the primary 80% of flows
- Redundant labels (if the group is "Auth" and node is "JWT", don't label the edge "auth edge")
- Version numbers unless the diagram is specifically about versions

### Legend
Add a legend when the diagram uses more than 3 color groups or non-obvious shapes.
Place it bottom-right, 10px font, neutral background.
