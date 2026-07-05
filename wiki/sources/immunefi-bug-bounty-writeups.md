---
type: source
title: "Immunefi Bug Bounty Writeups List"
source_type: url
author: "sayan011"
date_published: 2025
url: "https://github.com/sayan011/Immunefi-bug-bounty-writeups-list"
confidence: high
tags: [security, web3, bug-bounty, immunefi]
related:
  - "[[Trust Security]]"
  - "[[pwning.eth]]"
  - "[[riptide]]"
  - "[[Ashiq Amien]]"
status: active
---

# Immunefi Bug Bounty Writeups List

A curated, community-maintained list of bug bounty writeups from **Immunefi**, the leading web3 bug bounty platform. Compiled by [sayan011](https://github.com/sayan011), it aggregates reports from 2021 onwards spanning over 100 protocols.

## Scope and Coverage

The list contains writeups from multiple platforms (Immunefi primary, plus Code4rena, Sherlock, and others). It covers:

- **Smart contract vulnerabilities**: reentrancy, logic errors, access control, oracle manipulation, rounding errors
- **Infrastructure / chain-level**: consensus bypass, block production shutdown, proof verification flaws
- **Bridge vulnerabilities**: uninitialized proxies, message spoofing, signature verification

## Key Data Points

- **Highest bounty**: $10M (Wormhole — uninitialized proxy bug)
- **Total bounty range**: $500 to $10M
- **Most active whitehat**: Trust Security (30+ reports across multiple protocols)
- **Most targeted protocol category**: DeFi (lending, AMM, yield aggregators)

## Notable Reports

| Bounty | Protocol | Vulnerability Class | Whitehat |
|--------|----------|-------------------|----------|
| $10M | Wormhole | Uninitialized proxy | satya0x |
| $6M | Aurora | Infinite spend | pwning.eth |
| $2M | Optimism | Money duplication | saurik |
| $1M+ | Moonbeam | Cross-chain logic | pwning.eth |
| $1M | Scroll | Message spoofing in bridge | WhiteHatMage |
| $1M | Balancer V2 | Logic error | kankodu |
| $800K | Fei | Flashloan vulnerability | Bobface |
| $560K | Redacted Cartel | Custom approval logic | Tommaso Pifferi |
| $505K | Raydium | Tick manipulation | riproprip |

## Vulnerability Patterns Observed

1. **Logic errors** (most common): incorrect state transitions, missing edge cases in business logic
2. **Access control**: missing privilege checks, unprotected administrative functions
3. **Rounding/precision**: financial calculation errors leading to inflation or theft
4. **Bridge/verification**: insufficient proof validation, signature spoofing
5. **Initialization**: uninitialized proxies, missing constructor calls

## Whitehat Profiles

The list reveals distinct researcher profiles:
- **Specialized hunters**: pwning.eth (infra/cross-chain), Trust Security (lending protocols)
- **Generalists**: riptide, Ashiq Amien (diverse protocol types)
- **Teams**: Zellic, ChainLight, Verichains (structured security firms)
- **Emerging**: GregoAI, Marco Nunes (2025+ active)

## Source

- Repository: https://github.com/sayan011/Immunefi-bug-bounty-writeups-list
- Immunefi: https://immunefi.com
- Bug fix reviews: https://reports.immunefi.com
