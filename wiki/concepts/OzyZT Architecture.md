---
type: concept
title: "OzyZT Architecture"
domain: security
created: 2026-05-13
updated: 2026-05-13
tags: [security, zero-trust, architecture, hexagonal]
status: mature
---

# OzyZT Architecture

OzyZT is designed as a **Security Control Plane** that enforces Zero-Trust principles across the entire software development lifecycle (SDLC). It leverages a hexagonal architecture to remain agnostic to specific scanning tools while providing a unified governance layer.

## 🏗️ High-Level Design

### 1. The Core Engine (Go)
The central orchestrator that manages:
- **Identity Context**: Maps developers, agents, and service accounts to specific permissions.
- **Policy Engine**: Executes [[Policy-as-Code]] (OPA) to determine if a gate should be opened.
- **Audit Ledger**: A tamper-proof log of all security decisions and findings.

### 2. Security Gates (The Interceptors)
Interceptors are placed at critical transition points in the pipeline:
- **Pre-commit Interceptor**: Prevents secrets from being committed local-first.
- **CI/CD Interceptor**: Injects [[Trivy]] and [[react-doctor]] scans. Generates [[Supply-Chain-Security]] attestations (SLSA).
- **Deployment Interceptor**: Verifies image signatures and SBOMs before allowing a K8s deployment.

## 🚀 Advanced Capabilities (2026)

### Agentic AppSec Integration
OzyZT supports "Agentic Remediation". When a vulnerability is found:
1. An autonomous agent is spawned to verify the exploitability.
2. If verified, the agent generates a PR with the fix.
3. OzyZT marks the finding as "Verified & Remediation Pending".

### Drift Detection & Runtime Correlation
OzyZT doesn't stop at deployment. It continuously monitors the runtime (K8s/Cloud) to ensure:
- **No manual changes**: If someone modifies a security group manually, OzyZT alerts and triggers a GitOps resync.
- **Runtime Vulnerability Correlation**: If a new CVE is discovered for a library already in production, OzyZT calculates the "Blast Radius" based on actual traffic flows.

## 🛡️ Data Flow

```text
[Dev/Agent] ──▶ [Git Commit] ──▶ [OzyZT Commit Gate]
                                       │
                                       ▼
[Build Artifact] ◀── [OzyZT CI Gate] ◀── [Trivy/Linter]
       │
       ▼
[Attestation Signed] ──▶ [Deployment Gate] ──▶ [Runtime]
                                                  │
                                                  ▼
                                         [Drift Detection]
```

## 🛠️ Stack Components
- **Language**: Go 1.26 (for performance and concurrency).
- **Database**: PostgreSQL (Relational audit trails) + Redis (Real-time gate state).
- **Orchestration**: [[n8n]] for complex alert flows and external integrations.
- **Policies**: Rego (OPA).

## Related
- [[OzyZT]] (Project Page)
- [[Policy-as-Code]]
- [[Supply-Chain-Security]]
- [[Trivy]]
- [[Security-Guardrails]]
