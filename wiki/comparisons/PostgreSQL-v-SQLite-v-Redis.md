---
type: comparison
title: "PostgreSQL v SQLite v Redis"
updated: 2026-05-14
status: active
---

# PostgreSQL vs SQLite vs Redis

> ¿Cuándo usar cada base de datos?

## Comparativa Rápida

| Criterio | PostgreSQL | SQLite | Redis |
|----------|------------|--------|-------|
| **Tipo** | Relational | Embedded DB | In-memory K-V |
| **Size** | 100MB+ | < 1MB | Variable |
| **Concurrency** | ✅ Multi-process | ❌ File lock | ✅ Event loop |
| **Joins** | ✅ Full | ✅ Full | ❌ No |
| **Persistence** | Disk | Disk | Optional (RDB/AOF) |
| **Speed** | 1-10ms | 0.1-1ms | < 1ms |

## Cuándo PostgreSQL

- **Datos estructurados** — usuarios, orders, inventory
- **Queries complejos** — joins, agregaciones, full-text search
- **Concurrent access** — múltiples clientes simultáneos
- **Migrations** — schema versioning
- **Producción** — high availability, replication

```sql
-- Ejemplo: PostgreSQL shine
SELECT u.name, COUNT(o.id) as orders
FROM users u
LEFT JOIN orders o ON u.id = o.user_id
WHERE o.created_at > NOW() - INTERVAL '30 days'
GROUP BY u.id
HAVING COUNT(o.id) > 5;
```

## Cuándo SQLite

- **Prototyping** — velocidad de desarrollo
- **Local apps** — desktop, mobile
- **Embedded devices** — IoT, sensors
- **Read-heavy** — blogs, CMS simples
- **Testing** — fixtures, e2e tests

```bash
# Ejemplo: shipped with app
./my-app --db=./data.db
```

## Cuándo Redis

- **Caching** — sessions, API responses
- **Rate limiting** — counters at scale
- **Pub/Sub** — real-time notifications
- **Leaderboards** — sorted sets
- **Session store** — fast user lookup

```javascript
// Ejemplo: Redis shine
await redis.setex(`session:${userId}`, 3600, JSON.stringify(session));
await redis.incr(`rate-limit:${ip}`);
```

## La Arquitectura Común

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Client    │────▶│   Redis     │────▶│  PostgreSQL │
│             │     │  (cache)    │     │  (source)  │
└─────────────┘     └─────────────┘     └─────────────┘
```

1. **Redis** first → si está en cache, retornalo
2. **PostgreSQL** → si no, consultá y guardá en cache
3. **Write** → siempre a PostgreSQL, invalidá cache

## Regla de Thumb

```
¿Relaciones entre datos? → PostgreSQL
¿Solo clave-valor? → Redis
¿Aplicación standalone/simple? → SQLite
```

## En Este Proyecto

- **PostgreSQL** → datos persistentes (usuarios, proyectos, ADRs)
- **Redis** → cache de Engram, rate limiting
- **SQLite** → testing local

> 💡 Empezá con PostgreSQL. Agregá Redis solo cuando tengas problema de perf.