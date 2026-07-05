---
type: concept
title: Security and Hardening
domain: security
created: 2026-05-16
updated: 2026-05-16
tags: [security, OWASP, hardening]
sources: [[agent-skills]]
---

# Security and Hardening

## Definition
Security-first development. Treat every external input as hostile, every secret as sacred, and every authorization check as mandatory. Security is a constraint on every line of code that touches user data, auth, or external systems.

## The Three-Tier Boundary System
- **Always Do**: Validate all input, parameterize queries, encode output, HTTPS, hash passwords, security headers, httpOnly cookies, npm audit
- **Ask First**: New auth flows, storing new sensitive data, new integrations, CORS changes, file uploads, rate limit changes
- **Never Do**: Commit secrets, log sensitive data, trust client-side validation, disable security headers, use eval/innerHTML with user data, store sessions in localStorage

## OWASP Top 10 Prevention
- Injection → parameterized queries
- Broken Auth → bcrypt + secure sessions
- XSS → output encoding + CSP
- CSRF → sameSite cookies + tokens
- Security Misconfiguration → security headers, no stack traces

## Red Flags
- Secrets in version control
- SQL concatenation with user input
- Missing auth checks on endpoints
- No security headers configured
- Client-side validation as sole security boundary

## Source
- [[agent-skills]] — Security and Hardening skill
