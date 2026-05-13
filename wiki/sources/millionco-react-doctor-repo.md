---
type: source
title: "Million.co - React Doctor Repository"
source_url: "https://github.com/millionco/react-doctor"
created: 2026-05-13
updated: 2026-05-13
tags: [react, tooling, static-analysis, dx]
---

# Million.co - React Doctor Repository

[React Doctor](https://github.com/millionco/react-doctor) is a diagnostic tool designed to scan React codebases and output a **0 to 100 health score** with actionable diagnostics. It aims to catch "bad React" patterns and teach coding agents best practices.

## Key Takeaways
- **Health Score**: Calculates a 0-100 score based on unique rules triggered (`100 - (errors * 1.5) - (warnings * 0.75)`).
- **Agent Integration**: Specifically supports teaching agents (OpenCode, Claude Code, Cursor) via `npx react-doctor install`, which adds `.cursorrules`, `SKILL.md`, or `AGENTS.md` files.
- **Ruleset**: Covers state & effects, performance, architecture, security, accessibility, and dead code.
- **Tooling**: Works with Next.js, Vite, and React Native. Integrates with GitHub Actions for PR comments.
- **Customization**: Supports `react-doctor.config.json` for ignoring rules or files.

## Integration for OpenCode
The command `npx -y react-doctor@latest install` detects OpenCode and installs local rule files to improve the agent's React output.

## Source
- [millionco/react-doctor](https://github.com/millionco/react-doctor)
- [[million-co]] (Organization)
- [[react-doctor]] (Product)
