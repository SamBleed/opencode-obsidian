---
type: concept
title: "Vector Database Patterns"
domain: tech-stack
created: 2026-05-13
updated: 2026-05-13
tags: [ai, vector-db, pgvector, hybrid-search, rag]
status: mature
---

# Vector Database Patterns (v2026)

Vector databases are the foundation of modern Retrieval-Augmented Generation (RAG). In the Bunker, we prioritize **Hybrid Search** and **PostgreSQL-native** vector storage using `pgvector` to ensure ACID compliance and operational simplicity.

## 🚀 The 2026 Gold Standard: Hybrid Search (RRF)

Semantic search (dense vectors) is great for concepts but terrible for rare tokens (e.g., "A-700-X"). The solution is **Reciprocal Rank Fusion (RRF)**, which combines Vector Search with Full-Text Search (BM25).

### SQL Implementation (Postgres + pgvector)

```sql
WITH semantic_search AS (
    SELECT id, RANK () OVER (ORDER BY embedding <=> $1) AS rank
    FROM documents
    ORDER BY embedding <=> $1
    LIMIT 20
),
keyword_search AS (
    SELECT id, RANK () OVER (ORDER BY ts_rank_cd(to_tsvector('english', content), query) DESC)
    FROM documents, plainto_tsquery('english', $2) query
    WHERE to_tsvector('english', content) @@ query
    ORDER BY ts_rank_cd(to_tsvector('english', content), query) DESC
    LIMIT 20
)
SELECT
    COALESCE(semantic_search.id, keyword_search.id) AS id,
    COALESCE(1.0 / (60 + semantic_search.rank), 0.0) + 
    COALESCE(1.0 / (60 + keyword_search.rank), 0.0) AS score
FROM semantic_search
FULL OUTER JOIN keyword_search ON semantic_search.id = keyword_search.id
ORDER BY score DESC
LIMIT 10;
```

## 📐 Data Modeling & Indexing

### 1. HNSW vs IVFFlat
- **HNSW (Hierarchical Navigable Small World)**: The default for production. Higher recall and faster queries at the cost of memory and slower index builds.
  - *Tuning*: Increase `M` (connections) and `ef_construction` for datasets > 1M rows.
- **IVFFlat**: Faster to build, less memory, but requires a warm-up dataset to define clusters (Voronoi cells).

### 2. Memory Optimization with `halfvec`
As of `pgvector` 0.7+, use `halfvec` (16-bit floats) to cut memory usage by 50% with negligible loss in accuracy.

```sql
ALTER TABLE items ADD COLUMN embedding_half halfvec(1536);
CREATE INDEX ON items USING hnsw (embedding_half halfvec_l2_ops);
```

## 🛠️ Go Implementation (pgx)

Use `github.com/pgvector/pgvector-go` to handle vector types in Go.

```go
config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
    return pgxvec.RegisterTypes(ctx, conn)
}
// Query
rows, _ := pool.Query(ctx, "SELECT id FROM items ORDER BY embedding <-> $1 LIMIT 5", pgvector.NewVector(queryVec))
```

## ⚖️ Trade-offs: Dedicated vs Relational
- **Dedicated (Pinecone, Weaviate)**: Better for 100M+ vectors, sub-10ms p99 latency, and managed scaling.
- **Relational (Postgres + pgvector)**: Perfect for <10M vectors. Single source of truth, atomic deletes, and zero extra infrastructure.

## Related
- [[PostgreSQL]]
- [[LLM Wiki Pattern]]
- [[Agentic-Workflows]]
- [[performance-optimization]]
