---
tags: [ozyrecon, arquitectura]
created: 2026-04-19
---

# Arquitectura OzyRecon v4.0

## Visión General

OzyRecon es una **plataforma de reconocimiento ofensivo local-first** que se especializa en:
- Descubrimiento de superficie de ataque
- Detección de cambios entre escaneos (Diff Engine reactivo)
- Inteligencia adaptativa basada en memoria
- OPSEC de grado militar
- Export normalizado para OzyAudit

## Diferencia Clave: Plataforma vs Wrapper

| Wrapper | OzyRecon (Plataforma) |
|---------|----------------------|
| Ejecuta herramientas | Toma decisiones basadas en estado |
| Sin memoria | Persiste reputation, WAF, hallazgos |
| Pipeline fijo | Reacciona al delta automáticamente |
| Output heterogéneo | Schema normalizado único |
| Sin inteligencia | Genera IntelligenceBrief |

## Arquitectura de 8 Pilares

### 1. tool_manager + Capabilities
- `asset_discovery`
- `service_discovery`
- `template_scan`
- `port_scan`

### 2. Modos Operativos
- **HUNT**: Ejecución exhaustiva, línea base
- **CONTINUOUS**: Monitoreo reactivo basado en delta
- **RESEARCH**: Escaneo dirigido por tecnología/CVE
- **CAMPAIGN**: Escalado de patrones
- **FORENSIC**: Análisis histórico
- **SERVICIO**: Reportes ejecutivos

### 3. Memoria Táctica
- `AgentMemory` (razonamientos)
- `host_reputation` (historial de hallazgos)
- `waf_detected` (presencia de WAF)

### 4. Diff Engine Reactivo
- Detecta nuevos subdominios → Escanea
- Detecta cambios de versión → Investiga
- Detecta puertos cerrados → Limpia memoria

### 5. Priority Engine
- Scoring dinámico basado en:
  - Reputación histórica del host
  - Novedad del activo
  - Patrones detectados

### 6. OPSEC Integrado
- Pre-flight WAF detection
- Rate adaptation automática
- Kill-switch automático

### 7. Inteligencia Generada
- `IntelligenceBrief`:
  - Surface delta %
  - Nuevos endpoints críticos
  - Patrones de vulnerabilidad
  - Recomendaciones

### 8. Output Normalizado
- `ScanResult` schema:
  - JSON para OzyAudit
  - Markdown para clientes
  - CSV para Excel
  - Burp SAR para importar

## Flujo de Datos

```
Input → Modo → validation_preconditions
         ↓
    OPSEC (pre_flight_check) → Ajusta estrategia
         ↓
    Discovery (asset_discovery) → Subdomains
         ↓
    Priority (score_hosts) → Ordena por reputación
         ↓
    Service Discovery → Puertos/Servicios
         ↓
    Vulnerability Scan → Findings
         ↓
    Diff Engine (get_diff) → Detecta cambios
         ↓
    REACCIÓN AUTOMÁTICA → Escanea lo nuevo
         ↓
    Intelligence Brief → Recomendaciones
         ↓
    Export → ScanResult
```

## Estructura de Directorios

```
src/
├── core/
│   ├── tool_manager.py
│   ├── providers/
│   │   ├── base.py
│   │   ├── nuclei.py
│   │   ├── subfinder.py
│   │   ├── nmap.py
│   │   └── naabu.py
│   └── config.py
├── modes/
│   ├── base.py
│   ├── hunt.py
│   ├── continuous.py
│   ├── research.py
│   ├── campaign.py
│   ├── forensic.py
│   └── servicio.py
├── storage/
│   ├── models.py
│   ├── queries.py
│   ├── diff.py
│   └── database.py
├── intelligence/
│   ├── priority.py
│   ├── analyzer.py
│   ├── brief.py
│   └── enrichment.py
├── opsec/
│   ├── manager.py
│   ├── rate_limiter.py
│   ├── waf_detector.py
│   ├── kill_switch.py
│   └── jitter.py
└── export/
    ├── schema.py
    ├── normalizer.py
    └── platforms.py
```

##OPSEC: Comportamiento Adaptativo

```python
# Pre-flight
if waf_detected:
    threads = 10
    delay = 3
    strategy = "stealth"

# Rate adaptation
if 403/429 > 10:
    rpm /= 2
if consecutive_errors > 50:
    kill_switch.trigger()
```

## Output: Schema Normalizado

```json
{
  "type": "scan-result",
  "source": "ozy-recon",
  "version": "4.0",
  "session_id": "abc123",
  "target": "example.com",
  "mode": "hunt",
  "assets": [...],
  "services": [...],
  "findings": [...],
  "diff": [...],
  "intelligence": {
    "surface_delta_pct": 15.2,
    "new_critical_endpoints": ["api.internal"],
    "recommendations": [...]
  }
}
```

## Fase 2: Aprendizaje Reflexivo

| Capa | Componente | Función |
|------|----------|---------|
| 1 | DecisionLog | Persistir decisiones |
| 2 | OutcomeEvaluator | Medir resultado |
| 3 | FeedbackEngine | Recalibrar scoring |
| 4 | FalsePositiveMemory | Aprender de ruido |

## Links
- [[Fase1-Cierre]]
- [[modes]]
- [[opsec]]
- [[inteligencia]]