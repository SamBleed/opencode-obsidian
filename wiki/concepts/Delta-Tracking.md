---
type: concept
title: "Delta Tracking"
complexity: intermediate
domain: knowledge-management
created: 2026-04-08
updated: 2026-04-08
tags:
  - concept
  - optimization
  - performance
status: mature
related:
  - "[[LLM Wiki Pattern]]"
  - "[[Ar9av-obsidian-wiki]]"
  - "[[opencode-obsidian-ecosystem-research]]"
sources:
---

# Delta Tracking

A performance optimization for wiki systems that prevents re-processing unchanged sources.

---

## The Problem

Traditional wiki systems process every source on every ingest, even if nothing changed. This wastes:
- Token usage
- API calls
- Processing time
- Creates duplicate pages

## The Solution

Delta tracking maintains a **manifest** (`.manifest.json`) per source:
- File hash (SHA256)
- Timestamp
- List of pages produced

Before processing:
1. Compute current hash
2. Compare with manifest
3. Skip if matches (source unchanged)

## Implementation

```json
{
  "sources": {
    "mcp-security.txt": {
      "hash": "sha256...",
      "processed": "2026-04-20T10:00:00Z",
      "pages": ["mcp-security", "security-guardrails"]
    }
  }
}
```

## Tools Using This Pattern

- [[Ar9av-obsidian-wiki]]
- rvvk7895/llm-knowledge-bases

## Backlog for opencode-obsidian

Planned for v1.5.0:
- Add `.manifest.json` per `.raw/` source
- Update wiki-ingest skill to check hash before processing
- Provide `/force` flag to bypass