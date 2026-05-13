---
type: blueprint
title: "React Web Blueprint"
domain: tech-stack
created: 2026-05-13
updated: 2026-05-13
tags: [react, frontend, blueprint, architecture, nextjs]
status: mature
---

# React Web Blueprint (v2026)

This blueprint defines the standard for Next.js 15+ applications, focusing on Server Components, modern styling, and type-safe state management.

## 📂 Directory Structure

```text
.
├── app/                    # App Router (Next.js 15+)
│   ├── (auth)/             # Auth route group
│   ├── (dashboard)/        # Dashboard route group
│   └── api/                # Route Handlers
├── components/
│   ├── ui/                 # Atomic elements (shadcn/ui style)
│   ├── forms/              # Controlled/Uncontrolled forms
│   └── layouts/            # Shared layouts
├── lib/
│   ├── api/                # Fetch clients / Actions
│   ├── utils/              # Helper functions (cn, formatters)
│   └── hooks/              # Custom reusable hooks
├── types/                  # Global TS interfaces
├── styles/                 # Global CSS (Tailwind 4)
└── next.config.js
```

## 🏗️ Design Principles

1. **Server First**: Favor Server Components (RSC) for data fetching. Use Client Components only for interactivity.
2. **Actions for Mutations**: Use React 19 **Actions** for form submissions and state changes.
3. **Optimistic UI**: Implement `useOptimistic` for all high-frequency user actions.
4. **Type Safety**: Zero `any` usage. Strict TypeScript for all props and state.

## 🛠️ Mandatory Stack

- **Framework**: Next.js 15+ (App Router).
- **Styling**: Tailwind CSS 4 + `class-variance-authority`.
- **State**: React Context for UI state, TanStack Query for server state cache.
- **Forms**: `react-hook-form` + `zod` for validation.
- **Icons**: Lucide React.

## 🧪 Testing Checklist

- [ ] **Component Tests**: Playwright or Vitest for unit tests.
- [ ] **E2E Tests**: Playwright for critical user journeys (Login, Checkout).
- [ ] **Accessibility**: Use `axe-core` to verify WCAG compliance.
- [ ] **Linting**: Strict ESLint config + [[react-doctor]] for best practices.

## 🚀 Deployment Pattern

- **Platform**: Vercel or Self-hosted Docker (Standalone mode).
- **CI**: GitHub Actions for linting, type-checking, and E2E.
- **Environment**: Strict separation of public and secret env variables.

## Related
- [[React]]
- [[DESIGN-TEMPLATE]]
- [[tailwind-4]]
- [[react-doctor]]
