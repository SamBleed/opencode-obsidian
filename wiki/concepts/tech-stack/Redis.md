# Redis Patterns (2026)

**Status**: Core Standard
**Source**: [[Context7]] (v7+)
**Parent**: [[_index]]

## Caching Strategies
- **Cache-Aside**: The most common pattern. Check cache -> Fetch from DB -> Set cache.
- **Invalidation**: Always delete/update the cache key when the underlying DB data changes.

## Persistence Options
- **RDB (Redis Database)**: Point-in-time snapshots. Good for backups, minimal performance impact.
- **AOF (Append Only File)**: Logs every write. High durability, but can lead to larger files.
- **Hybrid**: Recommended. Use both for maximum data safety.

## Data Structures
Use the right tool for the job:
- **Strings**: Simple key-value.
- **Hashes**: Objects/Fields.
- **Lists/Sets**: Queues and unique collections.

## Memory Management
- Set an explicit `maxmemory` policy (e.g., `allkeys-lru`) for caching-only instances.

---
**Last Updated**: 2026-04-19
