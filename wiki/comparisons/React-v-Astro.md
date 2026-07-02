---
type: comparison
title: "React v Astro"
updated: 2026-05-14
status: active
---

# React vs Astro para Websites

> ¿Cuándo usar React (SPA) vs Astro (Multi-Page)?

## Comparativa Rápida

| Criterio | React | Astro |
|----------|-------|-------|
| **Philosophy** | Everything JS | Islands Architecture |
| **Initial bundle** | 100-200KB | 0KB (static) |
| **Hydration** | Full | Parcial (on interaction) |
| **SEO** | JS-dependent | Built-in |
| **Interactivity** | Global | Per-component |
| **Learning curve** | Medio | Bajo |

## Cuándo React

- **Web apps complejas** — dashboards, calculators
- **Estado global** — múltiples usuarios, auth
- **Real-time** — websockets, live updates
- **SPA feel** — transitions fluidas
- **El usuario está logged-in siempre**

Ejemplos: Gmail, Linear, Figma

## Cuándo Astro

- **Marketing sites** — landing pages, blogs
- **Content-heavy** — docs, portfolios
- **SEO priority** — search indexing
- **Performance** — Lighthouse score > 95
- **Mostly static** — 90% read, 10% write

Ejemplos: Docusaurus, Stripe docs, blogs

## La Hibrid Approach (Astro + React)

```astro
---
// Ejemplo: Astro con islands de React
---
<Layout>
  <!-- Static: 0KB JS -->
  <Hero />
  <Features />
  
  <!-- Island: solo cuando el usuario(scroll) -->
  <Search client:visible />
  
  <!-- Island: solo cuando hace click -->
  <Newsletter client:idle />
</Layout>
```

## Regla de Thumb

```
¿El usuario VA a interactuar mucho? → React
¿El usuario VA a LEER más que interactura? → Astro
```

## En Este Proyecto

- **Astro** → Blog, docs, marketing
- **React** → Solo si necesito estado complejo

> 💡 La mayoría de los sitios pueden ser 90% Astro + 10% React.