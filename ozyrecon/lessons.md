---
tags: [ozyrecon, lessons, arquitectura]
created: 2026-04-19
---

# Lecciones Aprendidas - OzyRecon v4.0

## La Diferencia Real: Wrapper vs Plataforma

> La diferencia entre un wrapper y una plataforma está en: **memoria + decisión + reacción**.

### Wrapper
- Ejecuta herramientas secuencialmente
- Sin estado entre ejecuciones
- Output heterogéneo
- Sin inteligencia

### Plataforma
- Toma decisiones basadas en estado previo
- Persiste memoria entre ejecuciones
- Reacciona al delta automáticamente
- Genera inteligencia accionable

## Errores Comunes evitados

### 1. Abstracción falsas
- NO: "Usar capacidades" mean solo envolver comandos
- SÍ: El ToolManager resuelve qué proveedor usar basado en disponibilidad

### 2. Modos como包装
- NO: "El mismo pipeline con flags distintos"
- SÍ: Cada modo tiene intención operativa distinta y precondiciones únicas

### 3. Diff decorativo
- NO: "Imprimir cambios al final"
- SÍ: El Diff debe DISPARAR acciones automáticamente

### 4. OPSEC de configuración
- NO: "Jitter = random sleep"
- SÍ: Rate adaptation + WAF detection + kill-switch automático

### 5. Output sin schema
- NO: "Cada herramienta su formato"
- SÍ: Schema común (ScanResult) primero

## Lo que funcionó

### Arquitectura modular real
- Proveedores en capa neutral (`src/core/providers/`)
- Modos heredan de `BaseMode` con contrato explícito
- DiffEngine retorna `DiffReport` actionables
- IntelligenceBrief con recomendaciones concretas

### Memoria integrada
- `AgentMemory` para decisiones tácticas
- `host_reputation` para scoring histórico
- `seen_count` para deduplicación

### Reacción automática
```python
if diff_report.new_subdomains:
    _scan_new_assets()  # Automático
if diff_report.changed_services:
    _investigate_changed()  # Automático
```

## Roadmap correcto

| Fase | Enfoque | Estado |
|------|--------|--------|
| 1 | Adaptativo (reacciona) | ✅ Completado |
| 2 | Reflexivo (evalúa) | 🔲 Pendiente |

## Criterios de validación

1. Coherencia interna: todo Lo declarado está implementado
2. Conexión: todo implementado está conectado
3. Impacto: todo conectado modifica comportamiento

## Links
- [[Fase1-Cierre]]
- [[arquitectura]]