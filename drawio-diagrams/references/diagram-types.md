# Diagram Types — Layout Rules

Read the section matching your diagram type before laying out nodes.

---

## 1. Flowchart / Pipeline Flow

**Orientation**: Top-to-bottom (TD) for step-by-step processes. Left-to-right (LR) for data pipelines where the "source → transform → sink" mental model applies.

**Layout rules**:
- Entry point at top-left corner (TD) or leftmost column (LR)
- Final states (APPROVED, BLOCKED, DONE) at bottom or rightmost column
- Parallel branches at the same vertical (TD) or horizontal (LR) position
- Decision diamonds must have exactly 2 exits: "yes" and "no" (or equivalent)
- Use swimlane containers for logical phases: "Input", "Processing", "Output"
- Keep each swimlane roughly equal width; vary height based on content
- Place retry/error paths as secondary arrows curving back — not as new columns

**Node sizing**:
- Swimlane containers: wide enough to hold 2–3 nodes side by side with padding
- Process nodes: 120×40px minimum
- Decision diamonds: 80×60px
- Terminal nodes (start/end): 100×30px stadium shape

**Edge routing**:
- Happy path: straight arrows, primary color
- Error/retry path: dashed, curved, secondary color
- Never route edges through unrelated swimlanes

**Example structure** (TD pipeline):
```
[Input Group]
    [Trigger A]   [Trigger B]
         ↓              ↓
[Processing Group]
    [Step 1] → [Step 2] → [Decision?]
                              ↓ yes        ↓ no
[Output Group]           [Success]     [Failure]
```

---

## 2. Architecture Diagram

**Orientation**: Left-to-right (LR). Clients on the left, core services in the center, infrastructure/storage on the right.

**Layout rules**:
- 3–4 vertical columns maximum: Clients | Gateway/API | Core Services | Infra
- Group services into swimlane containers by deployment unit or responsibility
- Message buses / queues go between the layers they connect, not at the edge
- Show data flow direction with arrows — avoid bidirectional arrows unless the protocol is truly symmetric (e.g. WebSocket)
- Auth/security layer: place as a vertical band intersecting the API column, or as a small group adjacent to API
- Observability (Prometheus, logs): bottom of diagram or side panel

**Node sizing**:
- Service boxes: 140×50px
- Database cylinders: 100×60px
- Message queues: 120×40px parallelogram
- Group containers: size to content + 30px padding all sides

**What to show vs omit**:
- SHOW: all services, their communication protocol, storage systems, external integrations
- OMIT: internal implementation details (goroutines, memory layout), deployment configs unless relevant
- LABEL edges with protocol or method: "PUB scans.new", "REST :8080", "OTLP traces"

**Example structure** (LR architecture):
```
[Clients]          [API Layer]         [Core Services]      [Infra]
  Web UI     →    API Gateway    →       Worker         →   PostgreSQL
  CLI        →    (Auth/RBAC)    →       Notifier       →   NATS
  Webhook    →                           ↕ NATS bus     →   S3
                                                         →   Ollama
```

---

## 3. Sequence Diagram

**Orientation**: Top-to-bottom. Participants as columns, time flows downward.

**Layout rules**:
- Place participants in order of first interaction, left to right
- Lifeline bars (vertical dashed lines) span the full height of the diagram
- Activation boxes on lifelines show when a participant is actively processing
- Synchronous calls: solid arrow with filled head
- Async / fire-and-forget: dashed arrow
- Return values: dashed arrow in reverse direction, labeled with return value
- Group related exchanges in boxes (loop, alt, opt, par)

**Node sizing**:
- Participant boxes: 120×40px
- Message labels: keep to < 30 chars; use abbreviations
- Activation boxes: 10px wide, variable height

**Common patterns to represent correctly**:
- Request/response: solid arrow right → activation box → dashed arrow left
- Async publish: solid arrow → no return arrow (fire and forget)
- Error handling: "alt" frame with [success] and [error] branches

---

## 4. Entity-Relationship (ER) Diagram

**Orientation**: No fixed direction. Start with the most connected entity at center.

**Layout rules**:
- Place the central/most-referenced entity in the center of the canvas
- Radiate related entities outward
- Use crow's foot notation for cardinality (1:1, 1:N, M:N)
- Show only foreign key relationships as edges — not every field
- Use different colors for different entity categories (user data, system data, config)

**Node sizing**:
- Entity boxes: wide enough for the longest field name + 20px padding
- Show: entity name (bold), primary key (underlined or PK label), important fields
- Omit: `created_at`, `updated_at` unless part of the diagram's purpose

**Cardinality labels**:
- Place near the entity end of the edge
- "1", "N", "0..1", "0..N" — keep it simple

---

## 5. State Machine

**Orientation**: Left-to-right generally, but circular arrangements work for small machines (< 8 states).

**Layout rules**:
- Initial state: filled circle or rounded rectangle with bold border, leftmost
- Terminal state: double-border circle or rounded rectangle, rightmost
- Happy path: straight arrows left-to-right
- Error/retry transitions: arcs curving above or below the happy path
- Self-transitions (loops): small arc looping back to same node, labeled with trigger
- Group states into swimlanes if they belong to a lifecycle phase

**Node sizing**:
- State nodes: 120×40px rounded rectangle
- Initial/terminal: 60×60px circle
- Transition labels: must include trigger event + optional [guard condition] + /action

**Example**:
```
● → [PENDING] → [SCANNING] → [EVALUATING] → [APPROVED]
                     ↓ error                      ↑
                 [FAILED] ←──────── [BLOCKED] ────┘
```

---

## 6. Network / Infrastructure Diagram

**Orientation**: Usually top-to-bottom with network zones as horizontal bands.

**Layout rules**:
- Zones as horizontal swimlanes: Internet | DMZ | Private Network | Database Network
- Place load balancers / gateways at zone boundaries
- Use icons for physical/logical devices (router, firewall, server, cloud)
- Show ports on edges: `:8080`, `:5432`, `:4222`
- Dashed lines for logical/virtual connections, solid for physical

**Node sizing**:
- Server/instance: 80×60px with icon
- Network zone: full-width swimlane, 150px+ height
- Firewall/LB: 60×40px diamond or trapezoid

---

## Choosing Between Types

| Scenario | Type |
|----------|------|
| "How does a request flow through the system?" | Flowchart (TD) |
| "What services exist and how do they connect?" | Architecture (LR) |
| "What messages are exchanged between components?" | Sequence |
| "What does the database schema look like?" | ER |
| "What states can a scan be in?" | State Machine |
| "How is the infrastructure organized?" | Network |
| "Show me both the flow AND the services" | Architecture with flow annotations |
