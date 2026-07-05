---
type: entity
title: "Decepticon"
updated: 2026-05-14
status: active
---

# Decepticon (PurpleAILAB)

**Type**: Autonomous Red Team Framework
**Category**: Offensive Security / Multi-Agent
**Status**: Researching / Pattern Extraction
**Parent**: [[_index]]

## 📝 Overview

**Decepticon** is an open-source autonomous multi-agent framework by **PurpleAILAB** designed for automated penetration testing and Red Team operations. It mimics human attacker behavior by chaining reconnaissance, exploitation, and post-exploitation phases.

## 🚀 Key Features

- **Multi-Agent Orchestration**: Specialized agents for Recon, Exploit, and Post-Exploit.
- **Interactive Shell Handling**: Uses persistent sessions (via `tmux`) to manage interactive terminal prompts.
- **Full Kill Chain**: Capable of pivoting and lateral movement within networks.
- **Sandboxed Execution**: Operates within isolated Kali Linux containers.

## 🧠 Connection to opencode-obsidian
Decepticon is the "Force" of the Agentic Trident, utilizing tools like [[Trivy]] (Shield) and [[Sherlock]] (Recon).

While opencode-obsidian is a knowledge base, Decepticon provides a blueprint for **autonomous execution**. 

### Patterns to Adopt:
1. **Task Separation (Orchestration)**: Refining our `sdd` workflow to separate high-level planning from atomic terminal executions.
2. **Persistence Management**: Using tmux-like persistence for long-running scripts or verification loops.
3. **Offensive Vaccine**: Using the same logic to proactively "attack" our code (Security Auditing) before it reaches production.

## 🔗 Resources
- **GitHub**: https://github.com/PurpleAILAB/Decepticon
- **Developer**: PurpleAILAB
- **Parent**: [[index]]

---
**Last Updated**: 2026-04-19
