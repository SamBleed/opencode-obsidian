#!/bin/bash
set -euo pipefail
# Bunker OS: safe wiki sync helper.
# Default mode is dry-run. Use --apply to create the commit.

APPLY=0
MESSAGE="chore(wiki): compounding knowledge update $(date +'%Y-%m-%d %H:%M')"

while [[ $# -gt 0 ]]; do
  case "$1" in
    --apply)
      APPLY=1
      shift
      ;;
    -m|--message)
      MESSAGE="${2:-}"
      shift 2
      ;;
    -h|--help)
      echo "Usage: $0 [--apply] [-m|--message <commit message>]"
      exit 0
      ;;
    *)
      echo "Unknown argument: $1" >&2
      exit 2
      ;;
  esac
done

BUNKER_HOME="${BUNKER_HOME:-$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)}"
cd "$BUNKER_HOME"

if ! git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
  if [[ "$APPLY" -eq 1 ]]; then
    echo "❌ Not a git repository: $BUNKER_HOME" >&2
    echo "Initialize git first or run from a cloned repository." >&2
    exit 1
  fi
  echo "DRY RUN: $BUNKER_HOME is not a git repository yet."
  echo "DRY RUN: would inspect changes after git init/clone."
  exit 0
fi

if [[ -z "$(git status -s)" ]]; then
  echo "No changes to sync."
  exit 0
fi

echo "Changes detected in $BUNKER_HOME:"
git status --short

if [[ "$APPLY" -ne 1 ]]; then
  echo "DRY RUN: would run: git add wiki/ _templates/ bin/ docs/ *.md"
  echo "DRY RUN: would commit with message: $MESSAGE"
  echo "Run with --apply to commit."
  exit 0
fi

git add wiki/ _templates/ bin/ docs/ *.md .gitignore .obsidian/community-plugins.json .obsidian/plugins/obsidian-local-rest-api/SECURITY-NOTE.md .obsidian/plugins/obsidian-local-rest-api/data.example.json 2>/dev/null || true
git commit -m "$MESSAGE"
echo "Sync commit created. Push manually after review."
