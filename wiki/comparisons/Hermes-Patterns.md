---
type: comparison
title: "Hermes Patterns"
updated: 2026-05-14
status: active
---

# Adopción de Patrones: Hermes Agent vs opencode-obsidian

> Análisis de qué conceptos de Hermes podemos integrar en nuestro flujo de trabajo para hacerlo más "autónomo" y "persistente".

## 🛠️ Patrones para Integrar

### 1. El Bucle de "Auto-Skill"
**Concepto Hermes**: El agente detecta que resolvió algo difícil y escribe un manual para la próxima vez.
**Integración Bunker**: 
- Implementar un trigger en el `bunker close` (handover) que pregunte: "¿Aprendimos un procedimiento nuevo hoy?"
- Si la respuesta es sí, crear automáticamente un borrador en `wiki/concepts/`.

### 2. Evaluación de Procedimiento (Self-Critique)
**Concepto Hermes**: El agente no solo termina, sino que juzga si lo hizo bien.
**Integración Bunker**: 
- Usar el skill `sdd-verify` de manera más agresiva.
- Guardar el "Verify Report" no solo como log, sino como lección aprendida en `wiki/meta/lessons-learned.md`.

### 3. Modelado de Entorno (Environment Sync)
**Concepto Hermes**: Sabe qué herramientas tenés instaladas y cómo usarlas.
**Integración Bunker**:
- Mejorar el `/sdd-init` para que el escaneo de testing y stacks sea más profundo.
- Mantener un archivo `wiki/meta/env-status.md` con las versiones de Go, Node, Docker, etc., para evitar errores de compatibilidad.

---

## 📈 Roadmap de Mejora

| Patrón | Dificultad | Impacto | Estado |
| :--- | :--- | :--- | :--- |
| **Auto-FAQ** | Baja | Alto | ✅ Implementado hoy |
| **Env-Status** | Media | Medio | ⏳ Pendiente |
| **Self-Critique** | Alta | Crítico | 🚀 En progreso (via sdd-verify) |

---
> 💡 **Conclusión**: Hermes nos enseña que el conocimiento no es solo *datos*, sino *procedimientos*. Debemos transformar más de nuestras charlas en **procedimientos accionables**.
