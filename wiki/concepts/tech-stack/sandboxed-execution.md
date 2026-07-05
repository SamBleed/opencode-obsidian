---
type: concept
domain: tech-stack
title: "Sandboxed Execution"
description: "Technique to execute code in isolated environments to prevent malicious scripts from escaping and to optimize context window usage"
created: 2026-04-22
updated: 2026-04-22
tags: [sandbox, security, execution, context-optimization]
status: seed
related: [context-mode, fts5-search]
sources: [context-mode]
---

# Sandboxed Execution

## Definition

Technique to execute code in isolated environments that restrict what the executed code can access. In the context of LLM-assisted development, sandboxed execution prevents the model from dumping raw data into the context window by processing it externally.

## How It Works

1. **Code is written to a temporary file** in an isolated temp directory
2. **Environment variables are sanitized** — dangerous vars like `LD_PRELOAD`, `NODE_OPTIONS`, `PYTHONSTARTUP` are blocked
3. **The code is executed** with restricted permissions
4. **Only the result is returned** — stdout/stderr, not the raw data
5. **Temp files are cleaned up** after execution

## Key Security: DENIED Environment Variables

The most critical part of sandboxed execution is blocking dangerous environment variables:

```typescript
const DENIED = new Set([
  // Shell — auto-execute scripts, override builtins
  "BASH_ENV", "ENV", "PROMPT_COMMAND", "SHELLOPTS",
  
  // Node.js — require injection, inspector
  "NODE_OPTIONS", "NODE_PATH",
  
  // Python — stdlib override, startup injection
  "PYTHONSTARTUP", "PYTHONHOME", "PYTHONBREAKPOINT",
  
  // Ruby — option/module injection
  "RUBYOPT", "RUBYLIB",
  
  // Dynamic linker — shared library injection
  "LD_PRELOAD",          // Linux
  "DYLD_INSERT_LIBRARIES", // macOS
  
  // Compiler — binary substitution
  "CC", "CXX", "AR",
  
  // Git — command injection via hooks/config
  "GIT_SSH", "GIT_ASKPASS",
  
  // ... and many more
]);
```

## Supported Languages (context-mode)

- JavaScript / TypeScript
- Python
- Go
- Rust
- Ruby
- PHP
- Perl
- R
- Elixir
- Shell

## Example Use Case

Instead of reading a 10MB log file into context:

```javascript
// DON'T: Read entire file into context
const content = fs.readFileSync("huge.log", "utf-8");
// Process mentally... ❌ Context blown

// DO: Execute in sandbox, return only the answer
// Code sent to sandbox:
const lines = require("fs").readFileSync("huge.log", "utf-8").split("\n");
const errors = lines.filter(l => l.includes("ERROR"));
console.log(`Found ${errors.length} errors`);
console.log(errors.slice(0, 5).join("\n"));
```

Output returned to context:
```
Found 47 errors
Error: Connection refused at line 123
Error: Timeout at line 456
...
```

## Related Concepts

- [[context-window-optimization]]
- [[fts5-search]] — Full-text search for knowledge bases
- [[mcp-protocol]] — Model Context Protocol

## Tools Implementing This

- [[context-mode]] — MCP plugin with polyglot sandbox executor