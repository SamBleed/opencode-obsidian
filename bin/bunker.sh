#!/bin/bash
# opencode-obsidian: Orchestrator script

COMMAND=$1
PROJECT_NAME=$2

case $COMMAND in
  init)
    echo "Initializing bunker workflow..."
    # Logic to be handled by the agent after reading this output
    echo "STATUS: READ_HOT_AND_HANDOVER"
    ;;
  save)
    echo "Saving checkpoint..."
    /home/sam/opencode-obsidian/bin/wiki-sync.sh
    ;;
  close)
    echo "Closing session and generating handover..."
    # Logic to be handled by the agent
    echo "STATUS: GENERATE_HANDOVER_AND_PUSH"
    /home/sam/opencode-obsidian/bin/wiki-sync.sh
    ;;
  *)
    echo "Usage: bunker [init|save|close]"
    ;;
esac
