#!/bin/bash
set -euo pipefail

# Bunker Pulse v4.0: El Motor de Autonomía
# Escanea el inbox, crea issues, registra memorias y AUTO-DOCUMENTA el proyecto.

INBOX_DIR="/home/sam/opencode-obsidian/wiki/inbox"
ARCHIVE_DIR="/home/sam/opencode-obsidian/wiki/archive"
RULES_FILE="/home/sam/opencode-obsidian/BUNKER_RULES.md"

mkdir -p "$ARCHIVE_DIR"

echo "💓 Bunker Pulse v4.0 en marcha..."

# --- Función para reconstruir el manual de reglas ---
rebuild_rules() {
    echo "📚 Reconstruyendo BUNKER_RULES.md desde AgentMemory..."
    
    cat <<EOF > "$RULES_FILE"
# 🛡️ Manual de Vuelo del Bunker
*Este archivo se genera automáticamente mediante el Bunker Pulse v4.0 y AgentMemory.*

## 📐 Estándares de Arquitectura y Decisiones
EOF

    # Consultamos a AgentMemory por todas las memorias de tipo architecture y decision
    # Usamos el endpoint de export y filtramos con jq
    curl -s http://localhost:3111/agentmemory/export | jq -r '
      .memories | 
      map(select(.type == "architecture" or .type == "fact" or .type == "decision")) | 
      sort_by(.createdAt) | 
      .[] | 
      "### " + .title + "\n> " + .content + "\n"
    ' >> "$RULES_FILE"

    echo "✅ BUNKER_RULES.md actualizado."
}

CHANGE_DETECTED=false

for file in "$INBOX_DIR"/*.md; do
    [ -e "$file" ] || continue
    
    filename=$(basename "$file")
    echo "Procesando $filename..."

    # 1. Procesar Bugs (GitHub)
    if grep -q "type: bug\|tags: \[bug\]" "$file"; then
        echo "🐞 Bug detectado! (Simulando Issue...)"
        mv "$file" "$ARCHIVE_DIR/$filename"
        CHANGE_DETECTED=true

    # 2. Procesar Decisiones (Memoria + Documentación)
    elif grep -q "type: decision\|tags: \[decision\]" "$file"; then
        title=$(grep -v "^---" "$file" | grep -v "^type:\|^created:\|^source:\|^tags:" | head -n 1 | sed 's/^[ \t]*//')
        content=$(cat "$file")
        
        echo "📐 Nueva decisión detectada! Registrando..."
        curl -s -X POST http://localhost:3111/agentmemory/remember \
        -H "Content-Type: application/json" \
        -d "{
          \"title\": \"$title\",
          \"content\": \"$content\",
          \"type\": \"decision\"
        }" > /dev/null
        
        mv "$file" "$ARCHIVE_DIR/$filename"
        CHANGE_DETECTED=true
    fi
done

# Si hubo cambios, reconstruimos las reglas
if [ "$CHANGE_DETECTED" = true ]; then
    rebuild_rules
fi

echo "🏁 Ciclo completado."
