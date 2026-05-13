---
type: project
title: "OzyRecon"
domain: security
stack: "Python"
status: active
created: 2026-05-08
updated: 2026-05-13
tags: [osint, reconnaissance, security, automation]
---

# OzyRecon

OzyRecon is the reconnaissance and asset discovery engine for the [[OzyZT]] platform. It automates the initial phase of the security lifecycle by mapping the attack surface of internal and external infrastructure.

## 🕵️ Capabilities

### 1. Asset Discovery
Automated identification of active services, open ports, and subdomains.
- **Tools**: Integrates with `nmap`, `subfinder`, and `masscan`.

### 2. Identity OSINT
Scans for leaked credentials, exposed `.env` files, and public repository mentions related to the organization.
- **Source Monitoring**: GitHub, Pastebin, and dark-web leak databases.

### 3. Drift Monitoring
Detects changes in the public-facing infrastructure that were not authorized via IAC (Infrastructure as Code).

## 🏗️ Architecture

OzyRecon is built using **Python 3.12** for its extensive library support for networking and OSINT tasks. It uses a **Task Queue** model to handle long-running scans.

- **Queue**: Redis based tasks.
- **Storage**: Findings are stored in Postgres via the central [[OzyZT]] API.
- **Reporting**: Generates visual "Attack Surface Maps" in the OzyWeb dashboard.

## 🚀 Patterns for 2026

### Continuous Recon
Instead of point-in-time scans, OzyRecon performs "Micro-Scans" every hour on critical assets to identify new exposure windows immediately.

### LLM-Enhanced Profiling
Uses agents to analyze the results of port scans and service banners to infer the specific software versions and potential vulnerabilities without triggering heavy DAST tools.

## Related
- [[OzyZT]]
- [[OzyAudit]]
- [[Pentest-Swarm-AI]]
- [[Security-Guardrails]]
