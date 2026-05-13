---
type: project
title: "OzyAudit"
domain: security
stack: "Node.js/TS"
status: active
created: 2026-05-08
updated: 2026-05-13
tags: [security, auditing, sast, dast, automated-review]
---

# OzyAudit

OzyAudit is the dedicated security auditing engine of the [[OzyZT]] platform. Its primary mission is to perform deep static and dynamic analysis on codebases and containers to identify vulnerabilities before they reach production.

## 🚀 Key Integration Modules

### 1. React Health Auditor
Integrates directly with [[react-doctor]] to scan frontend repositories.
- **Goal**: Maintain a [[react-health-score]] above 80.
- **Automation**: Runs as a CI check and reports findings back to the OzyZT Dashboard.

### 2. Infrastructure Auditor
Uses [[Trivy]] to scan Dockerfiles, Terraform manifests, and Kubernetes YAMLs.
- **Scope**: Misconfigurations (e.g., running as root, open S3 buckets) and known CVEs in base images.

### 3. Agentic Reviewer
A specialized agentic workflow that reads the results from static analysis and:
1. Filters out false positives.
2. Explains the security impact of a finding.
3. Suggests a fix following [[Security-Guardrails]].

## 🏗️ Technical Workflow

```text
[Source Code] ──▶ [OzyAudit Trigger]
                        │
         ┌──────────────┴──────────────┐
         ▼                             ▼
  [Static Scan]                 [Dependency Scan]
 (react-doctor)                  (Trivy/Snyk)
         │                             │
         └──────────────┬──────────────┘
                        ▼
            [Agentic Reconciliation]
            (LLM-based verification)
                        │
                        ▼
             [Final Audit Report] ──▶ [[OzyZT]] Dashboard
```

## 🧪 Operational Status
- **Current Stack**: Node.js 22, TypeScript 5, Docker.
- **Path**: `/home/sam/Proyectos/OzyAudit`
- **Audit Ledger**: All findings are persisted in the Relational Audit Trail managed by [[OzyZT Architecture]].

## Related
- [[OzyZT]]
- [[Trivy]]
- [[react-doctor]]
- [[react-health-score]]
- [[Security-Guardrails]]
