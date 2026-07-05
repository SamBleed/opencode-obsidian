---
type: entity
title: "Trust Security"
domain: security
tags: [entity, organization, security, web3]
status: active
sources:
  - "[[immunefi-bug-bounty-writeups]]"
---

# Trust Security

A web3 security research firm. One of the most active whitehat teams on Immunefi with 30+ reported vulnerabilities across lending protocols, bridges, and DeFi platforms.

## Notable Findings

- Compound: liquidators seizing non-collateral assets (Critical)
- Morpho: logic contract destruction via delegatecall (Medium)
- Iron Bank: incorrect collateral calculations (High)
- LayerZero: Stargate reward distribution (High)
- RAI: returndata bombing in liquidation engine (Critical)
- Brahma: fee collection ignoring previous losses (Critical)

## Approach

Trust specializes in finding logic errors in DeFi lending protocols. Their writeups often focus on edge cases in accounting, collateral calculations, and liquidation mechanisms.

## Source

- [[immunefi-bug-bounty-writeups]]
- https://www.trust-security.xyz
