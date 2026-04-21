---
type: concept
title: "AI Analysis Workflow"
complexity: intermediate
domain: ai-project-management
created: 2026-04-21
updated: 2026-04-21
tags:
  - concept
  - workflow
  - analysis
status: active
related:
  - "[[LLM Wiki Pattern]]"
  - "[[SDD-Workflow]]"
sources:
---

# AI Analysis Workflow

Patrón de análisis para requests de nuevos proyectos antes de implementar.

---

## Definition

Flujo estructurado que el agente sigue para entender un request antes de escribir código. Evita trabajar sin contexto claro.

## The 5 Steps

| Step | Action | Output |
|------|--------|-------|
| 1 | Analizar | Tipo, funcionalidad, complejidad |
| 2 | Buscar | Memorias, wiki, skills |
| 3 | Identificar stack | Qué existe |
| 4 | Identificar recursos | Wiki, concepts, entities |
| 5 | **Preguntar** | Pregunta crítica |

## Pregunta Crítica

Una pregunta **clara y compleja** que el usuario debe responder para definir la solución.

> Ejemplo: "quiero un agent que monitoree servidores"
> 
> Pregunta: "¿Cómo decides que algo Falló?"
> - Solo up/down?
> - Métricas (CPU > 90%)?
> - Errores en logs?
> - IA que detecte anomalías?

## Por qué funciona

- Evita assumptions incorrectas
- El usuario define la solución
- Stack = basado en decisiones, no en suposiciones
- Wiki pages = contexto compilado

## Regla

> NUNCA proceder sin la pregunta crítica respondida

Primero entender, después codificar.

## Ejemplo Aplicado

Request: "quiero un AI agent que alerte por Telegram"

1. Analizar → AI Agent + Alerts + Telegram
2. Buscar → tengo n8n, Telegram Bot, Hermes-Agent
3. Stack existente → Go API, n8n, Telegram configurado
4. Recursos → wiki pages para el nuevo proyecto
5. Preguntar → "¿Qué constituye una 'falla'?"

---

## Source

- Workflow desarrollado 2026-04-21 con el usuario