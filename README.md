# opencode-obsidian

<p align="center">
  <img src="wiki/meta/claude-obsidian-gif-cover-16x9.gif" alt="opencode-obsidian" width="100%" />
</p>

[![GitHub stars](https://img.shields.io/github/stars/SamBleed/opencode-obsidian?style=flat&color=e8734a)](https://github.com/SamBleed/opencode-obsidian/stargazers)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![OpenCode](https://img.shields.io/badge/OpenCode-ready-8B5CF6)](https://opencode.ai)

Wiki Vault adaptado para **OpenCode** basado en el patrón LLM Wiki de Karpathy.

---

## Instalación Rápida

```bash
# Un comando instala todo
bash <(curl -sL https://raw.githubusercontent.com/SamBleed/opencode-obsidian/main/bin/setup-opencode.sh)
```

O manual:

```bash
git clone https://github.com/SamBleed/opencode-obsidian.git ~/opencode-obsidian
cd ~/opencode-obsidian
chmod +x bin/setup-opencode.sh
./bin/setup-opencode.sh
```

---

## Requisitos

- [OpenCode](https://opencode.ai) instalado
- [Obsidian](https://obsidian.md) instalado
- Node.js (para el MCP)

---

## Estructura

```
opencode-obsidian/
├── .raw/           # Fuentes originales (nunca modificar)
├── wiki/           # Knowledge base
├── bin/            # Scripts de setup
├── skills/         # Skills adaptados para OpenCode
└── _templates/     # Templates de Obsidian
```

---

## Comandos

| Comando | Descripción |
|---------|-----------|
| `ingest [archivo]` | Procesar fuente y crear notas wiki |
| `query [tema]` | Buscar en el wiki |
| `what do you know about X?` | Responder desde el conocimiento |
| `save this` | Guardar conversación como nota |
| `lint` | Health check del wiki |

---

## MCP

El MCP `obsidian-vault` permite acceso directo al vault desde OpenCode.

Configuración automática via `setup-opencode.sh`.

O manual:

```bash
# En opencode.json:
"mcp": {
  "obsidian-vault": {
    "type": "stdio",
    "command": "npx",
    "args": ["-y", "@bitbonsai/mcpvault@latest", "~/opencode-obsidian"]
  }
}
```

---

## Cross-Project

Para usar este wiki desde otro proyecto, agregar a su AGENTS.md:

```markdown
## Wiki Knowledge Base
Path: ~/opencode-obsidian

Cuando necesites contexto:
1. Leé wiki/hot.md primero (~500 palabras)
2. Si no alcanza, leé wiki/index.md
3. Solo entonces leé páginas individuales
```

---

## Diferencias con claude-obsidian

| Característica | claude-obsidian | opencode-obsidian |
|---------------|-----------------|-------------------|
| Plataforma | Claude Code | OpenCode |
| Skills | Formato Claude | Formato OpenCode |
| MCP | Configuración CLAUDE.md | Configuración opencode.json |
| Instalación | Plugin CLI | Script bash |

---

## Créditos

- [Karpathy](https://github.com/karpathy) - LLM Wiki pattern
- [AgriciDaniel](https://github.com/AgriciDaniel/claude-obsidian) - Original claude-obsidian

---

MIT License