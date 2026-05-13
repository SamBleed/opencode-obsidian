---
type: meta
title: "Hot Cache"
updated: 2026-05-06
tags:
  - meta
  - hot-cache
status: evergreen
related:
  - "[[index]]"
  - "[[log]]"
  - "[[Wiki Map]]"
  - "[[getting-started]]"
  - "[[opencode-obsidian-v1.4-release-session]]"
---

# Recent Context

Navigation: [[index]] | [[log]] | [[overview]]

## Vault Flow (Para Nuevas Sesiones)

El vault es un wiki persistente en `/home/sam/opencode-obsidian/wiki/`. Tres operaciones principales:

| Operacion | Cuándo | Cómo |
|----------|--------|------|
| Guardar | Aprení algo importante | `mem_save` + wiki page |
| Consultar | Necesito contexto | `mem_search` o wiki-query |
| Ingestar | Recurso nuevo | workflow de 7 pasos |

Workflow de Ingestion (7 pasos):
1. Crear source → `wiki/sources/[slug].md`
2. Crear entidad → `wiki/entities/[name].md` (si aplica)
3. Crear concepto → `wiki/concepts/[name].md`
4. Incluir CÓDIGO — no solo descripción
5. Actualizar `wiki/index.md`
6. Actualizar `wiki/log.md`
7. Vincular en `wiki/Wiki Map.canvas`

Véase [[vault-flow]] para guía completa.

## Last Updated
2026-05-13: **Expansion Quadrant (v2.0)**: Ingestados 4 grandes dominios: Vector Databases (RRF), IaC Pro (OpenTofu), Tailwind 4 Mastery y Gobernanza de Decisiones (AgDR).
2026-05-13: **Graph Health & Connectivity (v1.9)**: Reducida la deuda técnica del wiki. Fusionadas notas de performance, corregidos huérfanos core (Comparativas, FAQ, Blueprints) y expandidos motores de OzyZT.
2026-05-13: **System Hardening & Observability (v1.8)**: Implementado Blueprint de Observabilidad de Agentes y Gobernanza MCP v2026. Script de integridad operativo.
2026-05-13: **NotebookLM Integration (v1.7)**: Ingestados patrones avanzados de Workflows Agénticos, AI Swarms, Policy-as-Code y Seguridad de Supply Chain desde notebooks de investigación.
2026-05-13: **Systemic Capacity Upgrade (v1.6)**: Todas las notas core (Go, React, Postgres, Docker, Security) expandidas a capacidad de Arquitecto Senior (100+ líneas, patterns 2026).
2026-05-13: Ingestado [[millionco-react-doctor-repo]], [[react-doctor]], [[million-co]] para diagnóstico de React y mejores prácticas de agentes.
2026-05-06: Ingestados 5 repos recientes: [[PicoClaw]], [[qwen-code]], [[Pentest-Swarm-AI]], [[NousResearch-hermes-agent]], [[onyx-dot-app-onyx]].
2026-05-06: Ingestados [[affaan-m-everything-claude-code]], [[claude-code-harness]], [[aaif-goose-goose]] para ecosistema Claude Code.
2026-05-05: Ingestado [[zakirkun-deep-eye]], [[ai-powered-pentesting]], [[Muhammad-Zakir-Ramadhan]] para AI pentesting.
2026-05-04: NotebookLM MCP CLI instalado y autenticado. 3 notebooks activos.
2026-05-04: Ingestada fuente [[notebooklm]] y entidad [[notebooklm-mcp-cli]] para procesamiento de contenido vault.
2026-04-25: Integrado [[Awesome-DESIGN-md]] y actualizado concepto [[DESIGN-MD]].
2026-04-22: Vault flow documentado para nuevas sesiones. MCP Governance operativo.

## Active Threads
- **AI Agent Ecosystem**: 5 nuevos agentes ingeridos (PicoClaw, qwen-code, Pentest-Swarm, Hermes, Onyx). Estudiando arquitecturas para OzyZT.
- **Claude Code Harness**: [[affaan-m/everything-claude-code]] integrado. 16+ agents, 65+ skills disponibles para OpenCode.
- **AI-Powered Pentesting**: [[Pentest-Swarm-AI]] es el primer swarm real (no pipeline). Crítico para OzyZT + Owl Perú.
- **Awesome DESIGN.md Integration**: Ya puedo usar especificaciones de diseño de marcas top (Stripe, Vercel, Linear) para generar UIs pixel-perfect.
- **OzyZT Platform**: Modular Zero-Trust security engine con Slack y OzyAudit integrations.
- **n8n Ingest Pipeline**: n8n polling → Go server :9090 → wiki/inbox/ (debugging connection).

- v1.5.0 backlog: `/adopt` command, vault graph analysis in wiki-lint, semantic search via qmd, Marp output
- `community` remote (`avalonreset-pro/opencode-obsidian`) still has pre-write history. Force-push needed next time that remote is configured.

## Repo Locations
- Working: `~/Desktop/opencode-obsidian/`
- Public: https://github.com/AgriciDaniel/opencode-obsidian
- Community (private): https://github.com/avalonreset-pro/opencode-obsidian
