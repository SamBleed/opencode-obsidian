---
type: concept
title: "Agentic FinOps Patterns"
domain: tech-stack
created: 2026-05-13
updated: 2026-05-13
tags: [finops, ai, economics, budget, tokens, optimization]
status: mature
---

# Agentic FinOps Patterns (v2026)

Agentic FinOps is the practice of managing the economic lifecycle of autonomous AI agents. Since agents can trigger recursive call chains and expensive tool use, financial guardrails must be baked into the **Orchestration Layer**, not just the billing dashboard.

## 📊 Core Metric: CAPO
**Cost-per-Accepted-Outcome (CAPO)**: Instead of measuring raw token counts, we measure the total cost required to reach a verified result.

## 🚀 Optimization Strategies

### 1. Model Routing & Cascading
Don't use a frontier model (GPT-4/Sonnet) for everything.
- **Tier 1 (Flash/Haiku)**: Routing, summarization, simple extraction.
- **Tier 2 (Pro/Sonnet)**: Planning, complex reasoning, code generation.
- **Pattern**: If Tier 1 fails a confidence check, escalate to Tier 2.

### 2. Logic Density (Token Efficiency)
- **Prompt Compression**: Distill system instructions into dense "Policy-as-Code."
- **Prefix Caching**: Keep common context (e.g., Project Index) in the cached prefix to reduce input token costs by 50-80%.
- **Conciseness Enforcement**: Force agents to respond in structured JSON or brief markdown to minimize output token spend.

### 3. Temporal Arbitrage (Spot Inference)
- **Instant Inference**: Premium price for real-time user requests.
- **Batch/Spot Inference**: Defer non-urgent tasks (e.g., weekly wiki linting) to 24-hour batch windows for up to 90% discount.

## 🛡️ Budget-as-Code Guardrails

Implement hard limits in the orchestration code to prevent "loop explosions."

```json
{
  "agent_policy": {
    "max_steps_per_task": 15,
    "max_token_budget_per_run": 50000,
    "high_cost_tool_gate": 5.0,
    "approval_required": "budget_exceeded"
  }
}
```

## 🏗️ The 3-Agent FinOps Architecture

| Agent Type | Responsibility |
|------------|----------------|
| **Forecasting Agent** | Predicts spend spikes based on queue depth and throughput. |
| **Policy Agent** | Applies circuit breakers and budget checks before execution. |
| **Execution Agent** | Rightsizes GPU instances and moves workloads to Spot capacity. |

## ⚠️ Known Spend Drivers
- **Recursive Loops**: Agents getting stuck on a failing test and retrying 50 times.
- **Context Bloat**: Carrying 200,000 tokens of history for a 5-token answer.
- **Shadow Agents**: Background processes that were never terminated.

## Related
- [[Agentic-Workflows]]
- [[AI-Analysis-Workflow]]
- [[Observability-MD]]
- [[AGENT-OBSERVABILITY-BLUEPRINT]]
