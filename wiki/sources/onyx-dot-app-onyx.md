---
type: source
title: "onyx-dot-app/onyx"
domain: ai
created: 2026-05-06
updated: 2026-05-06
tags: [onyx, ai-platform, rag, llm, enterprise]
status: active
related: [[PicoClaw], [hermes-agent], [qwen-code]]
---

# onyx-dot-app/onyx

## Definition
Plataforma AI open-source que actúa como la "capa de aplicación" para LLMs. Trae una interfaz rica que puede ser hospedada por cualquiera, con RAG de clase mundial, web search, ejecución de código y generación de artefactos.

## Key Features
- **Agentic RAG**: Calidad de búsqueda y respuesta top (benchmark pendiente vs. Perplexity)
- **Deep Research**: Flujo multi-paso para informes profundos (Top del leaderboard Feb 2026)
- **Custom Agents**: Construí agentes con instrucciones, conocimiento y acciones únicas
- **50+ Connectors**: Indexa via MCP y otros conectores
- **Artifacts**: Genera documentos, gráficos e imágenes descargables
- **Voice Mode**: Text-to-speech y speech-to-text
- **Enterprise Ready**: SSO, RBAC, SCIM, auditoría de consultas

## Quick Start
```bash
# One-line install
curl -fsSL https://onyx.app/install_onyx.sh | bash

# Or deploy manually with Docker
# See: https://docs.onyx.app/deployment/overview

# Visit: https://www.onyx.app (free tier available)
```

## Editions
| Feature | Community (CE) | Enterprise (EE) |
|---------|-----------------|-------------------|
| Chat UI + Agents | ✅ | ✅ |
| RAG (Vector + Keyword) | ✅ | ✅ |
| Web Search + Code Execution | ✅ | ✅ |
| SSO (Google/OIDC/SAML) | ❌ | ✅ |
| RBAC + SCIM | ❌ | ✅ |
| Analytics + Query History | ❌ | ✅ |
| Whitelabeling | ❌ | ✅ |

## Source
- GitHub: https://github.com/onyx-dot-app/onyx
- Docs: https://docs.onyx.app/
- Discord: https://discord.gg/TDJ59cGV2X
- Website: https://www.onyx.app
- Stars: 29K+
- License: MIT (CE), Commercial (EE)
