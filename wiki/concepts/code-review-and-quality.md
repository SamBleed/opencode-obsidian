---
type: concept
title: "Code Review and Quality"
domain: engineering
created: 2026-05-16
updated: 2026-05-16
tags: [code review, quality, engineering]
---

# Code Review and Quality

## Overview
Enforces five‑axis review standards: change size (~100 lines), severity labeling (Nit/Optional/FYI), review speed norms, and effective splitting strategies. Guarantees high‑quality code before merge.

## Key Practices
- **Atomic Commits**: Keep changes small (~100 lines) for focused review.
- **Severity Labels**: Categorize feedback (Nit for style, Optional for enhancements, FYI for informational).
- **Review Speed**: Target timely reviews; understand norms for your team.
- **Splitting Strategies**: Break large changes into logical slices, each with clear acceptance criteria.
- **Quality Gates**: Ensure tests pass, lint passes, and no new warnings are introduced.

## Related Concepts
- [[spec-driven-development]] – Define acceptance criteria.
- [[incremental-implementation]] – Build changes incrementally.
- [[test-driven-development]] – Verify correctness before review.
- [[git-workflow-and-versioning]] – Versioning practices for reviewable units.

## Verification Checklist
- [ ] Change size within ~100 lines.
- [ ] All tests pass and linting clean.
- [ ] Review assigned with appropriate severity labels.
- [ ] Review completed within agreed speed norm.
- [ ] No unresolved comments remain before merge.