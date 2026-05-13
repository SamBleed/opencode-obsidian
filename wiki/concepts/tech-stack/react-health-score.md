---
type: concept
title: "React Health Score"
domain: tech-stack
created: 2026-05-13
updated: 2026-05-13
tags: [metrics, react, quality]
status: seed
---

# React Health Score

The React Health Score is a metric used by [[react-doctor]] to quantify the quality of a React codebase. It focuses on unique rules triggered rather than total instances of violations.

## Formula
The score is calculated as:
`100 - (unique_error_rules * 1.5) - (unique_warning_rules * 0.75)`

## Labels
- **75+**: Great
- **50 to 74**: Needs work
- **Under 50**: Critical

## Philosophy
The scoring system encourages eliminating categories of bugs rather than just individual instances. Fixing 100 instances of the same warning rule yields the same score improvement as fixing 1 instance, as long as it's the *only* instance left. This prioritizes systemic fixes.

## Source
- [[millionco-react-doctor-repo]]
