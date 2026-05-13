---
type: blueprint
title: "Go API Blueprint"
domain: tech-stack
created: 2026-05-13
updated: 2026-05-13
tags: [go, api, blueprint, architecture, backend]
status: mature
---

# Go API Blueprint (v2026)

This blueprint defines the standard structure and conventions for all Go-based backend services in the Bunker. It prioritizes maintainability, type safety, and structured logging.

## 📂 Directory Structure

```text
.
├── cmd/
│   └── api/                # Entry point
├── internal/
│   ├── domain/             # Entities and interfaces (Hexagonal Core)
│   ├── service/            # Business logic
│   ├── repository/         # DB implementation (Postgres/Redis)
│   ├── handler/            # HTTP/gRPC handlers
│   └── config/             # Typed configuration
├── pkg/                    # Shared utilities (safe for external use)
├── api/                    # OpenAPI/Proto definitions
├── scripts/                # Task scripts (bash/go)
├── deployments/            # Docker/K8s manifests
├── Makefile                # Task runner
└── go.mod
```

## 🏗️ Design Principles

1. **Dependency Injection**: Use constructors to inject dependencies. No global variables for DB/Logger.
2. **Context Propagation**: Every function that performs I/O must accept `context.Context`.
3. **Structured Errors**: Wrap errors with context: `fmt.Errorf("service: %w", err)`.
4. **Middleware-First**: Handle Auth, Logging, and Rate Limiting at the middleware layer.

## 🛠️ Mandatory Stack

- **Router**: `chi` or `echo` (v5+).
- **Driver**: `pgx/v5`.
- **Logging**: `slog` (Standard library).
- **Validation**: `validator/v10`.
- **Documentation**: Swagger/OpenAPI 3.1.

## 🧪 Testing Checklist

- [ ] **Unit Tests**: Domain logic must have 90%+ coverage.
- [ ] **Integration Tests**: Handlers tested against a real Postgres container (Testcontainers).
- [ ] **Benchmarking**: Critical paths (e.g., JSON parsing) must have benchmarks.
- [ ] **Mocks**: Use `moq` or `mockery` for interface-based mocking.

## 🚀 Deployment Pattern

- **Image**: Multi-stage Dockerfile (Alpine/Distroless).
- **Config**: Environment variables with `.env` support for local dev.
- **Health**: `/healthz` and `/readyz` endpoints for K8s.

## Related
- [[Go]]
- [[PostgreSQL]]
- [[deployment-patterns]]
- [[OzyZT Architecture]]
