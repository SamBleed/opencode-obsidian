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

## 5. Red Team Patterns (Inspired by [[Decepticon]])
- **Autonomous Auditing**: Treat every new feature as a target. Run automated recon on new endpoints before deployment.
- **Persistent Session Audits**: Use persistent shells for security scanning to ensure interactive prompts are handled.
- **Kill-Chain Thinking**: Don't just check for "bugs"; check how multiple small bugs could be chained for an exploit.

- **OSINT Recon**: Use tools like [[Sherlock]] to map the digital footprint of a target (usernames, social profiles) before attempting technical exploitation.
- **Data Anonymization**: When logging OSINT results, ensure that personal data is handled according to [[Ethics-Compliance]] standards.

## 6. Blue Team Patterns (Inspired by [[Trivy]] & [[A-poc]])
- **Continuous Vulnerability Scanning**: Run `trivy fs .` regularly to detect CVEs and secrets.
- **Incident Response Plan**: Refer to the [[BlueTeam-Tools]] collection for forensics and recovery procedures.
- **Supply Chain Security**: Verify all external dependencies and tool versions before integration.

## 7. Agentic Security (Inspired by [[Agentic-Security-OWASP]])
- **Least Agency**: Grant the agent only the minimum autonomy required. Avoid "super-agents" with access to everything.
- **Goal Guarding**: Implement intent-gates to prevent "Agent Goal Hijacking" (ASI01).
- **Tool Sandbox**: Restrict tool usage based on the agent's active profile in [[PROMPTS]].
- **Behavioral Monitoring**: Audit tool calls and cross-agent communication via [[Observability-MD]].

---
**Last Updated**: 2026-04-19
