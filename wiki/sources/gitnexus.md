---
type: source
title: "GitNexus"
description: "Building nervous system for agent context. Indexes any codebase into a knowledge graph — every dependency, call chain, cluster, and execution flow — then exposes it through smart tools so AI agents never miss code."
author: "Abhigyan Patwari"
url: "https://github.com/abhigyanpatwari/GitNexus"
created: 2026-04-22
updated: 2026-04-22
tags: [knowledge-graph, mcp, code-analysis, tree-sitter, agentic]
status: active
---

# GitNexus

## Overview

**"Building nervous system for agent context."**

GitNexus indexes any codebase into a knowledge graph — every dependency, call chain, cluster, and execution flow — then exposes it through smart tools so AI agents never miss code.

### The Problem It Solves

Tools like Cursor, Claude Code, Codex, Cline, Roo Code, and Windsurf are powerful — but they don't truly know your codebase structure:

1. AI edits `UserService.validate()`
2. Doesn't know 47 functions depend on its return type
3. **Breaking changes ship**

### How It Works

GitNexus builds a complete knowledge graph through a multi-phase indexing pipeline:

1. **Structure** — Walks the file tree and maps folder/file relationships
2. **Parsing** — Extracts functions, classes, methods, and interfaces using Tree-sitter ASTs
3. **Resolution** — Resolves imports, function calls, heritage, constructor inference
4. **Clustering** — Groups related symbols into functional communities (Leiden algorithm)
5. **Processes** — Traces execution flows from entry points through call chains
6. **Search** — Builds hybrid search indexes (BM25 + semantic + RRF)

## Two Ways to Use

| | CLI + MCP | Web UI |
|---|---|---|
| **What** | Index repos locally, connect AI agents via MCP | Visual graph explorer + AI chat in browser |
| **For** | Daily development with Cursor, Claude Code, Codex, OpenCode | Quick exploration, demos |
| **Scale** | Full repos, any size | Limited by browser memory (~5k files) |
| **Install** | `npm install -g gitnexus` | No install — [gitnexus.vercel.app](https://gitnexus.vercel.app) |
| **Storage** | LadybugDB native (fast, persistent) | LadybugDB WASM (in-memory) |

## MCP Tools (16 total)

| Tool | What It Does |
|------|-------------|
| `query` | Process-grouped hybrid search (BM25 + semantic + RRF) |
| `context` | 360-degree symbol view — categorized refs, process participation |
| `impact` | Blast radius analysis with depth grouping and confidence |
| `detect_changes` | Git-diff impact — maps changed lines to affected processes |
| `rename` | Multi-file coordinated rename with graph + text search |
| `cypher` | Raw Cypher graph queries |
| `list_repos` | Discover all indexed repositories |

## Supported Languages

TypeScript, JavaScript, Python, Java, Kotlin, C#, Go, Rust, PHP, Ruby, Swift, C, C++, Dart

## Tech Stack

- **Parsing**: Tree-sitter (native bindings for CLI, WASM for web)
- **Database**: LadybugDB (embedded graph database with vector support)
- **Search**: BM25 + semantic + RRF hybrid
- **Agent Interface**: MCP (stdio)
- **Clustering**: Graphology (Leiden community detection)

## Key Features

- **Precomputed Relational Intelligence** — LLM can't miss context, it's already in the tool response
- **Token efficiency** — No 10-query chains to understand one function
- **Model democratization** — Smaller LLMs work because tools do the heavy lifting
- **Multi-repo support** — One MCP server serves multiple indexed repos
- **Wiki generation** — Auto-generate documentation from knowledge graph

## Security & Privacy

- **CLI**: Everything runs locally. No network calls. Index stored in `.gitnexus/`
- **Web**: Everything in browser. No code uploaded to any server.

## Related Concepts

- [[knowledge-graph-code]]
- [[mcp-protocol]]
- [[tree-sitter]]
- [[code-analysis]]

## Related Entities

- [[Abhigyan-Patwari]] — Author
- [[Akonlabs]] — Enterprise offering