---
type: blueprint
title: "AUTOMATION STRATEGY"
updated: 2026-05-14
status: active
---

# Automation Strategy Blueprint

> Blueprint para la orquestación de flujos de trabajo autónomos y automatizaciones low-code.

## 🏗️ Arquitectura de Automatización
```text
/automation
├── workflows/     # Archivos .json de n8n
├── custom-nodes/  # Código JS para nodos personalizados
└── audit/         # Reportes de seguridad (AI-BOM)
```

## 🛠️ Workflow de Integración Bunker + n8n
1. **Trigger**: Un evento de [[Event-Driven]] (ej. `wiki.updated`) llega al Broker (NATS/Redis).
2. **Listen**: Un nodo de n8n escucha el evento y activa el flujo.
3. **Action**: n8n se comunica con APIs externas (Slack, Discord, Email).
4. **Log**: El resultado se devuelve al Bunker para actualizar el [[hot.md]].

## 🔒 Security Guardrails
- **Secret Management**: Usar variables de entorno de n8n; NUNCA hardcodear keys en el JSON.
- **Audit**: Correr un escaneo de seguridad antes de importar cualquier workflow de la comunidad.
- **Isolation**: n8n debe correr en una red aislada con acceso restringido al sistema de archivos del Bunker.

## MCP Exposure Order

Para decidir qué flujos abrir primero a `n8n-mcp`, ver [[n8n-mcp-exposure-plan]].

---
**Inspired by**: [[n8n]] and [[Zie619]]
