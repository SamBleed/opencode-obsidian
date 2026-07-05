---
type: source
title: "QwenLM/qwen-code"
domain: ai
created: 2026-05-06
updated: 2026-05-06
tags: [qwen-code, ai-agent, terminal, coding, alibaba]
status: active
related: [[PicoClaw], [hermes-agent], [Pentest-Swarm-AI]]
---

# QwenLM/qwen-code

## Definition
Agente AI de código abierto que vive en tu terminal, optimizado para modelos de la serie Qwen. Alternativa directa a Claude Code desde China (Alibaba Cloud).

## Key Features
- **Multi-protocol**: Soporta OpenAI / Anthropic / Gemini-compatible APIs
- **Open-source co-evolving**: Framework y modelo Qwen3-Coder evolucionan juntos
- **Agentic workflow**: Skills integradas, SubAgents, experiencia tipo Claude Code
- **Terminal-first**: Diseñado para desarrolladores que viven en la línea de comandos
- **IDE integration**: Soporta VS Code, Zed, JetBrains IDEs

## Quick Start
```bash
# Install (Linux/macOS)
bash -c "$(curl -fsSL https://qwen-code-assets.oss-cn-hangzhou.aliyuncs.com/installation/install-qwen.sh)"

# Windows (Run as Administrator)
powershell -Command "Invoke-WebRequest 'https://qwen-code-assets.oss-cn-hangzhou.aliyuncs.com/installation/install-qwen.bat' -OutFile (Join-Path $env:TEMP 'install-qwen.bat'); & (Join-Path $env:TEMP 'install-qwen.bat')"

# NPM
npm install -g @qwen-code/qwen-code@latest

# Start
qwen
```

## Configuration Example
```json
{
  "modelProviders": {
    "openai": [
      {
        "id": "qwen3.6-plus",
        "name": "qwen3.6-plus",
        "baseUrl": "https://dashscope.aliyuncs.com/compatible-mode/v1",
        "envKey": "DASHSCOPE_API_KEY"
      }
    ]
  },
  "env": {
    "DASHSCOPE_API_KEY": "sk-xxxxxxxxxxxxx"
  },
  "security": {
    "auth": { "selectedType": "openai" }
  },
  "model": { "name": "qwen3.6-plus" }
}
```

## Source
- GitHub: https://github.com/QwenLM/qwen-code
- Docs: https://qwenlm.github.io/qwen-code-docs/
- Stars: 24K+
- License: Open Source
- Language: TypeScript (Node.js 20+)
