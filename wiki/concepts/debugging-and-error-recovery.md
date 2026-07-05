---
type: concept
title: "Debugging and Error Recovery"
domain: engineering
created: 2026-05-16
updated: 2026-05-16
tags: [debugging, error handling, engineering]
---

# Debugging and Error Recovery

## Overview
Systematic triage process for test failures, broken builds, and runtime bugs. Emphasizes stopping further changes, preserving evidence, reproducing the issue, localizing the failure, reducing to a minimal case, fixing the root cause, and guarding against recurrence.

## Key Practices
- **Stop‑the‑Line Rule**: Halt all other work when an unexpected error occurs.
- **Reproduce First**: Ensure the failure can be reliably reproduced before attempting a fix.
- **Bisection**: Use `git bisect` to identify the commit that introduced the regression.
- **Minimal Reproduction**: Strip away unrelated code to isolate the bug.
- **Root‑Cause Fix**: Address the underlying issue, not just symptoms.
- **Regression Test**: Add a test that fails without the fix to prevent future regressions.
- **Guarded Recovery**: Implement safe fallbacks and instrument logging only when necessary.

## Common Rationalizations
| Rationalization | Reality |
|-----------------|--------|
| “I know the bug, I’ll just fix it.” | You might be right 70 % of the time; the remaining 30 % costs hours. Reproduce first. |
| “The test is probably wrong.” | Verify the assumption; if the test is wrong, fix the test, don’t skip it. |
| “It works on my machine.” | Environments differ; verify in CI, check configs, dependencies. |
| “I’ll fix it in the next commit.” | Fix it now; the next commit will introduce new bugs on top of the current one. |
| “This is a flaky test, ignore it.” | Flaky tests mask real bugs; fix the flakiness or understand why it’s intermittent. |

## Verification Checklist
- [ ] Issue reproduced reliably.
- [ ] Root cause identified and documented.
- [ ] Fix addresses the root cause, not just symptoms.
- [ ] Regression test added that fails without the fix.
- [ ] All existing tests pass.
- [ ] Build succeeds.
- [ ] Original scenario verified end‑to‑end.