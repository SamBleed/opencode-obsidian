# Blueprint: Go API (2026)

**Target**: Backend Microservices
**Stack**: Go 1.24+, PostgreSQL 18, Redis 7

## Directory Structure
```text
/
├── cmd/
│   └── api/            # Entry point
├── internal/
│   ├── api/            # Handlers & Routes
│   ├── service/        # Business Logic
│   ├── store/          # DB Access (Repository pattern)
│   └── model/          # Entities
├── pkg/                # Reusable utilities
├── scripts/            # Build/Dev tools
├── Dockerfile          # Multi-stage build
└── docker-compose.yml
```

## Mandatory Checks
- [ ] Structured logging initialized.
- [ ] Graceful shutdown implemented.
- [ ] Health check endpoint (`/health`).
- [ ] Migrations for SQL.

---
**Template Type**: Project Blueprint
