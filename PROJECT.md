# opencode-obsidian — Bunker OS v2.3.1

This repository is an OpenCode agent environment and an Obsidian wiki vault.

**Version:** 2.3.1 (Senior Architect Release)
**Status:** Stable / Production Ready
**Vault path:** This directory (open in Obsidian directly)

## 🏗️ Core Philosophy: "Deep Capacity"
Bunker OS follows the **Deep Capacity** standard. All core technical notes must exceed 100 lines and include senior-level patterns, troubleshooting guides, and modern (2026) conventions.

## 📂 Vault Structure

```text
.
├── .raw/               # Immutable source documents (articles, PDFs)
├── bin/                # Operational tools (scaffold, integrity, sync)
├── docs/               # Executive reports and briefings
├── wiki/               # Structured knowledge base
│   ├── blueprints/     # Executable architecture templates
│   ├── concepts/       # Tech stack and domain patterns
│   ├── entities/       # People, orgs, products
│   └── meta/           # Log, Hot Cache, and Handovers
└── BUNKER-OS.canvas    # Visual Mission Control
```

## 🛠️ Key Capabilities

| Tool | Action | Description |
|------|--------|-------------|
| **Ingest** | `ingest [file]` | Process sources into cross-linked wiki pages. |
| **Scaffold** | `bunker-scaffold` | Generate Go/React projects using v2.1 Blueprints. |
| **Audit** | `wiki-integrity` | Detect knowledge debt (orphans/seeds). |
| **Governance** | `ADR/AgDR` | Document and enforce architectural decisions. |
| **Grounding** | `NotebookLM` | Deep research via 10+ connected notebooks. |

## 🚀 How to Use

1. **Initialize**: Read `wiki/hot.md` to get current context.
2. **Build**: Use `./bin/bunker-scaffold.sh` to start a new service.
3. **Research**: Drop URLs into `.raw/` and ingest them to expand the brain.
4. **Govern**: Save every major decision as an `AgDR` in the wiki.

## 🌐 Cross-Project Access
To use this brain in another project, add to its `AGENTS.md`:
```markdown
## Wiki Knowledge Base
Path: /home/sam/opencode-obsidian
Follow the protocols in [[vault-flow]].
```

---
MIT License © 2026 | **SamBleed**
