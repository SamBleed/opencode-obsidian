#!/bin/bash
set -euo pipefail
# Bunker OS orchestrator script.

COMMAND="${1:-}"

BUNKER_HOME="${BUNKER_HOME:-$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)}"

case "$COMMAND" in
  init)
    echo "Initializing bunker workflow..."
    echo "STATUS: READ_HOT_AND_HANDOVER"
    ;;
  save)
    echo "Saving checkpoint..."
    "$BUNKER_HOME/bin/wiki-sync.sh"
    ;;
  refresh)
    echo "Refreshing context and re-reading core pillars..."
    echo "STATUS: RE-READING_HOT_INDEX_AND_ADRS"
    ;;
  close)
    echo "Closing session and generating handover..."
    echo "STATUS: GENERATE_HANDOVER_AND_PUSH"
    "$BUNKER_HOME/bin/wiki-sync.sh"
    ;;
  *)
    echo "Usage: bunker [init|save|refresh|close]"
    ;;
esac
