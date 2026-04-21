---
type: source
title: "OpenCode + Obsidian Ecosystem Research"
source_url: null
created: 2026-04-08
updated: 2026-04-08
tags:
  - ai
  - obsidian
  - knowledge-management
  - ecosystem
status: complete
related:
  - "[[LLM Wiki Pattern]]"
  - "[[index]]"
---

# OpenCode + Obsidian Ecosystem Research

> Researched: 2026-04-08 | Goal: cherry-pick best features for opencode-obsidian

Research into 16+ active projects combining AI assistants with Obsidian vaults. Organized into 4 categories.

---

## Summary

The ecosystem has matured significantly with three distinct approaches:
1. **LLM Wiki Pattern plugins** (OpenCode Code skill-based)
2. **Native Obsidian plugins** (embedded UI)
3. **MCP servers** (bridge protocols)
4. **Traditional plugins** (community store)

Key finding: Most successful projects implement delta tracking and auto-commit hooks.

---

## Key Takeaways

- **Delta tracking** (Ar9av/obsidian-wiki) prevents re-processing sources
- **Multi-depth queries** (rvk7895/llm-knowledge-bases) provide quick/standard/deep modes
- **Goal cascade** (ballred/obsidian-claude-pkm) integrates productivity with knowledge
- **Auto-commit hooks** keep vault in version control automatically
- **Multi-agent compatibility** reduces lock-in risk
- **Hybrid search** (BM25 + vector) improves recall
- **Vision ingestion** enables image-based sources
- **Output formats** include Marp slides and charts

## Top Competitors Analysis

| Project | Stars | Key Differentiator |
|---------|-------|---------------------|
| logancyang/obsidian-copilot | 5,776 | Multi-provider AI chat |
| brianpetro/obsidian-smart-connections | 4,357 | Vector embeddings |
| nhaouari/obsidian-textgenerator-plugin | 1,837 | Text generation |
| bramses/chatgpt-md | 1,229 | Chat in markdown |

## Design Patterns to Consider Adopting

1. **Delta tracking manifest** — `.manifest.json` per source
2. **Multi-depth query modes** — Quick/Standard/Deep tiers
3. **Goal cascade** — Daily → Weekly → Monthly reviews
4. **PostToolUse auto-commit** — Git commit on every edit
5. **Hybrid search** — qmd with BM25 + vector
6. **Vision support** — Image ingestion capability

---

## Sources Referenced

- github.com/AgriciDaniel/opencode-obsidian
- github.com/heyitsnoah/claudesidian
- github.com/ProfSynapse/claudesidian-mcp
- github.com/YishenTu/claudian
- github.com/kepano/obsidian-skills
- github.com/Ar9av/obsidian-wiki
- github.com/rvk7895/llm-knowledge-bases
- github.com/ballred/obsidian-claude-pkm
- gist.github.com/karpathy/442a6bf555914893e9891c11519de94f