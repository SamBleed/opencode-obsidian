#!/bin/bash
# wiki-integrity.sh — Scan for orphans and broken links in the Bunker
# Usage: ./bin/wiki-integrity.sh

VAULT_DIR="wiki"
REPORT_FILE="wiki/meta/INTEGRITY-REPORT.md"

echo "# 🔍 Wiki Integrity Report" > $REPORT_FILE
echo "Generated: $(date)" >> $REPORT_FILE
echo "" >> $REPORT_FILE

echo "## 🕸️ Orphan Notes (No incoming links)" >> $REPORT_FILE
# Find all files, then for each file, grep for its name in all other files
# Exclude index, log, hot, and meta files
for file in $(find $VAULT_DIR -name "*.md" -not -path "*/meta/*" -not -name "index.md" -not -name "log.md" -not -name "hot.md"); do
    filename=$(basename "$file" .md)
    count=$(grep -r "\[\[$filename\]\]" $VAULT_DIR | grep -v "$file" | wc -l)
    if [ $count -eq 0 ]; then
        echo "- [[$filename]]" >> $REPORT_FILE
    fi
done

echo "" >> $REPORT_FILE
echo "## 🛠️ Short Notes (Seed candidates < 30 lines)" >> $REPORT_FILE
find $VAULT_DIR -name "*.md" -not -path "*/archive/*" -not -path "*/inbox/*" -exec wc -l {} + | sort -n | while read line; do
    count=$(echo $line | awk '{print $1}')
    filepath=$(echo $line | awk '{print $2}')
    if [ $count -lt 30 ]; then
        filename=$(basename "$filepath" .md)
        echo "- [[$filename]] ($count lines)" >> $REPORT_FILE
    fi
done

echo "" >> $REPORT_FILE
echo "## 🚩 Broken Links (Pending implementation)" >> $REPORT_FILE
echo "Check manually for [[ ]] that don't point to a file." >> $REPORT_FILE

echo "Integrity report saved to $REPORT_FILE"
