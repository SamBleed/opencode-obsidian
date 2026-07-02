---
type: concept
title: "Test-Driven Development"
domain: engineering
created: 2026-05-16
updated: 2026-05-16
tags: [TDD, testing, engineering]
---

# Test-Driven Development (TDD)

## Overview
Write a failing test before implementing the code that makes it pass. Tests serve as executable specifications and protect against regressions. TDD enforces a strict red‑green‑refactor cycle.

## Core Cycle
1. **RED** – Create a test that fails; run the test suite.
2. **GREEN** – Write minimal code to make the test pass.
3. **REFACTOR** – Clean up implementation while keeping all tests green.

## When to Use
- Implementing new logic, bug fixes, or behavior changes.
- Any change where correctness must be proven before shipping.
- For pure configuration or documentation changes, skip TDD.

## Key Practices
- **Write the test first** – It must fail before any production code.
- **Keep tests small and focused** – One assertion per behavior.
- **Run the full suite after each step** – Ensure no regressions.
- **Refactor safely** – Only after the test passes; run tests again post‑refactor.
- **Prefer real implementations over mocks** – Use fakes or in‑memory substitutes when appropriate.

## Common Anti‑Patterns
| Anti‑Pattern | Why It's Harmful | Fix |
|--------------|------------------|-----|
| Skipping the failing test | Guarantees no regression guard | Write the test first; let it fail. |
| Over‑mocking | Tests pass but production may break | Use real implementations or lightweight fakes; reserve mocks for external services. |
| Skipping refactor step | Code becomes messy over time | Allocate time for cleanup after each green test. |
| Large, monolithic tests | Hard to isolate failures | Split into small, descriptive test cases. |
| Ignoring test failures | Bugs compound, harder to debug | Stop and fix failures before proceeding. |

## Verification Checklist
- [ ] Every new behavior has a corresponding test.
- [ ] All tests pass (`npm test` or equivalent).
- [ ] Refactor does not break existing tests.
- [ ] Test names clearly describe the behavior being verified.
- [ ] No tests are skipped or disabled.