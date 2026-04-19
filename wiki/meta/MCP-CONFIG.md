# Model Context Protocol (MCP) Configuration

> Este documento detalla cómo "enchufar" herramientas externas al cerebro de los agentes para que operen de forma nativa sobre el Bunker.

## 🛡️ Trivy (Security Audit)
Permite al agente realizar escaneos de vulnerabilidades y secretos en lenguaje natural.

**Configuración (JSON)**:
```json
{
  "mcpServers": {
    "trivy": {
      "command": "npx",
      "args": ["-y", "@aquasecurity/mcp-server-trivy"]
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

## 📡 n8n (Advanced AI Orchestration - MCP)
**Estado**: 🛠️ Pendiente de compatibilidad con la versión instalada de OpenCode

**Nota**: la versión instalada en este entorno valida `mcp.n8n-mcp` como un servidor MCP local o deshabilitado. El bloque HTTP anterior rompe el arranque. Dejalo deshabilitado hasta tener soporte explícito para `type: "remote"` en tu binario.

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

---
**Last Updated**: 2026-04-19
