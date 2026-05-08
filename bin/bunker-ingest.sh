#!/bin/bash
set -euo pipefail

# bunker-ingest.sh — Procesa una fuente y la integra al Bunker
# Uso: ./bin/bunker-ingest.sh <ruta_al_archivo_o_url>

SOURCE=$1

if [ -z "$SOURCE" ]; then
    echo "❌ Error: Debes proporcionar una fuente (archivo o URL)."
    echo "Uso: $0 <archivo_o_url>"
    exit 1
fi

echo "🧠 Iniciando ingestión de: $SOURCE"

# 1. Mover a .raw si es un archivo local
if [ -f "$SOURCE" ]; then
    FILENAME=$(basename "$SOURCE")
    cp "$SOURCE" .raw/
    echo "📥 Archivo copiado a .raw/$FILENAME"
fi

# 2. Notificar al Agente (Triggering logic)
# Aquí invocamos el comando interno del agente para procesar
echo "🤖 Solicitando análisis agéntico..."

./bin/bunker-alert.sh "Bunker-Ingest" "INFO" "📥 NUEVA FUENTE DETECTADA: Se ha iniciado la ingestión de $SOURCE. Procesando entidades y conceptos..."

echo "✅ Proceso de ingestión iniciado. Esperando síntesis..."
