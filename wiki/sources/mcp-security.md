---
type: source
title: "MCP Security"
created: 2026-05-14
updated: 2026-05-14
status: active
---

# MCP Security

## Definition
Seguridad en servidores MCP: validar siempre el origen del Bearer Token y restringir el acceso a localhost.

## Source
- `.raw/mcp-security.txt`

## Patterns
- Local-only bind for MCP-related servers
- Bearer token verification before ingest or tool calls
- Keep shared artifacts free of secrets
