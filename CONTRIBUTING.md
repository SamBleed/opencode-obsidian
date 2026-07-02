# Contribuyendo a Bunker OS

PRs bienvenidas. Leé esto antes de abrir una.

---

## Workflow

1. Forkeá el repo y creá una branch desde `main`
2. Hacé tus cambios
3. Corré `make test` — todas las suites deben pasar
4. Si cambiás skills, actualizá los tests correspondientes
5. Commit con mensaje descriptivo
6. Abrí un PR a `main`

## Convenciones de Commits

Usamos [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: añadir /autoresearch
fix: corregir conexión en Dead Letter Queue
docs: actualizar README con comparativa
test: agregar test de workflows n8n
chore: actualizar docker-compose
```

## Self-Review Checklist

Antes de pedir review:

- [ ] `make test` pasa (429 tests, 0 failures)
- [ ] Si cambiás skills: actualizaste SKILL.md?
- [ ] Si cambiás scripts: actualizaste tests?
- [ ] Si cambiás estructura: actualizaste README.md / PROJECT.md / WIKI.md?
- [ ] No hay secretos hardcodeados
- [ ] Los mensajes de error son útiles en español
- [ ] Los cambios son backward-compatible o están documentados

## Tests

Todos los PRs deben pasar `make test`. El CI lo verifica automáticamente.

```bash
make test
```

## Reportar Issues

Usá GitHub Issues para:
- Bugs en skills o scripts
- Workflows n8n que fallan
- Documentación faltante o incorrecta
- Sugerencias de features

Incluí:
- Versión del Bunker (ver `PROJECT.md`)
- OpenCode version
- Pasos para reproducir
- Logs o mensajes de error relevantes

## Código de Conducta

Sé respetuoso. Este es un proyecto personal, no una corporación.
