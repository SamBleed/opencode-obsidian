---
type: concept
title: "DESIGN.md"
complexity: intermediate
domain: design-systems
created: 2026-04-20
updated: 2026-04-20
tags:
  - concept
  - design-system
  - ai
status: mature
related:
  - "[[Design-System-MD]]"
  - "[[awesome-claude-design]]"
  - "[[VoltAgent]]"
sources:
---

# DESIGN.md

Archivo markdown que describe la identidad visual de una marca en formato que agentes IA pueden leer y actuar.

---

## Definition

Un archivo markdown estructurado que define tokens de diseño (colors, typography, spacing, components) de manera que un agente IA puede interpretar y generar código de UI automáticamente.

## How It Works

1. Escribes un archivo `DESIGN.md` con secciones específicas
2. Lo pasas a un agente de diseño (ej: Claude Design)
3. El agente parsea el archivo y scaffoldea un design system completo

## The 9 Sections

| # | Section | What Defines |
|---|---------|-------------|
| 1 | Visual Theme & Atmosphere | Tone, density, mood |
| 2 | Color Palette & Roles | CSS variables + semantic names |
| 3 | Typography Rules | Type scale, Google Fonts |
| 4 | Component Stylings | Buttons, inputs, cards, nav |
| 5 | Layout Principles | Spacing, grid, rhythm |
| 6 | Depth & Elevation | Shadows, surface hierarchy |
| 7 | Do's and Don'ts | Guardrails para el agente |
| 8 | Responsive Behavior | Breakpoints, touch targets |
| 9 | Agent Prompt Guide | Prompts reutilizables |

## Comparison

| File | Who Reads | What Defines |
|------|----------|--------------|
| AGENTS.md | Coding agents | Cómo construir el proyecto |
| DESIGN.md | Design agents | Cómo se ve y siente el proyecto |

## Tools Using This

- Claude Design (claude.ai/design)
- Google Stitch
- [[VoltAgent]]/awesome-claude-design (colección de 68 ejemplos)

## Examples

Los 68 ejemplos de awesome-claude-design incluyen:
- AI/LLM: Claude, Cohere, ElevenLabs, Ollama
- Developer Tools: Cursor, Vercel, Raycast
- Fintech: Stripe, Coinbase, Revolut
- E-commerce: Airbnb, Nike, Shopify

## Example DESIGN.md Template

```markdown
# Brand Name

## Visual Theme & Atmosphere
- Tone: Premium, minimalist
- Density: Comfortable, ample whitespace
- Mood: Trustworthy, modern

## Color Palette & Roles
- Primary: #0A0A0A (void black)
- Accent: #10B981 (emerald)
- Surface: #FFFFFF / #0A0A0A (light/dark)
- Text: #1F2937 / #F9FAFB

## Typography Rules
- Display: Geist (Vercel)
- Body: Inter
- Scale: 12/14/16/20/24/32/48/64

## Component Stylings
- Border radius: 8px (buttons), 12px (cards)
- Shadows: 0 4px 6px -1px rgba(0,0,0,0.1)
- States: hover (scale 1.02), focus (ring 2px)

## Do's and Don'ts
- DO: Use semantic color roles (primary, not blue)
- DON'T: Mix more than 3 accent colors
```

## Patterns

- **Token-based**: Colors as CSS variables, not hardcoded
- **Semantic naming**: `bg-primary`, `text-accent` not `#fff`, `blue`
- **Component tokens**: Button, Input, Card, Nav with variants

## Source

- https://github.com/VoltAgent/awesome-claude-design
- https://getdesign.md/what-is-design-md
- [[Context7]]