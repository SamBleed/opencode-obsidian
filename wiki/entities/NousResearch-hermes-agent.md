---
type: entity
title: "NousResearch/hermes-agent"
domain: ai
created: 2026-05-06
updated: 2026-05-06
tags: [hermes-agent, nous-research, self-improving, memory, mcp]
status: active
related: [[PicoClaw], [qwen-code], [affaan-m-everything-claude-code]]
---

# NousResearch/hermes-agent

## Definition
Agente AI construido por Nous Research que se **mejora a sí mismo**. Único con un bucle de aprendizaje integrado: crea skills desde la experiencia, se mejora durante el uso, y construye un modelo profundo de quién sos vos a través de sesiones.

## Key Features
- **Self-improving**: Crea skills desde la experiencia, mejora las existentes
- **Closed learning loop**: Memoria con nudges periódicos, búsqueda FTS5 con resúmenes LLM
- **Runs anywhere**: $5 VPS, GPU cluster, o serverless (Daytona/Modal) — hiberna cuando está inactivo
- **Multi-platform**: Telegram, Discord, Slack, WhatsApp, Signal, CLI
- **Scheduled automations**: Cron nativo con entrega a cualquier plataforma
- **Parallelization**: Spawnea subagentes aislados para trabajo paralelo
- **MCP Integration**: Conecta cualquier servidor MCP para extender capacidades

## Quick Start
```bash
# Install
curl -fsSL https://raw.githubusercontent.com/NousResearch/hermes-agent/main/scripts/install.sh | bash

# Start chatting
hermes

# Configure
hermes setup        # Full wizard
hermes model          # Switch models
hermes gateway        # Start messaging gateway
hermes claw migrate   # Migrate from OpenClaw
```

## Tech Stack
| Component | Technology |
|-----------|-----------|
| Platform | TypeScript (Node.js) |
| CLI | Cobra + bubbletea |
| Memory | FTS5 session search + LLM summarization |
| Backends | Docker, SSH, Singularity, Modal, Daytona, Vercel Sandbox |
| Skills | agentskills.io open standard |

## Source
- GitHub: https://github.com/NousResearch/hermes-agent
- Docs: https://hermes-agent.nousresearch.com/docs/
- Discord: https://discord.gg/NousResearch
- Stars: 135K+
- License: MIT
