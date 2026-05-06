# opencode-obsidian — OpenCode + Obsidian Wiki Vault

This folder is both a OpenCode plugin and an Obsidian vault.

**Plugin name:** `opencode-obsidian`
**Skills:** `/wiki`, `/wiki-ingest`, `/wiki-query`, `/wiki-lint`
**Vault path:** This directory (open in Obsidian directly)

## What This Vault Is For

This vault demonstrates the LLM Wiki pattern — a persistent, compounding knowledge base for OpenCode + Obsidian. Drop any source, ask any question, and the wiki grows richer with every session.

## Vault Structure

```
.raw/           source documents — immutable, OpenCode reads but never modifies
wiki/           OpenCode-generated knowledge base
_templates/     Obsidian Templater templates
_attachments/   images and PDFs referenced by wiki pages
```

## How to Use

Drop a source file into `.raw/`, then tell OpenCode: "ingest [filename]".

Ask any question. OpenCode reads the index first, then drills into relevant pages.

Run `/wiki` to scaffold a new vault or check setup status.

Run "lint the wiki" every 10-15 ingests to catch orphans and gaps.

## Cross-Project Access

To reference this wiki from another OpenCode project, add to that project's AGENTS.md:

```markdown
## Wiki Knowledge Base
Path: /path/to/this/vault

When you need context not already in this project:
1. Read wiki/hot.md first (recent context, ~500 words)
2. If not enough, read wiki/index.md
3. If you need domain specifics, read wiki/<domain>/_index.md
4. Only then read individual wiki pages

Do NOT read the wiki for general coding questions or things already in this project.
```

## Plugin Skills

| Skill | Trigger |
|-------|---------|
| `bunker init` | Workflow Start: Lee hot.md y handover. |
| `bunker refresh` | Re-alinea: Re-lee pilares para evitar context drift. |
| `bunker save` | Checkpoint: Commits rápidos de la wiki. |
| `bunker close` | Session End: Handover, sync y push. |
| `/wiki` | Setup, scaffold, route to sub-skills |
| `ingest [source]` | Single or batch source ingestion |
| `query: [question]` | Answer from wiki content |
| `lint the wiki` | Health check |
| `/save` | File the current conversation as a structured wiki note |
| `/autoresearch [topic]` | Autonomous research loop: search, fetch, synthesize, file |
| `/canvas` | Visual layer: add images, PDFs, notes to Obsidian canvas |
## MCP (Optional)

If you configured the MCP server, OpenCode can read and write vault notes directly.
See `skills/wiki/references/mcp-setup.md` for setup instructions.
