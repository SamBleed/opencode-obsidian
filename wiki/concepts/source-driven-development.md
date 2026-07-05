---
type: concept
title: Source-Driven Development
domain: engineering
created: 2026-05-16
updated: 2026-05-16
tags: [documentation, verification, best-practices]
sources: [[agent-skills]]
---

# Source-Driven Development

## Definition
Every framework-specific code decision must be backed by official documentation. Don't implement from memory — verify, cite, and let the user see your sources. Training data goes stale, APIs get deprecated.

## The Process
```
DETECT → FETCH → IMPLEMENT → CITE
  │         │          │         │
  ▼         ▼          ▼         ▼
What     Get the    Follow    Show your
stack?   relevant   documented sources
         docs       patterns
```

## Source Hierarchy (by authority)
1. Official documentation (react.dev, docs.djangoproject.com)
2. Official blog/changelog
3. Web standards references (MDN, web.dev)
4. Browser/runtime compatibility (caniuse.com)

**Never cite as primary**: Stack Overflow, blog posts, AI-generated docs.

## Key Patterns
- Detect exact versions from package.json/requirements.txt before fetching docs
- Use Context7 for up-to-date library documentation
- Cite sources inline with implementation
- Ask user when versions are ambiguous — don't guess

## Red Flags
- Implementing framework patterns from memory
- Citing blog posts over official docs
- Not verifying which version is in use
- Using outdated patterns from training data

## Source
- [[agent-skills]] — Source-Driven Development skill
