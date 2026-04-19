---
type: session
title: "n8n Alerter Registration and Validation"
created: 2026-04-19
updated: 2026-04-19
tags:
  - session
  - n8n
  - mcp
  - automation
status: active
related:
  - "[[n8n]]"
  - "[[MCP-CONFIG]]"
  - "[[n8n-mcp-exposure-plan]]"
  - "[[opencode-obsidian]]"
---

# n8n Alerter Registration and Validation

## Mission Status

- Goal: document and register the `Bunker Ultimate Alerter` workflow in both the vault and Engram, then validate the end-to-end alert path.
- Outcome: success.
- Current Blockers: none. The only transient issue was the sandbox not reaching `localhost:5678`, which was resolved by rerunning the alert outside the sandbox.

## Context Anchors

- Active Files:
  - `~/n8n-lab/workflows/bunker-ultimate-alerter.json`
  - `~/opencode-obsidian/bin/bunker-alert.sh`
  - `~/opencode-obsidian/wiki/meta/MCP-CONFIG.md`
  - `~/opencode-obsidian/wiki/blueprints/n8n-mcp-exposure-plan.md`
  - `~/opencode-obsidian/wiki/entities/n8n.md`
- New Entities:
  - `Bunker Ultimate Alerter` is confirmed active in n8n with ID `wq08EHiWdTmUQjXL`.
  - `n8n-mcp` is connected in OpenCode.
- Decision Log:
  - `n8n-mcp` should expose read-only workflows first, then alerting flows, then write-heavy remediation flows.
  - `Bunker Ultimate Alerter` is now the validated alert path.

## Procedural Memory

- Gotchas:
  - `bunker-alert.sh` must be run outside the sandbox to reach `localhost:5678` reliably in this environment.
  - The n8n instance already had the workflow in its database even though the repo export initially looked disconnected.
  - `Bunker Ultimate Alerter` needed a real `chatId` and activation in n8n before it could send alerts.
- Patterns:
  - Keep alert payloads JSON-safe before POSTing them.
  - Verify the live n8n export when a workflow exists both in the repo and in the instance.
  - Document the MCP exposure order so future workflows are added in the safest sequence.

## Next Actions

- [ ] Create a read-only n8n workflow for status and health snapshots.
- [ ] Add a `Respond to Webhook` or explicit success path if you want better feedback from alerts.
- [ ] Consider moving the Telegram credential and any API keys fully out of inline config.

