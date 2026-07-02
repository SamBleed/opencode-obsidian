# Changelog

## v1.3.1 — Feature Port desde claude-obsidian v1.9 — 2026-07-02

### Added

**Skills**
- **Autoresearch**: Research loop autónomo de 3 rondas con Exa + webfetch + filing a wiki. Porteado de claude-obsidian v1.9.2 y adaptado a OpenCode.
- **wiki-retrieve**: BM25 index (211 chunks, 6.544 términos) + cosine rerank vía ollama nomic-embed-text. Búsqueda híbrida por significado.
- **/think**: Framework de 10 principios (OBSERVE-OBSERVE-LISTEN-THINK-CONNECT-CONNECT-FEEL-ACCEPT-CREATE-GROW) para decisiones arquitectónicas. En español.

**Testing**
- **Makefile** con 5 targets: workflows n8n, wiki integrity, scripts, YAML, retrieve
- **test_workflow_connections.py**: 344 tests, valida nodos y conexiones de todos los workflows n8n
- **test_wiki_integrity.sh**: 21 tests, archivos esenciales, frontmatter, docker status
- **test_scripts.sh**: 61 tests, sintaxis bash, shebang, secretos hardcodeados, go vet
- **GitHub Actions CI** (`.github/workflows/test.yml`): 5 suites en cada push/PR a main

**Automatización n8n**
- **Dead Letter Queue**: Error trigger global que atrapa errores de todos los workflows y persiste en staticData
- **AOC v4 Enterprise**: 37 nodos con AI triage (OpenRouter/Mistral), idempotencia Redis, GitHub issues + Slack/Telegram/Discord
- Docker tuning: `N8N_CONCURRENCY_PRODUCTION_LIMIT`, `EXECUTIONS_DATA_PRUNE`, `N8N_METRICS`, `N8N_LOG_FORMAT=json`
- Limpieza de 5 workflows zombies del repo

**Infraestructura**
- `.env` saneado: removidas API keys muertas, agregadas vars del AOC v3, dry-run por defecto
- docker-compose: removido `version:` deprecado, agregado production tuning
- Hooks mejorados: PostCompact con `cat` command, Stop con prompt auto-update hot.md, PostToolUse con stage de tests
- Token n8n actualizado, Excalidraw MCP eliminado

### Changed
- README.md actualizado con tabla de características, comparativa vs claude-obsidian, badges de CI
- PROJECT.md actualizado a v3.1, 13 skills, nueva arquitectura
- n8n reiniciado con nueva configuración de entorno

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
