---
type: concept
title: "Performance Optimization Master Standard"
domain: tech-stack
created: 2026-05-13
updated: 2026-05-13
tags: [performance, profiling, optimization, staff-engineer, scalability]
status: mature
---

# Performance Optimization Master Standard

Performance optimization is the methodical practice of diagnosing and resolving system bottlenecks based on empirical data (profiling) rather than assumptions. In the Bunker, we prioritize efficiency to minimize token usage, latency, and infrastructure costs.

## 🏛️ The Three Axioms of Performance

1. **Measure Before You Optimize**: "Premature optimization is the root of all evil." Always identify the bottleneck using profilers before changing a single line of code.
2. **The 90/10 Rule**: 90% of execution time is typically spent in 10% of the code. Focus your energy only on that "Hot Path".
3. **Mandatory Trade-offs**: Every optimization is an exchange. You usually trade **Code Complexity** or **Memory** for **Execution Speed**.

## 🚀 Optimization Layers (The Order of Attack)

| Layer | Impact | Strategy |
|-------|--------|----------|
| **1. Algorithms** | Massive | Change O(n²) to O(n log n) or O(1). |
| **2. I/O & DB** | High | Fix N+1 queries, add missing indexes, use connection pooling. |
| **3. Caching** | Medium | Implement LRU, Redis, or local memoization for expensive results. |
| **4. Concurrency** | Medium | Use `errgroup` in Go or Worker Threads in Node.js for CPU-bound tasks. |

## 🛠️ Senior Troubleshooting Guide

### 1. The "N+1" Database Trap
**Problem**: Fetching 100 users, then making 100 separate queries to fetch their roles.
**Solution**: Use an `IN` clause or a `JOIN`.
```sql
-- BAD
SELECT * FROM roles WHERE user_id = 1; -- x100

-- GOOD
SELECT r.* FROM roles r JOIN users u ON r.user_id = u.id; -- x1
```

### 2. Go Profiling (pprof)
Use `net/http/pprof` to visualize CPU and Memory usage.
```go
import _ "net/http/pprof"
go func() {
    log.Println(http.ListenAndServe("localhost:6060", nil))
}()
```
Command to analyze: `go tool pprof -http=:8080 http://localhost:6060/debug/pprof/profile`

### 3. React Rendering Optimization
Avoid "Cascading Renders" by keeping state close to where it's used.
- **Pattern**: Use `useDeferredValue` for heavy UI updates.
- **Tool**: [[react-doctor]] to catch expensive `useEffect` synchronizations.

## 📦 Implementation in the Bunker

- **Batching**: Our `ingest_server.go` uses body limiters and could be expanded to support batch processing of multiple files.
- **Streaming**: Large file reads must use `io.Reader` (Go) or `fs.createReadStream` (Node) to maintain constant memory O(1).
- **Concurrency Control**: Use [[mcp-governance]] to prevent too many parallel tool calls from overwhelming the LLM's context.

## 🧪 Benchmarking Checklist

- [ ] Does the change reduce `duration_ms` in a reproducible way?
- [ ] Have you checked for memory leaks using a heap profiler?
- [ ] Is the code still readable/maintainable?
- [ ] Did you add a regression benchmark test?

## Related
- [[Go]]
- [[React]]
- [[PostgreSQL]]
- [[Redis]]
- [[react-health-score]]
- [[agentmemory]]
