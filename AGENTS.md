# OpenCode-Bunker — Wiki Vault

> [!important] INSTRUCCIONES DE INICIO (MANDATORIO)
> Al iniciar cualquier sesión en este repositorio, el agente DEBE:
> 1. Leer `wiki/hot.md` inmediatamente para cargar el contexto activo.
> 2. Leer el último archivo de handover en `wiki/meta/handovers/` para entender el estado de la última sesión.
> 3. Actuar bajo el **Protocolo de Proactividad** detallado abajo.

Este directorio es un vault de Obsidian con skills adaptados para OpenCode.

## Vault Path
```
${BUNKER_HOME:-<path-to-opencode-obsidian>}
```

## Estructura
```
.raw/           source documents — immutable, nunca modificar
wiki/           knowledge base generado por IA
_templates/     templates de Obsidian
_attachments/   imágenes y PDFs
```

## Comportamiento Proactivo (MANDATORIO)

Como agente de este bunker, NO esperes órdenes para tareas de mantenimiento. Debes actuar proactivamente en los siguientes casos:

1.  **Detección de Fuentes**: Al iniciar, revisa `.raw/`. Si hay archivos nuevos no registrados en `wiki/meta/ingestion-log.md`, ofrece procesarlos inmediatamente.
2.  **Auto-Sincronización**: Tras completar una tarea significativa (cambio de código, nueva nota wiki), ejecuta `bin/wiki-sync.sh` sin preguntar.
3.  **Prevención de Deriva**: Si la conversación se vuelve ambigua, re-lee `wiki/hot.md` para realinear objetivos.
4.  **Cierre Inteligente**: Al detectar el fin de la sesión, genera el Handover y el resumen de Engram proactivamente.
5.  **Mantenimiento de Links**: Si creas una nota, busca proactivamente otras notas relacionadas para crear vínculos `[[ ]]`.

## Workflow de Ingestión (MANDATORIO)

Cuando recibas un recurso (URL, archivo, tema) para el vault:

### Ubicación exacta por tipo

| Tipo de recurso | Folder destino | Archivo ejemplo |
|---------------|--------------|-----------------|
| Entidad (persona, org, producto) | `wiki/entities/` | `wiki/entities/[name].md` |
| Concepto técnico con código | `wiki/concepts/tech-stack/` | `wiki/concepts/tech-stack/[name].md` |
| Concepto general | `wiki/concepts/` | `wiki/concepts/[name].md` |
| Fuente externa | `wiki/sources/` | `wiki/sources/[slug].md` |

### Pasos obligatorios

1. **Crear página source** en `wiki/sources/` (resumen del recurso)
2. **Crear entidad** en `wiki/entities/` si es persona/org/producto
3. **Crear concepto** en folder correcto (`tech-stack/` si tiene código/patrones)
4. **Incluir ejemplos de código** — no solo descripción
5. **Actualizar `wiki/index.md`** — agregar entrada en sección correspondiente
6. **Actualizar `wiki/log.md`** — prepend operation con fecha
7. **Agregar al Canvas** — crear nodo y edge en `wiki/Wiki Map.canvas`

### NO hacer

- ❌ Guardar solo en Engram (`mem_save`) — eso se pierde entre sesiones
- ❌ Crear en `wiki/concepts/` cuando corresponde `wiki/concepts/tech-stack/`
- ⡪omitir el canvas — el grafo visual es parte del vault
- ❌ Solo descripción — incluir código/patrones prácticos

### Formato de Wiki Pages

```markdown
---
type: [entity|concept|source]
title: "[Title]"
domain: [tech-stack|design|ai|etc]
created: YYYY-MM-DD
updated: YYYY-MM-DD
tags: []
status: [active|mature|seed]
related: []
sources: []
---

# Title

## Definition
[Qué es, en una línea]

## How It Works
[Mecanismo]

## Example
\`\`\`lenguaje
código aquí
\`\`\`

## Patterns
- [Patrón 1]
- [Patrón 2]

## Source
- URL
- [[Context7]]
```

### Regla de oro

> NO考虑一下 si "esta página será consultable por otra sesión en 3 meses". Si la respuesta es sí, crear la wiki page correctamente.

## Skills Disponibles

| Skill | Trigger |
|-------|---------|
| `wiki` | Setup, scaffold, verificar estado |
| `ingest [archivo]` | Ingestión de fuentes |
| `query [pregunta]` | Responder desde el wiki |
| `search the wiki` | Buscar en el vault |
| `lint` | Health check |
| `save this` | Guardar conversación como nota |
| `bunker init` | **Workflow Start**: Lee hot.md y handover. |
| `bunker refresh` | **Re-align**: Re-lee pilares para evitar context drift. |
| `bunker save` | **Checkpoint**: Commits rápidos de la wiki. |
| `bunker close` | **Session End**: Handover, sync y push. |

## Cómo Usar

1. **Ingestión**: "ingest ~/Documents/fuente.pdf"
2. **Consulta**: "what do you know about X?"
3. **Guardar**: "save this" o "/save"
4. **Mantenimiento**: "lint the wiki"

## Acceso Cross-Project

Para referenciar este wiki desde otro proyecto, agregar al AGENTS.md de ese proyecto:

```markdown
## Wiki Knowledge Base
Path: ${BUNKER_HOME:-<path-to-opencode-obsidian>}

Cuando necesites contexto:
1. Leé wiki/hot.md primero (~500 palabras)
2. Si no alcanza, leé wiki/index.md
3. Solo entonces leé páginas individuales

NO leer el wiki para preguntas de coding generales.
```

## MCP (Configurados)

1. **obsidian-vault** (`~/.config/opencode/opencode.json`):
   - `list_notes`, `read_note`, `search_notes`, `write_note`
2. **n8n-mcp**:
   - Expone flujos de `/automation/n8n-lab` a OpenCode para ejecución segura y asíncrona.

## Tool Governance (MANDATORIO)

Para evitar solapación y confusión cognitiva entre herramientas MCP:

1.  **Búsqueda e Ingestión**:
    - **Context7**: Fuente PRIMARIA para documentación de librerías (React, Go, etc.).
    - **Exa**: Fuente SECUNDARIA para noticias, tendencias o conceptos genéricos.
    - **Firecrawl** (si se instala): SOLO para extraer contenido profundo de una URL específica cuando Exa no tiene el detalle.
2.  **Acción y Estado**:
    - **BrowserMCP**: SOLO para interactuar con la UI (clics, formularios) o validar visualmente.
    - **GitHub MCP**: Para gestión de repo (PRs, Issues, Estructura). No usar bash para esto si el MCP puede hacerlo.
    - **n8n MCP**: Para disparar webhooks o flujos asíncronos (alertas, remediación asistida). NUNCA hardcodear bash scripts para tareas que n8n ya resuelve.
3.  **Seguridad y Auditoría**:
    - **Trivy**: Usar SOLO bajo demanda explícita de "Auditoría de Seguridad" o mediante el script `bin/audit-repo.sh`.
4.  **Memoria**:
    - **Engram**: Guardar TODA decisión arquitectónica, fix de bug o patrón establecido.

## Referencias

- skills adaptados: `~/.config/opencode/skills/wiki*`
- vault config: `~/.config/opencode/skills/wiki/references/vault-config.md`
