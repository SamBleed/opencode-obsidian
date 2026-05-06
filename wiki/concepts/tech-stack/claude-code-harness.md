---
type: concept
title: "Claude Code Harness"
domain: tech-stack
created: 2026-05-06
updated: 2026-05-06
tags: [claude-code, agent-harness, ai, skills, memory]
status: active
related: [[affaan-m-everything-claude-code], [ai-powered-pentesting]]
---

# Claude Code Harness

## Definition
Sistema de configuración y orquestación para agentes de código (Claude Code, OpenCode, Cursor) que provee agents especializados, skills reutilizables y gestión de memoria.

## How It Works
Un "harness" es una capa de configuración que define:
1. **Agents**: Especialistas con personalidad, procesos y entregables probados
2. **Skills**: Patrones comprobados para tareas específicas
3. **Memory**: Sistema de memoria persistente
4. **Security**: Configuraciones de seguridad integradas

## Example (Agent Definition)
```bash
Agent: "Security Auditor"
Role: Analiza código y Dockerfiles
Skills: [security-review, trivy-scan, owasp-check]
Personality: "Ve vulnerabilidades donde otros ven código"
Deliverables: [security-report.md, fixed-code.diff]
```

## Patterns
- **Specialized Agents**: Cada agente tiene un propósito único
- **Skill Composition**: Combina skills para tareas complejas
- **Memory Persistence**: El agente "crece con you"
- **Research-First**: Investiga antes de implementar

## Tools Using This Pattern
- **affaan-m/everything-claude-code**: 16+ agents, 65+ skills
- **OpenCode**: Skills system nativo
- **Cursor**: .cursorrules integration

## Source
- [[affaan-m-everything-claude-code]]
- [[skill-creator]] — Para crear nuevos skills
