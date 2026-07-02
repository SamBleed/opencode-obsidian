---
type: concept
title: API and Interface Design
domain: api-design
created: 2026-05-16
---

# API and Interface Design

## Overview
Design stable, well-documented interfaces that are hard to misuse.

## Core Principles
- **Hyrum's Law**: Public behavior becomes de facto contract.
- **One-Version Rule**: Extend, don't fork; keep a single version.
- **Contract First**: Define types before implementation.
- **Consistent Error Semantics**: Unified error format.
- **Validate at Boundaries**: Only at system edges.
- **Add Over Modify**: Add fields optionally.
- **Predictable Naming**: REST conventions, camelCase, etc.

## Red Flags
- Inconsistent error formats.
- Breaking changes.
- No pagination.
- Verbs in URLs.

## Verification Checklist
- Typed input/output schemas.
- Unified error responses.
- Validation at boundaries only.
- List endpoints paginated.
- New fields additive.
- Consistent naming.