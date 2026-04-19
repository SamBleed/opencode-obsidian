# Research: LLM Wiki Pattern by Andrej Karpathy

The **LLM Wiki** is a pattern proposed by Andrej Karpathy in early 2026. It suggests a shift from stateless RAG to a persistent, compounding knowledge graph.

## Core Concepts
- **Persistence**: Unlike RAG, knowledge is synthesized into structured Markdown pages.
- **Compounding**: Every session adds to the existing knowledge, making it smarter over time.
- **Human-Readable**: The knowledge is stored in standard Markdown, accessible by both humans and LLMs.
- **Efficiency**: At personal scales, the index can fit in context, bypassing vector DB needs.

## Architecture
- `raw/`: Unstructured source data.
- `wiki/`: Structured, LLM-generated Markdown files.
- `hot.md`: High-priority context loaded every session.
- `index.md`: The map of all knowledge.
