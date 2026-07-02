---
type: concept
title: "Spec-Driven Development"
domain: engineering
created: 2026-05-16
updated: 2026-05-16
tags: [spec, development, engineering]
---

# Spec-Driven Development

## Overview
Contract-first design where a structured specification precedes any code. The spec defines what we are building, why, acceptance criteria, and success metrics. It creates a shared source of truth, surfaces assumptions early, and prevents guesswork.

## Key Elements
1. **Objective** – Clear user stories and success criteria.
2. **Tech Stack** – Framework, language, dependencies.
3. **Commands** – Executable build/test/lint commands.
4. **Project Structure** – Directory layout.
5. **Code Style** – Naming conventions and formatting examples.
6. **Testing Strategy** – Test framework and coverage expectations.
7. **Boundaries** – Always/Ask First/Never rules.
8. **Success Criteria** – Concrete, testable conditions.

## Workflow
1. **Specify** – Draft spec with assumptions and get human approval.
2. **Plan** – Generate technical plan from spec.
3. **Tasks** – Break plan into discrete tasks.
4. **Implement** – Execute tasks incrementally with verification.

## Why It Matters
- Prevents code without clear requirements.
- Surfaces misunderstandings early (assumptions, edge cases).
- Provides a baseline for scope changes.
- Links documentation directly to implementation.

## Red Flags
- Starting to code without a spec.
- Ignoring or postponing spec drafting.
- Making architectural decisions without documenting them.
- Skipping spec updates when scope evolves.

## Related Concepts
- [[incremental-implementation]] – Build code incrementally.
- [[test-driven-development]] – Write tests before implementation.
- [[spec-driven-development]] – This page.

## Verification
- [ ] Spec covers objective, tech stack, commands, structure, style, testing, boundaries.
- [ ] Human has reviewed and approved the spec.
- [ ] Success criteria are specific and Testable.
- [ ] Spec is version‑controlled alongside code.