# Changelog

## v3.1 — Async Automation & Hard Detach — 2026-07-02

### Added

- **n8n Automation Engine** integration (`/automation` folder).
- `bunker-health-check.json` and `bunker-assisted-remediation.json` workflow templates.
- Explicit `n8n MCP` documentation in architecture files.

### Changed

- **Hard Detach**: Purged legacy brand nomenclature (Claude Code) from core architectural patterns, replacing with agnostic `AGENTS.md` standards.
- Re-aligned `WIKI.md` schema documentation from v2.3.1 to v3.0.
- Resolved broken links and updated `testing-design.md` to `mature` status.

## v3.0 — Local Knowledge Operating System — 2026-05-14

### Added

- `bin/bunker-check.sh` one-command local health check.
- `wiki/meta/knowledge-supply-chain.md`.
- `wiki/meta/workflows.md`.
- `wiki/meta/governance.md`.
- Release notes under `docs/releases/`.

### Changed

- Updated README and PROJECT to describe Bunker OS as a local-first knowledge operating system.
- Dashboard now acts as command center.

## v2.9 — Knowledge Supply Chain — 2026-05-14

### Added

- Formal source-to-handover workflow.
- Governance model for evidence, decisions, and automation.

## v2.8 — Command Center Dashboard — 2026-05-14

### Changed

- Rebuilt `wiki/meta/dashboard.md` around active context, health, evidence, agents, governance, and safe operations.

## v2.7 — Agent Operating Layer — 2026-05-14

### Added

- `agents/evidence-indexer.md`.
- `skills/evidence-index/SKILL.md`.
- `commands/evidence.md`.
- `wiki/meta/agent-queue.md`.

## v2.6 — Evidence Vault — 2026-05-14

### Added

- `bin/evidence-index.sh`.
- `wiki/meta/evidence-index.md`.
- `wiki/meta/evidence-manifest.json`.

## v2.5 — Integrity Engine — 2026-05-14

### Changed

- Rebuilt `bin/wiki-integrity.sh` to emit Markdown and JSON health reports.

## v2.4 — Hardening Release — 2026-05-14

### Security

- Removed shared Local REST API `data.json` from distributable tree.
- Added safe Local REST API example and security note.
- Ingest server now binds to localhost by default and supports token protection.

### Changed

- Scripts now resolve root through `BUNKER_HOME`.
- External-effect scripts default to dry-run.
- Removed missing `calendar-beta` from declared plugins.
