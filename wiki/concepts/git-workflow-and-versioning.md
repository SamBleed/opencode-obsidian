---
type: concept
title: Git Workflow and Versioning
domain: engineering
created: 2026-05-16
updated: 2026-05-16
tags: [git, version-control, workflow]
sources: [[agent-skills]]
---

# Git Workflow and Versioning

## Definition
Git is your safety net. Treat commits as save points, branches as sandboxes, and history as documentation. Disciplined version control keeps changes manageable, reviewable, and reversible.

## Core Principles
- **Trunk-Based Development**: Keep `main` always deployable; short-lived feature branches (1-3 days)
- **Commit Early, Commit Often**: Each increment gets its own commit
- **Atomic Commits**: Each commit does one logical thing
- **Descriptive Messages**: Explain the *why*, not just the *what*
- **Conventional Commits**: `feat:`, `fix:`, `refactor:`, `chore:`, etc.

## Key Patterns
- Feature flags > long-lived branches
- Rebase for local cleanup, merge for integration
- Resolve conflicts by understanding both sides, not by blindly accepting
- Branch protection: required reviews, required status checks

## Red Flags
- Giant commits mixing multiple concerns
- Long-lived branches diverging from main
- Commit messages that just say "update" or "fix"
- Force-pushing to shared branches

## Source
- [[agent-skills]] — Git Workflow and Versioning skill
