---
type: concept
domain: ai-development
title: "Context Window Optimization"
description: "Strategies to maximize the effective use of LLM context windows by reducing noise and processing data externally"
created: 2026-04-22
updated: 2026-04-22
tags: [context-window, llm, optimization, tokens]
status: seed
related: [sandboxed-execution, fts5-search]
sources: [context-mode]
---

# Context Window Optimization

## Definition

Strategies and techniques to maximize the effective use of LLM context windows. Instead of filling context with raw data (logs, file contents, search results), the model processes data externally and receives only the relevant answers.

## The Problem

When working with LLMs, tools often dump raw data into the context:
- Full file contents when only a count is needed
- Complete command outputs when only a summary is required
- Search results with all metadata when only matches matter

This consumes tokens rapidly and reduces the model's capacity to reason.

## Core Principle

> **"Think in Code — MANDATORY"**: When you need to analyze, count, filter, search, parse or process data, **write code** that does the work and return only the answer.

## Techniques

### 1. Sandboxed Code Execution

Execute code in an isolated environment and return only the result. See [[sandboxed-execution]].

### 2. Full-Text Search (FTS5)

Use SQLite FTS5 for fast searching without loading all data into context. See [[fts5-search]].

### 3. Intent-Driven Search

Design search tools that return processed results, not raw data.

### 4. Batch Processing

Execute multiple commands in one call, auto-index results, then search the index.

### 5. Hot Cache

Keep critical context in a lightweight cache (like `hot.md`) that survives between sessions.

## Example Transformations

| Instead of... | Do this... |
|--------------|------------|
| Read 5000 lines of log | Execute: `grep -c ERROR huge.log` → return only the count |
| Read entire JSON file | Execute: `JSON.parse().filter(...)` → return matches |
| Dump search results | Return: "Found 47 matches, top 5 are..." |

## Benefits

- **Saves up to 98% of context window** in certain scenarios
- **Faster responses** — less data to process
- **Better reasoning** — model focuses on analysis, not data scanning

## Related Concepts

- [[sandboxed-execution]]
- [[fts5-search]]
- [[mcp-protocol]]

## Tools

- [[context-mode]] — MCP plugin implementing these techniques