---
type: concept
title: "Web3 Bug Bounty Methodology"
domain: security
tags: [concept, security, web3, bug-bounty]
status: seed
sources:
  - "[[immunefi-bug-bounty-writeups]]"
---

# Web3 Bug Bounty Methodology

Common patterns and approaches used by whitehat researchers to find vulnerabilities in web3 protocols, based on analysis of 100+ Immunefi writeups.

## Common Vulnerability Classes

### Logic Errors (most frequent)
Incorrect state transitions, missing business logic edge cases, or flawed accounting assumptions. Often found in lending protocols where collateral calculations, liquidation thresholds, or interest rate models have edge cases.

### Access Control Flaws
Missing `onlyOwner` modifiers, unprotected administrative functions, or incorrect permission checks. Common in upgradeable contracts where initialization functions can be called by anyone.

### Precision / Rounding Errors
Financial calculations using integer division that round in the attacker's favor. Found in AMMs, yield aggregators, and any protocol handling fractional amounts.

### Bridge Vulnerabilities
Uninitialized proxy contracts, insufficient proof verification, message spoofing, or signature replay across chains.

## Researcher Approaches

1. **Code review first**: Manual audit of critical functions (accounting, transfers, permissions)
2. **Integration testing**: Testing interactions between contracts (lender + liquidator, etc.)
3. **Economic analysis**: Finding arbitrage opportunities that indicate financial logic flaws
4. **Fuzzing**: Input validation edge cases in parsing functions

## Methodology Evolution

| Period | Focus | Example |
|--------|-------|---------|
| 2021-2022 | Basic access control, reentrancy | 88mph, Zapper |
| 2022-2023 | Complex DeFi logic errors | Fei, Balancer, Enzyme |
| 2023-2024 | Cross-chain, bridge, L2 | zkSync, Optimism, Polygon |
| 2024-2025 | Infrastructure, chain-level | Sei, Story, Movement |

## Source

- [[immunefi-bug-bounty-writeups]]
