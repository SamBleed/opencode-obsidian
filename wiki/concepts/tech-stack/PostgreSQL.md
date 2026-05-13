---
type: concept
title: "PostgreSQL"
domain: tech-stack
created: 2026-05-13
updated: 2026-05-13
tags: [database, sql, relational, storage]
status: mature
---

# PostgreSQL

PostgreSQL is the world's most advanced open-source relational database. In the Bunker's ecosystem, it serves as the primary source of truth for persistent data.

## 🚀 Advanced Patterns

### JSONB vs Relational
Postgres handles semi-structured data with excellence. Use `jsonb` for flexible schemas while maintaining performance via GIN indexes.

```sql
-- Creating a GIN index on a jsonb column
CREATE INDEX idx_user_metadata ON users USING GIN (metadata);

-- Querying inside jsonb
SELECT * FROM users WHERE metadata @> '{"role": "admin"}';
```

### Zero-Downtime Migrations
Changing schemas in high-traffic systems requires surgical precision.
1. **Adding a Column**: Always use `DEFAULT NULL` or add the default in a separate step.
2. **Adding an Index**: Use `CREATE INDEX CONCURRENTLY`.
3. **Renaming a Column**: Use a VIEW to support both old and new names.

### Partitioning & Sharding
For tables with millions of rows (e.g., `audit_logs`), use declarative partitioning. For extreme scale, consider **Citus** for sharding across multiple nodes.

## ⚙️ Maintenance & Operations

### Vacuuming & Bloat
Postgres uses MVCC (Multi-Version Concurrency Control), which leaves "dead" rows after updates/deletes.
- **Autovacuum**: Ensure it's tuned correctly to prevent table bloat and transaction ID wraparound.
- **VACUUM FULL**: Avoid in production as it locks the entire table. Use `pg_repack` instead.

### Performance Monitoring
- **EXPLAIN ANALYZE**: The ultimate tool for debugging slow queries. Look for `Seq Scan` (sequential scans) on large tables and replace them with indexes.
- **pg_stat_statements**: A module that tracks execution statistics of all SQL statements executed.

## 🛡️ Security & Hardening

- **Principle of Least Privilege**: Never use the `postgres` superuser for the app.
- **Row-Level Security (RLS)**: Enforce security at the database level.
  ```sql
  ALTER TABLE documents ENABLE ROW LEVEL SECURITY;
  CREATE POLICY user_docs ON documents USING (user_id = current_user_id());
  ```
- **SSL/TLS**: Force encrypted connections.

## 🛠️ Tooling
- **Migration Engine**: We prefer `atlas` or `golang-migrate` for Go projects.
- **ORM vs SQL**: Use `sqlc` (type-safe SQL) or `pgx` (native driver).

## 🧪 Testing with Postgres
Use **Testcontainers** or **Docker Compose** to run ephemeral instances. Avoid mocking; test against the real engine.

## Related
- [[GO-API-BLUEPRINT]]
- [[Data-Governance]]
- [[Trivy]]
- [[PostgreSQL-v-SQLite-v-Redis]]
