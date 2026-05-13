---
type: concept
title: "ADR Lifecycle & Decision Governance"
domain: governance
created: 2026-05-13
updated: 2026-05-13
tags: [governance, adr, documentation, process, senior-architect]
status: mature
---

# ADR Lifecycle & Decision Governance (v2026)

An **Architecture Decision Record (ADR)** is a document that captures an important architectural decision made along with its context and consequences. In the Bunker, ADRs are not ceremonial; they are active filters for both humans and agents.

## 🔄 The AgDR (Agent Decision Record) Extension

For 2026, we extend ADRs to support **Agentic Autonomy**. When an agent makes a significant choice (e.g., selecting a persistence library), it must document it as an **AgDR**.

### The AgDR Template
- **Context**: The situation prompting the decision.
- **Y-Statement**: "In the context of [situation], we decided to [option], resulting in [consequence], to achieve [goal]."
- **Agent Metadata**: Model version, session ID, and prompt trigger.
- **Human Approval**: Signature of the "Accountable Operator."

## 🚦 ADR Triggers (When to write one?)

CI/CD pipelines should block PRs that cross these boundaries without an ADR:
1. **New Dependency**: Adding a third-party library to the core stack.
2. **Contract Change**: Altering a public API or event schema.
3. **Persistence Strategy**: Changing how data is stored or cached.
4. **State Machine Logic**: Modifying agentic control flows.

## 🛡️ Governance Principles

### 1. One Agent, One Human
Every autonomous process is bound to a single identifiable human who is accountable for its decisions. If no human is available, the agent must halt at high-impact decision gates.

### 2. Causal Traceability
Every architectural change must be reconstructible. An incident in production should lead directly to the ADR that authorized the pattern.

### 3. Decisions as Context
ADRs are fed back to agents as **First-Class Context**. The same document that disciplines humans constrains the model's choices.

## 🚑 Incident Post-mortems (The Detection Chain)

When something fails, we perform a post-mortem focusing on the **MTTD-for-Agents** (Mean Time To Detection).

### Retro Sections for Agents
- **Context Engineering**: Was the agent given the right data to succeed?
- **Tool Design**: Did the tool surface make a wrong call too easy?
- **Guardrail Effectiveness**: Why did [[OzyZT]] or [[Security-Guardrails]] not block the action?
- **Recurrence Prevention**: The *single* change that prevents this specific incident from happening again.

## 🛠️ Tooling Stack
- **Storage**: Markdown files in `docs/adr/`.
- **Workflow**: `/decide` command for agents to scaffold an AgDR.
- **Monitoring**: OzyAudit to scan for "Shadow Architecture" (undocumented patterns).

## Related
- [[Security-Guardrails]]
- [[AI-Analysis-Workflow]]
- [[Agentic-Workflows]]
- [[OzyZT Architecture]]
