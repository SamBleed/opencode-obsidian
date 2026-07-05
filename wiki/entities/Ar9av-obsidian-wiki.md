---
type: project
title: "Ar9av/obsidian-wiki"
domain: ai-knowledge-management
created: 2026-04-08
updated: 2026-04-08
tags:
  - project
  - obsidian
  - multi-agent
status: active
related:
  - "[[opencode-obsidian-ecosystem-research]]"
  - "[[Delta Tracking]]"
  - "[[Vision Ingestion]]"
sources:
---

# Ar9av/obsidian-wiki

Multi-agent compatible wiki system by **Ar9av**. One of the most feature-complete LLM Wiki implementations.

---

## Key Features

- **Multi-agent compatibility**: OpenCode Code, Cursor, Windsurf, Codex, Gemini CLI, OpenClaw, GitHub Copilot
- `setup.sh` — auto-configures all agents via symlinks
- **Delta tracking manifest** (`.manifest.json`) — hash-based change detection
- **Vision support** — ingest images, screenshots, whiteboard photos
- 4-stage pipeline: Ingest → Extract → Resolve → Schema
- Schema emerges from sources (not fixed upfront)
- Each page gets `summary:` frontmatter for preview
- Multi-format: PDFs, JSONL, conversation exports, transcripts

## Comparison with opencode-obsidian

| Feature | Ar9av | opencode-obsidian |
|---------|-------|------------------|
| Multi-agent | Yes (7 agents) | No |
| Delta tracking | Yes | No (v1.5.0 backlog) |
| Vision ingestion | Yes | No |
| Emerging schema | Yes | No (fixed structure) |
| Setup script | Yes | No |

## Source

- URL: https://github.com/Ar9av/obsidian-wiki