---
type: concept
title: "Policy-as-Code (PaC)"
domain: tech-stack
created: 2026-05-13
updated: 2026-05-13
tags: [security, devops, opa, governance, automation]
status: mature
---

# Policy-as-Code (PaC)

Policy-as-Code is the approach of managing and enforcing security, compliance, and operational rules using machine-readable code. It eliminates manual "gatekeeping" and human error by making policies executable and version-controlled.

## 🏗️ Core Mechanisms

### 1. Decoupled Policy Decision
Policies are separated from the application logic. The application (or CI/CD pipeline) asks a **Policy Engine** for a decision: "Is this deployment allowed?".

### 2. The Rego Language (OPA)
**Open Policy Agent (OPA)** is the industry standard for PaC. It uses the **Rego** language to define declarative policies.

```rego
# Example: Deny deployment if image is 'latest'
package bunker.deployment

deny[msg] {
    input.image == "latest"
    msg := "Deploying with 'latest' tag is forbidden in the Bunker."
}
```

## 🚀 Use Cases in the Bunker

### CI/CD Security Gates
Automated checks that block PRs or deployments if they violate governance rules:
- **Image Signature**: Must be signed by a trusted authority.
- **Vulnerability Threshold**: No "Critical" or "High" vulnerabilities allowed.
- **Review Requirement**: At least 2 human (or verified AI) reviews.

### Infrastructure Guardrails (IAC)
Scanning Terraform or CloudFormation files before provisioning:
- **Public S3**: Deny if an S3 bucket is set to public access.
- **Unencrypted DB**: Deny if a database doesn't have encryption enabled.

### Cost Control
- **Instance Limits**: Deny if a developer tries to spin up an expensive GPU instance without an approved Jira ticket.

## 🛡️ Implementation Stack

- **Engine**: Open Policy Agent (OPA) or HashiCorp Sentinel.
- **Integration**: Spacelift, Harness, or custom GitHub Actions.
- **Verification**: [[Trivy]] for scanning IAC files against OPA policies.

## ⚖️ Senior Architect Considerations

- **Policy Drift**: Ensure that policies in Git match the active policies in the engine.
- **Dry Run Mode**: Always deploy new policies in "Audit" mode (logging violations without blocking) before turning on "Enforcement" mode.
- **Exceptions**: Design a structured way to handle policy exemptions with a mandatory expiration date.

## Related
- [[Security-Guardrails]]
- [[OzyZT Architecture]]
- [[Supply-Chain-Security]]
- [[Trivy]]
