# n8n (Orchestration Engine)

**Type**: Low-Code Automation / Orchestrator
**Category**: Infrastructure / Workflow
**Status**: Researching / Integration Reference
**Parent**: [[_index]]

## 📝 Overview

**n8n** is a powerful, fair-code workflow automation tool. It allows for the creation of complex, multi-step processes by connecting hundreds of apps and APIs. The [[Zie619]] repository provides 2,000+ community templates to accelerate adoption.

## 🚀 Key Features

- **Fair-Code & Self-Hosted**: Privacy-first, runs locally or in the cloud.
- **AI-Agent Nodes**: Native support for LangChain-style AI chains and agents.
- **JSON Based**: Workflows are stored as portable JSON files.
- **Infinite Extensibility**: Custom JS nodes for logic that goes beyond standard integrations.

## 🧠 Connection to opencode-obsidian

n8n acts as the **Central Nervous System**. It can trigger agent actions, sync the wiki to external databases, or send alerts based on [[Observability-MD]] events.

### Patterns to Adopt:
1. **Workflow-as-Code**: Storing n8n JSON templates in the bunker for version control.
2. **AI-BOM Auditing**: Using security tools to ensure no secrets are hardcoded in automations.
3. **Event-Triggered Agents**: Using n8n to listen to webhooks and wake up specialized agents (e.g., [[Decepticon]] or [[DeepTutor]]).

## MCP Exposure Order

For the current local setup, see [[n8n-mcp-exposure-plan]] for the recommended order to expose workflows to `n8n-mcp`.

## 🔗 Resources
- **Repository**: https://github.com/Zie619/n8n-workflows
- **Templates**: https://zie619.github.io/n8n-workflows/
- **Official Site**: https://n8n.io

---
**Last Updated**: 2026-04-19
