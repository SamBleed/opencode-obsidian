# Blueprints Engine

**Status**: Infrastructure
**Type**: Template Index
**Parent**: [[_index]]

> Repositorio de arquitecturas probadas y seguras. Los Blueprints son el "punto de partida" para cualquier nuevo desarrollo en el Bunker.

## 📐 Blueprints Disponibles

### [Go API Blueprint](wiki/blueprints/GO-API-BLUEPRINT.md)
- **Stack**: Go 1.26, Chi/Fiber, PostgreSQL.
- **Patterns**: Hexagonal, Error Wrapping, Structured Logging.
- **Security**: RBAC, JWT, Rate Limiting.

### [React Web Blueprint](wiki/blueprints/REACT-WEB-BLUEPRINT.md)
- **Stack**: React 19, Vite, Tailwind 4.
- **Patterns**: Atomic Design, [[Design-System-MD]].
- **Design System**: Refer to [[DESIGN-TEMPLATE.md]] for visual guardrails.

### [Messaging Strategy Blueprint](wiki/blueprints/MESSAGING-STRATEGY.md)
- **Stack**: NATS, Redis Streams.
- **Patterns**: [[Event-Driven]], Pub/Sub, CloudEvents.
- **Goal**: Decoupled multi-agent collaboration.

## 🚀 Cómo usar un Blueprint
1. El usuario elige un blueprint.
2. El agente ejecuta `/sdd-init` basado en el blueprint.
3. Se generan los archivos base automáticamente con [[Security-Guardrails]] aplicados.

---
**Last Updated**: 2026-04-19
