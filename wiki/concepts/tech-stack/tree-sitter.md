---
type: concept
domain: tech-stack
title: "Tree-sitter"
description: "Parser generator tool and incremental parsing library for building fast, accurate syntax trees"
created: 2026-04-22
updated: 2026-04-22
tags: [parser, ast, tree-sitter, code-analysis]
status: seed
related: [knowledge-graph-code]
sources: [gitnexus]
---

# Tree-sitter

## Definition

**Tree-sitter** is a parser generator tool and incremental parsing library that generates parsers for programming languages from grammar definitions. It produces **Abstract Syntax Trees (ASTs)** that are used for code analysis, syntax highlighting, code navigation, and more.

## How It Works

1. **Grammar Definition** — Define language syntax in a JSON-like grammar file
2. **Parser Generation** — Tree-sitter generates a parser from the grammar
3. **Parsing** — The parser builds a syntax tree from source code
4. **Incremental Updates** — Can re-parse only changed portions (very fast for edits)

## Key Features

| Feature | Description |
|---------|-------------|
| **Incremental Parsing** | Re-parse only changed portions, not entire file |
| **Multi-language** | Supports 30+ languages (TS, JS, Python, Go, Rust, etc.) |
| **Accurate ASTs** | Produces precise syntax trees for code analysis |
| **Fast** | Written in C, designed for performance |
| **Bindings** | Node.js, Python, Rust, Go, and more |

## Use in Code Analysis

Tree-sitter is fundamental for building knowledge graphs:

```
Source Code → Tree-sitter AST → Knowledge Graph
                                    ↓
                       - Function definitions
                       - Class definitions
                       - Import statements
                       - Function calls
                       - Type annotations
```

## Example: Extracting Functions from TypeScript

```javascript
const parser = new Parser();
parser.setLanguage(tsGrammar);

const sourceCode = `
function greet(name: string): string {
  return "Hello, " + name;
}
`;

const tree = parser.parse(sourceCode);

// Navigate the tree
tree.rootNode.children.forEach(node => {
  if (node.type === 'function_declaration') {
    console.log(`Function: ${node.childForField('name').text}`);
    console.log(`Params: ${node.childForField('parameters').text}`);
  }
});
```

## Tree-sitter in GitNexus

GitNexus uses Tree-sitter to extract:
- Functions, classes, methods, interfaces
- Import/export statements
- Type annotations
- Function calls

Both as:
- **Native bindings** — for CLI (fast, full features)
- **WASM** — for web UI (runs in browser)

## Related Concepts

- [[knowledge-graph-code]] — Uses Tree-sitter for AST extraction
- [[code-analysis]] — Broader field of analyzing code structure