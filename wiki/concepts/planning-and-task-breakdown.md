---
type: concept
title: "Planning and Task Breakdown"
domain: engineering
created: 2026-05-16
updated: 2026-05-16
tags: [planning, task breakdown, engineering]
---

# Planning and Task Breakdown

## Overview
Decompose a spec or requirements into small, verifiable tasks with clear acceptance criteria. Proper planning ensures implementable units, correct ordering, risk mitigation, and effective parallelization.

## Key Practices
1. **Enter Plan Mode** – Read-only review before coding; no implementation until a plan doc is drafted.
2. **Map Dependencies** – Visualize component/graph dependencies; build foundations first.
3. **Slice Vertically** – Deliver end‑to‑end user flows (e.g., create account) rather than horizontal layers.
4. **Write Tasks** – Each task includes description, acceptance criteria, verification steps, dependencies, and estimated scope.
5. **Order & Checkpoint** – Arrange tasks by dependency, add checkpoints, keep tasks small (S/M size), break large tasks further.
6. **Document Risks & Questions** – Capture open items and mitigation strategies; obtain human approval before proceeding.

## Task Template
```markdown
## Task [N]: [Title]
**Description:** ...
**Acceptance criteria:**
- [ ] ...
- [ ] ...
**Verification:**
- [ ] Tests pass: `npm test -- --grep "..."`.
- [ ] Build succeeds: `npm run build`.
- [ ] Manual check: ...
**Dependencies:** [...]
**Files likely touched:** [...]
**Estimated scope:** [XS/S/M/L/XL]
```

## Parallelization
- Safe to parallelize: independent feature slices, documentation, tests for completed work.
- Sequential: database migrations, shared state changes, API contract updates.
- Coordination needed: shared contracts; define first, then parallelize.

## Common Pitfalls
- Starting implementation without a written task list.
- Tasks lacking acceptance criteria or verification steps.
- All tasks being XL‑size; break them down.
- Ignoring dependency ordering or checkpoints.

## Verification
- [ ] Every task has clear acceptance criteria and verification.
- [ ] Dependencies are identified and ordered.
- [ ] No task exceeds ~5 files.
- [ ] Checkpoints exist between phases.
- [ ] Human has reviewed and approved the plan.