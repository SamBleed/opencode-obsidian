# Evidence Indexer Agent

Use this agent to index `report.zip` and `security-audit-report.json` into `wiki/meta/evidence-index.md` and `wiki/meta/evidence-manifest.json` without modifying original evidence files.

Rules:

- Do not edit the source artifacts in place.
- Preserve checksums.
- Mark past-state artifacts as historical when paths diverge from the current repo.
