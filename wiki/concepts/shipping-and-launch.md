---
type: concept
title: Shipping and Launch
domain: engineering
created: 2026-05-16
updated: 2026-05-16
tags: [deployment, launch, production]
sources: [[agent-skills]]
---

# Shipping and Launch

## Definition
Ship with confidence. Every launch should be reversible, observable, and incremental. Deploy safely with monitoring, rollback plan, and clear success criteria.

## Pre-Launch Checklist
- **Code Quality**: Tests pass, build clean, lint/typecheck pass, reviewed
- **Security**: No secrets in code, npm audit clean, input validation, auth checks, security headers, rate limiting
- **Performance**: Core Web Vitals good, no N+1 queries, images optimized, bundle within budget
- **Accessibility**: Keyboard nav, screen reader, color contrast, focus management
- **Infrastructure**: Env vars set, migrations ready, DNS/SSL, CDN, logging, health check
- **Documentation**: README, API docs, ADRs, changelog updated

## Key Patterns
- Feature flags to decouple deployment from release
- Staged rollouts: staging → production with monitoring window
- Rollback plan for every deployment
- Monitor for errors post-deploy (15-minute window)

## Red Flags
- Deploying without tests passing
- No rollback mechanism
- No monitoring or alerting configured
- Secrets in code or CI config

## Source
- [[agent-skills]] — Shipping and Launch skill
