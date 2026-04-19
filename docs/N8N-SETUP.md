# n8n Orchestrator Setup Guide

> Esta guía explica cómo configurar el cerebro de automatización del Bunker usando n8n en Docker.

## 🏗️ 1. Levantar el Laboratorio (Docker)
El Bunker utiliza una instancia local de n8n para orquestar alertas y tareas en segundo plano.

### Requisitos:
- Docker y Docker Compose instalados.
- Arch Linux o sistema compatible con Bash.

### Pasos:
1. Navega a la carpeta de tu lab: `cd ~/n8n-lab/workflows`.
2. Ejecuta el script de inicio: `./setup.sh`.
3. Accede a `http://localhost:5678`.

## 📡 2. Conectar el Bunker (Webhook)
Para que el agente pueda enviarte alertas:
1. Crea un nuevo Workflow en n8n.
2. Importa el archivo JSON: `wiki/blueprints/N8N-SECURITY-ALERT.md` (o el Ultimate Alerter).
3. Configura tus credenciales de Telegram o Slack.
4. **IMPORTANTE**: Pon el switch del workflow en **ACTIVE** (arriba a la derecha).

## 🔌 3. Configurar el Servidor MCP
Para que el agente "vea" tus flujos de n8n de forma nativa:
1. En n8n, ve a **Settings** -> **AI** y habilita **"Enable instance-level MCP access"**.
2. Genera un API Key en **Settings** -> **Public API**.
3. Añade el siguiente bloque a tu `opencode.json`:

```json
"n8n-mcp": {
  "type": "http",
  "url": "http://localhost:5678/mcp-server/http",
  "headers": {
    "Authorization": "Bearer <TU_TOKEN>"
  }
}
```

## 🛠️ Troubleshooting (Auto-Healing)
Si te olvidas la contraseña de administrador de n8n:
```bash
docker exec -u node n8n_meco_lab n8n user-management:reset
docker restart n8n_meco_lab
```

---
**Last Updated**: 2026-04-19
