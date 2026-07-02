---
type: concept
title: Documentation and ADRs
domain: engineering
created: 2026-05-16
updated: 2026-05-16
tags: [documentation, ADR, architecture]
sources: [[agent-skills]]
---

# Documentation and ADRs

## Definition
Document decisions, not just code. The most valuable documentation captures the *why* — context, constraints, and trade-offs that led to a decision.

## When to Write an ADR
- Choosing a framework, library, or major dependency
- Designing a data model or database schema
- Selecting an authentication strategy
- Any decision expensive to reverse

## ADR Template
- **Status**: Accepted | Superseded | Deprecated
- **Context**: What is the issue? Requirements and constraints.
- **Decision**: What was chosen and why.
- **Alternatives Considered**: What else was evaluated.
- **Consequences**: What results from this decision.

## Key Patterns
- Store ADRs in `docs/decisions/` with sequential numbering
- Document the *why*, not the *what* — code shows what
- Write ADRs for architectural decisions, API changes, and feature ships
- Don't document obvious code

## Red Flags
- No ADRs for major technical decisions
- Documentation that restates what code already says
- Decisions made without recording alternatives considered

## Source
- [[agent-skills]] — Documentation and ADRs skill
