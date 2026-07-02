---
type: source
title: "CocoIndex Code — Semantic Code Search Engine"
source_url: https://github.com/cocoindex-io/cocoindex-code
created: 2026-05-19
updated: 2026-05-19
tags:
  - source
  - developer-tools
  - semantic-search
  - mcp
  - rust
status: active
related:
  - "[[cocoindex-io]]"
  - "[[cocoindex-code]]"
---

# CocoIndex Code

An open-source, AST-based semantic code search engine optimized for AI coding agents.

- **Developer**: [[cocoindex-io]]
- **Repository**: [cocoindex-io/cocoindex-code](https://github.com/cocoindex-io/cocoindex-code)
- **Language**: Python, Rust
- **Licence**: Apache 2.0
- **Ecosystem**: Model Context Protocol (MCP)

## Key Takeaways

- Designed specifically to feed clean, structured context to LLM coding agents.
- **AST-Aware Chunking**: Uses Tree-sitter to parse code into logical AST nodes (functions, classes) instead of line-based splits.
- **Incremental Indexing**: Powered by CocoIndex (Rust-based engine) which updates indexes on-the-fly when files change.
- **Local-First & Remote**: Supports local SentenceTransformer embedding models or remote providers via LiteLLM.
- **Zero Infra**: Local SQLite-based vector storage requiring no complex DB servers.
