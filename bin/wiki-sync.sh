#!/bin/bash
# opencode-obsidian: Auto-sync script for the wiki vault

PROJECT_DIR="/home/sam/opencode-obsidian"
cd "$PROJECT_DIR"

# Check if there are changes
if [[ -n $(git status -s) ]]; then
    echo "Changes detected. Syncing..."
    git add .
    git commit -m "chore(wiki): compounding knowledge update $(date +'%Y-%m-%d %H:%M')"
    # Uncomment the next line if you want auto-push
    # git push origin main
    echo "Sync complete."
else
    echo "No changes to sync."
fi
