---
type: concept
title: "Performance Optimization Standard"
domain: tech-stack
created: 2026-05-08
updated: 2026-05-08
tags: [performance, profiling, optimization, staff-engineer]
status: active
related: [Bunker OS, AgentMemory, n8n-lab]
---

# Performance Optimization Standard

## Definition
Estándar metódico para diagnosticar y resolver problemas de rendimiento basado en el principio: **Medir primero, optimizar después.**

## The Three Axioms
1. **Medir antes de optimizar:** La suposición es el enemigo.
2. **Regla 90/10:** El 90% del tiempo se gasta en el 10% del código.
3. **Trade-offs obligatorios:** Cada mejora tiene un costo (Memoria, Complejidad, Consistencia).

## Optimization Layers (Order of Impact)
1. **Algoritmos y Estructuras:** O(n²) -> O(n log n).
2. **DB e I/O:** N+1 queries, falta de índices, carga masiva en RAM.
3. **Cache:** LRU, Redis, Caching de resultados costosos.
4. **Concurrencia:** Async/Await para I/O, Multiprocessing para CPU.

## Implementation in the Bunker
- **n8n:** Monitoreo de `durationMs` en el Audit Log.
- **AgentMemory:** Uso de local embeddings optimizados.
- **Scripts:** Hardening con `set -euo pipefail` para evitar procesos zombie.

## Tools
- **Go:** `pprof`, `benchmarks`.
- **Node.js:** `clinic.js`, `performance.now()`.
- **SQL:** `EXPLAIN ANALYZE`.
