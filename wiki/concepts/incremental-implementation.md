---
type: concept
title: "Incremental Implementation"
domain: engineering
created: 2026-05-16
updated: 2026-05-16
tags: [AI, engineering, incremental-implementation]
---

# Incremental Implementation

## Overview
Build changes incrementally using thin vertical slices. Implement one small, complete piece, test it, verify it, commit, then proceed to the next slice. This prevents large, untested code accumulations and ensures each step leaves the system in a working state.

## Key Practices
- **Slice vertically**: Deliver end‑to‑end functionality in each increment (e.g., DB + API + UI).
- **Risk‑first slicing**: Tackle the highest‑risk component early.
- **Scope discipline**: Only touch what the current task requires; do not refactor unrelated code.
- **Feature flags & safe defaults**: Keep incomplete work hidden until ready.
- **Rollback‑friendly changes**: Design each increment to be independently revertable.

## When to Use
- Multi‑file changes or new features.
- Refactoring existing code.
- Any situation where you’d write more than ~100 lines before testing.

## Common Rationalizations (and why they’re dangerous)
| Rationalization | Reality |
|-----------------|---------|
| “I’ll test it all at the end.” | Bugs compound; each slice must be verified. |
| “It’s faster to do it all at once.” | Large commits hide bugs and make rollbacks painful. |
| “These small fixes are too minor to commit.” | Small commits are free; they keep the codebase healthy. |

## Verification Checklist
- [ ] Each increment is tested and committed.
- [ ] All existing tests still pass.
- [ ] The build succeeds.
- [ ] Type checking and linting pass.
- [ ] The new functionality works as expected.
- [ ] No uncommitted changes remain.

**Remember:** Incremental Implementation keeps changes small, safe, and continuously verifiable.