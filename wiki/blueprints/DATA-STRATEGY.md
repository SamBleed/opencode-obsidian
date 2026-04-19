# Data Strategy Blueprint (SQL/NoSQL)

> Blueprint para implementar persistencia robusta y compatible con agentes.

## 🏗️ Estructura de Directorios (Standard)
```text
/db
├── migrations/    # Archivos .sql o scripts de migración
├── seeds/         # Datos de prueba (JSON/CSV/SQL)
├── schema/        # Diagramas ER y documentación de tablas
└── docker-compose.db.yml # Definición del contenedor de DB
```

## 🛠️ Herramientas Recomendadas
- **Go**: `golang-migrate` para migraciones.
- **Node**: `Prisma` o `TypeORM`.
- **Python**: `Alembic` (SQLAlchemy).

## 🚀 Workflow de Datos para el Agente
1. **Init**: Agente lee `db/schema/` para entender la DB.
2. **Setup**: Agente corre `docker-compose up -d` para levantar la instancia.
3. **Migrate**: Agente aplica `db/migrations/` para versionar.
4. **Seed**: Agente carga `db/seeds/` para tener un entorno de pruebas funcional.

## 🔒 Security Check (Decepticon Style)
- ¿Se están exponiendo puertos de DB al exterior? → **NO**.
- ¿Las credenciales están en `.env` (ignoradas por git)? → **SÍ**.
- ¿El usuario de la DB tiene privilegios mínimos (Principle of Least Privilege)? → **SÍ**.

---
**Inspired by**: [[Data-Governance]]
