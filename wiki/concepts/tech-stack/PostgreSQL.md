# PostgreSQL Patterns (2026)

**Status**: Core Standard
**Source**: [[Context7]] (v18)
**Parent**: [[_index]]

## Core Features (v18+)
- **Asynchronous I/O**: Improved performance for sequential and bitmap heap scans.
- **UUID v7**: Built-in support for timestamp-ordered UUIDs (ideal for primary keys).
- **Temporal Constraints**: Support for primary, unique, and foreign keys with time ranges.

## Design & Performance
- **Indexing**: Use GIN indexes for `jsonb` fields. Use "Skip Scan" lookups for multicolumn B-tree indexes.
- **JSONB vs Relational**: Use `jsonb` for flexible schemas but enforce relational constraints for core business entities.
- **Connection Pooling**: Always use a pooler (like PgBouncer) for high-concurrency applications.

## Network & Connectivity
- **Keepalives**: Enable TCP keepalives (`keepalives=1`) to prevent connection drops.
- **Timeouts**: Always set `connect_timeout` to avoid hanging backends.

---
**Last Updated**: 2026-04-19
