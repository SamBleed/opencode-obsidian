---
type: concept
domain: tech-stack
title: "FTS5 - Full-Text Search 5"
description: "SQLite's high-performance full-text search engine for building knowledge bases"
created: 2026-04-22
updated: 2026-04-22
tags: [sqlite, fts5, full-text-search, database, knowledge-base]
status: seed
related: [context-window-optimization, sandboxed-execution]
sources: [context-mode]
---

# FTS5 - Full-Text Search 5

## Definition

**FTS5** is SQLite's high-performance full-text search engine. It's a virtual table module that provides fast text searching capabilities without loading entire datasets into memory or context.

## Why FTS5?

In LLM-assisted development, searching through large knowledge bases or documentation can flood the context window with results. FTS5 solves this by:

1. **Indexing** content in a specialized format
2. **Searching** the index directly
3. **Returning only relevant results** — not raw data

## Basic Usage

```sql
-- Create an FTS5 virtual table
CREATE VIRTUAL TABLE documents USING fts5(
  title,
  content,
  tokenize='porter'
);

-- Insert data
INSERT INTO documents(title, content) VALUES 
  ('Getting Started', 'This guide helps you install and configure...'),
  ('API Reference', 'The API provides endpoints for...');

-- Search with relevance ranking
SELECT title, snippet(documents, 1, '<mark>', '</mark>', '...', 32) AS highlight
FROM documents
WHERE documents MATCH 'install OR configure'
ORDER BY rank;
```

## Key Features

| Feature | Description |
|---------|-------------|
| **Porter Stemming** | Tokenizer that reduces words to root form (running → run) |
| **Relevance Ranking** | BM25 algorithm for sorting results by relevance |
| **Snippets** | Return context around matches with highlights |
| **Fast Prefix Search** | `install*` matches "installing", "installed", etc. |
| **Low Memory** | Indexes stay on disk, not in RAM |

## FTS5 in context-mode

context-mode uses FTS5 to build a knowledge base that can be searched without loading raw data into context:

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│   Raw Data      │────▶│   FTS5 Index    │────▶│  Search Query   │
│ (docs, logs,    │     │ (SQLite file)   │     │ (returns only   │
│  wikis, etc.)   │     │                 │     │   matches)      │
└─────────────────┘     └─────────────────┘     └─────────────────┘
```

## Comparison: FTS5 vs LIKE

| Aspect | `LIKE '%query%'` | FTS5 |
|--------|------------------|------|
| Performance | O(n) full scan | O(log n) index |
| Memory | Loads all data | Index on disk |
| Stemming | No | Yes (porter) |
| Ranking | None | BM25 relevance |
| Prefix search | Slow | Fast (`term*`) |

## Practical Example

```javascript
// Instead of reading 1000 wiki pages into context:
// 1. Index them into FTS5
const documents = getAllWikiPages();
db.exec(`CREATE VIRTUAL TABLE wiki USING fts5(title, content)`);
for (const doc of documents) {
  db.prepare(`INSERT INTO wiki(title, content) VALUES (?, ?)`).run(doc.title, doc.content);
}

// 2. Search the index (returns only matches)
const results = db.prepare(`
  SELECT title, snippet(wiki, 1, '**', '**', '...', 64) AS context
  FROM wiki
  WHERE wiki MATCH ?
  ORDER BY rank
  LIMIT 5
`).all('authentication OR login');

// 3. Return only results to context
// "Found 3 matches: title1 [context...], title2 [context...]"
```

## Related Concepts

- [[context-window-optimization]]
- [[sandboxed-execution]]
- [[mcp-protocol]]

## Tools Using This

- [[context-mode]] — Uses FTS5 for knowledge base searching