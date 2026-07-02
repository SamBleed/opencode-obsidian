---
type: concept
title: Doubt-Driven Development
domain: engineering
created: 2026-05-16
updated: 2026-05-16
tags: [review, verification, adversarial]
sources: [[agent-skills]]
---

# Doubt-Driven Development

## Definition
A confident answer is not a correct one. Materialize a fresh-context reviewer — biased to **disprove**, not approve — before any non-trivial output stands. In-flight cross-examination while course-correction is still cheap.

## When to Use
A decision is non-trivial when:
- It introduces or modifies branching logic
- It crosses a module or service boundary
- It asserts a property the type system cannot verify
- Its correctness depends on invisible context
- Its blast radius is irreversible (deploy, migration, API change)

## The Process (5 Steps)
1. **CLAIM** — Write the claim + why it matters
2. **EXTRACT** — Isolate artifact + contract, stripped of reasoning
3. **DOUBT** — Invoke fresh-context reviewer with adversarial prompt
4. **RECONCILE** — Classify every finding against the artifact
5. **STOP** — Trivial findings only, 3 cycles max, or user override

## Key Patterns
- Not `/review` — this is in-flight, not post-hoc
- Fresh-context review means the reviewer has NO prior session context
- Don't add to persona skills list (avoids orchestration anti-pattern)

## Red Flags
- Skipping doubt for "obvious" non-trivial decisions
- Self-questioning fallback without flagging as degraded
- Approving without adversarial framing

## Source
- [[agent-skills]] — Doubt-Driven Development skill
