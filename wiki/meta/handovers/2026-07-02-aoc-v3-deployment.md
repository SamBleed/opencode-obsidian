---
date: 2026-07-02
status: "Completed"
type: "handover"
tags: [n8n, aoc, architecture, docker]
---

# Handover: AOC v3 Deployment & V4 Blueprint

## 🎯 Objetivo de la Sesión
Desplegar la infraestructura local del laboratorio de n8n, estabilizar el repositorio y plantear la hoja de ruta para la automatización de nivel Enterprise del Bunker OS.

## ✅ Tareas Completadas
1. **Despliegue de Docker**: Se descargó e inicializó exitosamente la imagen de n8n (`n8n_meco_lab`) en el puerto `:5678`.
2. **Reestructuración de Directorios**: Se aplanó y corrigió el árbol de directorios de `/automation/n8n-lab/`, moviendo los archivos `.env` y `docker-compose.yml` a la raíz del laboratorio y consolidando los flujos JSON.
3. **Mejoras Visuales (Branding)**: 
   - Se inyectó un nuevo banner espectacular en el `README.md`.
   - Se generó e integró un diagrama de flujo (infográfico) visual para representar la arquitectura del Bunker OS.
4. **Seguridad del Repositorio**: Se creó el script `bin/bunker-github-setup.sh` para aplicar protección de ramas (PRs obligatorios) y configurar el "About" usando `gh cli` vía Infraestructura como Código.
5. **Auditoría de Workflows**: Se auditaron los workflows de n8n, identificando el `bunker-health-check.json` (Read-only) y el `bunker-aoc-v3-enterprise.json` (Pipeline completo).
6. **Diseño de Arquitectura V4**: Se propuso el diseño conceptual para la V4 del AOC, incluyendo deduplicación con Redis, Triage con IA (OpenRouter/Mistral) y escalado con Twilio (WhatsApp).

## 🚀 Próximos Pasos (Next Session)
- **Implementación Física del V4**: Pegar los nodos del *Upgrade Pack V4* directamente en el canvas de n8n para migrar el flujo V3.
- **Configuración de Credenciales**: Cargar las API Keys de OpenRouter, Redis y Twilio dentro del vault de n8n para que el AOC V4 pueda operar.
- **Pruebas E2E**: Ejecutar un evento de vulnerabilidad simulado para ver si la IA atrapa el error y dispara el workflow de remediación automáticamente.

## 🧠 Contexto para el Próximo Agente
El orquestador n8n está operativo en `localhost:5678`. Para cualquier prueba futura, arrancar verificando el flujo `bunker-health-check.json`. El V3 Enterprise es demasiado complejo para modificar vía JSON crudo, **SIEMPRE** pedirle al humano que copie y pegue los nuevos nodos en la interfaz gráfica.
