#!/bin/bash

# bunker-alert.sh — Dispara una alerta de seguridad al orquestador n8n
# Uso: ./bunker-alert.sh <target> <level> <message>

TARGET=$1
LEVEL=$2
MESSAGE=$3

# Configuración (Ajustar según MCP-CONFIG.md)
N8N_WEBHOOK_URL="http://localhost:5678/webhook/bunker-security-alert"

if [ -z "$TARGET" ] || [ -z "$LEVEL" ] || [ -z "$MESSAGE" ]; then
    echo "Uso: $0 <target> <level> <message>"
    exit 1
fi

echo "📡 Enviando alerta de seguridad a n8n..."

curl -X POST "$N8N_WEBHOOK_URL" \
-H "Content-Type: application/json" \
-d "{
  \"target\": \"$TARGET\",
  \"level\": \"$LEVEL\",
  \"message\": \"$MESSAGE\"
}"

echo -e "\n✅ Alerta enviada exitosamente."
