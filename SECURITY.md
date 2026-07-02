# Política de Seguridad de Bunker OS

## Reportar Vulnerabilidades

Si encontrás una vulnerabilidad de seguridad, **no abras un issue público**.

En cambio, contactá directamente:
- GitHub: [SamBleed](https://github.com/SamBleed)
- Repo: [opencode-obsidian](https://github.com/SamBleed/opencode-obsidian)

Las vulnerabilidades se atenderán dentro de las 48 horas hábiles.

## Alcance

Esta política cubre:
- Skills de OpenCode en `skills/`
- Scripts en `bin/` y `scripts/`
- Workflows n8n en `automation/n8n-lab/`
- Hooks en `hooks/hooks.json`
- Tests en `tests/`
- Configuración en `opencode.json`

## Medidas de Seguridad Existentes

### Secretos
- No hay API keys hardcodeadas en scripts (verificado por `make test-scripts`)
- Las keys viven en `.env` (gitignored) o en el vault de credenciales de n8n
- El test suite escanea patrones de secretos en cada corrida

### n8n
- El webhook del AOC v4 soporta `x-aoc-secret` header
- Rate limiting por source implementado en el pipeline
- Redacción de secrets en logs y payloads
- Dry-run mode por defecto (`AOC_DRY_RUN=true`)

### Red
- Ingest server bindea a `127.0.0.1:9090` por defecto
- n8n expone solo en localhost
- Autoresearch valida y sanitiza URLs antes de fetch
- Content sanitization: evita inyección de wikilinks y scripts en fuentes externas

### Dependencias
- n8n container con versión pinneada
- Escaneo de secretos en test suite
- npm audit como práctica recomendada

## Buenas Prácticas

1. **Nunca commitees** `.env`, API keys, tokens, o credenciales
2. **Usá variables de entorno** para configuración sensible
3. **Corré `make test`** antes de pushear (incluye escaneo de secretos)
4. **Mantené n8n actualizado** (docker pull regular)
5. **No exponás n8n a internet** sin reverse proxy con SSL
6. **Configurá `x-aoc-secret`** antes de activar el AOC v4

## Entorno Seguro por Defecto

```bash
# Verificar que no hay secretos en el repo
make test-scripts

# Verificar integrity del vault
./bin/bunker-check.sh

# Verificar que n8n no está expuesto
ss -tlnp | grep 5678  # Debería mostrar 127.0.0.1:5678 o 0.0.0.0:5678
```
