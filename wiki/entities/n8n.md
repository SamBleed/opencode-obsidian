---
type: entity
title: "n8n"
domain: automation
created: 2026-05-13
updated: 2026-05-13
tags: [automation, workflow, low-code, nodejs]
status: mature
---

# n8n

n8n is an extendable workflow automation tool that enables connecting various applications and services. In the Bunker, n8n acts as the **central nervous system** for asynchronous operations and external integrations.

## 🚀 Key Roles in the Bunker

### 1. Data Ingestion Pipeline
n8n polls external sources (RSS, GitHub, Slack) and pushes them to the [[vault-flow]] for ingestion.
- **Path**: `External Source` → `n8n Webhook/Cron` → `Go Ingest Server (:9090)` → `Wiki Inbox`.

### 2. Event Orchestration
Triggers security scans or notifications based on events.
- **Example**: A new PR in GitHub triggers a [[Trivy]] scan via n8n, which then posts the results back to the PR and the wiki.

### 3. Agent Tooling (MCP)
n8n can expose its workflows as **MCP Tools**, allowing agents to execute complex, multi-step automations with a single natural language command.

## 🏗️ Best Practices

- **Atomic Workflows**: Keep workflows small and focused on one task. Use "Execute Workflow" node for composition.
- **Error Handling**: Always use "Error Trigger" nodes to catch and log failures to the wiki's `log.md`.
- **Credential Security**: Never hardcode API keys. Use n8n's native credential management.
- **Version Control**: Export workflows as JSON and commit them to the `automation/` folder in the repo.

## ⚠️ Gotchas
- **Memory Usage**: Complex workflows with large data sets can crash the n8n container if memory limits aren't set.
- **Execution Mode**: Use `queue` mode for high-volume production environments to avoid blocking the main process.

## Related
- [[AUTOMATION-STRATEGY]]
- [[n8n-mcp-exposure-plan]]
- [[vault-flow]]
- [[OzyZT Architecture]] (Uses n8n for alerts)
