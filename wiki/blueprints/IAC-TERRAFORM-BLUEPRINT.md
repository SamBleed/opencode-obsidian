---
type: blueprint
title: "IaC Terraform & OpenTofu Blueprint"
domain: tech-stack
created: 2026-05-13
updated: 2026-05-13
tags: [iac, terraform, opentofu, blueprint, infrastructure]
status: mature
---

# IaC Terraform & OpenTofu Blueprint (v2026)

This blueprint defines the standard for managing cloud infrastructure using Terraform or OpenTofu. It implements the **Franchise Model** and **Zero-Trust State** patterns.

## 📂 Directory Structure (Standard)

```text
.
├── modules/                # Reusable, single-responsibility components
│   ├── vpc/
│   ├── rds/
│   └── eks/
├── platforms/              # Foundation layer (Shared Infra)
│   ├── networking/
│   └── identity/
├── services/               # Product configuration layer (App-specific)
│   ├── api-backend/
│   └── web-frontend/
├── scripts/                # Task scripts (tf-plan, tf-apply)
├── terragrunt.hcl          # Root configuration (DRY)
└── README.md
```

## 🏗️ Design Principles

1. **Flat Module Tree**: Avoid deep nesting of modules. Use composition at the root layer.
2. **State Encryption**: Always use OpenTofu's native encryption for state and plans.
3. **Data-Only Modules**: Create modules that only use `data` sources to fetch existing infra info.
4. **One-Directional Flow**: Foundation outputs → Service modules consume → Product configs compose.

## 🛡️ Security Guardrails

- **Zero Secrets**: Use `trufflehog` to scan for secrets before every commit.
- **Policy-as-Code**: Enforce [[Policy-as-Code]] (OPA) using [[Trivy]] scans in CI.
- **RBAC**: Each state file must have its own IAM role with minimum required permissions.

## 🧪 Testing Checklist

- [ ] **Static Analysis**: `tofu validate` and `tflint`.
- [ ] **Security Scan**: `trivy config .`.
- [ ] **Unit Tests**: `terratest` for complex logic in modules.
- [ ] **Drift Detection**: Daily automated plans to detect manual changes.

## 🚀 Execution Pattern

Use **Terragrunt Stacks** to manage dependencies between layers:
```hcl
# services/api-backend/terragrunt.hcl
dependency "vpc" {
  config_path = "../../platforms/networking"
}
inputs = {
  vpc_id = dependency.vpc.outputs.vpc_id
}
```

## Related
- [[Terraform-OpenTofu-Senior-Patterns]]
- [[GO-API-BLUEPRINT]]
- [[Security-Guardrails]]
- [[Policy-as-Code]]
