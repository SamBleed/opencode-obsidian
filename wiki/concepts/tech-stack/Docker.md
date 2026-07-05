---
type: concept
title: "Docker & Containerization"
domain: tech-stack
created: 2026-05-13
updated: 2026-05-13
tags: [devops, containers, docker, infrastructure]
status: mature
---

# Docker & Containerization

Docker revolutionized software delivery by encapsulating applications and their dependencies into immutable artifacts.

## 🏗️ Production-Grade Dockerfiles

### Multi-Stage Builds
Crucial for keeping images small and secure. The build environment should never be in the final image.

```dockerfile
# Stage 1: Build
FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Stage 2: Final
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
# Non-root user for security
RUN adduser -D appuser
USER appuser
CMD ["./main"]
```

### Security Hardening
1. **Rootless Execution**: Never run your app as `root` inside the container.
2. **Read-Only Root FS**: Use `--read-only` flag at runtime.
3. **Scan Regularly**: Use [[Trivy]] to find vulnerabilities.
4. **Base Image Selection**: Favor `alpine` or `distroless`.
5. **Registry Security**: Use private registries (GHCR, ECR) and enable image signing (Cosign).

## 🌐 Kubernetes Context
While Docker is the runtime, Kubernetes (K8s) is the orchestrator.
- **Liveness/Readiness Probes**: Crucial for self-healing.
- **Resource Limits**: Always set `cpu` and `memory` limits to prevent a single container from starving the host.
- **ConfigMaps & Secrets**: Don't bake configuration into the image. Inject it at runtime.

## 🛠️ Docker Compose for Development
We use `docker-compose.yml` to orchestrate local environments.

```yaml
services:
  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DB_URL=postgres://user:pass@db:5432/db
    depends_on:
      db:
        condition: service_healthy

  db:
    image: postgres:17-alpine
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U postgres"]
      interval: 5s
      timeout: 5s
      retries: 5
```

## 🔄 Lifecycle Management
- **Pruning**: Regularly run `docker system prune`.
- **Logs**: Use `json-file` driver with `max-size` and `max-file` limits to prevent disk exhaustion.

## Related
- [[docker-patterns]]
- [[deployment-patterns]]
- [[Trivy]]
- [[OzyZT Architecture]]
