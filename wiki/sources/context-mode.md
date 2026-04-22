---
type: source
title: "context-mode"
description: "MCP plugin that saves 98% of context window through sandboxed code execution, FTS5 knowledge base, and intent-driven search"
author: "Mert Koseoğlu"
url: "https://github.com/mksglu/context-mode"
created: 2026-04-22
updated: 2026-04-22
tags: [mcp, context-window, sandbox, code-execution, fts5]
status: active
---

# context-mode

## Overview

MCP (Model Context Protocol) plugin developed by Mert Koseoğlu that solves one of the biggest problems in LLM-assisted development: **wasting the context window**.

### The Problem It Solves

When working with Claude Code, Gemini CLI, VS Code Copilot, OpenCode or Codex CLI, tools tend to dump raw data into the context (logs, command outputs, search results). This massively consumes tokens and reduces the model's reasoning capacity.

### How It Works

1. **Sandboxed Code Execution**: Instead of reading large files/outputs into context, the plugin allows executing JavaScript code in a sandbox and returning only the processed answer.

2. **FTS5 Knowledge Base**: Uses SQLite with FTS5 (Full-Text Search 5) to index and search information without loading everything into context.

3. **Intent-Driven Search**: Optimized search tools that return processed results, not raw data.

4. **Batch Processing**: A single call can execute multiple commands, replacing dozens of individual calls.

### Core Philosophy

> **"Think in Code — MANDATORY"**: When you need to analyze, count, filter, search, parse or process data, **write code** that does the work and return only the answer. Do NOT mentally process raw data that the model throws into context.

### Main Tools

- `batch_execute`: Execute multiple commands, auto-index and search
- `execute`: Execute JavaScript code (only Node.js built-ins: fs, path, child_process)
- `fetch_and_index`: Web fetch + index for subsequent searching
- `search`: Intent-driven search

### Compatibility

Works with:
- Claude Code (`.claude-plugin`)
- OpenCode (`./build/opencode-plugin.js`)
- OpenClaw (`./build/openclaw-plugin.js`)
- Gemini CLI
- VS Code Copilot
- Codex CLI

### Stats

- Version: 1.0.89
- License: Elastic-2.0
- Main dependencies: `@modelcontextprotocol/sdk`, `better-sqlite3` (optional), `zod`

### Benchmarks

According to BENCHMARK.md, the plugin promises to save up to **98% of the context window** in certain usage scenarios.

---

## Key Source Files

| File | Purpose |
|------|---------|
| `src/executor.ts` | Polyglot sandbox executor (JS, TS, Python, Go, Rust, etc.) |
| `src/runtime.ts` | Automatic runtime detection |
| `src/security.ts` | Environment variable denylist for sandbox security |
| `src/db-base.ts` | SQLite FTS5 infrastructure |
| `src/server.ts` | MCP server implementation |
| `CLAUDE.md` | Instructions for AI agents using the project |

---

## Relevant Concepts

- [[sandboxed-execution]]
- [[context-window-optimization]]
- [[fts5-search]]
- [[mcp-protocol]]

## Related

- Author: [[mert-koseoglu]]