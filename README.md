# opencode-obsidian

<p align="center">
  <img src="wiki/meta/banner-opencode-obsidian.png" alt="opencode-obsidian" width="100%" />
</p>

[![GitHub stars](https://img.shields.io/github/stars/SamBleed/opencode-obsidian?style=flat&color=e8734a)](https://github.com/SamBleed/opencode-obsidian/stargazers)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![OpenCode](https://img.shields.io/badge/OpenCode-ready-8B5CF6)](https://opencode.ai)

**opencode-obsidian** es un Bunker de Conocimiento Agéntico y Ecosistema de Desarrollo Full-Stack optimizado para **OpenCode**. Basado en el patrón **LLM Wiki** de Andrej Karpathy, este sistema permite que la inteligencia artificial aprenda, documente y opere de forma autónoma en tu entorno local.

---

## ⚡ Quick Start

### 1. Prerrequisitos
Este ecosistema requiere tener instalado **Docker** y el orquestador **n8n** (MECO_Framework Lab). Consulta la [Guía de Setup de n8n](docs/N8N-SETUP.md).

### 2. Instalación del Bunker
```bash
# Levanta el entorno y lanza el agente
bunker-up
```

### 3. Ingestión de Conocimiento
```bash
# Dale una fuente (PDF, TXT, URL) y el Bunker la destilará solo
bunker-ingest "https://articulo-tecnico.com"
```

---

## 🚀 Características Principales

- **Self-Documenting IA**: El sistema ingiere fuentes externas y crea notas estructuradas automáticamente.
- **Agentic Full-Stack**: Incluye una API (Go 1.26) y un Dashboard (React 19) sincronizados en tiempo real.
- **Security-First**: Auditoría nativa con **Trivy MCP** y mitigación de riesgos **OWASP ASI 2026**.
- **Local Sovereignty**: Todo el conocimiento y los datos (PostgreSQL) residen en tu máquina.
- **Live Feedback**: Alertas automáticas al celular vía n8n y Telegram.

---

## 🏢 Ecosistema de Proyectos (/projects)

El Bunker no es solo teoría; incluye implementaciones de referencia "Agent-Native":

- **[OZY-API](projects/ozy-api/)**: Backend robusto (Hexagonal, JWT, Postgres).
- **[OZY-WEB](projects/ozy-web/)**: Frontend moderno (React 19, Tailwind 4).

---

## 🤖 Comandos del Sistema

| Comando | Descripción |
|---------|-----------|
| `bunker-up` | Despierta todo el ecosistema (Docker + Agente). |
| `bunker-ingest [fuente]` | Ingiere y destila archivos o URLs a la wiki. |
| `bunker-push "[msg]"` | Sincroniza Git y te avisa al celular si hubo éxito. |
| `bunker-close` | Cierre de sesión: Handover + sync + push. |
| `/save` | Guarda el insight de la charla actual en la wiki. |
| `trivy scan` | Realiza una auditoría de seguridad sobre el código. |
| `bunker-init` | Workflow Start: Lee hot.md y handover. |
| `bunker-refresh` | Re-alinea: Re-lee pilares para evitar context drift. |

---

## 📖 Documentación Interna

- [Manual de Conexión MCP](wiki/meta/MCP-CONFIG.md) — Cómo enchufar n8n y Trivy.
- [Protocolo de Handover](wiki/meta/HANDOVER.md) — Estándar de transferencia entre agentes.
- [Security Guardrails](wiki/concepts/Security-Guardrails.md) — La constitución del Bunker.

---

## 📜 Créditos

- [Karpathy](https://github.com/karpathy) - Patrón LLM Wiki original.
- **SamBleed** - Arquitectura, Full-Stack y Automatización Agéntica.

---

MIT License © 2026
