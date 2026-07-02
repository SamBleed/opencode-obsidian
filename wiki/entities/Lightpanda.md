---
type: entity
title: "Lightpanda"
updated: 2026-05-14
status: active
---

# Lightpanda

**Type**: Headless Browser / Web Automation
**Category**: AI-Native / Performance
**Status**: Researching / Integration Reference
**Parent**: [[_index]]

## 📝 Overview

**Lightpanda** is an open-source, purpose-built headless browser written in **Zig**. Unlike Chromium-based engines, it is designed strictly for machines, omitting visual rendering, CSS layout engines, and media subsystems to achieve extreme performance and low memory footprint (~20MB/instance).

## 🚀 Key Features

- **Built in Zig**: Leverages manual memory management and high-performance system programming.
- **CDP Compatible**: Works as a drop-in replacement for Playwright, Puppeteer, and Chromedp.
- **AI-Native**: Optimized for DOM extraction and structured data gathering rather than visual display.
- **Extreme Concurrency**: Allows running 10x-15x more parallel sessions on the same hardware compared to Chrome.

## 🧠 Connection to opencode-obsidian

Lightpanda represents the **Sensing Layer**. It is the "fast eye" that allows the bunker to ingest web knowledge at scale.

### Patterns to Adopt:
1. **Machine-First Browsing**: Prioritizing structured DOM data over visual screenshots for research notes.
2. **High-Speed Ingestion**: Using lightweight tools to feed the `.raw/` directory faster.
3. **Performance Awareness**: Preferring tools written in Zig/Rust for local automation tasks on the Arch Linux stack.

## 🔗 Resources
- **GitHub**: https://github.com/lightpanda-io/browser
- **Website**: https://lightpanda.io

---
**Last Updated**: 2026-04-19
