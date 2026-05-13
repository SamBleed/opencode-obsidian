---
type: concept
title: "Tailwind CSS 4 Mastery"
domain: tech-stack
created: 2026-05-13
updated: 2026-05-13
tags: [frontend, styling, tailwind, css, design-system]
status: mature
---

# Tailwind CSS 4 Mastery (v2026)

Tailwind CSS v4 is a ground-up rewrite featuring an engine powered by Rust (Oxide) and a **CSS-first configuration** paradigm. It eliminates the need for `tailwind.config.js` and embraces native CSS modern features.

## 🎨 The CSS-First Revolution

In v4, your configuration lives directly in your CSS file using the `@theme` directive. This bridges the gap between design tokens and utility classes.

```css
@import "tailwindcss";

@theme {
  --color-brand-primary: oklch(0.6 0.2 250);
  --color-brand-secondary: oklch(0.8 0.1 150);
  
  --font-display: "Satoshi", "sans-serif";
  --breakpoint-3xl: 1920px;
}
```

### OKLCH Colors
We favor **OKLCH** over RGB/HEX. It provides perceptually uniform brightness, making it easier to create accessible color palettes and dark mode variants.

## 🏗️ Senior Patterns: Atomic UI with CVA

To handle complex component variants without "class name soup," we use **Class Variance Authority (CVA)**.

```tsx
const button = cva("rounded-md px-4 py-2 transition-all", {
  variants: {
    intent: {
      primary: "bg-brand-primary text-white hover:bg-opacity-90",
      ghost: "bg-transparent border-2 border-brand-primary",
    },
    size: {
      sm: "text-sm",
      lg: "text-lg",
    },
  },
  defaultVariants: {
    intent: "primary",
  },
});
```

## 📦 Micro-Frontends & Shared Tokens

In a distributed frontend architecture:
1. **Shared Tokens Package**: Publish a `@org/tokens` package containing `tokens.css` and a Tailwind v4 preset.
2. **Runtime Theming**: Inject CSS variables at the shell level. Remotes consume `var(--color-brand-primary)` automatically.
3. **Isolation**: Use Tailwind prefixes for remotes to avoid style collisions.

## ⚡ Next-Gen Layout Features

### Container Queries
No more plugins. Use the `@container` class and `@sm:`, `@md:` variants to style components based on their parent's width, not the viewport.
```html
<div class="@container">
  <div class="grid grid-cols-1 @md:grid-cols-3">...</div>
</div>
```

### Native Nesting
Tailwind v4 supports native CSS nesting out of the box. Use it to keep your custom styles clean and co-located.

## 🛡️ Performance Checklist
- [ ] **Oxide Engine**: Instant HMR (Hot Module Replacement).
- [ ] **No Unused Classes**: Oxide's compiler is 10x faster at purging.
- [ ] **Preflight Control**: Disable preflight in remotes, keep it only in the shell.

## Related
- [[React]]
- [[REACT-WEB-BLUEPRINT]]
- [[DESIGN-TEMPLATE]]
- [[performance-optimization]]
