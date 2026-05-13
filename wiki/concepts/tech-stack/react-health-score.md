---
type: concept
title: "React Health Score"
domain: tech-stack
created: 2026-05-13
updated: 2026-05-13
tags: [metrics, react, quality, ci-cd]
status: mature
---

# React Health Score

The React Health Score is a metric used by [[react-doctor]] to quantify the quality of a React codebase. It focuses on unique rules triggered rather than total instances of violations, promoting systemic fixes over repetitive patching.

## 📐 The Formula
The score is calculated as:
`100 - (unique_error_rules * 1.5) - (unique_warning_rules * 0.75)`

### Why unique rules?
If you have 100 instances of `react-doctor/no-fetch-in-effect`, you only lose **1.5 points**. Why? Because it represents a **systemic misunderstanding** of the pattern. Fixing 99 instances doesn't change the score; fixing the *last one* (meaning the team finally learned the pattern) gains the points back.

## 📊 Score Interpretation

| Range | Label | Team Action |
|-------|-------|-------------|
| **75 - 100** | Great | Maintain hygiene. PRs should not decrease the score. |
| **50 - 74** | Needs work | Allocate "Tech Debt" days to eliminate specific rule categories. |
| **0 - 49** | Critical | Stop feature development. Core architecture is likely compromised. |

## 🚀 CI/CD Integration
The health score is most effective when used as a **Quality Gate** in GitHub Actions.

```yaml
- uses: millionco/react-doctor@main
  with:
    fail-on: error
    # You can set a threshold for the score output
    github-token: ${{ secrets.GITHUB_TOKEN }}
```

## 🧠 Cultural Impact
- **Gamification**: Teams compete to have the highest health score.
- **Learning Loop**: Instead of vague "code review", the score provides objective goals. "Let's fix the `useEffect` rules this sprint."
- **Standardization**: Ensures that even juniors follow senior-level patterns through automated feedback.

## Related
- [[react-doctor]]
- [[React]]
- [[coding-agent-integration]]
- [[AI-Analysis-Workflow]]
