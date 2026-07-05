---
type: entity
title: "Trivy"
updated: 2026-05-14
status: active
---

# Trivy (Aqua Security)

**Type**: Vulnerability Scanner / Security Audit
**Category**: Blue Team / DevSecOps
**Status**: Researching / MCP Integration
**Parent**: [[_index]]

## 📝 Overview

**Trivy** is the most popular open-source security scanner. It is reliable, fast, and easy to use. It detects vulnerabilities (CVEs), misconfigurations (IaC), secrets, and software bill of materials (SBOM).

## 🚀 Key Features

- **Multi-Target Scanning**: Containers, Filesystems, Git Repositories, Kubernetes, and Cloud.
- **MCP Server Support**: Native integration with AI agents via Model Context Protocol.
- **Vulnerability DB**: Updated continuously with the latest CVEs.
- **IaC Scanning**: Support for Terraform, Dockerfile, and Kubernetes manifests.

## 🧠 Connection to opencode-obsidian

Trivy provides the **Defensive Shield**. It complements the offensive logic of [[Decepticon]] with proactive scanning and compliance.

### Patterns to Adopt:
1. **Continuous Auditing**: Integrating Trivy scans into the [[sdd-verify]] phase.
2. **Secret Detection**: Preventing API keys from entering the repository via pre-commit Trivy hooks.
3. **AI-Driven Remediation**: Using Trivy MCP to get natural language explanations of vulnerabilities and fix suggestions.

## 🔗 Resources
- **GitHub**: https://github.com/aquasecurity/trivy
- **Documentation**: https://aquasecurity.github.io/trivy/
- **Parent**: [[index]]

---
**Last Updated**: 2026-04-19
