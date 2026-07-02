# SVG Export — Draw.io XML Format, SVG Output, README Embedding

---

## Draw.io XML Format

A `.drawio` file is an XML file with this root structure:

```xml
<mxfile host="Claude" modified="2024-01-01T00:00:00.000Z" version="21.0.0">
  <diagram id="diagram-id" name="Diagram Name">
    <mxGraphModel dx="1422" dy="762" grid="1" gridSize="10" guides="1"
                  tooltips="1" connect="1" arrows="1" fold="1"
                  page="1" pageScale="1" pageWidth="1654" pageHeight="1169"
                  math="0" shadow="0">
      <root>
        <mxCell id="0" />
        <mxCell id="1" parent="0" />
        <!-- All nodes and edges go here, as children of cell id="1" -->
      </root>
    </mxGraphModel>
  </diagram>
</mxfile>
```

IDs 0 and 1 are always reserved. Start your content nodes at id="2".

---

## Node XML Examples

### Service Rectangle (rounded)

```xml
<mxCell id="2" value="⚙️ API Gateway&#xa;Go · Chi · :8080"
        style="rounded=1;whiteSpace=wrap;html=1;
               fillColor=#E3F2FD;strokeColor=#1565C0;fontColor=#0D47A1;
               fontSize=12;fontFamily=Helvetica;arcSize=12;
               strokeWidth=2;"
        vertex="1" parent="1">
  <mxGeometry x="300" y="200" width="160" height="55" as="geometry" />
</mxCell>
```

Note: `&#xa;` is a newline inside the label.

### Database Cylinder

```xml
<mxCell id="3" value="🗄 PostgreSQL 17&#xa;+ pgvector"
        style="shape=mxgraph.flowchart.stored_data;whiteSpace=wrap;html=1;
               fillColor=#E0F2F1;strokeColor=#00695C;fontColor=#004D40;
               fontSize=12;fontFamily=Helvetica;
               strokeWidth=2;"
        vertex="1" parent="1">
  <mxGeometry x="520" y="200" width="120" height="70" as="geometry" />
</mxCell>
```

### Message Queue (parallelogram)

```xml
<mxCell id="4" value="📨 NATS JetStream&#xa;scans.new · results.*"
        style="shape=parallelogram;perimeter=parallelogramPerimeter;
               whiteSpace=wrap;html=1;fixedSize=1;size=20;
               fillColor=#E0F2F1;strokeColor=#00695C;fontColor=#004D40;
               fontSize=12;fontFamily=Helvetica;strokeWidth=2;"
        vertex="1" parent="1">
  <mxGeometry x="300" y="320" width="160" height="50" as="geometry" />
</mxCell>
```

### Decision Diamond

```xml
<mxCell id="5" value="¿Hash&#xa;existe?"
        style="rhombus;whiteSpace=wrap;html=1;
               fillColor=#FFF8E1;strokeColor=#F57F17;fontColor=#E65100;
               fontSize=12;fontFamily=Helvetica;strokeWidth=2;"
        vertex="1" parent="1">
  <mxGeometry x="300" y="440" width="100" height="75" as="geometry" />
</mxCell>
```

### Terminal / Result Node

```xml
<mxCell id="6" value="✅ APPROVED"
        style="rounded=1;whiteSpace=wrap;html=1;
               fillColor=#E8F5E9;strokeColor=#2E7D32;fontColor=#1B5E20;
               fontSize=13;fontFamily=Helvetica;fontStyle=1;
               strokeWidth=3;arcSize=50;"
        vertex="1" parent="1">
  <mxGeometry x="520" y="560" width="140" height="45" as="geometry" />
</mxCell>
```

### Swimlane Container (group)

```xml
<mxCell id="7" value="Worker"
        style="swimlane;fontStyle=1;align=center;
               startSize=30;
               fillColor=#FAFAFA;strokeColor=#BDBDBD;fontColor=#424242;
               fontSize=13;fontFamily=Helvetica;
               strokeWidth=1.5;"
        vertex="1" parent="1">
  <mxGeometry x="200" y="120" width="500" height="280" as="geometry" />
</mxCell>
```

Nodes inside a swimlane must have `parent="7"` (the swimlane's id) and use
**relative coordinates** within the container:

```xml
<mxCell id="8" value="🔧 Scanner Pool"
        style="rounded=1;whiteSpace=wrap;html=1;
               fillColor=#E8F5E9;strokeColor=#2E7D32;fontColor=#1B5E20;
               fontSize=12;fontFamily=Helvetica;strokeWidth=2;"
        vertex="1" parent="7">
  <mxGeometry x="30" y="80" width="140" height="50" as="geometry" />
</mxCell>
```

---

## Edge XML Examples

### Standard directed edge

```xml
<mxCell id="20" value="PUB scans.new"
        style="edgeStyle=orthogonalEdgeStyle;
               strokeColor=#1565C0;strokeWidth=2;
               fontColor=#424242;fontSize=10;fontFamily=Helvetica;
               exitX=1;exitY=0.5;exitDx=0;exitDy=0;
               entryX=0;entryY=0.5;entryDx=0;entryDy=0;"
        edge="1" source="2" target="4" parent="1">
  <mxGeometry relative="1" as="geometry" />
</mxCell>
```

### Dashed async edge

```xml
<mxCell id="21" value="async"
        style="edgeStyle=orthogonalEdgeStyle;dashed=1;
               strokeColor=#757575;strokeWidth=1.5;
               fontColor=#424242;fontSize=10;fontFamily=Helvetica;"
        edge="1" source="4" target="8" parent="1">
  <mxGeometry relative="1" as="geometry" />
</mxCell>
```

### Error / alert edge (red)

```xml
<mxCell id="22" value="⚠️ alerta"
        style="edgeStyle=orthogonalEdgeStyle;
               strokeColor=#C62828;strokeWidth=2;
               fontColor=#C62828;fontSize=10;fontFamily=Helvetica;"
        edge="1" source="6" target="2" parent="1">
  <mxGeometry relative="1" as="geometry" />
</mxCell>
```

### Curved edge with waypoints

```xml
<mxCell id="23" value="retry"
        style="edgeStyle=elbowEdgeStyle;elbow=orthogonal;dashed=1;
               strokeColor=#E65100;strokeWidth=1.5;
               fontColor=#E65100;fontSize=10;fontFamily=Helvetica;"
        edge="1" source="9" target="8" parent="1">
  <mxGeometry relative="1" as="geometry">
    <Array as="points">
      <mxPoint x="150" y="500" />
      <mxPoint x="150" y="400" />
    </Array>
  </mxGeometry>
</mxCell>
```

---

## SVG Output Format

After generating the `.drawio` XML, produce the SVG as a standalone file.

### SVG Root Element

```xml
<svg xmlns="http://www.w3.org/2000/svg"
     xmlns:xlink="http://www.w3.org/1999/xlink"
     viewBox="0 0 1200 800"
     style="max-width:100%;height:auto;font-family:Helvetica Neue,Helvetica,Arial,sans-serif;">
```

Key rules:
- **Always use `viewBox`**, never fixed `width`/`height` on the root element
- **`max-width:100%`** so it scales down on small screens
- **No `background` or `fill` on root** — let the page/viewer decide
- Font family on root so all text inherits it

### Text Rendering

```xml
<text x="350" y="230"
      text-anchor="middle" dominant-baseline="middle"
      font-size="12" font-family="Helvetica Neue,Helvetica,Arial,sans-serif"
      fill="#0D47A1">
  API Gateway
</text>
<text x="350" y="248"
      text-anchor="middle" dominant-baseline="middle"
      font-size="10" font-family="Helvetica Neue,Helvetica,Arial,sans-serif"
      fill="#757575">
  Go · Chi · :8080
</text>
```

Multi-line labels: use separate `<text>` elements at staggered y positions (16px apart for 12px font).

### Shape Rendering

Rectangles (rounded):
```xml
<rect x="270" y="197" width="160" height="55" rx="8" ry="8"
      fill="#E3F2FD" stroke="#1565C0" stroke-width="2"/>
```

Cylinders (database): use ellipse top + rectangle body + ellipse bottom:
```xml
<ellipse cx="580" cy="207" rx="60" ry="12" fill="#E0F2F1" stroke="#00695C" stroke-width="2"/>
<rect x="520" y="207" width="120" height="46" fill="#E0F2F1" stroke="#00695C" stroke-width="2"/>
<ellipse cx="580" cy="253" rx="60" ry="12" fill="#E0F2F1" stroke="#00695C" stroke-width="2"/>
```

Diamonds (decision):
```xml
<polygon points="350,440 410,477 350,515 290,477"
         fill="#FFF8E1" stroke="#F57F17" stroke-width="2"/>
```

Group containers (swimlane):
```xml
<!-- Container border -->
<rect x="200" y="120" width="500" height="280" rx="4"
      fill="#FAFAFA" stroke="#BDBDBD" stroke-width="1.5"/>
<!-- Header bar -->
<rect x="200" y="120" width="500" height="30" rx="4"
      fill="#EEEEEE" stroke="#BDBDBD" stroke-width="1.5"/>
<!-- Header text -->
<text x="450" y="140" text-anchor="middle" dominant-baseline="middle"
      font-size="13" font-weight="bold" fill="#424242">Worker</text>
```

### Arrow Markers

Define markers in `<defs>`:
```xml
<defs>
  <!-- Standard arrow -->
  <marker id="arrow-primary" markerWidth="10" markerHeight="7"
          refX="10" refY="3.5" orient="auto">
    <polygon points="0 0, 10 3.5, 0 7" fill="#1565C0"/>
  </marker>

  <!-- Error arrow (red) -->
  <marker id="arrow-danger" markerWidth="10" markerHeight="7"
          refX="10" refY="3.5" orient="auto">
    <polygon points="0 0, 10 3.5, 0 7" fill="#C62828"/>
  </marker>

  <!-- Neutral arrow (gray) -->
  <marker id="arrow-neutral" markerWidth="10" markerHeight="7"
          refX="10" refY="3.5" orient="auto">
    <polygon points="0 0, 10 3.5, 0 7" fill="#757575"/>
  </marker>
</defs>
```

Using a marker on an edge:
```xml
<line x1="430" y1="225" x2="520" y2="225"
      stroke="#1565C0" stroke-width="2" marker-end="url(#arrow-primary)"/>
```

Dashed edge:
```xml
<line x1="430" y1="345" x2="520" y2="345"
      stroke="#757575" stroke-width="1.5" stroke-dasharray="6,4"
      marker-end="url(#arrow-neutral)"/>
```

Edge with label:
```xml
<line x1="430" y1="225" x2="520" y2="225"
      stroke="#1565C0" stroke-width="2" marker-end="url(#arrow-primary)"/>
<rect x="455" y="210" width="85" height="16" rx="2" fill="white" opacity="0.85"/>
<text x="497" y="222" text-anchor="middle" font-size="10" fill="#424242">PUB scans.new</text>
```

(The white rect behind the label prevents overlap with the edge line.)

---

## File Naming and Delivery

### Naming Convention

```
docs/diagrams/
├── {slug}.drawio
└── {slug}.svg
```

Slugs: lowercase, hyphens, descriptive.
- `pipeline-flow.drawio` + `pipeline-flow.svg`
- `service-architecture.drawio` + `service-architecture.svg`
- `dedup-layers.drawio` + `dedup-layers.svg`

### README Embedding

```markdown
## Cómo Funciona

![Pipeline Flow](docs/diagrams/pipeline-flow.svg)
```

Rules:
- Use a relative path from the README location
- Always provide alt text (used by screen readers and when the image fails to load)
- Do NOT wrap in a code block
- Do NOT use HTML `<img>` tag unless you need width control:
  ```html
  <!-- Only if you need to constrain the width -->
  <img src="docs/diagrams/pipeline-flow.svg" alt="Pipeline Flow" width="900"/>
  ```

### Commit Message Convention

```
docs: add pipeline flow diagram (Draw.io SVG)

- Add docs/diagrams/pipeline-flow.drawio (editable source)
- Add docs/diagrams/pipeline-flow.svg (rendered, embedded in README)
- Replace Mermaid block in README with SVG image reference
```

---

## Validation Before Delivery

Run through this mentally before outputting files:

**XML validation**:
- [ ] All `<mxCell>` tags are properly closed
- [ ] No duplicate `id` attributes
- [ ] All edge `source` and `target` reference existing node IDs
- [ ] Swimlane children have the swimlane's `id` as their `parent`
- [ ] Root has exactly `id="0"` and `id="1"` cells before content

**SVG validation**:
- [ ] `viewBox` is set on root `<svg>`
- [ ] No `fill="white"` on root
- [ ] All `<marker>` definitions are inside `<defs>`
- [ ] Every `marker-end` reference matches a defined marker `id`
- [ ] Text has explicit `font-size` and `fill` — no reliance on CSS cascade
- [ ] Multi-line labels use separate `<text>` elements, not `<tspan>` with `dy` (more compatible)

**Visual validation** (mental check at 800px width):
- [ ] No text overlaps with edges or other nodes
- [ ] Edge labels have background rect to ensure readability
- [ ] Arrow tips are visible (not hidden under target node)
- [ ] Group labels are readable against their header background
