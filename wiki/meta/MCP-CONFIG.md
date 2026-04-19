# Model Context Protocol (MCP) Configuration

> Este documento detalla cómo "enchufar" herramientas externas al cerebro de los agentes para que operen de forma nativa sobre el Bunker.

## 🛡️ Trivy (Security Audit)
Permite al agente realizar escaneos de vulnerabilidades y secretos en lenguaje natural.

**Configuración (JSON)**:
```json
{
  "mcpServers": {
    "trivy": {
      "command": "trivy",
      "args": ["mcp"]
    }
  }
}
```

## 🛠️ Cómo Activar
1. Ubica tu archivo de configuración del agente (ej. `~/.config/opencode/opencode.json` o la configuración de Goose).
2. Copia y pega el bloque `mcpServers` de arriba.
3. Reinicia el agente.
4. Para verificar, dile al agente: *"Trivy, haz un escaneo del sistema de archivos en el directorio actual"*.

## 🚀 Próximos Servidores MCP Detectados
- **n8n MCP**: Para disparar workflows directamente.
- **GitHub MCP**: Para gestión avanzada de PRs e Issues.
- **Postgres MCP**: Para consultas directas a la base de datos de producción/staging.

## Orden Recomendado de Exposición
Ver [[n8n-mcp-exposure-plan]] para decidir qué workflows abrir primero a `n8n-mcp`.

## 📡 n8n (Advanced AI Orchestration - MCP)
**Estado**: deshabilitado en la config local

**Nota**: si `localhost:5678/mcp-server/http` no responde, OpenCode puede mostrar `SSE error: Non-200 status code (404)` o fallar por conexión cerrada. Dejalo deshabilitado hasta que el servicio n8n esté realmente escuchando en ese endpoint y el transporte sea compatible con tu binario.

**Copia y pega este bloque en tu `opencode.json` local**:
```json
{
  "mcpServers": {
    "n8n-mcp": {
      "enabled": false
    }
  }
}
```

## Observación sobre Trivy
El bloque anterior que usaba `npx -y @aquasecurity/mcp-server-trivy` quedó obsoleto. El repositorio oficial actual de Aqua publica el plugin MCP en `aquasecurity/trivy-mcp`, y el arranque documentado es `trivy plugin install mcp` seguido de `trivy mcp`.

---
**Last Updated**: 2026-04-19
