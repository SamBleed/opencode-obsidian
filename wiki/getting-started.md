---
type: meta
title: "Getting Started"
updated: 2026-05-13
tags: [meta, onboarding, quickstart]
status: evergreen
related:
  - "[[index]]"
  - "[[overview]]"
  - "[[BUNKER-OS.canvas]]"
  - "[[LLM Wiki Pattern]]"
---

# Getting Started with Bunker OS

Welcome to the **Bunker**. This is more than a wiki; it's a persistent second brain and operational toolkit designed for Senior Architects and AI Agents.

## 🚀 Three-Step Quick Start

### 1. Ingest Knowledge
Drop any technical resource (PDF, URL, Transcript) into `.raw/` and run:
```bash
ingest [filename]
```
The system will create 8-15 cross-linked pages, updating the index and graph automatically.

### 2. Scaffold a Project
Turn theory into code. Use the operational command center to start a new service:
```bash
./bin/bunker-scaffold [project-name] --blueprint [go|react]
```
This generates a production-grade folder structure, Dockerfiles, and security checklists.

### 3. Maintain Integrity
Ensure your knowledge graph stays healthy by running the integrity script:
```bash
./bin/wiki-integrity.sh
```
Check `wiki/meta/INTEGRITY-REPORT.md` for orphans or knowledge gaps.

## 🛠️ Key Commands

| Action | Command | Purpose |
|--------|---------|---------|
| **Ingest** | `ingest [file]` | Process new sources into the wiki. |
| **Query** | `what do you know about X?` | Synthesize answers from vault context. |
| **Scaffold** | `bunker-scaffold` | Create projects based on standards. |
| **Lint** | `wiki-integrity.sh` | Find orphans and broken links. |
| **Sync** | `bin/wiki-sync.sh` | Persist and push all changes. |
| **Save** | `/save` | File a conversation as a wiki note. |

## 🗺️ Navigation Hub
- **[[BUNKER-OS.canvas]]** — Visual center of the whole system.
- **[[index]]** — Master catalog of all projects, concepts, and sources.
- **[[hot]]** — Current focus and recent context summary.
- **[[Security-Guardrails]]** — The rules of the game.
- **[[FAQ]]** — Common questions about the Bunker.

---
*Built on the [LLM Wiki pattern](https://github.com/karpathy) by Andrej Karpathy.*
