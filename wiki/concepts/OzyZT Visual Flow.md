---
type: concept
title: "OzyZT Visual Flow"
created: 2026-04-22
updated: 2026-04-22
tags:
  - visualization
  - flow
  - architecture
status: finished
related:
  - "[[OzyZT Architecture]]"
---

# OzyZT Visual Flow

This diagram visualizes the end-to-end journey of a code scan in the [[OzyZT]] platform.

```mermaid
graph TB
    subgraph "Delivery Interfaces"
        CLI[Go CLI]
        API[FastAPI]
        WEB[Dashboard]
    end

    subgraph "OzyZT Engine (The Brain)"
        direction TB
        subgraph "Execution Layer"
            ORCH[Orchestrator]
            COV[Coverage Tracker]
        end

        subgraph "Verification Gates"
            G1[Source Code]
            G2[Secrets]
            G3[Dependencies]
            G4[IaC Security]
            G5[Container]
            G6[Signature]
            G7[Runtime]
        end

        subgraph "Policy & Logic"
            POL[Policy Engine]
            WAV[Waiver Manager]
            SCR[Scoring Engine]
        end
    end

    subgraph "External Ecosystem"
        SLK[Slack Webhook]
        AUD[OzyAudit]
        DB[(SQLite / Postgres)]
    end

    %% Flow
    CLI -->|Stream .tar.gz| API
    API --> ORCH
    ORCH --> G1 & G2 & G3 & G4 & G5 & G6 & G7
    G1 & G2 & G3 & G4 & G5 & G6 & G7 --> POL
    POL --> WAV
    WAV --> SCR
    SCR --> API
    API --> SLK
    API --> AUD
    API --> DB
    API -->|Live Progress| WEB

    %% Styling
    style ORCH fill:#f9f,stroke:#333,stroke-width:2px
    style POL fill:#bbf,stroke:#333,stroke-width:2px
    style G1,G2,G3,G4,G5,G6,G7 fill:#dfd,stroke:#333
    style SLK fill:#eb4,stroke:#333
    style AUD fill:#eb4,stroke:#333
```

## Functional Breakdown
1. **Ingestion**: The CLI packs code and sends it via HTTP streaming.
2. **Orchestration**: The engine decides the execution mode (Sequential/Parallel) and tracks tool availability.
3. **Evaluation**: Findings are filtered through the **Allowlist** and active **Waivers**.
4. **Conclusion**: A final **Security Score** is computed, and alerts are dispatched.
