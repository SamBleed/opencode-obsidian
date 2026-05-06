---
type: project
title: "OzyZT - Zero-Trust Enforcement Platform"
created: 2026-04-22
updated: 2026-04-22
tags:
  - security
  - cybersecurity
  - zero-trust
  - architecture
  - backend
  - cli
status: active
related:
  - "[[OzyZT Architecture]]"
  - "[[OzyAudit]]"
---

# OzyZT

**OzyZT** (formerly Ozy Zero Trust) is a high-performance, real-time security scanning platform with a [[Zero-Trust]] architecture. It validates every stage of the [[SDLC]] (Software Development Life Cycle) through a system of 7 specialized [[Security Gates]].

## Core Mission
To provide a modular, extensible, and secure enforcement layer that prevents vulnerabilities, secrets, and insecure configurations from reaching production environments.

## Key Features
- **7 Security Gates**: Source Code, Secrets, Dependencies, IaC, Container, Signature, and Runtime.
- **Provider Pattern**: Decoupled execution where gates use multiple interchangeable tools (Trivy, Gitleaks, Semgrep, etc.).
- **Policy Packs**: Built-in configurations for different environments (Default, Strict, Reduced Coverage).
- **Normalization Layer**: Unified finding format for cross-tool reporting and [[OzyAudit]] integration.
- **Real-time Monitoring**: WebSocket streaming for live scan progress.
- **Notification System**: Built-in support for [[Slack]] alerts and centralized auditing.

## Architecture
The system is divided into a specialized **Engine** and a thin **CLI/API** layer:
1. **Engine**: The core logic living in `engine/`.
   - `core/`: Protocols and scoring models.
   - `gates/`: Individual gate implementations.
   - `policy/`: Rule evaluation and waivers.
   - `normalization/`: Data translation.
2. **CLI (Go)**: Lightweight client that streams code to the engine.
3. **API (FastAPI)**: REST interface for orchestration and integration.
4. **Dashboard (React)**: Visual management of scans and trends.

## Integration
OzyZT is designed to be the "enforcement arm" of a larger security ecosystem, pushing results to [[OzyAudit]] for long-term storage and compliance analysis.

## Development History
- **2026-04-22**: Major architectural overhaul. Transformed from a collection of scripts into a modular platform. Implemented providers, policy packs, waivers, and reinforced security against Tar Bomb attacks.

---
(Source: [[Session OzyZT Transformation]])
