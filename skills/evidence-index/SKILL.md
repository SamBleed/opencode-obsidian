---
name: evidence-index
description: Index immutable audit evidence into the Bunker OS evidence manifest and evidence index.
---

# Evidence Index Skill

Use this skill to preserve audit outputs as immutable evidence, calculate checksums, and link them to interpretive wiki notes.

## Rules

- Keep original artifacts unchanged.
- Index `report.zip` and `security-audit-report.json`.
- Write only to `wiki/meta/evidence-index.md` and `wiki/meta/evidence-manifest.json`.
