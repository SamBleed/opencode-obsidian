---
type: concept
title: "LLM Wiki Pattern"
tags: [concept, architecture, pattern]
status: active
sources:
  - "https://gist.github.com/karpathy/442a6bf555914893e9891c11519de94f"
---

# LLM Wiki Pattern

A pattern for persistent, compounding knowledge bases using AI agents. Created by Andrej Karpathy.

Each session reads `hot.md` for context, writes updates before closing. Every source ingested compounds the knowledge base.
