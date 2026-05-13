---
type: concept
title: "Coding Agent Integration (React Doctor)"
domain: ai
created: 2026-05-13
updated: 2026-05-13
tags: [ai, agents, prompt-engineering, react]
status: active
---

# Coding Agent Integration (React Doctor)

[[react-doctor]] provides a mechanism to "teach" AI coding agents (like OpenCode, Claude Code, or Cursor) about React best practices by injecting rule-based instructions into the project's local context.

## Mechanism
When running the install command:
```bash
npx -y react-doctor@latest install
```
The tool detects the active agent environment and writes specific files that the agent is programmed to read:
- **OpenCode**: `AGENTS.md` / `SKILL.md`
- **Claude Code**: `SKILL.md`
- **Cursor**: `.cursorrules`

## Why it matters
Instead of just fixing errors after they happen, these files provide "upstream" guidance. When an agent reads these patterns before writing code, it avoids common pitfalls like:
- Fetching inside `useEffect` without proper handling.
- Using derived state in `useEffect`.
- Excessive re-renders.

## Patterns
The generated files typically contain a curated list of React Doctor's most critical rules formatted as agent instructions.

## Source
- [[millionco-react-doctor-repo]]
