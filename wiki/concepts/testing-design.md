---
type: concept
title: "Testing Design"
domain: engineering
created: 2026-07-02
updated: 2026-07-02
tags: [testing, architecture, frontend, ui]
status: mature
related:
  - "[[frontend-ui-engineering]]"
---

# Testing Design

## Overview
Testing Design refers to the architectural strategy of how to test applications, specifically focusing on UI and frontend components. It defines the boundaries between unit, integration, visual regression, and end-to-end (E2E) tests.

## Testing Layers

1. **Unit Testing**: Testing isolated functions, hooks, or pure components. Mocking external dependencies.
2. **Component/Integration Testing**: Testing components in context, ensuring they interact correctly with the DOM and state management (e.g., using Testing Library).
3. **Visual Regression Testing**: Capturing screenshots of UI components and comparing them against baselines to catch unintended styling drift.
4. **End-to-End (E2E) Testing**: Simulating full user flows across the complete stack (e.g., Playwright, Cypress).

## Best Practices
- Test behavior, not implementation details.
- Use semantic queries (`getByRole`, `getByText`) instead of test IDs when possible, as this also acts as an accessibility check.
- Maintain a stable testing pyramid (many unit tests, moderate integration, few high-value E2E).

## Senior Gotchas & Anti-patterns
- **The Implementation Trap**: Asserting that a specific CSS class exists instead of verifying the visual or structural outcome (e.g., checking if a button has `.bg-red-500` instead of `aria-invalid="true"`).
- **Flaky E2E Tests**: Relying on arbitrary `setTimeout` instead of awaiting state changes, element visibility, or network idle.
- **Over-mocking**: Mocking too many internal dependencies in integration tests, resulting in tests that pass even when the actual component is fundamentally broken in production.
- **Testing the Framework**: Writing tests to verify that React updates state when a setter is called, rather than testing the business logic or user outcome.
