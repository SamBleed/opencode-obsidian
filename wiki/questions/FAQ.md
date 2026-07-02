---
type: question
title: "FAQ"
updated: 2026-05-14
status: active
---

# FAQ — Preguntas Frecuentes

> Colección de preguntas recurrentes en sesiones de pair-programming. El agente consulta este archivo antes de responder preguntas comunes.

---

## 🧠 Contexto & Memoria

### ¿Cómo reinicio el contexto de Engram?

```bash
# Busca el último session_summary
mem_search(query: "session_summary", project: "opencode-obsidian")

# O regenera desde el último handover
mem_context()
```

### ¿Cómo sabe el agente qué hicimos antes?

1. Lee `wiki/hot.md` → contexto activo de la sesión actual
2. Consulta Engram via `mem_search` → decisiones pasadas
3. Revisa `wiki/meta/handovers/` si necesita detalle de sesiones anteriores

---

## ⚡ SDD Workflow

### ¿Cuándo usar SDD vs código directo?

| SDD | Código directo |
|-----|---------------|
| Features nuevas | Quick fixes |
| Cambios multi-archivo | Typo fixes |
| Experimental | Config única |
| **Cuando hay spec clara** | **Cuando ya sabés lo que hay que hacer** |

### Las fases de SDD en orden?

```
Explore → Propose → Spec → Design → Tasks → Apply → Verify → Archive
         ↑                            ↑
      (opcional)                (REQUIRED)
```

### ¿Cómo hacer fast-forward a Apply directamente?

```
/sdd-ff <nombre-del-cambio>
```

Esto salta: proposal → specs → design → tasks → apply en un solo paso. Útil cuando ya tenés todo claro.

---

## 🛠️ Skills & delegation

### ¿Cuándo usar skills vs delegations?

| Skills | Delegations |
|--------|------------|
| Patterns establecidos | research/exploración |
| Testing | multi-archivo changes |
| Configuración | síntesis de resultados |

### ¿Cómo agrego un skill nuevo?

1. Creá `~/.config/opencode/skills/<nombre>/SKILL.md`
2. Registralo en el skill registry: `/skill-registry`
3. OpenCode lo detecta automáticamente

### Lista de skills disponibles

- `coding-standards` — TypeScript/JS patterns
- `sdd-*` — SDD workflow phases
- `git-workflow` — Git conventions
- `docker-patterns` — Container best practices
- `security-review` — OWASP checklist
- `wiki-*` — Wiki management
- → Ver todos en `skill` tool autocomplete

---

## 🐛 Debug

### El agente está repetiendo lo mismo

1. Buscá en Engram si hay contexto conflicting
2. Reseed el `wiki/hot.md` con el objetivo claro
3. Verificá que no hay "remember me" implícito en el prompt

### El workflow SDD está trabado

```
/sdd-continue <change-name>
```

O si necesita reset:
```
/sdd-new <change-name>
```

### Los tests no pasan

1. Verificá strict TDD mode: `mem_search(query: "sdd-init")`
2. Corré `npm test` o `go test` inline
3. Si es flake → re-ejecutá 1 vez antes de assume failure

---

## 📦 Instalación & Config

### ¿Cómo configuro el MCP de Obsidian?

1. Install plugin "OpenCode" en Obsidian
2. Copiá el config de `~/.config/opencode/opencode.json`
3. Reiniciá OpenCode

### ¿Dónde está la wiki desde otro proyecto?

```markdown
## Wiki Knowledge Base
Path: ~/opencode-obsidian
```

Agregá esto al `AGENTS.md` del otro repo.

---

## 🔄 Workflows

### ¿Cómo hago un handover entre sesiones?

1. El agente genera `wiki/meta/handovers/<date>.md` automáticamente
2. Incluye: goal, accomplishments, next steps, files relevantes
3. La próxima sesión lo lee como primer paso

### ¿Cómo sync la wiki a Git?

```bash
./bin/wiki-sync.sh
```

O desde Obsidian: comando `> OpenCode: Sync`

---

## 🎯 Shortcuts

| Comando | Descripción |
|---------|-------------|
| `/sdd-init` | Inicializar SDD en proyecto |
| `/sdd-new <name>` | Nuevo cambio SDD |
| `/save this` | Guardar conversación |
| `lint` | Health check del vault |
| `query <topic>` | Buscar en wiki |

---

> 💡 **Tip**: Usá `Ctrl+K` en Obsidian para search rápido en el vault.