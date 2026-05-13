---
type: project
title: "OzyZT"
domain: security
created: 2026-05-13
updated: 2026-05-13
tags: [security, zero-trust, platform, go, nextjs]
status: active
---

# OzyZT (Zero-Trust Platform)

OzyZT is a modular Zero-Trust security platform designed to enforce security gates across the development lifecycle. It evolved from a simple security toolkit into a full-scale orchestration engine.

## 🏗️ Architecture

OzyZT follows a **Modular Security Engine** approach, where different "adapters" can be plugged in to provide specific security capabilities.

### Core Components
1. **OzyAPI (Go)**: The central brain that orchestrates scans, manages policies, and stores audit trails.
2. **OzyWeb (Next.js)**: The administrative dashboard for visualizing security posture and managing gates.
3. **OzyRecon**: An engine for automated asset discovery and OSINT.
4. **OzyAudit**: Integration with tools like [[Trivy]] and [[react-doctor]] for static and dynamic analysis.

## 🛡️ Zero-Trust Enforcement Gates

| Gate | Stage | Action |
|------|-------|--------|
| **Commit Gate** | Pre-commit | Runs local linting and secret detection. |
| **PR Gate** | CI/CD | Runs [[react-health-score]] and vulnerability scans. |
| **Deploy Gate** | Pre-deployment | Verifies image signatures and production config. |
| **Runtime Gate** | Post-deployment | Continuous monitoring and automated drift detection. |

## 🛠️ Tech Stack
- **Backend**: Go 1.26, PostgreSQL, Redis.
- **Frontend**: React 19, Tailwind CSS 4, Lucide Icons.
- **Communication**: gRPC for internal engines, REST/JSON for external consumers.
- **Automation**: [[n8n]] for alert routing and third-party integrations.

## 🚀 Current Roadmap
- ✅ **Phase 1**: Core API and Dashboard MVP.
- ✅ **Phase 2**: [[Trivy]] integration and GitHub Actions support.
- 🔲 **Phase 3**: Multi-tenant support and Slack/Discord alerting.
- 🔲 **Phase 4**: Automated remediation (Self-healing security).

## 🧪 Setup & Development

```bash
# Clone the monorepo
git clone https://github.com/ozy-platform/ozy-zt
cd ozy-zt

# Setup local infrastructure
docker-compose up -d db redis

# Start the API
go run ./cmd/api
```

## Related
- [[OzyZT Architecture]]
- [[OzyZT Visual Flow]]
- [[GO-API-BLUEPRINT]]
- [[DESIGN-TEMPLATE]]
- [[n8n]]
