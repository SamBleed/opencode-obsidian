---
type: concept
title: "Offensive Vaccine"
updated: 2026-05-14
status: active
---

# Offensive Vaccine (Active Auditing)

**Status**: Core Standard
**Type**: Security Pattern
**Parent**: [[_index]]

> Patrón de seguridad inspirado en [[Decepticon]] y [[PayloadsAllTheThings]], donde el agente "ataca" su propio código preventivamente usando payloads reales para validar las defensas.

## 🧪 El Proceso
1. **Target Identification**: Identificar los puntos de entrada (inputs, endpoints, formularios).
2. **Payload Selection**: Consultar [[PayloadsAllTheThings]] para obtener los payloads más relevantes para el stack (ej. SQLi para Go/PostgreSQL).
3. **Automated Attack**: Correr una subrutina que inyecte estos payloads en un entorno de staging/local (vía [[Lightpanda]] o [[Trivy]]).
4. **Validation**: Si el sistema falla o permite la ejecución, se documenta como vulnerabilidad crítica y se bloquea el despliegue.

## 🛡️ Beneficios
- **Defensa Basada en la Realidad**: No solo sigues guías, te proteges contra ataques que existen en el mundo real.
- **Zero Hallucinations**: Al usar payloads verificados, el agente tiene una base sólida de comparación.
- **Continuous Improvement**: Cada nuevo bypass descubierto se integra en el ciclo de la "vacuna".

---
**Last Updated**: 2026-04-19
