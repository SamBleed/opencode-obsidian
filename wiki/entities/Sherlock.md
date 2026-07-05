---
type: entity
title: "Sherlock"
updated: 2026-05-14
status: active
---

# Sherlock (OSINT Project)

**Type**: OSINT Tool / Reconnaissance
**Category**: Security / Information Gathering
**Status**: Researching / Integration Reference
**Parent**: [[_index]]

## 📝 Overview

**Sherlock** is an open-source Python-based tool designed to find usernames across over 400 social networks and websites. It is a fundamental tool for OSINT (Open Source Intelligence) and the initial phases of penetration testing.

## 🚀 Key Features

- **Massive Coverage**: Scans 400+ platforms (Twitter, GitHub, Reddit, etc.) in seconds.
- **Deterministic**: Reliable HTTP response-based verification.
- **Export Formats**: Supports CSV, JSON, and Excel for easy integration with agents.
- **Privacy Support**: Can use Tor or Proxies for anonymous lookups.

## 🧠 Connection to opencode-obsidian

Sherlock serves as the **Sensing Layer for Identity**. It provides the raw data that specialized agents (like [[Decepticon]]) use to build target profiles.

### Patterns to Adopt:
1. **Username Mapping**: Automating the lookup of project contributors or targets to assess security exposure.
2. **Data Pipeline**: Parsing Sherlock's JSON output directly into wiki `entities/` pages.
3. **External Tool Integration**: Using the [[ENV-STATUS]] to manage Python dependencies required for Sherlock.

## 🔗 Resources
- **GitHub**: https://github.com/sherlock-project/sherlock
- **Parent**: [[index]]

---
**Last Updated**: 2026-04-19
