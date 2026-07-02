---
type: concept
title: Performance Optimization
domain: engineering
created: 2026-05-16
updated: 2026-05-16
tags: [performance, web-vitals, optimization]
sources: [[agent-skills]]
---

# Performance Optimization

## Definition
Measure before optimizing. Profile first, identify the actual bottleneck, fix it, measure again. Optimize only what measurements prove matters.

## Core Web Vitals Targets
| Metric | Good | Needs Improvement | Poor |
|--------|------|-------------------|------|
| LCP | ≤ 2.5s | ≤ 4.0s | > 4.0s |
| INP | ≤ 200ms | ≤ 500ms | > 500ms |
| CLS | ≤ 0.1 | ≤ 0.25 | > 0.25 |

## The Workflow
1. **MEASURE** → Establish baseline with real data (synthetic + RUM)
2. **IDENTIFY** → Find the actual bottleneck (not assumed)
3. **FIX** → Address the specific bottleneck
4. **VERIFY** → Measure again, confirm improvement
5. **GUARD** → Add monitoring or tests to prevent regression

## Key Patterns
- Use both synthetic (Lighthouse) and RUM (web-vitals library)
- Optimize bundle size, images, queries, caching
- Code splitting and lazy loading for frontend
- N+1 query detection and indexing for backend
- Never optimize without evidence of a problem

## Red Flags
- Optimizing without measuring first
- No Core Web Vitals monitoring in production
- Large unoptimized images
- No code splitting in SPA

## Source
- [[agent-skills]] — Performance Optimization skill
