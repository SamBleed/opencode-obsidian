---
type: concept
title: "MCP Governance"
domain: tech-stack
created: 2026-04-20
updated: 2026-04-20
tags:
  - mcp
  - governance
  - tech-stack
status: seed
related:
  - "[[index]]"
  - "[[MCP-CONFIG]]"
sources:
  - "[[opencode-obsidian/AGENTS.md]]"
---

# MCP Governance

## Definition
Marco de trabajo para la gestión y orquestación de múltiples servidores Model Context Protocol (MCP) evitando solapación de herramientas y fatiga cognitiva del agente.

## How It Works
Se basa en la jerarquía de herramientas definida en `AGENTS.md` y la delegación de tareas a agentes especializados. El objetivo es que el orquestador sepa exactamente qué herramienta es la más eficiente para cada tipo de tarea (Búsqueda, Ingestión, Acción o Seguridad).

## Example
### Jerarquía de Búsqueda
1. **Context7**: Documentación oficial de librerías.
2. **Exa**: Noticias y tendencias.
3. **Browser**: Interacción humana.

## Patterns
- **Abstraction Layer**: No llamar a herramientas complejas directamente si existe un script en `bin/` que estandarice los flags.
- **Agent Segmentation**: Limitar las herramientas visibles para cada sub-agente según su propósito.
- **On-Demand Activation**: Mantener desactivados los servidores MCP de uso esporádico (ej. Trivy) en `opencode.json` hasta que sean necesarios.

## Source
- [[AGENTS.md]]
- [[MCP-CONFIG]]
