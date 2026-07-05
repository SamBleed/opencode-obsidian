---
type: concept
title: "Deprecation and Migration"
domain: engineering
created: 2026-05-16
updated: 2026-05-16
tags: [deprecation, migration, engineering]
---

# Deprecation and Migration

## Overview
Treat code as a liability and manage its lifecycle deliberately. Deprecate old APIs, services, or features with clear signals, versioned deprecation warnings, and migration paths. Remove only after safe removal, monitoring, and confirmation that usage has ceased.

## Key Practices
1. **Versioned Deprecation Warnings** – Add deprecation notices to public interfaces, documenting removal timeline and replacement.
2. **Gradual Removal** – Release in phases: mark as deprecated, allow two release cycles, then retire.
3. **Migration Guides** – Provide explicit instructions, code snippets, and support for moving to the new alternative.
4. **Telemetry & Monitoring** – Track usage of deprecated elements; stop supporting once usage hits zero.
5. **Safeguards** – Keep deprecated code behind feature flags or behind a “kill switch” to abort if unexpected usage appears.
6. **Documentation** – Record the deprecation schedule, reasons, and migration steps in the project wiki.

## When to Use
- Retiring an API, endpoint, or module.
- Removing an old feature flag.
- Replacing a third‑party dependency with an in‑house solution.
- Any change that eliminates previously public behavior.

## Verification Checklist
- [ ] Deprecation warning added to affected code or docs.
- [ ] Migration guide published and linked.
- [ ] Usage telemetry shows zero active calls before removal.
- [ ] Feature‑flag or kill‑switch in place for final removal.
- [ ] Documentation updated with deprecation timeline.