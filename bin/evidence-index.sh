#!/bin/bash
set -euo pipefail

BUNKER_HOME="${BUNKER_HOME:-$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)}"
cd "$BUNKER_HOME"

python3 - <<'PY'
import hashlib
import json
import zipfile
from pathlib import Path
from datetime import datetime, timezone

root = Path('.')
artifacts = []

def sha256_file(path: Path) -> str:
    h = hashlib.sha256()
    with path.open('rb') as f:
        for chunk in iter(lambda: f.read(1024 * 1024), b''):
            h.update(chunk)
    return h.hexdigest()

for path in ['report.zip', 'security-audit-report.json']:
    p = root / path
    if not p.exists():
        continue
    item = {
        'path': path,
        'type': 'zip-evidence' if p.suffix == '.zip' else 'security-audit-json',
        'size_bytes': p.stat().st_size,
        'sha256': sha256_file(p),
        'indexed_at': datetime.now(timezone.utc).isoformat(),
    }
    if p.suffix == '.zip':
        entries = []
        with zipfile.ZipFile(p) as z:
            for info in sorted(z.infolist(), key=lambda x: x.filename):
                if info.is_dir():
                    continue
                data = z.read(info.filename)
                entries.append({
                    'path': info.filename,
                    'size_bytes': info.file_size,
                    'sha256': hashlib.sha256(data).hexdigest(),
                })
        item['entries'] = entries
    else:
        try:
            data = json.loads(p.read_text())
            item['summary'] = {
                'top_level_keys': sorted(data.keys()) if isinstance(data, dict) else [],
                'note': 'Stored as original audit output; review before using as current truth.'
            }
        except Exception as exc:
            item['summary'] = {'error': str(exc)}
    artifacts.append(item)

manifest = {'generated_at': datetime.now(timezone.utc).isoformat(), 'artifacts': artifacts}
Path('wiki/meta/evidence-manifest.json').write_text(json.dumps(manifest, indent=2, ensure_ascii=False) + '\n')

lines = [
    '---',
    'type: evidence-index',
    'title: "Evidence Index"',
    f'updated: {datetime.now(timezone.utc).date().isoformat()}',
    'status: active',
    'tags: [evidence, audit, governance, bunker-os]',
    'related: [[dashboard]], [[INTEGRITY-REPORT]]',
    '---',
    '',
    '# Evidence Index',
    '',
    'This index preserves evidence as original artifacts and links them to wiki interpretation. Do not edit original evidence files in place.',
    ''
]
for art in artifacts:
    lines += [
        f"## `{art['path']}`",
        '',
        f"- Type: `{art['type']}`",
        f"- Size: `{art['size_bytes']}` bytes",
        f"- SHA256: `{art['sha256']}`",
        '- Manifest: `wiki/meta/evidence-manifest.json`',
    ]
    if 'entries' in art:
        lines += ['', '### Internal entries']
        for entry in art['entries']:
            lines.append(f"- `{entry['path']}` — {entry['size_bytes']} bytes — sha256 `{entry['sha256']}`")
    if 'summary' in art:
        lines += ['', '### Summary']
        for k, v in art['summary'].items():
            lines.append(f"- {k}: `{v}`")
    lines.append('')
lines += [
    '## Evidence Handling Policy',
    '',
    '- Original artifacts remain immutable unless a new version is explicitly created.',
    '- Notes may interpret evidence, but must link back to this index or the manifest.',
    '- Checksums are used to detect accidental changes.',
    '- If an artifact comes from a past repository state, mark it as historical instead of treating it as current truth.',
]
Path('wiki/meta/evidence-index.md').write_text('\n'.join(lines) + '\n')
print('Evidence index updated: wiki/meta/evidence-index.md')
print('Evidence manifest updated: wiki/meta/evidence-manifest.json')
PY
