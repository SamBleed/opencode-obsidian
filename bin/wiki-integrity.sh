#!/bin/bash
# wiki-integrity.sh — Scan for orphans and broken links in the Bunker
# Usage: ./bin/wiki-integrity.sh

VAULT_DIR="wiki"
REPORT_FILE="wiki/meta/INTEGRITY-REPORT.md"

# Use absolute path for VAULT_DIR to avoid confusion
ABS_VAULT_DIR=$(realpath "$VAULT_DIR")

echo "# 🔍 Wiki Integrity Report" > "$REPORT_FILE"
echo "Generated: $(date)" >> "$REPORT_FILE"
echo "" >> "$REPORT_FILE"

echo "## 🕸️ Orphan Notes (No incoming links)" >> "$REPORT_FILE"
# Find all md files except meta/ and the index/log/hot files
find "$ABS_VAULT_DIR" -name "*.md" -not -path "*/meta/*" -not -name "index.md" -not -name "log.md" -not -name "hot.md" | while read -r file; do
    filename=$(basename "$file" .md)
    # Search for [[filename]] or [[filename|alias]]
    # We use -F for fixed string to avoid regex issues with special chars in filenames
    # We grep in the whole vault directory
    count=$(grep -rF "[[${filename}]]" "$ABS_VAULT_DIR" | grep -v "$file" | wc -l)
    alias_count=$(grep -rF "[[${filename}|" "$ABS_VAULT_DIR" | grep -v "$file" | wc -l)
    
    total_links=$((count + alias_count))
    
    if [ $total_links -eq 0 ]; then
        # Double check with a more loose grep just in case
        loose_count=$(grep -r "\[\[${filename}[|\]]" "$ABS_VAULT_DIR" | grep -v "$file" | wc -l)
        if [ $loose_count -eq 0 ]; then
            echo "- [[$filename]]" >> "$REPORT_FILE"
        fi
    fi
done

echo "" >> "$REPORT_FILE"
echo "## 🛠️ Short Notes (Seed candidates < 30 lines)" >> "$REPORT_FILE"
find "$ABS_VAULT_DIR" -name "*.md" -not -path "*/archive/*" -not -path "*/inbox/*" -exec wc -l {} + | sort -n | while read -r line; do
    count=$(echo "$line" | awk '{print $1}')
    filepath=$(echo "$line" | awk '{print $2}')
    if [[ "$filepath" == "total" ]]; then continue; fi
    
    if [ "$count" -lt 30 ]; then
        filename=$(basename "$filepath" .md)
        echo "- [[$filename]] ($count lines)" >> "$REPORT_FILE"
    fi
done

echo "" >> "$REPORT_FILE"
echo "## 🚩 Broken Links (WikiLinks to non-existent files)" >> "$REPORT_FILE"
# Find all wikilinks, extract the target, and check if file exists
grep -rEoh "\[\[[^\]|]+\|?[^\]]*\]\]" "$ABS_VAULT_DIR" | sed 's/\[\[//;s/\]\]//;s/|.*//' | sort | uniq | while read -r link; do
    # Skip empty or weird links
    if [[ -z "$link" || "$link" == " " ]]; then continue; fi
    # Check for file existence in any subfolder
    found=$(find "$ABS_VAULT_DIR" -name "${link}.md" | wc -l)
    if [ "$found" -eq 0 ]; then
        if [[ "$link" == "Index" || "$link" == "log" || "$link" == "hot" ]]; then continue; fi
        echo "- [[$link]] (Target file not found)" >> "$REPORT_FILE"
    fi
done

echo "Integrity report updated in $REPORT_FILE"
