# Messaging Strategy Blueprint

> Blueprint para implementar comunicación asíncrona robusta entre agentes y servicios.

## 🏗️ Estructura del Mensaje (Standard)
```json
{
  "specversion": "1.0",
  "type": "com.bunker.agent.action",
  "source": "/agents/decepticon-01",
  "id": "A234-1234-1234",
  "time": "2026-04-19T17:31:00Z",
  "traceparent": "00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01",
  "data": {
    "action": "recon_complete",
    "details": { "host": "10.0.0.5", "ports": [80, 443] }
  }
}
```

## 🛠️ Stack Recomendado para Arch Linux
- **NATS**: Ultra-liviano, escrito en Go, perfecto para sistemas locales.
- **Redis Streams**: Si ya usás Redis para cache, es el camino más fácil.

## 🚀 Workflow Agéntico Asíncrono
1. **Emit**: Agente A termina una tarea y emite un evento al Broker.
2. **Route**: El Broker entrega el mensaje a todos los interesados.
3. **Trigger**: Agente B se activa automáticamente al recibir el mensaje.
4. **Log**: [[Observability-MD]] captura el flujo completo mediante el `trace_id`.

## 🔒 Security (Decepticon Style)
- **Topic Isolation**: Cada agente solo puede publicar/suscribirse a sus temas autorizados.
- **Payload Encryption**: Si el mensaje viaja por red no segura, cifrar el payload.

---
**Inspired by**: [[Event-Driven]]
