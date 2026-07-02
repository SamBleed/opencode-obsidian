---
type: concept
title: "Ethics Compliance"
updated: 2026-05-14
status: active
---

# Ethics & Compliance Guardrails

**Status**: Governance / Safety
**Type**: Policy
**Parent**: [[_index]]

> Reglas éticas y legales para la operación de agentes autónomos dentro del ecosistema del Bunker.

## 🛑 Hard Limits (Prohibiciones)
1. **Acceso no autorizado**: Los agentes NO pueden atacar, escanear o interactuar con sistemas fuera del scope autorizado.
2. **Filtración de Secrets**: Queda estrictamente prohibido enviar API Keys, passwords o datos sensibles en los prompts a LLMs externos.
3. **Persistencia Maliciosa**: Los agentes no pueden crear puertas traseras (backdoors) ni ocultar su rastro para evadir la supervisión humana.

## 🤝 Human-in-the-Loop (HITL)
Se requiere aprobación humana explícita para:
- Despliegue en entornos de producción.
- Acciones destructivas (borrar bases de datos, buckets, etc.).
- Compras o transacciones financieras.

## ⚖️ Responsabilidad
El operador humano es el responsable final de las acciones del agente. El agente debe siempre identificarse como tal en sus interacciones externas (`User-Agent: OpenCode-Bunker-Agent`).

---
**Last Updated**: 2026-04-19
