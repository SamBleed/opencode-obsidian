---
type: concept
title: "Redis"
domain: tech-stack
created: 2026-05-13
updated: 2026-05-13
tags: [cache, database, nosql, performance]
status: mature
---

# Redis

Redis is an in-memory data structure store used as a database, cache, message broker, and streaming engine.

## 🚀 Key Use Cases

### 1. Caching Strategies
- **Cache Aside**: Application checks Redis first, then DB, then updates Redis.
- **Write-Through**: Application writes to Redis, which then updates the DB.
- **TTL (Time To Live)**: Always set expiration for cache keys to avoid stale data and memory bloat.

### 2. Distributed Locking (Redlock)
Crucial for preventing race conditions in distributed systems (e.g., ensuring a job only runs once).

```go
// Using redsync in Go
mutex := rs.NewMutex("lock-key")
if err := mutex.Lock(); err != nil {
    return // Locked by another instance
}
defer mutex.Unlock()
```

### 3. Messaging (Pub/Sub vs Streams)
- **Pub/Sub**: Fire-and-forget. If the subscriber is down, the message is lost.
- **Streams**: Persistent message history. Supports consumer groups (similar to Kafka) for reliable processing.

## 🛠️ Data Structures
- **Strings**: Simple key-value.
- **Hashes**: Storing objects (e.g., user sessions).
- **Lists**: Queues.
- **Sets**: Unique collections (e.g., unique visitors).
- **Sorted Sets (ZSET)**: Leaderboards and rate-limiting.

## ⚠️ Anti-patterns
- **Large Keys**: Storing multi-MB strings blocks the single-threaded event loop.
- **Keys Pattern**: Never use `KEYS *` in production; use `SCAN` instead.
- **No Eviction Policy**: Configure `maxmemory-policy` (e.g., `allkeys-lru`) to prevent OOM.

## Related
- [[go-rate-limiting]] (Often uses Redis)
- [[GO-API-BLUEPRINT]]
- [[PostgreSQL-v-SQLite-v-Redis]]
