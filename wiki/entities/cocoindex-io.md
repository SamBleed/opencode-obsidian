---
type: organization
title: "CocoIndex"
domain: developer-tools
created: 2026-05-19
updated: 2026-05-19
tags:
  - entity
  - organization
  - semantic-search
  - rust
status: active
related:
  - "[[cocoindex-code]]"
---

# CocoIndex

The developer organization behind the **CocoIndex** data transformation and codebase semantic search ecosystem.

## Core Products

- **[[cocoindex-code]]** — Open-source AST-based semantic code search engine and MCP server designed for AI coding agents.
- **CocoIndex Core** — Rust-based high-performance data pipeline and transformation framework.

## Architecture Principles

- **Local-First**: Private, fast, offline-capable indexing.
- **AST-Aware**: Compilers and parsers (Tree-sitter) before naive text splitting.
- **Incrementalism**: Delta-based updates that scale gracefully with codebase sizes.
