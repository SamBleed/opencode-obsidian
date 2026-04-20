# OZY-API

**Type**: Project / Backend
**Status**: Active / Populated
**Stack**: Go 1.26, PostgreSQL, Redis, Docker
**Blueprint**: [[GO-API-BLUEPRINT]]

## 📝 Description
The core backend for the Bunker's ecosystem. Implements hexagonal architecture and secure identity management.

## 🚀 Key Features
- **JWT Auth**: Secure access via [[Go-JWT]].
- **Structured Logging**: AI-readable JSON logs via [[Observability-MD]].
- **SQL Data Layer**: Powered by Postgres and [[Data-Governance]].
- **CI/CD Ready**: Designed for automated [[CI-CD-Pipeline]] integration (Trivy scan + GitHub Actions).

## 🛠️ Commands
```bash
cd projects/ozy-api
docker compose up -d
go run cmd/api/main.go
```

---
**Last Updated**: 2026-04-19
