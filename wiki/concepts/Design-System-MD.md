# Design-System-MD (VoltAgent Pattern)

**Status**: Core Standard
**Type**: UI/UX Governance
**Parent**: [[_index]]

> Patrón para documentar sistemas de diseño en Markdown, optimizado para que agentes de IA generen interfaces visualmente coherentes.

## 🎨 El Concepto
En lugar de depender de herramientas visuales pesadas, el agente lee un archivo `DESIGN.md` que actúa como la **"Verdad Única"** del diseño.

## 📋 Estructura Recomendada
Un `DESIGN.md` efectivo debe contener:
1. **Core Palette**: Colores primarios, secundarios, fondos y estados (hex codes).
2. **Typography**: Fuentes, tamaños de escala (H1-H6, body) y line-heights.
3. **Spacing & Grid**: Sistema de 8px, paddings y margins estándar.
4. **Components**: Definición de botones, inputs y cards.
5. **Dark Mode**: Ajustes específicos para temas oscuros.

## 🤖 Beneficios para el Agente
- **Cero Ambigüedad**: El agente sabe exactamente qué hex code usar.
- **Consistencia**: Todos los componentes generados siguen la misma línea visual.
- **Eficiencia**: No requiere procesar imágenes o Figma tokens complejos.

---
**Last Updated**: 2026-04-19
