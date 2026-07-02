---
type: entity
title: "OpenBB"
updated: 2026-05-14
status: active
---

# OpenBB

**Type**: Financial Data Platform / Infrastructure
**Category**: Fintech / Multi-Agent / MCP Native
**Status**: Researching / Integration Reference
**Parent**: [[_index]]

## 📝 Overview

**OpenBB** is the leading open-source financial data platform. It has transitioned from a CLI terminal to a modular infrastructure (OpenBB Platform - ODP) that enables AI agents to consume financial data via Python, REST APIs, and **MCP servers**.

## 🚀 Key Features

- **Consolidated Data**: Unified access to equity, crypto, forex, options, and macro data.
- **MCP Servers**: Native support for the Model Context Protocol, allowing AI agents to query market data directly.
- **OpenBB Workspace**: An AI-native collaborative research environment with integrated Copilots.
- **SDK for Agents**: Specialized toolkits (`openbb-ai`) for building financial-specific agents.

## 🧠 Connection to opencode-obsidian

OpenBB validates the **Domain-Specific Vault** pattern. Just as this vault is optimized for *coding and engineering*, OpenBB is an external "vault" for *financial intelligence*.

### Patterns to Adopt:
1. **Infrastructure Abstraction**: Exposing local data via standardized protocols (like OpenBB does with ODP).
2. **Domain Specialization**: Recognizing that a general-purpose LLM needs specialized "anchors" (Vaults) to be reliable in complex fields like finance or systems engineering.
3. **Connect Once, Use Everywhere**: The "Connect Once" philosophy aligns with our [[ENV-STATUS]] and [[LEARNING-LOOP]] goals.

## 🔗 Resources
- **GitHub**: https://github.com/OpenBB-finance/OpenBB
- **Website**: https://openbb.co
- **Agents SDK**: https://github.com/OpenBB-finance/openbb-ai

---
**Last Updated**: 2026-04-19
