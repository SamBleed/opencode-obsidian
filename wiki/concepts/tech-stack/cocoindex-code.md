---
type: concept
title: "CocoIndex Code — Semantic Code Search Integration"
domain: tech-stack
created: 2026-05-19
updated: 2026-05-19
tags:
  - concept
  - tech-stack
  - semantic-search
  - mcp
  - indexing
status: active
related:
  - "[[cocoindex-code]]"
  - "[[cocoindex-io]]"
  - "[[Vector-Database-Patterns]]"
---

# CocoIndex Code Integration

## Definition
An Abstract Syntax Tree (AST) aware semantic code search engine that integrates with **OpenCode** via the Model Context Protocol (MCP) to provide high-fidelity codebase navigation for AI agents.

## How It Works
1. **Parser Layer**: Uses Tree-sitter to chunk code files by functional syntax blocks (classes, functions) rather than plain line lengths.
2. **Transform & Cache**: Uses the Rust-based CocoIndex core engine to compute incremental code hashes, detecting which files have changed.
3. **Embeddings Layer**: Integrates with LiteLLM to call cloud embedding models (e.g. Gemini, OpenAI) or uses SentenceTransformers for local embedding generation.
4. **Vector Storage**: Employs a zero-infrastructure local SQLite database (`.cocoindex_code/target_sqlite.db`) to store chunks and vector representations.
5. **Agent Interface**: Exposes a standard MCP stdio channel, allowing the AI agent to call the search tool semantically to answer technical queries.

## Integration Patterns

### 1. OpenCode MCP Configuration
Add the server globally at `~/.config/opencode/opencode.json` under `"mcp"`:

```json
"cocoindex-code": {
  "command": [
    "/home/sam/.local/bin/ccc",
    "mcp"
  ],
  "type": "local"
}
```

### 2. Global Configuration & Throttling
For free-tier APIs like Gemini's 100 requests/minute, concurrency must be throttled to prevent `429 Too Many Requests`. This is configured at `~/.cocoindex_code/global_settings.yml`:

```yaml
embedding:
  provider: litellm
  model: gemini/gemini-embedding-2

# Environment variables injected into the background daemon.
envs:
  GEMINI_API_KEY: AIzaSyAH...
  COCOINDEX_SOURCE_MAX_INFLIGHT_ROWS: "5"  # Prevents API Rate Limits
```

### 3. Git Management
Since the local SQLite DB contains hundreds of megabytes of compiled vectors, exclude the workspace folder from version control in `.gitignore`:

```
# CocoIndex Code (ccc)
/.cocoindex_code/
```

## Best Practices
- **Use Remote Embeddings on Modern Python**: On Python versions where binary wheels for heavy libraries (PyTorch, transformers) are not yet compiled, use remote LiteLLM models to bypass local compilation.
- **Throttle Free API Keys**: Always set `COCOINDEX_SOURCE_MAX_INFLIGHT_ROWS` in quotes as a string to avoid Python `TypeError` inside daemon environment loader.
- **Local Indexing**: Initialize the index with `ccc init` and build it with `ccc index` before running the agent to pre-warm the search cache.
