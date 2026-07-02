---
type: evidence-index
title: "Evidence Index"
updated: 2026-05-14
status: active
tags: [evidence, audit, governance, bunker-os]
related: [[dashboard]], [[INTEGRITY-REPORT]]
---

# Evidence Index

This index preserves evidence as original artifacts and links them to wiki interpretation. Do not edit original evidence files in place.

## `report.zip`

- Type: `zip-evidence`
- Size: `168729` bytes
- SHA256: `b7522227f9c2eb28dae5c2fa0fb73849adb1b22387e9f1dc3e54e0a22c2761ee`
- Manifest: `wiki/meta/evidence-manifest.json`

### Internal entries
- `report/E01_nslookup_whois.txt` — 519 bytes — sha256 `427ad740810442b8a26779b66276257815d03c30ce08c29d75733b37beaf1494`
- `report/E02_gobuster_subdominios.txt` — 554 bytes — sha256 `38c499dba0b24528786f71af89c0e48fdb6fc1a90c2fc2a94f78488bee2729d0`
- `report/E03_nmap_host1.txt` — 1273 bytes — sha256 `929c8dd0e339e32cf879ccd870a07319b84ad865bd2d22113137d238fecf8213`
- `report/E04_nmap_host2.txt` — 317 bytes — sha256 `7ae57b9012a0577fdb006696d6ead4e4fb468f44e7b6bcd0d4bb307611c75f19`
- `report/E05_whatweb_wordpress.txt` — 515 bytes — sha256 `30bc52f4733e5185ad16e9297eef721bb874a852e2d0ef49f2a10a45b72d5993`
- `report/E06_whatweb_moodle.txt` — 619 bytes — sha256 `a547346ef4f15d94c7767a997b155ef9333fcbcab70ce8de2426f3279b3887e0`
- `report/E07_cpanel.png` — 38326 bytes — sha256 `139387de2b862ad627b1d189a3df3d7edcfc7c13d091f6d7181efbeb034b62ab`
- `report/E08_webmail.png` — 44598 bytes — sha256 `bfe7a4d2db6540067841cb297e8715fd83da50d1d06a795f79e14e229d51804d`
- `report/E08_wp_users.txt` — 413 bytes — sha256 `cd0f903674e336ba600910d7cc753cc7c9c84a1897bac1ee0745f5e8ffa99318`
- `report/E09_moodle.png` — 85267 bytes — sha256 `c602b3d665e9a7f49c0f6821d1ab0fc78c67569ca9b167396f3980e333c7a73c`
- `report/entregable_colegio_integridad.md` — 25247 bytes — sha256 `26a6061984ab09e779199afb2d7538251b988326f3101ed5006a3130cc05d341`

## `security-audit-report.json`

- Type: `security-audit-json`
- Size: `74459` bytes
- SHA256: `de2de7bdbabcc0ffea0a17b05ca55a8b5e449725be3c32ef8723f68b79a3cf41`
- Manifest: `wiki/meta/evidence-manifest.json`

### Summary
- top_level_keys: `['ArtifactID', 'ArtifactName', 'ArtifactType', 'CreatedAt', 'Metadata', 'ReportID', 'Results', 'SchemaVersion', 'Trivy']`
- note: `Stored as original audit output; review before using as current truth.`

## Evidence Handling Policy

- Original artifacts remain immutable unless a new version is explicitly created.
- Notes may interpret evidence, but must link back to this index or the manifest.
- Checksums are used to detect accidental changes.
- If an artifact comes from a past repository state, mark it as historical instead of treating it as current truth.
