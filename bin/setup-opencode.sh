#!/bin/bash
set -euo pipefail
#
# opencode-obsidian setup
# Instala el vault + skills + MCP para OpenCode
#

set -e

VAULT_DIR="${VAULT_DIR:-$HOME/opencode-obsidian}"
SKILLS_DIR="$HOME/.config/opencode/skills"
OPENCODE_CONFIG="$HOME/.config/opencode/opencode.json"

echo "========================================"
echo "  opencode-obsidian Setup"
echo "========================================"

# 1. Clonar o actualizar el vault
echo "[1/5] Preparando vault..."
if [ -d "$VAULT_DIR" ]; then
    echo "  → Vault ya existe en $VAULT_DIR"
    cd "$VAULT_DIR"
    git pull origin main 2>/dev/null || echo "  → No hay updates remotos"
else
    git clone https://github.com/SamBleed/opencode-obsidian.git "$VAULT_DIR"
    cd "$VAULT_DIR"
fi

# 2. Copiar skills
echo "[2/5] Instalando skills..."
mkdir -p "$SKILLS_DIR"
rm -rf "$SKILLS_DIR/wiki" "$SKILLS_DIR/wiki-ingest" "$SKILLS_DIR/wiki-query" "$SKILLS_DIR/wiki-save" "$SKILLS_DIR/wiki-lint" 2>/dev/null || true

cp -r skills/wiki "$SKILLS_DIR/" 2>/dev/null || echo "  ⚠ skills/no found, skipping"
cp -r skills/wiki-ingest "$SKILLS_DIR/" 2>/dev/null || true
cp -r skills/wiki-query "$SKILLS_DIR/" 2>/dev/null || true
cp -r skills/wiki-save "$SKILLS_DIR/" 2>/dev/null || true
cp -r skills/wiki-lint "$SKILLS_DIR/" 2>/dev/null || true

# 3. Configurar MCP en opencode.json
echo "[3/5] Configurando MCP..."
if [ -f "$OPENCODE_CONFIG" ]; then
    # Backup
    cp "$OPENCODE_CONFIG" "$OPENCODE_CONFIG.backup"
    
    # Agregar MCP si no existe
    if ! grep -q '"obsidian-vault"' "$OPENCODE_CONFIG"; then
        python3 << 'PYTHON'
import json
import sys

config_file = sys.argv[1]
with open(config_file, 'r') as f:
    config = json.load(f)

config['mcp'] = config.get('mcp', {})
config['mcp']['obsidian-vault'] = {
    "type": "stdio",
    "command": "npx",
    "args": ["-y", "@bitbonsai/mcpvault@latest", "/home/sam/opencode-obsidian"]
}

with open(config_file, 'w') as f:
    json.dump(config, f, indent=2)
PYTHON
        "$OPENCODE_CONFIG"
        echo "  → MCP configurado"
    else
        echo "  → MCP ya existe"
    fi
else
    echo "  ⚠ opencode.json no encontrado"
    echo "  → Configura el MCP manualmente:"
    echo '  obsidian-vault: npx @bitbonsai/mcpvault@latest /home/sam/opencode-obsidian'
fi

# 4. Abrir en Obsidian (opcional)
echo "[4/5] Preparando Obsidian..."
if command -v obsidian &> /dev/null; then
    echo "  → Obsidian instalado"
    echo "  → Para abrir: 'obsidian \"$VAULT_DIR\"'"
else
    echo "  ⚠ Obsidian no instalado"
fi

# 5. Instrucciones finales
echo "[5/5] Listo!"
echo ""
echo "========================================"
echo "  Próximos pasos:"
echo "========================================"
echo "1. Reiniciá OpenCode"
echo "2. abrí Obsidian → $VAULT_DIR"
echo "3. Usá:"
echo "   - ingest [archivo]"
echo "   - what do you know about X?"
echo "   - save this"
echo "   - lint"
echo ""
echo "========================================"

exit 0