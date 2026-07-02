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

# Recent Context — Bunker OS v1.3.1

This is the active context loaded at the start of each session.

## Current Focus

Bunker OS has been evolved from v1.2.4 into **v1.3.1** with ported skills from claude-obsidian v1.9, a full test suite, GitHub Actions CI, n8n Dead Letter Queue, and AOC v4 Enterprise.

## Completed Upgrade Chain

- **v1.2.4 Hardening**: Local REST API secret file removed from distributable tree, scripts made portable, dry-run by default.
- **v1.2.5 Integrity Engine**: Automated vault health scripts with Markdown + JSON reports.
- **v1.2.6 Evidence Vault**: Evidence indexed with SHA256 checksums.
- **v1.2.7 Agent Operating Layer**: evidence-indexer agent, evidence skill, agent queue.
- **v1.2.8 Command Center**: Dashboard rebuilt as operational entry point.
- **v1.2.9 Knowledge Supply Chain**: source-to-handover workflow documented.
- **v1.3.0 Local Knowledge OS**: Full local health check, governance, workflows.
- **v1.3.1 Feature Port**: Autoresearch, wiki-retrieve, think framework, 431 tests, CI, DLQ, AOC v4.

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
