---
type: concept
title: "MCP (Model Context Protocol)"
complexity: intermediate
domain: mcp-protocol
created: 2026-04-21
updated: 2026-04-21
tags:
  - concept
  - mcp
  - protocol
status: active
related:
  - "[[BrowserMCP]]"
  - "[[MCP-CONFIG]]"
sources:
---

# MCP — Model Context Protocol

Protocolo que permite a aplicaciones AI controlar recursos externos (navegador, archivos, APIs).

---

## Definition

Standard de Anthropic para comunicar AI agents con herramientas externas. El agente puedeinvocar servidores MCP para ejecutar acciones.

## How It Works

1. AI agent se conecta a un MCP server
2. El servidor expone herramientas (tools)
3. El agente invoca herramientas vía protocolo
4. El servidor ejecuta y retorna resultados

## Tipos de MCP Servers

| Server | Función |
|--------|---------|
| BrowserMCP | Controlar navegador |
| File system | Leer/escribir archivos |
| GitHub | Issues, PRs |
| Database | Queries SQL |
| Custom | Cualquier recurso |

## Ejemplo de Herramientas

```json
{
  "name": "browser_navigate",
  "description": "Navega a una URL",
  "inputSchema": {
    "type": "object",
    "properties": {
      "url": {"type": "string"}
    }
  }
}
```

## Comparación

| | REST API | MCP |
|---------|-----|-----|
| Formato | JSON | JSON-RPC |
| State | Stateless | Stateful |
| Tools | No | Yes |
| Streaming | No | Sí |

## En Nuestro Setup

Actualmente tenemos:
- obsidian-vault MCP configurado
- n8n MCP en evaluación

## Source

- https://modelcontextprotocol.io
- https://github.com/BrowserMCP/mcp