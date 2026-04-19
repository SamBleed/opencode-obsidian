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

JSON_PAYLOAD=$(python3 -c 'import json, sys; print(json.dumps({"target": sys.argv[1], "level": sys.argv[2], "message": sys.argv[3]}))' "$TARGET" "$LEVEL" "$MESSAGE")
HTTP_RESPONSE=$(mktemp)
HTTP_CODE=$(curl -sS -o "$HTTP_RESPONSE" -w "%{http_code}" -X POST "$N8N_WEBHOOK_URL" \
  -H "Content-Type: application/json" \
  -d "$JSON_PAYLOAD")

if [ "$HTTP_CODE" -lt 200 ] || [ "$HTTP_CODE" -ge 300 ]; then
    echo "❌ Error al enviar la alerta. HTTP $HTTP_CODE"
    cat "$HTTP_RESPONSE"
    rm -f "$HTTP_RESPONSE"
    exit 1
fi

rm -f "$HTTP_RESPONSE"

echo -e "\n✅ Alerta enviada exitosamente."
