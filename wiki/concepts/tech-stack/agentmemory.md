---
type: concept
title: "AgentMemory"
domain: tech-stack
created: 2026-05-08
updated: 2026-05-08
tags: [memory, ai, mcp, iii-engine]
status: active
related: [Engram, OpenCode, iii-engine]
sources: [https://github.com/rohitg00/agentmemory]
---

# AgentMemory

## Definition
AgentMemory es el motor de memoria persistente de largo plazo para el Bunker. Utiliza el patrón **LLM Wiki** con una arquitectura de 4 niveles (Working, Episodic, Semantic, Procedural) para eliminar la amnesia de los agentes de IA.

## How It Works
Corre sobre el motor `iii-engine` y se comunica mediante un servidor **MCP (Model Context Protocol)**. Captura automáticamente el uso de herramientas, decisiones y bugs, indexándolos mediante búsqueda híbrida (BM25 + Vectores).

## Example (Integration in Go)
```go
// Llamada al binario MCP desde el Bunker Ingest Server
cmd := exec.Command("/home/sam/.npm-global/bin/agentmemory-mcp")
cmd.Stdin = bytes.NewBuffer(jsonReq)
// ... procesamiento de la respuesta semántica
```

## Patterns
- **Reactive Auditing (Sentinel):** Escaneo de código en el momento de la ingesta de una decisión.
- **Auto-Documentation (Pulse):** Reconstrucción automática de `BUNKER_RULES.md` basada en memorias de arquitectura.

## Implementation in this Bunker
- **Server Port:** 3111 (API), 3113 (Viewer)
- **MCP Binary:** `/home/sam/.npm-global/bin/agentmemory-mcp`
- **Integrations:** `ingest_server.go` (v2.1+), `bunker-pulse.sh` (v4.0+)

## Source
- [GitHub Repository](https://github.com/rohitg00/agentmemory)
- [[Engram]] (Memoria complementaria)
