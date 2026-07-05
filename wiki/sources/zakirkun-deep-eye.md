---
type: source
title: "zakirkun/deep-eye"
domain: security
created: 2026-05-05
updated: 2026-05-05
tags: [pentesting, ai, security-scanner, owasp, python]
status: active
related: [[ai-powered-pentesting], [Muhammad-Zakir-Ramadhan]]
---

# zakirkun/deep-eye

## Definition
Repositorio de GitHub que contiene **Deep Eye**, una herramienta de pentesting impulsada por IA que automatiza la detección de vulnerabilidades mediante múltiples proveedores (OpenAI, Grok, OLLAMA, Claude).

## Key Features
- **Multi-AI Provider**: OpenAI, Grok, OLLAMA, Claude, OpenRouter
- **45+ attack methods**: SQLi, XSS, SSRF, XXE, JWT, GraphQL, WebSockets
- **AI-powered payloads**: CVE-aware, context-sensitive generation
- **Professional reporting**: PDF (ReportLab), HTML, JSON con OSINT intelligence
- **Plugin system**: Extensible con escáneres personalizados
- **Multi-channel notifications**: Email, Slack, Discord

## Usage Example
```bash
# Clone and install
git clone https://github.com/zakirkun/deep-eye.git
cd deep-eye
pip install -r requirements.txt

# Configure AI provider
cp config/config.example.yaml config/config.yaml
# Edit config.yaml with API keys

# Run scan
python deep_eye.py -u https://example.com
```

## Configuration Example
```yaml
scanner:
  target_url: "https://example.com"
  ai_provider: "openai"
  default_depth: 2
  default_threads: 5
  enable_recon: true
  full_scan: false
reporting:
  enabled: true
  output_directory: "reports"
  default_format: "html"
```

## Source
- GitHub: https://github.com/zakirkun/deep-eye
- Author: [[Muhammad-Zakir-Ramadhan]]
- License: MIT
- Last commit: 2026-02-03
