---
type: entity
title: "aaif-goose/goose"
domain: ai
created: 2026-05-06
updated: 2026-05-06
tags: [goose, ai-agent, rust, mcp, automation]
status: active
related: [[n8n-io-n8n], [affaan-m-everything-claude-code]]
---

# aaif-goose/goose

## Definition
Agente AI nativo de propósito general desarrollado en Rust. Incluye app de escritorio, CLI completa y API embebible. Parte de la Agentic AI Foundation (AAIF) en la Linux Foundation.

## Key Features
- **Multi-interface**: Desktop app (macOS/Linux/Windows), CLI, API
- **15+ Providers**: Anthropic, OpenAI, Google, Ollama, OpenRouter, Azure, Bedrock
- **70+ Extensions**: Via Model Context Protocol (MCP) standard
- **Built in Rust**: Rendimiento y portabilidad
- **ACP Support**: Usa suscripciones existentes de Claude, ChatGPT o Gemini

## Quick Start
```bash
# Install CLI
curl -fsSL https://github.com/aaif-goose/goose/releases/download/stable/download_cli.sh | bash

# Or download desktop app
# https://goose-docs.ai/docs/getting-started/installation
```

## Architecture
```
goose (Rust core)
├── Desktop App (macOS, Linux, Windows)
├── CLI (Terminal workflows)
└── API (Embed anywhere)
    └── 70+ MCP Extensions
```

## Source
- GitHub: https://github.com/aaif-goose/goose
- Docs: https://goose-docs.ai
- Discord: https://discord.gg/goose-oss
- License: Apache 2.0
- Previously: block/goose (moved to AAIF/Linux Foundation)
