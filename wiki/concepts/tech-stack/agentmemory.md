---
type: concept
title: "AgentMemory"
domain: tech-stack
created: 2026-05-13
updated: 2026-05-13
tags: [ai, memory, persistence, go, sqlite]
status: mature
---

# AgentMemory

AgentMemory is the persistent memory engine that powers the Bunker's agents. It allows agents to retain context across sessions, recall past decisions, and learn from project history.

## 🧠 Core Philosophy: "Compound Knowledge"
Unlike standard RAG (Retrieval-Augmented Generation) which often pulls disjointed fragments, AgentMemory focuses on **Structured Experience**. It records not just *what* happened, but *why* and *what was learned*.

## 🏗️ Technical Architecture

### Storage Layer
- **SQLite (FTS5)**: Used for high-speed keyword and full-text search.
- **Vector Store (Optional)**: For semantic similarity search.
- **JSON Files**: For raw observation persistence.

### Data Model
Every memory is an **Observation** with the following metadata:
- **Title**: Searchable identifier.
- **Content**: The core insight (**What, Why, Where, Learned**).
- **Topic Key**: A stable identifier for evolving concepts (e.g., `architecture/auth`).
- **Scope**: `project` or `personal`.
- **Project**: Associated repository or initiative.

## 🛠️ Usage Patterns

### Proactive Saving
Agents are instructed to call `mem_save` after any significant event:
- A bug fix with a non-obvious root cause.
- An architectural decision with specific tradeoffs.
- A user preference discovered during conversation.

### Semantic Recall
When a user asks "How did we solve X before?", the agent:
1. Calls `mem_context` for recent session history.
2. Calls `mem_search` with keywords for deep historical search.
3. Synthesizes the answer using `mem_get_observation`.

## ⚡ Integration with [[vault-flow]]
AgentMemory and the Obsidian Wiki work together:
- **AgentMemory**: Ephemeral, fast, and session-oriented.
- **Obsidian Wiki**: Persistent, curated, and structural knowledge.
*Pattern: High-value observations in AgentMemory should eventually be "promoted" to full Wiki pages.*

## ⚠️ Known Gotchas
- **Context Drift**: If multiple agents write to the same topic key without coordination, the memory can become fragmented.
- **Token Usage**: Recalling too many large observations can exhaust the agent's context window.

## Related
- [[LLM Wiki Pattern]]
- [[Delta-Tracking]]
- [[vault-flow]]
- [[fts5-search]]
