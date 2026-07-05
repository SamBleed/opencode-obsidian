---
type: concept
title: "AI Analysis Workflow"
domain: ai
created: 2026-05-13
updated: 2026-05-13
tags: [workflow, ai, logic, analysis]
status: mature
---

# AI Analysis Workflow

The AI Analysis Workflow is the fundamental protocol for all agent operations within the Bunker. It ensures that agents do not execute commands blindly, but instead reason through the request, the context, and the security implications first.

## 🚀 The 5-Step Protocol

### 1. Requirement Deconstruction
Break the user's prompt into atomic requirements.
- What is the primary goal?
- What are the constraints (language, stack, performance)?
- Is there any ambiguous terminology?

### 2. Context Discovery (The "Search First" Rule)
Never assume the state of the codebase.
- Use `grep` and `glob` to find relevant files.
- Read existing implementations to follow conventions.
- Check `wiki/hot.md` for recent project context.

### 3. Technical Reasoning & Planning
Formulate a step-by-step plan before touching a single file.
- Propose the plan to the user.
- Identify potential breaking changes or security risks.
- Design the verification path (tests).

### 4. Implementation with "Agentic AppSec"
While implementing, the agent acts as its own security reviewer.
- **Sanitization**: Validate all user-controlled inputs.
- **Best Practices**: Adhere to [[coding-standards]] and [[React]] patterns.
- **Audit**: Every significant change must be logged in the [[agentmemory]].

### 5. Multi-Layer Verification
Prove the implementation works and is secure.
1. **Linter/Type-Check**: Run `npm run lint` or `go vet`.
2. **Unit Tests**: Verify logic.
3. **Security Gate**: Run [[react-doctor]] or [[Trivy]].
4. **Final Diff**: Review the final changes against the original intent.

## 🛡️ Governance & AI-BOM

In a mature agentic environment, we track the provenance of decisions:
- **Decision Record**: High-level choices are saved to [[agentmemory]].
- **AI-BOM**: We keep a log of which agent/model version performed which action to ensure accountability.

## 🧠 The "Critical Question" Rule
Before executing any high-impact or ambiguous task, the agent must ask **exactly one** clarifying question to the user. This prevents expensive mistakes and ensures alignment.

## Related
- [[Agentic-Workflows]]
- [[vault-flow]]
- [[Security-Guardrails]]
- [[coding-standards]]
