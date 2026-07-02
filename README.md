# 🏛️ opencode-obsidian (Bunker OS)

<p align="center">
  <img src="wiki/meta/banner-opencode-obsidian.png" alt="Bunker OS" width="100%" />
</p>

**Bunker OS** es un **sistema operativo de conocimiento local-first** construido sobre Obsidian. Convierte tus sesiones con IA, investigaciones, auditorías y decisiones en un **centro de comando operativo persistente**.

[![Obsidian](https://img.shields.io/badge/Obsidian-Plugin_Ready-7C3AED?style=flat-square&logo=obsidian)](https://obsidian.md)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](LICENSE)
[![Version](https://img.shields.io/badge/Version-v3.0-purple?style=flat-square)](https://github.com/SamBleed/opencode-obsidian)

---

## ✨ ¿Por qué Bunker OS?

- Transforma conocimiento efímero (chats con IA, notas sueltas) en **activos operativos**.
- Combina **Obsidian + Automatización (n8n) + Agentes + Governance**.
- Totalmente local, privado y bajo tu control.

<p align="center">
  <img src="wiki/meta/descripcion.png" alt="Bunker OS Workflow" width="85%" />
</p>

---

## 🚀 Características Principales (v3.0)

- **Deep Capacity Wiki** — Patrones avanzados y notas reutilizables
- **Async Automation Brain (n8n)** — Webhooks + MCP bridge para ejecución agentica
- **Operational Scaffolding** — `bunker-scaffold.sh` genera proyectos listos para producción
- **Evidence Vault** — Evidencia inmutable con checksums
- **Self-Healing Integrity** — Scripts automáticos de salud del vault
- **Command Center** — Dashboard + canvases + agent queue

---

## 🛠️ Cómo Empezar (Installation)

1. Clona el repositorio:
   ```bash
   git clone https://github.com/SamBleed/opencode-obsidian.git
   ```

2. Abre la carpeta como vault en **Obsidian**
3. (Opcional) Instala los plugins recomendados desde `.obsidian` folder
4. Explora `wiki/meta/dashboard.md` — es tu centro de operaciones

---

## 🛠️ Power Commands

| Comando                        | Propósito                          |
|-------------------------------|------------------------------------|
| `./bin/bunker-scaffold.sh`    | Generar nuevos proyectos           |
| `./bin/wiki-integrity.sh`     | Chequear salud del vault           |
| `./bin/bunker-check.sh`       | Definition of Done local           |
| `go run bin/ingest_server.go` | Iniciar Smart Ingest API           |

---

## 📚 Documentación

- **[[overview]]** → Resumen del sistema
- **[[BUNKER-OS.canvas]]** → Mapa visual completo
- **wiki/meta/dashboard.md** → Command Center

---

**MIT License © 2026** — [SamBleed](https://github.com/SamBleed) &  [opencode.ai](https://opencode.ai)
