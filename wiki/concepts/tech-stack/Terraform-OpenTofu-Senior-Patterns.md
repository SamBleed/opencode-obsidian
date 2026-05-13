---
type: concept
title: "Terraform & OpenTofu Senior Patterns"
domain: tech-stack
created: 2026-05-13
updated: 2026-05-13
tags: [iac, terraform, opentofu, devops, security]
status: mature
---

# Terraform & OpenTofu Senior Patterns (v2026)

Managing infrastructure as code at scale requires moving beyond flat files to a **Three-Layer Module Architecture** and adopting **Zero-Trust State Management**.

## 🏗️ The Franchise Model (Layered Architecture)

| Layer | Owner | Content |
|-------|-------|---------|
| **1. Foundation** | Platform Team | Shared VPCs, IAM Baselines, KMS Keys. Sacred and rare changes. |
| **2. Service Modules** | Platform Team | Opinionated wrappers (e.g., `secure-rds`). Compliance baked-in. |
| **3. Product Config** | App Teams | Orchestration of service modules for a specific app. No raw resources. |

## 🛡️ Zero-Trust State Management

### 1. State Encryption (OpenTofu Native)
OpenTofu supports encrypting state files at rest using providers like **AWS KMS** or **PBKDF2**. This is a critical upgrade over standard Terraform.

```hcl
terraform {
  encryption {
    key_provider "pbkdf2" "state_key" {
      passphrase = var.state_passphrase
    }
    method "aes_gcm" "method1" {
      keys = key_provider.pbkdf2.state_key
    }
    state {
      method = method.aes_gcm.method1
    }
  }
}
```

### 2. State Isolation
Minimize the **Blast Radius**. Separate state files by:
- **Environment**: dev, staging, prod.
- **Component**: network, database, application.

## 🚀 Advanced Module Composition

### Dependency Inversion
Instead of a module creating its own VPC, make it accept a `vpc_id`. This allows the same module to be reused across different network topologies.

### Module Versioning
Always pin internal modules to a specific version. Use **Semantic Versioning (SemVer)** to communicate breaking changes in infrastructure APIs.

## 🛠️ Terragrunt Stacks
For 2026, we use **Terragrunt Stacks** to manage 100+ micro-deployments. It handles the dependency graph and allows running `terragrunt stack run apply` to orchestrate multi-account rollouts.

## 🧪 Verification
- **Trivy**: Automated OPA policy checks against IAC files.
- **Terratest**: Go-based unit tests for infrastructure modules.
- **Drift Detection**: Automated `tofu plan` in CI to detect manual cloud changes.

## Related
- [[Docker]]
- [[Security-Guardrails]]
- [[Policy-as-Code]]
- [[Trivy]]
