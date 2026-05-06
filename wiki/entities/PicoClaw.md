---
type: entity
title: "PicoClaw"
domain: ai
created: 2026-05-06
updated: 2026-05-06
tags: [ai-agent, go, lightweight, mcp, hardware]
status: active
related: [[qwen-code], [hermes-agent], [Pentest-Swarm-AI]]
---

# PicoClaw

## Definition
Asistente AI ultra-ligero escrito en Go por Sipeed. Corre en hardware de $10 con <10MB RAM, 99% menos memoria que OpenClaw. Pensado para despliegues en RISC-V, ARM, MIPS y x86.

## Key Features
- **Ultra-lightweight**: Core <10MB RAM, 99% menor que OpenClaw
- **Lightning-fast boot**: <1s incluso en procesadores de 0.6GHz
- **Truly portable**: Binario único para múltiples arquitecturas
- **AI-bootstrapped**: 95% del código generado por un Agent
- **MCP support**: Integración nativa con cualquier servidor MCP
- **Vision pipeline**: Procesa imágenes y archivos (multimodal)
- **Smart routing**: Modelos livianos para consultas simples (ahorro de costos)

## Quick Install
```bash
# Download from picoclaw.io (auto-detecta plataforma)
# o instalación manual:
curl -fsSL https://github.com/sipeed/picoclaw/releases/download/stable/download_cli.sh | bash
```

## Supported Platforms
| Hardware | Price | Status |
|-----------|-------|--------|
| LicheeRV-Nano | $9.9 | ✅ Tested |
| NanoKVM | $30-50 | ✅ Tested |
| Raspberry Pi Zero 2 W | $15 | ✅ Tested |
| Android Phone | - | ✅ APK Available |

## Architecture
```
PicoClaw (Go core)
├── WebUI Launcher (Desktop)
├── CLI (Terminal workflows)
└── MCP Extensions (70+ via Model Context Protocol)
```

## Source
- GitHub: https://github.com/sipeed/picoclaw
- Website: https://picoclaw.io
- Discord: https://discord.gg/V4sAZ9XWpN
- License: MIT
- Stars: 28K+ (alcanzado en solo 17 días)
- Language: Go 1.25+
