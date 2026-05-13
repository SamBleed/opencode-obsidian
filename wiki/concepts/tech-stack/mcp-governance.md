---
type: concept
title: "MCP Governance"
domain: tech-stack
created: 2026-04-20
updated: 2026-05-13
tags: [mcp, governance, policy, security]
status: mature
---

# MCP Governance (v2026)

MCP Governance is the framework for managing, securing, and orchestrating multiple **Model Context Protocol (MCP)** servers. As the ecosystem has grown to 9,000+ servers, governance has shifted from simple tool management to protecting **Privileged Execution Environments**.

## 🛡️ The Zero-Trust MCP Model

In the Bunker, we treat every MCP server as a potential attack vector. Governance is enforced through the following layers:

### 1. The 5 Core Primitives
We categorize tool capabilities into the standard MCP primitives to ensure predictable agent behavior:
- **Tools**: Executable functions (e.g., `bash`, `edit`).
- **Resources**: Static data/files (e.g., `read_note`).
- **Prompts**: Pre-defined templates for specific tasks.
- **Sampling**: Allowing the server to request completions from the LLM.
- **Roots**: Defining the boundaries of the file system the agent can see.

### 2. Transport Security
- **Local (STDIO)**: Used for high-trust, local development tools.
- **Remote (HTTP/SSE)**: Mandatory **OAuth 2.1** for all remote tool connections.
- **Auditing**: Every request-response pair is logged in the [[Agent-Observability-Blueprint]].

### 3. Tool Hierarchy & Deduplication
To avoid "tool fatigue", agents follow a strict lookup order:
1. **Primary**: [[Context7]] (Technical Docs).
2. **Secondary**: [[Exa]] (Web Intelligence).
3. **Tertiary**: [[BrowserMCP]] (Visual Interaction).

## 🚀 Patterns for 2026

### MCP-as-a-Service
Instead of each agent having its own set of tools, the Bunker uses a **Central Tool Registry**. Agents query the registry to find the most appropriate server for a task.

### Bounded Roots
Agents are restricted to specific "roots" in the filesystem. For example, a "Frontend Agent" has root access to `src/components` but not to `wiki/meta/secrets`.

### Agent-to-Agent (A2A) Orchestration
MCP solves the tool-integration problem, while A2A protocols manage how agents hand over tasks.
- **Pattern**: `Orchestrator` → `MCP Query` → `Task Delegation` → `A2A Handover`.

## ⚠️ Risks & Mitigations
- **Tool Poisoning**: Malicious MCP servers returning biased data. *Mitigation*: Only use verified servers from the [[WATCHLIST]].
- **Context Exhaustion**: Too many tools in the context window. *Mitigation*: **On-Demand Activation** (Load tools only when needed).

## Related
- [[MCP-Browser]]
- [[MCP-CONFIG]]
- [[AGENTS.md]]
- [[Agent-Observability-Blueprint]]
- [[OzyZT Architecture]]
