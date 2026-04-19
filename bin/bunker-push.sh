#!/bin/bash

# bunker-push.sh — Sincroniza el Bunker y notifica vía n8n
# Uso: ./bin/bunker-push.sh "mensaje del commit"

MESSAGE=$1

if [ -z "$MESSAGE" ]; then
    echo "❌ Error: Debes proporcionar un mensaje para el commit."
    echo "Uso: $0 \"mensaje del commit\""
    exit 1
fi

echo "📦 Iniciando sincronización del Bunker..."

# 1. Git Add y Commit
git add -A
git commit -m "$MESSAGE"

# 2. Git Push
if git push; then
    echo "✅ Sincronización exitosa en GitHub."
    
    # 3. Notificar a n8n
    LAST_COMMIT_HASH=$(git rev-parse --short HEAD)
    ./bin/bunker-alert.sh "Bunker-Git" "SUCCESS" "📦 PUSH EXITOSO: El Bunker ha sido sincronizado.\n\nHash: $LAST_COMMIT_HASH\nMessage: $MESSAGE\n\n_System status: Protected_"
else
    echo "❌ Error al sincronizar con el remote."
    ./bin/bunker-alert.sh "Bunker-Git" "ERROR" "⚠️ FALLO DE SINCRONIZACIÓN: Error al intentar hacer push al repositorio."
    exit 1
fi
