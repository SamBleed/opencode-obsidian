# Observability-MD

**Status**: Core Standard
**Type**: Infrastructure / Monitoring
**Parent**: [[_index]]

> Patrón para documentar y estandarizar la observabilidad de los agentes y servicios, asegurando que los fallos sean detectables y trazables por IAs.

## 👁️ Los Tres Pilares (AI-Readable)

### 1. Structured Logging (JSON)
Los logs no deben ser texto plano. Deben ser objetos JSON que el agente pueda parsear:
```json
{
  "timestamp": "2026-04-19T...",
  "level": "ERROR",
  "agent_id": "decepticon-01",
  "trace_id": "abc-123",
  "message": "Exploit failed: connection timeout",
  "context": { "target": "10.0.0.5", "tool": "nmap" }
}
```

### 2. Trace-ID Propagation
Cada acción iniciada por un agente debe llevar un `trace_id` único. Este ID debe viajar a través de todos los servicios (API, DB, Worker) para reconstruir la historia del fallo.

### 3. Agent Execution Traces (Estilo VoltOps)
Documentar el "pensamiento" del agente:
- **Input**: El prompt original.
- **Plan**: Pasos decididos.
- **Tools**: Herramientas llamadas y sus outputs.
- **Outcome**: Resultado final.

## 🚨 Alerting Guardrails
- **Thresholds**: Definir qué es un "error crítico" vs. un "warn".
- **AI-Summary**: Al dispararse una alerta, el sistema debe generar un mini-resumen para que el humano entienda el contexto sin leer mil logs.

---
**Last Updated**: 2026-04-19
