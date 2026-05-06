---
type: product
title: "notebooklm-mcp-cli"
domain: ai
created: 2026-05-04
updated: 2026-05-04
tags: [mcp, notebooklm, python, automation]
status: active
source: https://github.com/jacob-bd/notebooklm-mcp-cli
related: ["[[notebooklm]]", "[[MCP-CONFIG]]", "[[vault-flow]]"]
---

# notebooklm-mcp-cli

## Definition
Python package that provides **programmatic access to Google NotebookLM** via CLI (`nlm`) and MCP Server (`notebooklm-mcp`). Reverse-engineered from internal Google APIs.

## Key Features
- **Dual interface**: CLI + MCP server in single package
- **35 MCP tools**: notebook CRUD, source management, studio content generation, cross-notebook queries
- **Multi-browser auth**: Chrome, Brave, Arc, Edge, Chromium (CDP-based cookie extraction)
- **Pro tier compatible**: No rate limits with Google Pro accounts
- **AI Skills**: Installable expert guides for Claude Code, OpenCode, Cursor, etc.

## Installation (Arch Linux)
```bash
pipx install notebooklm-mcp-cli
# Binaries: ~/.local/bin/nlm, ~/.local/bin/notebooklm-mcp
```

## Authentication
```bash
nlm login   # launches browser, extracts cookies
nlm login --check  # verify auth
```

## MCP Configuration (OpenCode)
Already added to `~/.config/opencode/opencode.json`:
```json
"notebooklm-mcp": {
  "command": ["/home/sam/.local/bin/notebooklm-mcp"],
  "type": "local"
}
```

## Available Notebooks (2026-05-04)
1. **Zero Trust Security and Automated Deployment for Owl Perú** (1 source)
2. **Hacker's Rapid Note-Taking System with Obsidian** (1 source)
3. **Untitled** (empty)

## Use Cases in the Bunker
- Process vault sources (n8n, Hermes) → generate FAQs/Study Guides
- Cross-notebook queries for knowledge synthesis
- Audio overviews of research notes
- Automated source ingestion via CLI

## Source
- GitHub: https://github.com/jacob-bd/notebooklm-mcp-cli
- PyPI: https://pypi.org/project/notebooklm-mcp-cli/
- Related: [[notebooklm]]
