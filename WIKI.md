# WIKI.md — LLM Wiki Schema (Bunker v2.3.1)

> This file is the reference for maintaining the Bunker OS Knowledge Engine.
> Based on Andrej Karpathy's LLM Wiki pattern, optimized for **Senior Architects**.

---

## 🏛️ Architecture v2.3.1

The Bunker is structured as a 3-layer cognitive engine:

### Layer 1: The Raw Input (`.raw/`)
Immutable source documents. We never modify them. We use them as the grounding truth.

### Layer 2: The Interconnected Brain (`wiki/`)
The active knowledge base. It is designed for **High Capacity** (100+ lines per note) and **High Connectivity** (Wikilinks).
- **Blueprints**: Executable standards for scaffolding.
- **Concepts**: Technical deep-dives (Go, React, Vector DB).
- **Meta**: Operational state (Hot Cache, Log, Handovers).

### Layer 3: The Operational Tools (`bin/`)
Scripts that turn wiki knowledge into real-world action.
- `bunker-scaffold.sh`: Implementation engine.
- `wiki-integrity.sh`: Hygiene engine.
- `wiki-sync.sh`: Persistence engine.

---

## 📏 The Deep Capacity Standard
A note is "Mature" only when it meets the following criteria:
1. **Length**: 100+ lines of substantive content.
2. **Patterns**: Includes 2026 industry standards and code examples.
3. **Troubleshooting**: Includes a section on "Senior Gotchas" or "Anti-patterns."
4. **Connectivity**: Linked from at least 2 other relevant pages.

---

## 🔄 Agentic Maintenance Loop
Agents in the Bunker are responsible for:
1. **Ingestion**: 7-step workflow for new resources.
2. **Reconciliation**: Solving contradictions between new data and old notes.
3. **Deciding**: Documenting AgDRs for every significant architectural change.
4. **Cleaning**: Running integrity reports and fixing orphans daily.

---

## 🚀 Operational Workflow
```bash
# 1. Start session
cat wiki/hot.md

# 2. Add resource
ingest https://some-docs.com

# 3. Build something
./bin/bunker-scaffold my-api --blueprint go

# 4. End session
./bin/wiki-sync.sh
```

---
*Built for 2026. Knowledge is not a document; it's a runtime.*
