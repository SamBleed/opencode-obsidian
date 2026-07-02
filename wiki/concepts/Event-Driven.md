---
type: concept
title: "Event Driven"
updated: 2026-05-14
status: active
---

# Event-Driven Architecture (EDA)

**Status**: Core Standard
**Type**: Communication Patterns
**Parent**: [[_index]]

> Sistema de comunicación asíncrona basado en eventos, permitiendo que múltiples agentes y microservicios colaboren de forma desacoplada.

## 📡 Conceptos Clave

### 1. El Evento (The Event)
Un hecho inmutable que ha ocurrido en el sistema.
- **Header**: Metadata (timestamp, source_id, event_type, trace_id).
- **Payload**: El dato de lo que ocurrió (ej. `user.created`, `exploit.found`).

### 2. Pub/Sub (Publisher/Subscriber)
- **Publisher**: El que emite el evento sin saber quién lo escucha.
- **Subscriber**: El que escucha temas (topics) específicos y reacciona.

### 3. Broker (El Mensajero)
El software que gestiona la entrega de mensajes (ej. NATS, Redis Streams, RabbitMQ).

## 🤖 Colaboración Agéntica
En este Bunker, los agentes se hablan mediante eventos:
- **Agente Recon** publica `target.mapped`.
- **Agente Exploit** escucha `target.mapped` y publica `exploit.attempted`.
- **Agente Guardian** escucha todo para asegurar el [[Ethics-Compliance]].

---
**Last Updated**: 2026-04-19
