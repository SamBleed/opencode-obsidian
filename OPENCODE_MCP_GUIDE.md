# OpenCode MCP Guide

This guide captures the exact MCP config shape that OpenCode accepts in this repo, plus the fix we used when `opencode` would not open.

## What OpenCode Accepts

OpenCode reads MCP servers from the root `mcp` block in `~/.config/opencode/opencode.json`.

Valid server shapes are:

- `type: "local"` with a `command` array.
- `type: "remote"` with a `url`.
- `enabled: true | false` is optional.
- `headers` is allowed only for `type: "remote"`.
- `oauth` is allowed only for `type: "remote"`.

### Working remote MCP example

```json
{
  "mcp": {
    "n8n-mcp": {
      "type": "remote",
      "url": "http://localhost:5678/mcp-server/http",
      "enabled": true,
      "oauth": false
    }
  }
}
```

If you want OpenCode to manage OAuth, leave the server as `type: "remote"` and use:

```bash
opencode mcp auth n8n-mcp
```

OpenCode will store the auth state separately. Do not use the old `type: "http"` shape. That is what broke startup.

## What Broke

The original config used this invalid shape:

```json
{
  "n8n-mcp": {
    "type": "http",
    "url": "http://localhost:5678/mcp-server/http",
    "headers": {
      "Authorization": "Bearer <token>"
    }
  }
}
```

OpenCode 1.3.17 rejected it with:

```text
Invalid input mcp.n8n-mcp
```

The schema in the installed SDK only accepts `local` or `remote`.

## How We Resolved It

1. Replaced `type: "http"` with `type: "remote"`.
2. Kept the server URL as `http://localhost:5678/mcp-server/http`.
3. Enabled the server with `enabled: true`.
4. Ran `opencode mcp auth n8n-mcp`.
5. Completed the OAuth login in the browser.
6. Verified the result with `opencode mcp auth list` and `opencode mcp debug n8n-mcp`.

### Result

- `opencode mcp auth list` showed `n8n-mcp authenticated`.
- `opencode mcp debug n8n-mcp` showed `Connection successful (already authenticated)`.

## Why You Saw a 401

The `401 Unauthorized` from `opencode mcp debug n8n-mcp` is expected during the first probe.

OpenCode first calls the server without credentials to detect the auth flow.

That is why the debug output shows:

- initial `401 Unauthorized`
- then the OAuth redirect
- then success after login

So `401` does not mean the MCP is broken if the final state is authenticated.

## If OpenCode Still Will Not Open

If `opencode` fails before it even gets to MCP, check the local SQLite state:

```bash
sqlite3 ~/.local/share/opencode/opencode.db 'PRAGMA wal_checkpoint(TRUNCATE);'
```

We used that to clear the WAL file when OpenCode was stuck on:

```text
Failed to run the query 'PRAGMA wal_checkpoint(PASSIVE)'
```

After that, the WAL file dropped to zero bytes and OpenCode started again.

## Useful Commands

```bash
opencode mcp list
opencode mcp auth list
opencode mcp debug n8n-mcp
opencode mcp auth n8n-mcp
```

## Related Notes

- `wiki/meta/MCP-CONFIG.md`
- `skills/wiki/references/mcp-setup.md`
- `wiki/meta/HANDOVER.md`
