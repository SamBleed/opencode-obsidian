# OpenCode-Bunker — Wiki Vault

Este directorio es un vault de Obsidian con skills adaptados para OpenCode.

## Vault Path
```
/home/sam/opencode-obsidian
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
Path: /home/sam/opencode-obsidian

Cuando necesites contexto:
1. Leé wiki/hot.md primero (~500 palabras)
2. Si no alcanza, leé wiki/index.md
3. Solo entonces leé páginas individuales

NO leer el wiki para preguntas de coding generales.
```

## MCP (Configurado)

El MCP `obsidian-vault` está configurado en `~/.config/opencode/opencode.json`.
Herramientas disponibles cuando esté activo:
- `list_notes`
- `read_note`
- `search_notes`
- `write_note`

## Referencias

- skills adaptados: `~/.config/opencode/skills/wiki*`
- vault config: `~/.config/opencode/skills/wiki/references/vault-config.md`