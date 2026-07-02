---
type: concept
title: "Data Governance"
updated: 2026-05-14
status: active
---

# Data Governance & Strategy

**Status**: Core Standard
**Type**: Data Infrastructure
**Parent**: [[_index]]

> Patrón para la gestión, versionado y seguridad de los datos dentro del ecosistema del Bunker.

## 💾 Ciclo de Vida del Dato

### 1. Schema Management (Migrations)
- **Declarativo**: El esquema de la DB debe estar definido en código (ej. SQL migrations, Prisma, SQLAlchemy).
- **Versionado**: Cada cambio en el esquema debe ser una migración inmutable con timestamp.
- **Rollback**: Siempre debe existir un camino de vuelta en caso de fallo.

### 2. Data Seeding (AI-Ready)
Para que un agente pueda testear, necesita datos de prueba realistas pero seguros.
- **Anonymization**: Prohibido usar datos reales de producción para seeding.
- **Deterministic**: Los seeds deben generar el mismo estado inicial siempre para tests reproducibles.

### 3. Data Sovereignty (Soberanía)
- **Local First**: Los datos residen en el entorno de Sam (Arch Linux).
- **Encryption at Rest**: Las DBs locales deben estar cifradas si contienen info sensible.
- **Backup Strategy**: Sincronización periódica de volúmenes de DB.

## 📊 Integración con Agentes
El agente debe tener acceso a:
- Esquema de la DB (DDL).
- Diccionario de datos (descripción de tablas y campos).
- Scripts de seeding para levantar entornos de staging rápidos.

---
**Last Updated**: 2026-04-19
