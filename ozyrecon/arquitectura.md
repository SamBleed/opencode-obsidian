---
tags:
  - ozyrecon
  - arquitectura
  - v5.7
created: '2026-04-24T00:00:00.000Z'
version: '5.7'
---

# Arquitectura OzyRecon v5.7

## Visión General

OzyRecon es una **plataforma de validación ofensiva asistida** que se especializa en:

- Descubrimiento inteligente de superficie de ataque
- Detección de cambios entre escaneos (Diff Engine reactivo)
- Generación de hipótesis con correlación de señales
- **Human-in-the-loop Control (Human Gate)** — autorización manual obligatoria
- **Evidence Engine** con integridad criptográfica (SHA256)
- **Knowledge Graph** para visualización de relaciones
- OPSEC de grado militar con kill-switch automático
- Export normalizado para OzyAudit

## Diferencia Clave: Scanner vs Plataforma de Validación

| Scanner Tradicional | OzyRecon (Plataforma de Validación) |
|--------------------|-------------------------------------|
| Genera ruido | Genera inteligencia validada |
| Sin memoria | Persiste estado, reputación, WAF |
| Sin evidencia | Captura screenshots con hash SHA256 |
| Sin control | Human Gate obligatorio |
| Output heterogéneo | Schema normalizado unico |

## Pipeline de Validación

```
Discovery → Hypothesis → Human Gate (Approval) → Validation → Evidence → Report
```

### Human Gate
Todas las hipótesis generadas requieren autorización manual antes de ejecución.

### Evidence Engine
Cada finding validado incluye:
- Screenshot como evidencia visual
- Hash SHA256 para integridad criptográfica
- Trazabilidad completa

### Knowledge Graph
Visualiza relaciones entre:
- Targets
- Endpoints descubiertos
- Patrones de vulnerabilidades
- Hallazgos históricos

## Arquitectura de 8 Pilares

### 1. Tool Manager + Capabilities
Abstracción de herramientas basada en capacidades:
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
- `seen_count` (recurrencia de findings)

### 4. Diff Engine Reactivo
- Detecta nuevos subdominios → Escanea
- Detecta cambios de versión → Investiga
- Detecta puertos cerrados → Limpia memoria

### 5. Priority Engine
Scoring dinámico basado en:
- Reputación histórica del host
- Novedad del activo
- Patrones detectados

### 6. OPSEC Integrado
- Pre-flight WAF detection
- Rate adaptation automática
- Kill-switch automático
- Identity rotation

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

## Estructura de Directorios (v5.7)

```
src/
├── core/
│   ├── api.py                    # API + Knowledge Graph endpoint
│   ├── tool_manager.py
│   ├── providers/                # subfinder, nmap, shodan, censys...
│   ├── config.py
│   └── ...
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
│   ├── enrichment.py
│   ├── learning_orchestrator.py   # v5.7: Feedback loop
│   ├── outcome_evaluator.py
│   ├── feedback_engine.py
│   ├── false_positive_memory.py
│   └── dashboard.py
├── opsec/
│   ├── manager.py
│   ├── rate_limiter.py
│   ├── waf_detector.py
│   ├── kill_switch.py
│   ├── jitter.py
│   └── identity_rotation.py
├── validation/
│   ├── base.py
│   ├── auth.py                   # v5.7: Auth Validator
│   ├── http.py
│   ├── web.py
│   ├── cms.py
│   ├── config.py
│   ├── infra.py
│   └── automation.py
├── evidence/
│   └── engine.py                 # v5.7: Evidence + Screenshots + SHA256
├── gate/
│   └── manager.py                # Human-in-the-loop control
├── discovery/
│   ├── assets/
│   ├── services/
│   ├── targets/
│   └── crawler.py
├── scanners/
│   ├── templates/
│   ├── web/
│   └── db/
├── export/
│   ├── schema.py
│   ├── normalizer.py
│   └── platforms.py
├── workflow/
│   ├── orchestrator.py
│   ├── engine.py
│   └── states.py
├── reporting/
│   └── engine.py
├── utils/
│   ├── visual.py                 # v5.7: Visual utilities
│   └── benchmark.py
└── notifications/
    ├── telegram.py
    └── notifier.py
```

## v5.7: Nuevas Capacidades

### Knowledge Graph
Endpoint `/api/graph` que proporciona visualización de:
- Relaciones entre targets
- Endpoints conectados
- Patrones de exposición

### Auth Validator
Validación no intrusiva de credenciales por defecto:
- Detecta potential credential leaks
- Filosofía "No Exploitation"
- No explota, solo valida exposición

### Visual Evidence Engine
- Captura screenshots automatizada
- Hash SHA256 por cada evidencia
- Integridad criptográfica verificable

## Output: Schema Normalizado v5.7

```json
{
  "type": "scan-result",
  "source": "ozy-recon",
  "version": "5.7",
  "session_id": "abc123",
  "target": "example.com",
  "mode": "hunt",
  "assets": [...],
  "services": [...],
  "findings": [
    {
      "finding": "Exposed admin endpoint",
      "severity": "HIGH",
      "confidence": 0.92,
      "confidence_reason": "Correlation of multiple independent signals",
      "evidence_id": "ev_8a2f1c9",
      "integrity_hash": "sha256:7f83b1..."
    }
  ],
  "diff": [...],
  "knowledge_graph": {...},
  "intelligence": {
    "surface_delta_pct": 15.2,
    "new_critical_endpoints": ["api.internal"],
    "recommendations": [...]
  }
}
```

## Filosofia: NO EXPLOITATION

- **Zero Impact**: No destructive payloads
- **Surgical Probes**: Confirm exposure; don't exploit
- **Audit Focused**: Verifiable proof for remediation teams
- **Privacy First**: Built-in anonymization and OPSEC guards

## Links
- [[Fase1-Cierre]] - Cierre Fase 1
- [[lessons]] - Lecciones aprendidas
