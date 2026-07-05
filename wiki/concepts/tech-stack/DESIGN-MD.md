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
  - "[[Awesome-DESIGN-md]]"
  - "[[VoltAgent]]"
sources:
  - "[[awesome-design-md]]"
---

# DESIGN.md

Archivo markdown que describe la identidad visual de una marca en formato que agentes IA pueden leer y actuar. Este estándar fue popularizado por Google Stitch y expandido por la comunidad.

---

## Definition

Un archivo markdown estructurado que define tokens de diseño (colors, typography, spacing, components) de manera que un agente IA puede interpretar y generar código de UI automáticamente. Es el "manual de marca" para la inteligencia artificial.

## How It Works

1. **Definición**: Escribes un archivo `DESIGN.md` con secciones específicas (Tokens, Colores, Tipografía).
2. **Contexto**: Lo incluyes en la raíz del proyecto junto a tu `AGENTS.md`.
3. **Generación**: Le pides al agente (ej: OpenCode, Claude Design): "Crea una página siguiendo la especificación de DESIGN.md".
4. **Resultado**: El agente genera código (React, Tailwind, etc.) que respeta los tokens y el estilo definido.

## The 9 Sections (Standard)

| # | Section | What Defines |
|---|---------|-------------|
| 1 | Visual Theme & Atmosphere | Tono, densidad, mood visual |
| 2 | Color Palette & Roles | Variables CSS + nombres semánticos |
| 3 | Typography Rules | Escala tipográfica, fuentes, pesos |
| 4 | Component Stylings | Botones, inputs, cards, navegación |
| 5 | Layout Principles | Spacing, grid, ritmo, márgenes |
| 6 | Depth & Elevation | Sombras, jerarquía de superficies |
| 7 | Do's and Don'ts | Guías y prohibiciones de diseño |
| 8 | Responsive Behavior | Breakpoints, comportamiento en móviles |
| 9 | Agent Prompt Guide | Prompts listos para usar |

## Comparison

| File | Who Reads | What Defines |
|------|----------|--------------|
| `AGENTS.md` | Coding agents | Cómo construir y operar el proyecto |
| `DESIGN.md` | Design agents | Cómo se ve y siente el proyecto |

## Tools & Ecosystem

- **Google Stitch**: El origen del concepto.
- **[[Awesome-DESIGN-md]]**: Colección curada de 69+ especificaciones de sitios reales (Stripe, Vercel, Apple, etc.).
- **getdesign.md**: Interfaz web para explorar el catálogo.

## Examples Categorized

Gracias a [[Awesome-DESIGN-md]], podemos imitar estilos de:
- **AI/LLM**: Claude, Cohere, Ollama, OpenCode AI.
- **Developer Tools**: Cursor, Vercel, Expo, Supabase.
- **Fintech**: Stripe, Revolut, Binance.
- **SaaS/Productivity**: Linear, Notion, Cal.com.
- **High-End**: Tesla, Ferrari, Apple, Nike.

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

## Sources

- [Awesome DESIGN.md GitHub](https://github.com/VoltAgent/awesome-design-md)
- [getdesign.md](https://getdesign.md)
- [[Context7]]
