# Security Guardrails

**Status**: MANDATORY
**Type**: Security Protocol
**Parent**: [[AI-PROTOCOL]]

## 1. Secret Management
- **NEVER** commit `.env`, `*.key`, `*.pem`, or `credentials.json` files.
- Use a `.gitignore` that explicitly blocks these patterns.
- If a secret is leaked, the agent must immediately report it and suggest a rotation.

## 2. Input Validation
- Always assume user input is malicious.
- Use libraries like `Zod` (Node) or native validation in Go.
- Sanitize all SQL/NoSQL queries to prevent injection.

## 3. Least Privilege
- Agents should only have access to the directories they need.
- Docker containers must run as a non-root user (see [[concepts/tech-stack/Docker]]).

## 4. API Safety
- Implement rate limiting on all public endpoints.
- Use HTTPS/TLS for all communications in production.

---
**Last Updated**: 2026-04-19
