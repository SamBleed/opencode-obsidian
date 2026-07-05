---
type: concept
title: "Agentic Security OWASP"
updated: 2026-05-14
status: active
---

# OWASP Top 10 for Agentic Applications (2026)

**Status**: Core Governance
**Type**: Security Framework
**Parent**: [[_index]]

> Framework definitivo para mitigar riesgos en aplicaciones basadas en agentes autónomos. Este documento rige la seguridad conductual de los agentes en el Bunker.

## 🚨 Riesgos ASI (Agentic Security Index)

### ASI01: Agent Goal Hijack
Manipulación de inputs para redefinir los objetivos del agente.
- **Mitigación**: Definición estricta de misión en [[AGENTS.md]].

### ASI02: Tool Misuse & Exploitation
Uso de herramientas legítimas para fines no deseados.
- **Mitigación**: Restricción de herramientas por perfil en [[PROMPTS]].

### ASI03: Identity & Privilege Abuse
Uso excesivo o malintencionado de credenciales.
- **Mitigación**: Principio de **Least Agency** (mínima autonomía).

### ASI04: Agentic Supply Chain
Vulnerabilidades en herramientas o templates de terceros.
- **Mitigación**: Auditoría de fuentes (ej. [[n8n]], [[Trivy]]).

### ASI05: Unexpected Code Execution (RCE)
El agente genera y ejecuta código malicioso de forma autónoma.
- **Mitigación**: Sandboxing y [[Security-Guardrails]].

## 🛠️ Implementación en el Bunker
1. **Threat Modeling**: Analizar cada nuevo blueprint bajo este Top 10.
2. **Behavioral Auditing**: Monitorizar no solo el output, sino las llamadas a herramientas (vía [[Observability-MD]]).
3. **Human-in-the-Loop (HITL)**: Requerido para riesgos ASI03 y ASI05.

---
**Source**: OWASP GenAI Security Project (2026)
**Last Updated**: 2026-04-19
