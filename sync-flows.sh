#!/bin/bash
echo "📤 Exportando flujos consolidados desde n8n..."
docker exec -u node n8n_meco_lab n8n export:workflow --all --output=/data/workflows/all-workflows.json
echo "✅ Flujos guardados en workflows/all-workflows.json"
