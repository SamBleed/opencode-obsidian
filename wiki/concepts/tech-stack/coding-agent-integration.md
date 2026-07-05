---
type: concept
title: "Coding Agent Integration"
domain: ai
created: 2026-05-13
updated: 2026-05-13
tags: [ai, agents, productivity, best-practices]
status: mature
---

# Coding Agent Integration

Coding Agent Integration refers to the process of injecting project-specific knowledge, conventions, and rules into the local context of AI assistants (e.g., OpenCode, Cursor, Claude Code). This ensures the agent writes code that adheres to the project's unique architecture and quality standards.

## 🛠️ Integration Mechanism

Agents read local files to understand their operating environment. [[react-doctor]] automates this via its `install` command.

### 1. OpenCode / Claude Code (`SKILL.md` / `AGENTS.md`)
We use **Skills** to encapsulate behavior. A skill like `react-best-practices` might contain:
- Patterns for RSC (React Server Components).
- Rules against fetching in `useEffect`.
- Naming conventions for hooks.

### 2. Cursor (`.cursorrules`)
Cursor reads this file on every prompt. It is the most "aggressive" form of guidance.
```markdown
# React Rules
- Never use 'useState' for values that can be derived from props.
- All server actions must be in 'actions.ts'.
- Use 'useOptimistic' for all form submissions.
```

## 🔄 The Feedback Loop

1. **Scan**: Run `react-doctor` to find current anti-patterns.
2. **Teach**: Update the agent's instructions (Skills/Rules) to specifically forbid those anti-patterns.
3. **Execute**: The agent now generates compliant code, preventing the same bugs from re-occurring.

## 🚀 Impact on DX (Developer Experience)
- **Zero-Day Onboarding**: A new agent (or human) reading these rules can contribute immediately without violating conventions.
- **Cognitive Load Reduction**: The agent handles the "boring" parts of architecture compliance.
- **Consistency**: The entire codebase looks like it was written by one person.

## ⚠️ Challenges
- **Context Bloat**: Adding too many rules can confuse the agent or exhaust the context window.
- **Stale Rules**: Rules must be updated when the stack evolves (e.g., migrating from React 18 to 19).

## Related
- [[react-doctor]]
- [[PROMPTS]]
- [[vault-flow]]
- [[LLM Wiki Pattern]]
