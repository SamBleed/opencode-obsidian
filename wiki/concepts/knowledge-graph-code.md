---
type: concept
domain: ai-development
title: "Knowledge Graph for Code"
description: "Building a graph representation of codebase structure (dependencies, call chains, clusters, execution flows) for AI agent context"
created: 2026-04-22
updated: 2026-04-22
tags: [knowledge-graph, code-analysis, tree-sitter, mcp]
status: seed
related: [tree-sitter, mcp-protocol, context-window-optimization]
sources: [gitnexus]
---

# Knowledge Graph for Code

## Definition

A graph representation of a codebase that captures:
- **Dependencies** — imports, references between files
- **Call chains** — how functions call each other
- **Clusters** — functional communities (grouped by code relationships)
- **Execution flows** — processes from entry points through call chains

## The Problem

AI coding assistants (Cursor, Claude Code, Codex, etc.) don't truly know codebase structure:

```
1. AI edits UserService.validate()
2. Doesn't know 47 functions depend on its return type
3. Breaking changes ship 💥
```

## Traditional Graph RAG vs Smart Tools

### Traditional Graph RAG

```
User: What depends on UserService?
→ LLM receives raw graph
→ Query 1: Find callers
→ Query 2: What files?
→ Query 3: Filter tests?
→ Query 4: High-risk?
→ Answer after 4+ queries
```

### Knowledge Graph Approach (GitNexus)

```
User: What depends on UserService?
→ impact({target: "UserService", direction: "upstream"})
→ Pre-structured response: "8 callers, 3 clusters, all 90%+ confidence"
→ Complete answer in 1 query
```

## Core Innovation: Precomputed Relational Intelligence

Instead of giving the LLM raw graph edges and hoping it explores enough, the system **precomputes structure at index time**:

1. **Clustering** — Groups related symbols into functional communities
2. **Tracing** — Follows execution flows from entry points
3. **Scoring** — Assigns confidence to relationships

Benefits:
- **Reliability** — LLM can't miss context, it's already in the tool response
- **Token efficiency** — No 10-query chains to understand one function
- **Model democratization** — Smaller LLMs work because tools do the heavy lifting

## Indexing Pipeline

1. **Structure** — Walk file tree, map folder/file relationships
2. **Parsing** — Extract functions, classes, methods using Tree-sitter ASTs
3. **Resolution** — Resolve imports, calls, heritage, constructor inference
4. **Clustering** — Group symbols into communities (Leiden algorithm)
5. **Processes** — Trace execution flows from entry points
6. **Search** — Build hybrid search indexes (BM25 + semantic + RRF)

## Tools That Use This

- [[gitnexus]] — Main implementation
- [[context-mode]] — For code execution and search

## Related Concepts

- [[tree-sitter]] — AST parsing
- [[mcp-protocol]] — Tool exposure
- [[context-window-optimization]] — Why this matters