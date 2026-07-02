# opencode-obsidian — Bunker OS v1.3.1

This repository is an OpenCode agent environment and an Obsidian wiki vault.

**Version:** 1.3.1  
**Status:** Local-first operational knowledge system with n8n automation  
**Tests:** 429 (Makefile + GitHub Actions CI)  
**Vault path:** current repository root. Use `BUNKER_HOME` to override script resolution.

---

## Core identity

Bunker OS is not a traditional application. It is a local-first knowledge operating system for:

- AI session continuity
- structured research (**autoresearch**, **/think** framework)
- audit evidence preservation (**SHA256 checksums**)
- project intelligence
- agent-guided workflow (**n8n AOC v4**, **Dead Letter Queue**)
- hybrid knowledge retrieval (**BM25 + ollama rerank**)
- governance and decisions
- visual command center operations

---

## Architecture layers

```text
Capture       → wiki/inbox, .raw, ingest server, handovers
Normalize     → _templates, skills/wiki-ingest
Classify      → wiki/sources, entities, concepts, projects, comparisons
Evidence      → report.zip, security-audit-report.json, evidence index
Governance    → BUNKER_RULES, ADR template, knowledge supply chain
Automation    → bin/ (Sync ops) + n8n (Async flows, alerts, DLQ)
Retrieval     → BM25 index + ollama nomic-embed-text (cosine rerank)
Testing       → 5 suites: workflows n8n, wiki integrity, scripts, YAML, retrieve
Agent Runtime → agents, commands, skills (13), hooks (7)
Visualization → dashboard, canvases, Obsidian graph
```

---

## v1.3.1 changelog

### Added
- **Autoresearch**: research loop autónomo de 3 rondas con Exa + webfetch
- **wiki-retrieve**: BM25 index (211 chunks, 6.544 términos) + cosine rerank via ollama
- **/think framework**: 10-principios para decisiones arquitectónicas (español)
- **GitHub Actions CI**: 5 suites corren en cada push/PR a main
- **Dead Letter Queue**: error trigger global para n8n
- **AOC v4 Enterprise**: 37 nodos, AI triage, GitHub issues + 3 canales de notificación

### Testing
- `Makefile` con 5 targets: workflows, wiki, scripts, YAML, retrieve
- 344 tests de conexiones n8n
- 61 tests de scripts (sintaxis, secretos, ejecutables)
- 21 tests de integridad del vault
- Test de YAML (CI + docker-compose)
- Test de BM25 + rerank

### Infrastructure
- docker-compose: production tuning (concurrency, metrics, log format, data prune)
- `.env`: vars del AOC v3, removed dead API keys
- n8n: pin version, MCP endpoint configurado, concurrency limit
- PostCompact hook mejorado con `cat` command
- Stop hook con prompt auto-update hot.md

---

## Skills (13)

| Skill | Descripción |
|-------|-------------|
| `autoresearch` | Investigación web autónoma 3 rondas |
| `wiki-retrieve` | BM25 + rerank semántico vía ollama |
| `think` | Framework 10 principios para decisiones |
| `wiki-ingest` | Ingestar fuentes en la wiki |
| `wiki-query` | Consultar la wiki |
| `wiki-lint` | Health check del vault |
| `save` | Guardar conversación como nota |
| `code-review` | Revisión de código |
| `security-review` | Revisión de seguridad |
| `infra-design` | Diseño de infraestructura |
| `tdd-workflow` | Desarrollo guiado por tests |
| `verification-loop` | Build/test/lint pre-PR |
| `work-unit-commits` | Commits organizados |

---

## Operational entry points

| Purpose | File/Command |
|---|---|
| Active context | `wiki/hot.md` |
| Main map | `wiki/index.md` |
| Dashboard | `wiki/meta/dashboard.md` |
| Agent queue | `wiki/meta/agent-queue.md` |
| Integrity scan | `./bin/wiki-integrity.sh` |
| Evidence indexing | `./bin/evidence-index.sh` |
| Full local check | `./bin/bunker-check.sh` |
| Full test suite | `make test` |
| Hybrid search | `python3 scripts/retrieve.py "query"` |
| Autoresearch | `/autoresearch [topic]` via OpenCode |
| n8n status | `docker ps --filter name=n8n` |

---

## Safety model

- Scripts resolve the vault through `BUNKER_HOME` or repository root.
- External-effect scripts default to dry-run.
- Obsidian Local REST API secrets stay local.
- Ingest server binds to `127.0.0.1:9090` by default.
- Evidence artifacts are indexed, not modified.
- n8n secrets managed via `.env` (gitignored) + n8n credential vault.

---

## Cross-project access

To use this brain from another project, point that project to the vault root and follow the read order:

1. `wiki/hot.md`
2. latest `wiki/meta/handovers/*`
3. `wiki/index.md`
4. specific notes only when needed

Avoid hard-coded absolute paths. Use `BUNKER_HOME` or project-local configuration.

---

MIT License © 2026 | **SamBleed**
