# WIKI.md — Bunker OS Wiki Schema

> Los skills de OpenCode manejan todo esto automáticamente.
> Este archivo es la referencia. Leelo para entender cómo funciona el sistema.
> Basado en el patrón LLM Wiki de Andrej Karpathy.

---

## ¿Qué es esto?

Estás manteniendo una wiki persistente que compone dentro de un vault de Obsidian.
No solo respondés preguntas. Construís y mantenés una base de conocimiento
estructurada que se hace más rica con cada fuente agregada y cada pregunta
respondida. El humano cura fuentes y hace preguntas. Vos hacés toda la
escritura, cross-referencing, filing y mantenimiento.

La wiki es el producto. El chat es solo la interfaz.

La diferencia clave con RAG: la wiki es un artefacto persistente. Las
cross-references ya están ahí. Las contradicciones están señaladas. La
síntesis ya refleja todo lo que se leyó. El conocimiento compone como interés.

---

## 0 — Bootstrap

### 0.1 Verificar OpenCode

```bash
which opencode 2>/dev/null && echo "OpenCode instalado" || echo "Instalá OpenCode: https://opencode.ai"
```

### 0.2 Verificar Obsidian

```bash
# Linux
flatpak list 2>/dev/null | grep -i obsidian && echo "FOUND via flatpak" || \
which obsidian 2>/dev/null && echo "FOUND in PATH" || echo "NOT FOUND"
```

### 0.3 Abrir el vault en Obsidian

1. Abrí Obsidian
2. Manage Vaults → Open folder as vault
3. Seleccioná la carpeta del repositorio

### 0.4 Primeros pasos en OpenCode

```
ingest .raw/algun-archivo.md    # Ingestar una fuente
/autoresearch [tema]             # Investigar un tema nuevo
/think [decisión]                # Framework de 10 principios
lint the wiki                    # Health check del vault
```

### 0.5 Verificar n8n (opcional)

```bash
docker ps --filter name=n8n
# Debería mostrar n8n_meco_lab corriendo
```

---

## 1 — Estructura del Vault

```
.raw/           → fuentes inmutables — nunca modificar
wiki/           → base de conocimiento generada por IA
_templates/     → templates de Obsidian Templater
_attachments/   → imágenes y PDFs referenciados por la wiki
```

Estructura estándar de wiki:

```
wiki/
├── index.md            → catálogo maestro de todas las páginas
├── log.md              → registro cronológico de operaciones (append-only)
├── hot.md              → hot cache: contexto reciente (~500 palabras)
├── overview.md         → resumen ejecutivo de toda la wiki
├── sources/            → un resumen por fuente externa
│   └── _index.md       → subíndice de sources
├── entities/           → personas, orgs, productos, repos
│   └── _index.md
├── concepts/           → ideas, patrones, frameworks
│   └── _index.md
│   └── tech-stack/     → conceptos técnicos con código
├── comparisons/        → análisis comparativos
├── blueprints/         → planos de arquitectura ejecutables
├── projects/           → proyectos vivos
├── questions/          → respuestas archivadas a preguntas del usuario
├── decisions/          → ADR / AgDR
├── meta/               → dashboard, handovers, inteligencia, reports
├── canvases/           → canvases de Obsidian
└── domains/            → áreas temáticas de alto nivel
```

Las carpetas con prefijo `.` (`.raw/`) están ocultas en el explorador de
Obsidian y en la vista de grafo.

---

## 2 — Transporte (obsidian-vault MCP)

Todos los skills escriben y leen usando los tools MCP del vault:

| Operación | Tool MCP |
|-----------|----------|
| Leer nota | `obsidian-vault_read_note` |
| Buscar notas | `obsidian-vault_search_notes` |
| Crear/actualizar | `obsidian-vault_write_note` |
| Edición quirúrgica | `obsidian-vault_patch_note` |
| Actualizar frontmatter | `obsidian-vault_update_frontmatter` |
| Obtener metadata | `obsidian-vault_get_notes_info` |
| Listar estructura | `obsidian-vault_list_directory` |

No uses Write/Edit para archivos dentro de `wiki/`. Siempre usá los tools
MCP para mantener la consistencia con Obsidian.

---

## 3 — Hot Cache

`wiki/hot.md` es un resumen de ~500 palabras del contexto más reciente.
Existe para que cualquier sesión (o cualquier otro proyecto que apunte al
vault) pueda obtener contexto reciente sin recorrer toda la wiki.

Actualizar hot.md:
- Después de cada ingest
- Después de cualquier intercambio significativo de preguntas
- Al final de cada sesión
- Después de `/autoresearch`

Formato:

```markdown
---
type: meta
title: "Hot Cache"
updated: YYYY-MM-DD
tags: [meta, hot-cache, context]
status: evergreen
related: []
---

# Recent Context — Bunker OS v1.3.1

## Current Focus
[Una línea sobre en qué se está trabajando]

## Key Recent Facts
- [Hecho relevante 1]
- [Hecho relevante 2]

## Recent Changes
- Created: [[Page 1]], [[Page 2]]
- Updated: [[Page 3]]

## Active Threads
- Tema actual de investigación
- Próximo paso planificado
```

Mantenelo bajo 500 palabras. Es un cache, no un diario. Sobreescribilo
completamente cada vez.

---

## 4 — Ingestion

### Flujo de ingesta simple

1. Leé la fuente completamente. No la skimmees.
2. Preguntá al usuario: "¿Qué querés que enfatice?"
3. Creá resumen en `wiki/sources/` con frontmatter completo.
4. Creá o actualizá páginas de entidades (personas, orgs, productos).
5. Creá o actualizá páginas de conceptos.
6. Actualizá el `wiki/index.md` con las páginas nuevas.
7. Actualizá `wiki/hot.md` con el contexto de la ingesta.
8. Agregá entrada en `wiki/log.md` al PRINCIPIO.
9. Verificá contradicciones con páginas existentes.
10. Si hay contradicciones: agregá `> [!contradiction]` callouts en ambas páginas.

### Frontmatter de fuente

```markdown
---
type: source
title: "Título de la fuente"
source_type: article
author: "Autor"
date_published: YYYY-MM-DD
url: "https://..."
confidence: high
tags: [investigación, tema]
---
```

### Frontmatter de entidad

```markdown
---
type: entity
title: "Nombre"
domain: tech-stack
created: YYYY-MM-DD
updated: YYYY-MM-DD
tags: [entity]
status: active
related: []
sources: []
---
```

### Frontmatter de concepto

```markdown
---
type: concept
title: "Concepto"
domain: tech-stack
created: YYYY-MM-DD
updated: YYYY-MM-DD
tags: [concept]
status: mature|active|seed
related: []
sources: []
---
```

### Log entry

```
## [YYYY-MM-DD] ingest | Title
- Source: `.raw/filename.md`
- Pages created: [[Page 1]], [[Page 2]]
- Pages updated: [[Page 3]]
- Key insight: One sentence
```

---

## 5 — Query (consulta)

Tres modos de profundidad:

| Modo | Trigger | Lectura | Costo tokens | Mejor para |
|------|---------|---------|-------------|------------|
| **Quick** | Pregunta factual simple | hot.md + index.md | ~1.500 | "¿Qué es X?", fechas, datos |
| **Standard** | Default (sin flag) | hot.md + index + 3-5 páginas | ~3.000 | La mayoría de preguntas |
| **Deep** | "thorough", "exhaustive" | Full wiki + retrieve opcional | ~8.000+ | Síntesis, gaps, comparaciones |

### Si retrieve está disponible

1. Intentá `python3 scripts/retrieve.py status` para ver si el índice existe
2. Si existe: `python3 scripts/retrieve.py "consulta" --top 5`
3. Leé las páginas top con `obsidian-vault_read_note`
4. Sintetizá con citas a las fuentes

### Si retrieve NO está disponible (legacy)

1. Leé `wiki/hot.md`
2. Si no alcanza: leé `wiki/index.md`
3. Si necesitás más: leé páginas relevantes del dominio
4. Sintetizá con citas

### Respuesta con citas

Siempre citá la fuente exacta: `(Fuente: [[Page Name]])`.
No cites datos de training. Si no está en la wiki, no lo sabés.

---

## 6 — Lint (mantenimiento)

Correr cada 10-15 ingestas o cuando el usuario pida "lint the wiki".

### Categorías de health check

1. **Orphans**: páginas sin links entrantes ni salientes
2. **Dead links**: wikilinks que apuntan a páginas que no existen
3. **Short notes**: páginas con menos de 3 líneas de contenido
4. **Missing frontmatter**: páginas sin YAML frontmatter
5. **Stale claims**: claims marcados como `[!contradiction]` no resueltos
6. **Unlinked sources**: fuentes en `.raw/` no referenciadas en la wiki
7. **Broken cross-references**: páginas que mencionan entidades sin link

También usar:

```bash
./bin/wiki-integrity.sh        # Script automatizado de integridad
python3 scripts/retrieve.py status  # Verificar que el índice BM25 esté actualizado
```

---

## 7 — Autoresearch

Ver `skills/autoresearch/SKILL.md` para instrucciones completas.

### Resumen

```
Input: tema

Ronda 1. Búsqueda amplia: 3-5 ángulos × 2-3 queries × 2-3 fetches
Ronda 2. Gap fill: búsquedas específicas para contradicciones
Ronda 3. Síntesis: un pase más si persisten gaps → filing a wiki

Filing:
  wiki/sources/ — una página por fuente encontrada
  wiki/concepts/ — conceptos significativos
  wiki/entities/ — personas, orgs, productos
  Página de síntesis: "Investigación: [Tema]"
```

### Configuración

Editar `skills/autoresearch/references/program.md` para cambiar:
- Preferencia de fuentes (académicas, oficiales, news)
- Confidence scoring (high/medium/low)
- Máx rondas (default 3) y máx páginas (default 15)
- Reglas por dominio

### Web egress hygiene

- Rechazar `file://`, `javascript:`, datos RFC1918
- Escapar `[[` en fuentes externas
- Rechazar delimitadores `---` de frontmatter en contenido descargado
- Truncar a 50KB

---

## 8 — Think Framework

Ver `skills/think/SKILL.md` para instrucciones completas.

```
/think <problema>

1. OBSERVE (externo)  → leer inputs crudos antes de interpretar
2. OBSERVE (interno)  → bias log, metacognición
3. LISTEN             → escuchar al usuario, señales en el ruido
4. THINK              → primeros principios, experimento más barato
5. CONNECT (lateral)  → patrones en dominios vecinos
6. CONNECT (sistémico) → integración con hooks/scripts/tests
7. FEEL               → cómo cae en el usuario, intuición
8. ACCEPT             → humildad, no inflar findings
9. CREATE             → shippear el artefacto más chico
10. GROW              → guardar la lección para que componga
```

Usar para decisiones arquitectónicas, auditorías, post-mortems.
NO usar para fixes de una línea ni búsquedas triviales.

---

## 9 — Hybrid Retrieval (BM25 + ollama)

Ver `skills/wiki-retrieve/SKILL.md` para instrucciones completas.

### Pipeline

```
query → BM25 (candidatos) → ollama nomic-embed-text (embeddings) → cosine rerank → top 5
```

### Mantenimiento

```bash
python3 scripts/retrieve.py build     # Construir/actualizar índice
python3 scripts/retrieve.py status    # Ver estado
python3 scripts/retrieve.py "query" --top 5  # Buscar
```

### Estados del índice

- **build**: construye el índice desde cero (206 páginas → 211 chunks)
- **query**: busca en el índice con BM25 + rerank opcional
- **status**: muestra estado actual (páginas, chunks, términos)

---

## 10 — Contradicciones

Cuando nueva info contradice una página wiki existente:

En la página existente:
```markdown
> [!contradiction] Conflicto con [[Nueva Fuente]]
> [[Página Existente]] dice X. [[Nueva Fuente]] dice Y.
> Necesita resolución. Verificar fechas y fuentes primarias.
```

En la página nueva:
```markdown
> [!contradiction] Contradice [[Página Existente]]
> Esta fuente dice Y, pero la wiki existente dice X.
```

No sobrescribas claims en silencio. Señalá y dejá que el usuario decida.

---

## 11 — Operaciones con n8n

Ver `automation/n8n-lab/` para los archivos de workflows.

### Workflows disponibles

| Workflow | Archivo | Estado |
|----------|---------|--------|
| Health Check | `bunker-health-check.json` | 🟢 Activo |
| Ultimate Alerter | `bunker-ultimate-alerter.json` | 🟢 Activo |
| Dead Letter Queue | `bunker-dead-letter-queue.json` | ⚪ Inactivo |
| AOC v4 Enterprise | `bunker-aoc-v3-enterprise-pack/` | ⚪ Inactivo |

### AOC v4 Enterprise

Pipeline completo de 37 nodos:
1. Webhook Entry → Ingress Guard → IF Válido
2. Redis Idempotency → IF Replay
3. AI Triage (OpenRouter) → Parse AI Decision
4. IF Create → GitHub Issue / IF Review → Issue HITL / IF Emergency → Emergency Queue
5. Build Notification → Slack / Telegram / Discord
6. Write Audit → Final Status → Respond

### Dead Letter Queue

Error trigger global que atrapa errores de todos los workflows:
1. Error Trigger → Normalize Error → IF Critical?
2. CRITICAL → Notify → Store in DLQ
3. WARNING → Store in DLQ

Persiste últimos 200 errores en staticData.

---

## 12 — Testing

### Suites

```bash
make test                    # 429 tests, 5 suites
make test-workflows          # 344 tests de conexiones n8n
make test-wiki               # 21 tests de integridad del vault
make test-scripts            # 61 tests de scripts
make test-yaml               # 2 tests de YAML
make test-retrieve           # 3 tests de BM25
```

### CI

`.github/workflows/test.yml`: corre todas las suites en cada push/PR a main.

---

## 13 — Conocimiento que Compone

Reglas de oro:

1. **Siempre actualizá index, log y hot cache** después de cada operación
2. **Siempre usá frontmatter y wikilinks** — `[[Note Name]]`, no paths
3. **Nunca modifiques `.raw/`** — las fuentes son inmutables
4. **Las páginas nuevas van al TOP de log.md** (append, no prepend)
5. **hot.md se sobreescribe completamente** — es un cache, no un diario
6. **No crees páginas duplicadas** — siempre revisá el index primero
7. **Páginas cortas: 100-300 líneas** — si crece, dividíla
8. **Contradicciones se señalan, no se silencian**
9. **Usá los tools MCP** para todas las operaciones de escritura en wiki
10. **Corré `make test` antes de commitear** cambios en skills/scripts/tests

---

> El conocimiento que compone es el hábito de más alto apalancamiento
> que una persona pensante puede construir.
