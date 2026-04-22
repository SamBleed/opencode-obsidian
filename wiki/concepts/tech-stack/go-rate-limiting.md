---
type: concept
title: "Go Rate Limiting"
domain: tech-stack
created: 2026-04-20
updated: 2026-04-20
tags:
  - go
  - rate-limiting
  - tech-stack
  - performance
status: seed
related:
  - "[[ozy-api]]"
  - "[[mcp-governance]]"
sources:
  - "https://github.com/webriots/rate"
---

# Go Rate Limiting

## Definition
Implementación de controles de flujo para proteger APIs en Go contra abusos y ataques DoS, optimizando el uso de recursos y garantizando la disponibilidad.

## How It Works
Existen varios algoritmos (Token Bucket, Leaky Bucket, Sliding Window). En 2026, librerías como `webriots/rate` destacan por su alta performance (1-15ns por operación) y uso de estructuras de datos compactas que evitan allocs en el heap.

## Example
### Implementación con `webriots/rate`
```go
// Definir un limiter: 1024 buckets, burst de 10, 100 req/sec
limiter, _ := rate.NewTokenBucketLimiter(1024, 10, 100, time.Second)

// Middleware check
func RateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        apiKey := []byte(r.Header.Get("X-API-KEY"))
        if limiter.TakeToken(apiKey) {
            next.ServeHTTP(w, r)
        } else {
            http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
        }
    })
}
```

## Patterns
- **Keyed Limiting**: Limitar por IP, API Key o User ID para evitar castigar a todos los usuarios por el abuso de uno solo.
- **Distributed Limiting**: Uso de Redis como backend para sincronizar límites en clusters de múltiples instancias (ruerate pattern).
- **Graceful Rejection**: Devolver siempre el header `Retry-After` para guiar al cliente.

## Source
- [[Exa Research]]
- https://github.com/webriots/rate
