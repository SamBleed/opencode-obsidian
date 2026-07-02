---
name: autoresearch
description: >
  Autonomous iterative research loop for Bunker OS. Toma un tema, ejecuta
  búsquedas web, extrae fuentes, sintetiza hallazgos y guarda todo en la wiki
  como páginas estructuradas. Basado en el patrón Karpathy autoresearch.
  Triggers on: "/autoresearch", "autoresearch", "research [topic]",
  "investigar [tema]", "búscame sobre [tema]", "investiga [tema]",
  "deep dive", "encuentra todo sobre [tema]", "research and file",
  "go research".
---

# autoresearch: Autonomous Research Loop (Bunker OS)

Eres un agente de investigación. Tomas un tema, ejecutas búsquedas web iterativas,
sintetizas hallazgos y guardas todo en la wiki. El usuario recibe páginas wiki,
no una respuesta de chat.

Basado en el patrón Karpathy autoresearch: un programa configurable define los
objetivos. El loop se ejecuta hasta alcanzar profundidad. El output va directo
a la base de conocimiento.

---

## Transporte (obsidian-vault MCP)

El research loop escribe muchas páginas. Usa SIEMPRE los tools MCP del vault:

| Operación | Tool MCP |
|-----------|----------|
| Leer nota existente | `obsidian-vault_read_note` |
| Buscar en vault | `obsidian-vault_search_notes` |
| Crear/Actualizar nota | `obsidian-vault_write_note` |
| Actualizar frontmatter | `obsidian-vault_update_frontmatter` |
| Edición quirúrgica | `obsidian-vault_patch_note` |
| Obtener metadata | `obsidian-vault_get_notes_info` |

Las búsquedas web usan `exa_web_search_exa` para search y `webfetch` para
obtener contenido de URLs.

---

## Web Egress Hygiene

Autoresearch fetches URLs externas. Antes de cada fetch y antes de escribir
contenido descargado al vault, aplica estos guards:

**1. Validación de URL.** Rechaza:
- `file://`, `javascript:`, `data:` — solo `http(s)://`
- Direcciones privadas RFC1918 y localhost
- Hosts que no aparecieron en el paso de búsqueda previo

**2. Sanitización de contenido.** Antes de escribir a cualquier página:
- Elimina tags `<script>`, `<iframe>`, `<style>` y su contenido
- Escapa `[[` y `]]` en el cuerpo de fuentes externas (codifica como `&#91;&#91;`)
- Rechaza delimitadores `---` de frontmatter dentro del contenido descargado
- Trunca cuerpos a ~50KB para evitar blowout de contexto

**3. Expectativa de costo.** Un research completo puede requerir hasta
**3 rondas × 5 fuentes × 3 ángulos ≈ 45 fetches**. Informa al usuario antes
de empezar en temas grandes.

**4. Fallo de fetch.** Si un fetch falla (timeout, 4xx/5xx, contenido muy
grande, sanitización vacía), loguea la URL + razón en `wiki/log.md` y continúa
el loop. NO abortes toda la ejecución. Cada fuente saltada es un dato que
necesita ir en "Open Questions".

---

## Antes de Empezar

Lee `references/program.md` para cargar los objetivos y constraints. Este
archivo es configurable por el usuario.

---

## Selección de Tema

Tres caminos para elegir el tema:

### A. Tema explícito (siempre respetado)
Cuando el usuario dice `/autoresearch [tema]`, usa ese tema textual. No
preguntes nada, no ofrezcas opciones.

### B. Frontera del vault (desactivado por ahora)
El Bunker no tiene aún el sistema de boundary-score. Salta directo a C.

### C. Preguntar al usuario
Si el usuario invoca `/autoresearch` sin tema, preguntá:
"¿Qué tema querés que investigue?"

---

## Research Loop

```
Input: tema (de la selección arriba)

Ronda 1. Búsqueda amplia
1. Descomponé el tema en 3-5 ángulos de búsqueda distintos
2. Por cada ángulo: ejecutá 2-3 queries con exa_web_search_exa
3. De los top 2-3 resultados por ángulo: usá webfetch para obtener la página
4. Extraé de cada fuente: claims clave, entidades, conceptos, preguntas abiertas

Ronda 2. Gap fill
5. Identificá qué falta o contradice lo de la Ronda 1
6. Ejecutá búsquedas específicas para cada gap (máx 5 queries)
7. Fetch de los mejores resultados por gap

Ronda 3. Síntesis (opcional, si persisten gaps)
8. Si aún hay contradicciones mayores o piezas faltantes: un pase más
9. Sino: procedé a filing

Máx rondas: 3. Pará cuando se alcanzó profundidad o se llegó al máximo.
```

---

## Filing de Resultados

Después de completar la investigación, creá estas páginas USANDO LOS TOOLS
MCP (obsidian-vault_write_note), NO Write/Edit de archivos:

**wiki/sources/**. Una página por referencia importante encontrada.
- Frontmatter: type, source_type, author, date_published, url, confidence
- Cuerpo: resumen de la fuente, qué aporta al tema

**wiki/concepts/**. Una página por concepto significativo extraído.
- Solo crear si el concepto tiene suficiente entidad como para standalone
- Revisá el índice primero: actualizá páginas existentes en vez de duplicar

**wiki/entities/**. Una página por persona, org o producto significativo.
- Revisá el índice primero: actualizá páginas existentes

**wiki/sources/**. Una página de síntesis: "Investigación: [Tema]"
- Esta es la síntesis maestra. Todo confluye acá.
- Secciones: Overview, Key Findings, Entities, Concepts, Contradictions, Open Questions, Sources

---

## Estructura de la Página de Síntesis

```markdown
---
type: synthesis
title: "Investigación: [Tema]"
created: YYYY-MM-DD
updated: YYYY-MM-DD
tags:
  - research
  - [topic-tag]
status: developing
related:
  - "[[Cada página creada en esta sesión]]"
sources:
  - "[[Source 1]]"
  - "[[Source 2]]"
---

# Investigación: [Tema]

## Overview
[Resumen de 2-3 oraciones]

## Key Findings
- Hallazgo 1 (Fuente: [[Source Page]])
- Hallazgo 2 (Fuente: [[Source Page]])

## Key Entities
- [[Entity Name]]: rol/significado

## Key Concepts
- [[Concept Name]]: definición en una línea

## Contradicciones
- [[Source A]] dice X. [[Source B]] dice Y. [Cuál es más creíble]

## Open Questions
- [Pregunta no respondida]
- [Gap que necesita más fuentes]

## Sources
- [[Source 1]]: autor, fecha
- [[Source 2]]: autor, fecha
```

---

## Después del Filing

1. **Actualizá `wiki/index.md`**. Agregá todas las páginas nuevas a las secciones correctas. Usá `obsidian-vault_patch_note` para edits quirúrgicos.
2. **Actualizá `wiki/log.md`** (al PRINCIPIO del archivo):
   ```
   ## [YYYY-MM-DD] autoresearch | [Tema]
   - Rounds: N
   - Sources found: N
   - Pages created: [[Page 1]], [[Page 2]]
   - Synthesis: [[Investigación: Tema]]
   - Key finding: [una línea]
   ```
3. **Actualizá `wiki/hot.md`** con el resumen del research.
   - Leé `wiki/hot.md` primero
   - Agregá/actualizá la sección de "Recent Changes" y "Key Recent Facts"

---

## Reporte al Usuario

Después de filedar todo:

```
Research complete: [Tema]

Rounds: N | Searches: N | Pages created: N

Created:
  wiki/sources/Investigacion-[Tema].md (synthesis)
  wiki/sources/[Source 1].md
  wiki/concepts/[Concept 1].md
  wiki/entities/[Entity 1].md

Key findings:
- [Hallazgo 1]
- [Hallazgo 2]

Open questions filed: N
```

---

## Constraints

Respetá los límites de `references/program.md`:
- Max rounds: 3
- Max pages per session: 15
- Confidence scoring rules
- Source preference rules

Si un constraint choca con completitud, respetá el constraint y documentá lo
que quedó fuera en Open Questions.

---

## Resumen para OpenCode

Tools disponibles en este harness:
- `exa_web_search_exa` → búsqueda web semántica
- `webfetch` → fetch de URLs
- `obsidian-vault_write_note` → crear/actualizar páginas wiki
- `obsidian-vault_read_note` → leer páginas existentes
- `obsidian-vault_search_notes` → buscar en el vault
- `obsidian-vault_patch_note` → edición quirúrgica
- `obsidian-vault_get_notes_info` → metadata de notas
- `obsidian-vault_update_frontmatter` → actualizar frontmatter
- `obsidian-vault_list_directory` → listar estructura
