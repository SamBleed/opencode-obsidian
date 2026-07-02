# 🏛️ opencode-obsidian (Bunker OS)

<p align="center">
  <img src="wiki/meta/banner-opencode-obsidian.png" alt="Bunker OS" width="100%" />
</p>

[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](LICENSE)
[![Version](https://img.shields.io/badge/Version-v1.3.1-blue?style=flat-square)](PROJECT.md)
[![Tests](https://img.shields.io/badge/Tests-429_passing-brightgreen?style=flat-square)](Makefile)
[![CI](https://img.shields.io/badge/CI-GitHub_Actions-2088FF?style=flat-square&logo=githubactions)](.github/workflows/test.yml)
[![n8n](https://img.shields.io/badge/n8n-Automation-8B5CF6?style=flat-square)](automation/n8n-lab)
[![OpenCode](https://img.shields.io/badge/OpenCode-Ready-8B5CF6?style=flat-square)](https://opencode.ai)
[![Obsidian](https://img.shields.io/badge/Obsidian-Vault-7C3AED?style=flat-square&logo=obsidian)](https://obsidian.md)

**Bunker OS** es un sistema operativo de conocimiento local-first construido sobre Obsidian. Convierte sesiones con IA, investigaciones, auditorías, evidencia y decisiones en activos operativos persistentes. Automatización vía n8n, búsqueda híbrida vía BM25 + ollama, y un pipeline completo de orquestación.

Basado originalmente en [Andrej Karpathy's LLM Wiki pattern](https://gist.github.com/karpathy/442a6bf555914893e9891c11519de94f) y fork de [claude-obsidian](https://github.com/AgriciDaniel/claude-obsidian), evolucionado independientemente hacia un sistema de operaciones + automatización.

---

## 📋 Contenido

- [¿Por qué Bunker OS?](#-por-qué-bunker-os)
- [Características](#-características-principales)
- [vs claude-obsidian](#-vs-claude-obsidian)
- [Quick Start](#-quick-start)
- [Comandos](#-comandos)
- [Skills OpenCode](#-skills-opencode)
- [Automatización n8n](#-automatización-n8n)
- [Arquitectura](#-arquitectura)
- [Búsqueda Híbrida](#-búsqueda-híbrida-bm25--ollama)
- [Testing](#-testing)
- [Estructura del Repo](#-estructura-del-repositorio)
- [FAQ](#-faq)
- [Requisitos](#-requisitos)
- [Contribuir](#-contribuir)
- [Licencia](#-licencia)

---

## ✨ ¿Por qué Bunker OS?

| Capability | Bunker OS | claude-obsidian | Smart Connections |
|------------|-----------|-----------------|-------------------|
| **Wiki autoorganizada** | ✅ Crea entidades, conceptos, fuentes | ✅ | ❌ |
| **Automatización n8n** | ✅ Pipeline async + webhooks + DLQ | ❌ | ❌ |
| **Autoresearch** | ✅ 3 rondas web con gap-filling | ✅ | ❌ |
| **Búsqueda híbrida** | ✅ BM25 + cosine rerank (ollama) | ✅ BM25 + API | ❌ |
| **Thinking framework** | ✅ 10 principios en español | ✅ | ❌ |
| **Tests + CI** | ✅ 429 tests, GitHub Actions | ✅ | ❌ |
| **Dead Letter Queue** | ✅ Error handler global | ❌ | ❌ |
| **Evidence vault** | ✅ SHA256 checksums | ❌ | ❌ |
| **Alertas multicanal** | ✅ Slack / Telegram / Discord | ❌ | ❌ |
| **Health monitoring** | ✅ Workflow n8n activo | ❌ | ❌ |
| **Multi-metodología** | ❌ (solo generic) | ✅ LYT/PARA/Zettelkasten | ❌ |
| **Multi-writer safe** | ❌ (single-writer) | ✅ Locks por archivo | ❌ |
| **Open source** | ✅ MIT | ✅ MIT | ⚠️ Freemium |

> Bunker OS no compite con claude-obsidian — **lo complementa**. Donde claude-obsidian es un investigador, Bunker OS es un operador. Ambos pueden coexistir apuntando al mismo vault.

---

## 🚀 Características Principales

### 🧠 Conocimiento
- **Wiki persistente**: 200+ páginas entre conceptos, entidades, fuentes, blueprints y proyectos
- **Autoresearch**: Investigación web autónoma en 3 rondas. Descompone el tema, busca, contrasta fuentes, identifica contradicciones y guarda todo en la wiki
- **Hybrid Retrieval**: BM25 sobre 211 chunks + rerank semántico vía ollama (nomic-embed-text). Búsqueda por significado, no solo palabras exactas
- **/think**: Framework de 10 principios para decisiones arquitectónicas y auditorías (español)
- **Wiki-ingest**: Ingesta de fuentes con extracción automática de entidades y conceptos
- **Wiki-lint**: Health check con 8 categorías (orphans, dead links, frontmatter faltante)

### 🤖 Automatización (n8n)
- **AOC v4 Enterprise**: Pipeline: webhook → AI triage (OpenRouter) → GitHub issues → notificaciones multicanal
- **Dead Letter Queue**: Error trigger global que atrapa errores de TODOS los workflows y los persiste
- **Health Check**: Workflow activo que monitorea el estado del sistema cada 5 min
- **Ultimate Alerter**: Alertas multicanal (Slack / Telegram / Discord)
- **Emergency Reprocessor**: Reintenta automáticamente eventos fallidos cada 5 min

### 🛡️ Operaciones
- **429 tests**: 5 suites, Makefile, GitHub Actions en cada push/PR
- **Evidence Vault**: report.zip y security-audit-report.json indexados con SHA256
- **Integrity Engine**: Scripts automáticos de salud del vault con reportes Markdown + JSON
- **Command Center**: Dashboard + agent queue + handovers + governance
- **Hooks OpenCode**: SessionStart (carga hot.md), PostCompact (recarga tras compactación), Stop (auto-update hot.md)
- **13 skills OpenCode** listos para usar

---

## 🆚 vs claude-obsidian

| Dimensión | Bunker OS | claude-obsidian |
|-----------|-----------|-----------------|
| **Versión** | v1.3.1 | v1.9.2 |
| **Propósito** | Operaciones + automatización + seguridad | PKM / second brain |
| **Harness** | OpenCode | Claude Code (plugin oficial) |
| **Idioma** | Español | Inglés |
| **Skills** | 13 operativas | 15 de investigación |
| **Automatización asíncrona** | ✅ n8n en Docker | ❌ No tiene |
| **Hybrid Retrieval** | ✅ BM25 + ollama rerank | ✅ BM25 + Anthropic API + rerank |
| **Tests** | 429 (5 suites) | ~1.240 (9 suites) |
| **Evidencia con checksums** | ✅ | ❌ |
| **Dead Letter Queue** | ✅ | ❌ |
| **Alertas multicanal** | ✅ Slack/TG/Discord | ❌ |
| **Multi-metodología** | ❌ | ✅ LYT/PARA/Zettelkasten |
| **Multi-writer safe** | ❌ | ✅ |
| **Plugin marketplace** | ❌ | ✅ Claude Code plugin |
| **Stars** | Privado | 8.5k ⭐ |

---

## ⚡ Quick Start

### Opción 1: Clonar como vault (recomendado)

```bash
git clone https://github.com/SamBleed/opencode-obsidian.git
cd opencode-obsidian
```

Abrí la carpeta en **Obsidian**: Manage Vaults → Open folder as vault → seleccioná `opencode-obsidian/`.

Abrí **OpenCode** en la misma carpeta. Empezá con `/autoresearch` o `ingest [archivo]`.

### Opción 2: Usar desde otro proyecto

Agregá a tu `AGENTS.md`:

```markdown
## Wiki Knowledge Base
Path: /path/to/opencode-obsidian

Cuando necesites contexto:
1. Leé wiki/hot.md primero (contexto reciente)
2. Si no alcanza, leé wiki/index.md
3. Solo entonces leé páginas individuales
```

---

## 🎮 Comandos

### Skills OpenCode

| Comando | Qué hace |
|---------|----------|
| `/autoresearch [tema]` | Investigación web autónoma en 3 rondas. Descompone, busca, contrasta, sintetiza y guarda en la wiki |
| `/retrieve [búsqueda]` | Búsqueda híbrida: BM25 + rerank semántico vía ollama |
| `/think [problema]` | Framework de 10 principios para decisiones arquitectónicas |
| `ingest [archivo]` | Ingestar fuente: extrae entidades y conceptos, crea/actualiza páginas |
| `ingest all of these` | Ingesta por lotes con procesamiento paralelo |
| `what do you know about X?` | Consulta la wiki: lee hot → index → páginas → sintetiza con citas |
| `lint the wiki` | Health check: orphans, dead links, gaps, stale claims |
| `save this` | Guarda la conversación actual como nota wiki |
| `/canvas` | Abre o crea canvas visual |
| `/canvas add image [path]` | Agrega imagen al canvas |

### Scripts CLI

```bash
make test                              # 429 tests, 5 suites
./bin/bunker-check.sh                  # Definition of Done local
./bin/wiki-integrity.sh                # Escanear orphans y broken links
./bin/evidence-index.sh                # Indexar evidencia con SHA256
./bin/bunker.sh init                   # Health check del vault
./bin/wiki-sync.sh --apply             # Sync + commit de la wiki
python3 scripts/retrieve.py build      # Reconstruir índice BM25
python3 scripts/retrieve.py "query"    # Búsqueda semántica
```

### /autoresearch: research loop autónomo

Configurable en `skills/autoresearch/references/program.md`:

- Máx rondas: 3
- Máx páginas por sesión: 15
- Preferencia de fuentes (académicas, oficiales, news)
- Confidence scoring (high/medium/low)
- Reglas por dominio

El loop:

1. **Ronda 1, búsqueda amplia**: descompone en 3-5 ángulos, 2-3 búsquedas cada uno vía Exa, fetch de top resultados vía webfetch
2. **Ronda 2, gap fill**: búsquedas específicas para contradicciones y piezas faltantes
3. **Ronda 3, síntesis**: un pase más si persisten gaps. Luego filing a la wiki

URL validation + content sanitization: rechaza `file://`/`javascript://`/RFC1918, escapa `[[` en fuentes externas, trunca a 50KB.

### /think: framework de 10 principios

```
/think <planteo del problema>
```

Camina por 10 etapas: OBSERVE (externo) → OBSERVE (interno) → LISTEN → THINK → CONNECT (lateral) → CONNECT (sistémico) → FEEL → ACCEPT → CREATE → GROW

Usar para decisiones arquitectónicas no triviales, auditorías, post-mortems.

---

## 🧠 Skills OpenCode (13)

| Skill | Descripción | Estado |
|-------|-------------|--------|
| `autoresearch` | Investigación web 3 rondas con Exa + webfetch | ✅ Activo |
| `wiki-retrieve` | BM25 + rerank semántico vía ollama | ✅ Activo |
| `think` | Framework 10 principios para decisiones | ✅ Activo |
| `wiki-ingest` | Ingestar fuentes a la wiki | ✅ Activo |
| `wiki-query` | Consultar la wiki con síntesis | ✅ Activo |
| `wiki-lint` | Health check del vault | ✅ Activo |
| `save` | Guardar conversación como nota wiki | ✅ Activo |
| `code-review` | Revisión de código | ✅ Activo |
| `security-review` | Revisión de seguridad OWASP | ✅ Activo |
| `infra-design` | Diseño de infraestructura | ✅ Activo |
| `tdd-workflow` | TDD con red-green-refactor | ✅ Activo |
| `verification-loop` | Build + test + lint + security pre-PR | ✅ Activo |
| `work-unit-commits` | Commits organizados por unidad de trabajo | ✅ Activo |

---

## 🤖 Automatización n8n

n8n es el "sistema nervioso" asíncrono del Bunker. Corre en Docker y expone un MCP bridge para OpenCode.

### Workflows disponibles

| Workflow | Nodos | Estado | Descripción |
|----------|-------|--------|-------------|
| **Health Check** | 2 | ✅ **Activo** | Verifica estado del sistema cada 5 min |
| **Ultimate Alerter** | 2 | ✅ **Activo** | Alertas multicanal vía webhook |
| **Dead Letter Queue** | 5 | ⚪ Inactivo | Error trigger: captura errores de todos los workflows |
| **AOC v4 Enterprise** | 37 | ⚪ Inactivo | Pipeline: webhook → AI triage → GitHub → notificaciones |

### AOC v4 Enterprise Pipeline

```
Webhook → Ingress Guard → IF Válido
  ├── Inválido → Rejected Response
  └── Válido → Redis Idempotency → IF Replay
       ├── Replay → Cached Response
       └── Nuevo → AI Triage (OpenRouter)
            → Parse AI Decision → IF Create?
                 ├── Crear → GitHub Issue
                 ├── Review → Issue HITL
                 ├── Emergencia → Emergency Queue
                 └── Duplicado → (saltea)
                              ↓
                         Build Notification
                              ↓
                    IF Slack? → Slack / Skip
                         → IF Telegram? → TG / Skip
                         → IF Discord? → DC / Skip
                              ↓
                         Write Audit → Final Status
```

### Dead Letter Queue

```
Error Trigger (global) → Normalize Error → IF Critical?
  ├── CRITICAL → Notify (placeholder) → Store in DLQ
  └── WARNING  → Store in DLQ
```

Atrapa errores de **cualquier** workflow en la instancia. Clasifica severidad (CRITICAL por auth/permisos/timeout). Persiste últimos 200 errores en staticData.

### Infraestructura

```yaml
# docker-compose tuning
N8N_CONCURRENCY_PRODUCTION_LIMIT=10   # Evita OOM en picos
EXECUTIONS_DATA_PRUNE=true            # Limpieza automática
N8N_METRICS=true                      # Observabilidad
N8N_LOG_FORMAT=json                   # Logs estructurados
```

Ambiente: `POSTGRES recomendado`, modo `queue` con Redis para escalado horizontal.

---

## 🏗️ Arquitectura

### Flujo del vault

```
.raw/ (fuentes inmutables)
  │
  ▼
wiki-ingest (skill OpenCode)
  │  Extrae entidades, conceptos, fuentes
  │  Crea/actualiza páginas, cross-references
  ▼
wiki/ (conocimiento persistente)
  ├── hot.md    → contexto reciente (~500 palabras)
  ├── index.md  → catálogo maestro
  ├── log.md    → registro de operaciones
  ├── sources/  → resúmenes de fuentes
  ├── entities/ → personas, orgs, productos
  ├── concepts/ → ideas, patrones, frameworks
  ├── meta/     → dashboard, handovers, evidencia
  └── ...       → comparisons, questions, projects
```

### Búsqueda híbrida

```
query → scripts/retrieve.py "texto"
  │
  ├─ BM25 index (211 chunks, 6.544 términos)  → candidatos sparse
  ├─ ollama nomic-embed-text                   → embeddings densos
  ├─ cosine similarity (0.3 BM25 + 0.7 semántica) → rerank
  │
  ▼
resultados rankeados con: path, score, similarity, preview
```

### Automatización asíncrona

```
OpenCode (skills)
  │  n8n-mcp bridge
  ▼
n8n Docker (localhost:5678)
  │
  ├── Webhooks → AOC pipeline → GitHub Issues
  ├── Schedule → Health Check (5 min)
  ├── Schedule → Emergency Reprocessor (5 min)
  └── Error Trigger → Dead Letter Queue
       │
       ▼
  Slack / Telegram / Discord (notificaciones)
```

---

## 🔍 Búsqueda Híbrida (BM25 + ollama)

El Bunker incluye un sistema de retrieval híbrido que combina búsqueda textual (BM25) con búsqueda semántica (embeddings vía ollama).

### Pipeline

```
query → BM25 sparse (10-20 candidatos) → cosine rerank (nomic-embed-text) → top 5
```

### Rendimiento

- **206 páginas** indexadas
- **211 chunks** (~500 tokens cada uno)
- **6.544 términos** en el vocabulario
- **Rerank**: 0.3 BM25 + 0.7 similitud semántica

### Mantenimiento

```bash
# Construir/actualizar índice
python3 scripts/retrieve.py build

# Ver estado
python3 scripts/retrieve.py status

# Buscar
python3 scripts/retrieve.py "automatización n8n docker" --top 5
```

---

## 🧪 Testing

### Suites

| Suite | Tests | Qué valida |
|-------|-------|------------|
| `test-workflows` | 344 | Todos los JSONs de n8n parseables, conexiones válidas, sin nodos huérfanos |
| `test-wiki` | 21 | Archivos esenciales existen, frontmatter válido, container docker corriendo |
| `test-scripts` | 61 | Sintaxis bash, shebang, ejecutables, sin secretos hardcodeados, go vet |
| `test-yaml` | 2 | YAML del CI y docker-compose válidos |
| `test-retrieve` | 3 | Índice BM25 existe, búsqueda funciona, rerank semántico disponible |

### CI

```yaml
# .github/workflows/test.yml
on: push/PR a main → 5 suites → Python + Go + bash
```

```bash
make test    # 429 tests, 5 suites
```

### Secretos escaneados

El test suite revisa automáticamente:
- OpenAI keys (`sk-...`)
- GitHub PAT (`ghp_...`, `gho_...`)
- AWS keys (`AKIA...`)
- Google API keys (`AIza...`)
- Slack tokens (`xox[baprs]-...`)

---

## 📁 Estructura del Repositorio

```
opencode-obsidian/
├── skills/                       # 13 skills OpenCode (v1.3.1)
│   ├── autoresearch/             # Investigación autónoma 3 rondas
│   ├── wiki-retrieve/            # BM25 + rerank semántico
│   ├── think/                    # Framework 10 principios
│   ├── wiki-ingest/              # Ingesta de fuentes
│   ├── wiki-query/               # Consulta a la wiki
│   ├── wiki-lint/                # Health check del vault
│   ├── save/                     # Guardar conversación
│   ├── code-review/              # Revisión de código
│   ├── security-review/          # Revisión de seguridad
│   ├── infra-design/             # Diseño de infraestructura
│   ├── tdd-workflow/             # TDD red-green-refactor
│   ├── verification-loop/        # Pre-PR checklist
│   └── work-unit-commits/        # Commits por unidad de trabajo
├── agents/                       # 3 agentes OpenCode
│   ├── wiki-ingest.md            # Ingesta paralela
│   ├── wiki-lint.md              # Health check
│   └── evidence-indexer.md       # Indexación de evidencia
├── commands/                     # Slash commands
├── hooks/
│   └── hooks.json                # 7 hooks en 4 eventos
├── scripts/                      # Python + bash
│   ├── bm25-index.py             # BM25 indexer
│   ├── retrieve.py               # Orquestador de búsqueda
│   └── ...
├── tests/                        # 5 suites, 429 tests
│   ├── test_workflow_connections.py
│   ├── test_wiki_integrity.sh
│   ├── test_scripts.sh
│   └── Makefile
├── bin/                          # 18 scripts operativos
│   ├── bunker-check.sh           # Health check completo
│   ├── wiki-integrity.sh         # Escaneo de integridad
│   ├── evidence-index.sh         # Indexación de evidencia
│   └── ... (15 más)
├── automation/
│   └── n8n-lab/                  # Docker + workflows
│       ├── docker-compose.yml    # n8n container
│       ├── .env                  # Variables de entorno
│       └── workflows/            # JSONs de workflows
├── wiki/                         # Vault Obsidian (200+ páginas)
│   ├── hot.md                    # Contexto reciente
│   ├── index.md                  # Catálogo maestro
│   ├── log.md                    # Registro de operaciones
│   ├── sources/                  # Fuentes externas
│   ├── entities/                 # Personas, orgs, productos
│   ├── concepts/                 # Conceptos y patrones
│   ├── meta/                     # Dashboard, handovers, evidencia
│   ├── blueprints/               # Planos de arquitectura
│   └── ...                       # Comparisons, questions, projects
├── .raw/                         # Fuentes inmutables
├── .github/workflows/test.yml    # CI: corre en push/PR
├── Makefile                      # 5 targets de test
├── README.md                     # Este archivo
├── PROJECT.md                    # Documentación técnica
├── BUNKER_RULES.md               # Reglas de gobierno
├── AGENTS.md                     # Instrucciones para agentes
└── CHANGELOG.md                  # Historial de versiones
```

---

## ❓ FAQ

**¿Qué diferencia hay entre Bunker OS y claude-obsidian?**
Ambos proyectos comparten origen pero divergieron en propósito. claude-obsidian es un second brain orientado a PKM/investigación. Bunker OS es un sistema de operaciones con automatización (n8n), seguridad, alerts, evidencia y un pipeline completo de remediación. Se complementan.

**¿Puedo usar Bunker OS sin OpenCode?**
Técnicamente sí (los scripts son bash, los tests son Makefile, n8n es Docker), pero el valor real está en los skills de OpenCode. Sin OpenCode perdés autoresearch, think, wiki-ingest y wiki-query.

**¿Dónde se guardan los datos?**
Todo es local. El vault es una carpeta de archivos Markdown en tu disco. n8n corre en Docker local. No hay datos en la nube.

**¿Cómo sincronizo entre dispositivos?**
El vault es una carpeta de archivos. Usá Obsidian Sync, Syncthing, iCloud, Dropbox o git.

**¿Cómo agrego una fuente a la wiki?**
Dejá el archivo en `.raw/` y decí `ingest [filename]`. El skill extrae entidades, conceptos, crea páginas y actualiza índices.

**¿Qué es la Dead Letter Queue?**
Un workflow global que atrapa errores de cualquier otro workflow en n8n. Clasifica severidad y persiste en staticData para revisión posterior.

**¿Cómo activo el AOC v4 Enterprise?**
Requiere: API key de OpenRouter, webhook URL de Discord, y credenciales de GitHub en n8n. Configurá las env vars en `.env`, importá el JSON en la UI de n8n, y activalo.

**¿Cómo actualizo el índice BM25 después de agregar páginas?**
```bash
python3 scripts/retrieve.py build
```

**¿Dónde están los tests?**
```bash
make test    # 429 tests
```

**¿Qué hago si se pierde el contexto del agente después de una compactación?**
Los hooks PostCompact se encargan de recargar `hot.md` automáticamente.

---

## 📋 Requisitos

| Componente | Mínimo | Notas |
|------------|--------|-------|
| OpenCode | latest | https://opencode.ai |
| Obsidian | v1.6+ | Cualquier versión moderna |
| Python | 3.10+ | Para BM25 index + retrieve |
| Bash | 4.0+ (o zsh) | Para scripts operativos |
| Docker | latest | Para n8n |
| Git | cualquiera | Para versionado del vault |

**Opcional:**
- **ollama** + `nomic-embed-text` (para rerank semántico en retrieve)
- **n8n** en Docker (para automatización de workflows)
- **API key OpenRouter** (para AI triage en AOC v4)

---

## 🤝 Contribuir

PRs bienvenidas. Leé antes:

- [`PROJECT.md`](PROJECT.md): documentación técnica
- [`BUNKER_RULES.md`](BUNKER_RULES.md): governance y estándares
- [`CHANGELOG.md`](CHANGELOG.md): historial de versiones

```bash
make test    # Corré los tests antes de pushear
```

---

## 📜 Licencia

MIT License. Ver [LICENSE](LICENSE) para el texto completo. Libre para uso personal y comercial.

---

<p align="center">
  <i>Basado en el <a href="https://gist.github.com/karpathy/442a6bf555914893e9891c11519de94f">LLM Wiki pattern</a> de Andrej Karpathy.</i>
  <br>
  <i>Construido por <a href="https://github.com/SamBleed">SamBleed</a> para <a href="https://opencode.ai">OpenCode</a>.</i>
  <br>
  <i>El conocimiento que compone es el hábito de más alto apalancamiento que una persona puede construir.</i>
</p>
