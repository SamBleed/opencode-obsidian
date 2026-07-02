# opencode-obsidian — Bunker OS v3.0

This repository is an OpenCode agent environment and an Obsidian wiki vault.

**Version:** 3.0  
**Status:** Local-first operational knowledge system  
**Vault path:** current repository root. Use `BUNKER_HOME` to override script resolution.

## Core identity

Bunker OS is not a traditional application. It is a local-first knowledge operating system for:

- AI session continuity
- structured research
- audit evidence preservation
- project intelligence
- agent-guided workflows
- governance and decisions
- visual command center operations

## Architecture layers

```text
Capture       → wiki/inbox, .raw, ingest server, handovers
Normalize     → _templates, skills/wiki-ingest
Classify      → wiki/sources, entities, concepts, projects, comparisons
Evidence      → report.zip, security-audit-report.json, evidence index
Governance    → BUNKER_RULES, ADR template, knowledge supply chain
Automation    → bin/ (Sync ops) + n8n (Async flows & MCP)
Agent Runtime → agents, commands, skills, hooks
Visualization → dashboard, canvases, Obsidian graph
```

## Release state

v3.0 includes the completed sequence:

- v2.4 Hardening
- v2.5 Integrity Engine
- v2.6 Evidence Vault
- v2.7 Agent Operating Layer
- v2.8 Command Center Dashboard
- v2.9 Knowledge Supply Chain
- v3.0 Local Knowledge Operating System

## The n8n Orchestrator Role

n8n acts as the asynchronous execution engine and "nervous system" of the Bunker:
- **Security & Decoupling**: API keys and external integrations (GitHub, Telegram) live in n8n, keeping bash scripts clean and secure. Scripts just fire webhooks to `localhost:5678`.
- **Agentic MCP Bridge**: Gives OpenCode "hands" to trigger complex pipelines (e.g., assisted remediation) in a sandboxed manner without arbitrary code execution.
- **Complex Pipelines**: Manages retries, conditional logic, and LLM evaluations via visually orchestrable flows stored in `automation/n8n-lab`.

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
| Safe orchestrator | `./bin/bunker.sh` |

## Safety model

- Scripts resolve the vault through `BUNKER_HOME` or repository root.
- External-effect scripts default to dry-run.
- Obsidian Local REST API secrets stay local.
- Ingest server binds to `127.0.0.1:9090` by default.
- Evidence artifacts are indexed, not modified.

## Cross-project access

To use this brain from another project, point that project to the vault root and follow the read order:

1. `wiki/hot.md`
2. latest `wiki/meta/handovers/*`
3. `wiki/index.md`
4. specific notes only when needed

Avoid hard-coded absolute paths. Use `BUNKER_HOME` or project-local configuration.

---
MIT License © 2026 | **SamBleed**
