---
type: source
title: NotebookLM (Google)
source_url: 'https://notebooklm.google.com'
created: '2026-05-04'
updated: '2026-05-04'
tags:
  - ai
  - research
  - google
  - knowledge-management
  - mcp
domain: ai
status: active
related:
  - '[[n8n]]'
  - '[[Hermes-Agent]]'
  - '[[AI-Analysis-Workflow]]'
  - '[[notebooklm-mcp-cli]]'
---

# NotebookLM (Google)

## Definition
NotebookLM es una herramienta de investigación asistida por IA desarrollada por Google, diseñada para ayudar a los usuarios a comprender y sintetizar información a partir de sus propias fuentes.

## Key Features
- **Source Grounding**: Solo responde usando las fuentes que el usuario sube (PDFs, docs, notas).
- **FAQ & Study Guides**: Genera preguntas frecuentes y guías de estudio automáticamente.
- **Audio Summaries**: Convierte documentos en resúmenes hablados (Audio Overview).
- **Cross-Source Analysis**: Sintetiza información de múltiples documentos simultáneamente.

## Use Case in the Bunker
Se usa para procesar fuentes del vault (ej. n8n, Hermes, workflows) y generar:
- FAQs y Conceptos Clave
- Study Guides para nuevos agentes
- Análisis cruzado de múltiples notas del vault

## Integration with opencode-obsidian
Cuando el usuario genera contenido en NotebookLM, debe copiar el texto aquí para que el agente lo ingeste formalmente en el vault siguiendo el workflow de `wiki-ingest`.

## Source
- Official: https://notebooklm.google.com
- Related: [[n8n]], [[Hermes-Agent]], [[AI-Analysis-Workflow]]
