---
type: blueprint
title: "n8n mcp exposure plan"
updated: 2026-05-14
status: active
---

# n8n MCP Exposure Plan

> Plan de exposición gradual para n8n MCP en el Bunker. Prioriza utilidad alta y bajo riesgo primero.

## Objetivo

Definir qué workflows conviene exponer primero a `n8n-mcp` para que OpenCode pueda usarlos sin abrir superficies de riesgo innecesarias.

## Criterios

- Preferir workflows idempotentes o casi idempotentes.
- Evitar al inicio flujos que escriban en GitHub, Telegram u otros destinos externos sin validación.
- Priorizar flujos con entradas claras, salidas previsibles y fácil observabilidad.
- No exponer nada que dependa de credenciales todavía incompletas o placeholders.

## Orden Recomendado

### 1. Crear un workflow read-only de estado

- Rol: health check, inventario de workflows, status snapshot.
- Motivo: es la mejor primera superficie MCP porque no escribe nada.
- Riesgo: bajo.
- Estado: ✅ Creado como `bunker-health-check.json` en `automation/n8n-lab/workflows/workflows/`.

### 2. `Bunker Ultimate Alerter`

- Archivo: `automation/n8n-lab/workflows/workflows/bunker-ultimate-alerter.json`
- Rol: alerta de seguridad por webhook con salida a Telegram.
- Motivo: es el flujo más simple y más valioso para validación inicial.
- Riesgo: medio, porque toca mensajería externa, pero la lógica es lineal y controlable.
- Estado: `chatId` configurado y workflow activo localmente.
- Requisito previo: validar el webhook `bunker-security-alert` en una ejecución real.

### 3. Flujo de respuesta asistida para remediación

- Archivo: `automation/n8n-lab/workflows/workflows/bunker-assisted-remediation.json`
- Rol: análisis de Dockerfile y propuesta de parche.
- Motivo: sirve para probar el acoplamiento entre evento, IA y acción posterior.
- Riesgo: alto, porque hoy termina en edición de GitHub y tiene credenciales placeholder.
- Requisito previo: reemplazar `TU_ID_MISTRAL` y `TU_ID_GITHUB`, y añadir una revisión manual antes de cualquier commit automático.

### 4. Cualquier workflow de consulta o estado

- Rol: read-only, introspección, health checks, status snapshots.
- Motivo: son los mejores candidatos para MCP porque reducen riesgo y ayudan a diagnosticar.
- Riesgo: bajo.
- Requisito previo: ninguno más allá de exponer salida útil y consistente.

## No Exponer Primero

- Workflows que alteren repositorios sin aprobación explícita.
- Workflows con dependencias rotas o campos placeholder.
- Flujos que escriban a múltiples sistemas externos en una sola ejecución.

## Decisión

Para la configuración actual, el siguiente paso lógico es:

1. ~~crear un workflow read-only de estado~~ (✅ Completado)
2. terminar de endurecer `Bunker Ultimate Alerter`
3. configurar `bunker-assisted-remediation.json` en modo asistido con revisión humana

## Relacionados

- [[n8n]]
- [[MCP-CONFIG]]
- [[Automation Strategy Blueprint]]
- [[Security-Guardrails]]

---
**Last Updated**: 2026-07-02
