---
type: concept
title: Context Engineering
domain: ai-agent-engineering
created: 2026-05-16
updated: 2026-05-16
tags: [AI, context, agent-quality]
sources: [[agent-skills]]
---

# Context Engineering

## Definition
Feed agents the right information at the right time. Context is the single biggest lever for agent output quality — too little and the agent hallucinates, too much and it loses focus.

## The Context Hierarchy
1. **Rules Files** (AGENTS.md, SKILL.md) — Always loaded, project-wide
2. **Spec / Architecture Docs** — Loaded per feature/session
3. **Relevant Source Files** — Loaded per task
4. **Error Output / Test Results** — Loaded per iteration
5. **Conversation History** — Accumulates, compacts

## Key Patterns
- Create a rules file that persists across sessions (highest-leverage context)
- Structure context from most persistent to most transient
- Load specs per feature, not globally
- Only read source files relevant to the current task
- Feed CI failures directly back to agents

## Red Flags
- Agent hallucinating APIs or patterns
- Agent ignoring project conventions
- No rules file configured for the project
- Loading entire codebase when only a subset is needed

## Source
- [[agent-skills]] — Context Engineering skill
