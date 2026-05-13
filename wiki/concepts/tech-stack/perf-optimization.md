---
type: concept
title: "Performance Optimization (Standard)"
domain: performance
created: 2026-05-08
updated: 2026-05-08
tags: [optimization, profiling, bottlenecks, trade-offs, scalability]
status: active
related: [Bunker OS, AgentMemory, n8n-lab]
---

# Performance Optimization: The Methodical Standard

## Definition
Diagnóstico y resolución metódica de problemas de rendimiento basados en datos reales (profiling) y no en suposiciones.

## The Three Axioms of Performance
1. **Medir antes de optimizar:** La suposición es el enemigo. El cuello de botella rara vez está donde crees.
2. **La Regla 90/10:** El 90% del tiempo se gasta en el 10% del código. Optimiza ese 10%.
3. **Trade-offs Obligatorios:** Toda optimización intercambia algo (Velocidad vs Memoria, Throughput vs Latencia).

## Optimization Layers (Order of Attack)
1. **Algoritmos y Estructuras de Datos:** El mayor impacto (O(n²) -> O(n log n)).
2. **Base de Datos e I/O:** N+1 queries, falta de índices, carga masiva en RAM.
3. **Capa de Cache:** LRU, Redis, Cache-aside.
4. **Concurrencia y Paralelismo:** Async/Await (I/O) vs Multiprocessing (CPU).

## Diagnostic Tools by Stack
- **Go:** `pprof`, `go test -bench`.
- **Node.js:** `clinic.js`, `performance.now()`.
- **SQL:** `EXPLAIN ANALYZE`, `pg_stat_statements`.
- **Python:** `cProfile`, `py-spy`, `memory_profiler`.

## Patterns for the Bunker
- **Batching:** Reducir el overhead de llamadas individuales (ej: Ingest Server).
- **Streaming:** Procesar archivos por chunks para mantener memoria constante O(1).
- **Vectorization:** Uso de operaciones masivas en lugar de loops pesados.

---
*Este estándar rige para todos los proyectos del ecosistema SamBleed.*
