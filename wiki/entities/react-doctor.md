---
type: product
title: "React Doctor"
created: 2026-05-13
updated: 2026-05-13
tags: [tool, react, static-analysis]
status: active
---

# React Doctor

React Doctor is a CLI tool and lint plugin that analyzes React codebases to improve quality and performance. It provides a "health score" and identifies issues in state management, effects, and architecture.

## Features
- **0-100 Health Score**: A quantitative measure of codebase quality.
- **Diagnostic Rules**: Custom rules for common React anti-patterns (e.g., `no-fetch-in-effect`, `no-derived-state-effect`).
- **Agent Skills**: Can generate `SKILL.md` or `AGENTS.md` files for AI coding agents to follow React best practices.
- **Framework Support**: Next.js, Vite, and React Native.

## Usage
```bash
npx -y react-doctor@latest .
```

## Related
- Created by [[million-co]]
- Concept: [[react-health-score]]
- Concept: [[coding-agent-integration]]
