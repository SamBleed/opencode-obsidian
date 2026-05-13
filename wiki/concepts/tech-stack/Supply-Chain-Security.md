---
type: concept
title: "Supply Chain Security"
domain: security
created: 2026-05-13
updated: 2026-05-13
tags: [security, supply-chain, sbom, slsa, devsecops]
status: mature
---

# Supply Chain Security

Supply Chain Security focuses on the integrity and provenance of all components that make up a software system, from third-party libraries and base images to the build processes themselves.

## 🛡️ The SLSA Framework (Supply-chain Levels for Software Artifacts)
SLSA provides a set of incrementally increasing security levels to protect the software lifecycle:
- **Level 1**: Scripted build and SBOM generation.
- **Level 2**: Hosted build service with signed attestations.
- **Level 3**: Hardened build environment (ephemeral, isolated).

## 🚀 Critical Patterns

### 1. SBOM (Software Bill of Materials)
A comprehensive inventory of every package and dependency used in an application.
- **Format**: CycloneDX or SPDX.
- **Usage**: Scanned by [[Trivy]] to find known vulnerabilities (CVEs) in nested dependencies.

### 2. Cryptographic Attestations
Evidence generated during the build process that proves:
- The code was built from a specific Git commit.
- The build ran on a trusted CI/CD runner.
- All security tests passed.
*Pattern: The production environment should only accept artifacts with valid attestations.*

### 3. Ephemeral Secrets
Never use long-lived API tokens. Use **OIDC (OpenID Connect)** for "Trusted Publishing". The CI runner proves its identity to the cloud provider (AWS, GCP, NPM) and receives a short-lived token for that specific task.

### 4. GitOps as Security
Prohibit manual changes to runtime environments. Any deviation ("Drift") from the Git-defined state must be automatically detected and reverted. This ensures that an attacker cannot persist changes manually.

## 🧪 Implementation Stack

- **Scanning**: [[Trivy]], Snyk, or GitHub Dependabot.
- **Signing**: Cosign (for container images), Sigstore.
- **Attestation**: GitHub Actions + SLSA generator.
- **Governance**: [[Policy-as-Code]] (OPA) to enforce SBOM and SLSA requirements.

## 🧠 Senior Architect Considerations

- **Transitive Vulnerabilities**: 90% of risks come from libraries you didn't directly import. Deep-scan all layers.
- **AI-BOM**: In the era of agentic development, track which LLMs and agents contributed to each part of the codebase.

## Related
- [[Security-Guardrails]]
- [[OzyZT Architecture]]
- [[Policy-as-Code]]
- [[Trivy]]
