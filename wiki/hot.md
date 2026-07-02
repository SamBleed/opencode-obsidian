---
type: meta
title: "Hot Cache"
updated: 2026-05-14
tags: [meta, hot-cache, context]
status: evergreen
related:
  - "[[index]]"
  - "[[overview]]"
  - "[[dashboard]]"
  - "[[agent-queue]]"
  - "[[evidence-index]]"
  - "[[knowledge-supply-chain]]"
  - "[[log]]"
---

# Recent Context — Bunker OS v3.0

This is the active context loaded at the start of each session.

## Current Focus

Bunker OS has been evolved from v2.3.1 into **v3.0 Local Knowledge Operating System** without a destructive rewrite. The vault remains an Obsidian-first knowledge system, not a traditional web app.

## Completed Upgrade Chain

- **v2.4 Hardening**: Local REST API secret file removed from distributable tree, safe example added, scripts made portable through `BUNKER_HOME`, external-effect scripts made dry-run by default, ingest server hardened.
- **v2.5 Integrity Engine**: `bin/wiki-integrity.sh` now emits Markdown and JSON reports.
- **v2.6 Evidence Vault**: `report.zip` and `security-audit-report.json` indexed with SHA256 checksums.
- **v2.7 Agent Operating Layer**: evidence-indexer agent, evidence skill/command, and `wiki/meta/agent-queue.md` added.
- **v2.8 Command Center**: `wiki/meta/dashboard.md` rebuilt as the operational entry point.
- **v2.9 Knowledge Supply Chain**: source-to-handover workflow documented.
- **v3.0 Local Knowledge OS**: `bin/bunker-check.sh` validates scripts, Go build, secret policy, path portability, evidence index, and integrity report.

## Start Here

1. Open [[dashboard]].
2. Review [[agent-queue]].
3. Run `./bin/bunker-check.sh` before major changes.
4. Use `--apply` only when external effects are intentional.

## Active Risks / Known Debt

- `wiki/meta/INTEGRITY-REPORT.md` still lists active broken links, short notes, and missing frontmatter.
- Some `.raw/` files remain pending review against `wiki/meta/ingestion-log.md`.
- `report.zip` remains packed by design; extraction/index interpretation should be a deliberate future decision.

## Latest Ingestion

- **[[insforge]]** (source) — Agentic backend platform by [[InsForge]], 9.9k ⭐
- **[[InsForge]]** (entity) — Organization behind the platform
- **[[agentic-backend-platform]]** (concept) — Backend platforms designed for agent-native operation via MCP

## Safe Commands

```bash
./bin/bunker-check.sh
./bin/wiki-integrity.sh
./bin/evidence-index.sh
./bin/bunker.sh init
```

---
Use this hot cache as the first read for the next session.
