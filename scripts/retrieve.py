#!/usr/bin/env python3
"""retrieve.py — BM25 text retrieval for Bunker OS.

Pure BM25 index search. No embeddings, no LLM, no external services.
OpenCode (the agent) reads the ranked results and synthesizes answers.

Pipeline:
  query → BM25 (sparse candidates) → ranked results

Uso:
  python3 scripts/retrieve.py "consulta" --top 5
  python3 scripts/retrieve.py build
  python3 scripts/retrieve.py status

Deps: Python 3.10+ (stdlib only — no external packages required)
"""

import json
import os
import sys

SCRIPTS = os.path.dirname(os.path.abspath(__file__))
sys.path.insert(0, SCRIPTS)
import importlib
bm25_mod = importlib.import_module('bm25-index')

# Vault path: env var overrides default (repo root)
DEFAULT_VAULT = os.path.dirname(SCRIPTS)
VAULT_PATH = os.environ.get('BUNKER_HOME') or os.environ.get('VAULT_PATH') or DEFAULT_VAULT
if VAULT_PATH != DEFAULT_VAULT:
    bm25_mod.VAULT = VAULT_PATH

query_index = bm25_mod.query_index
build_index = bm25_mod.build_index


def main():
    args = sys.argv[1:]
    if not args or args[0] in ('-h', '--help'):
        print('Uso: python3 scripts/retrieve.py "query" [--top N]')
        print('      python3 scripts/retrieve.py build  (construir/actualizar índice)')
        print('      python3 scripts/retrieve.py status (estado del índice)')
        print('')
        print('Sin dependencias externas. BM25 puro. OpenCode sintetiza.')
        sys.exit(0)

    if args[0] == 'build':
        build_index()
        return
    elif args[0] == 'status':
        bm25_mod.main()
        return

    query = args[0]
    if not query.strip():
        print('Error: la consulta no puede estar vacía', file=sys.stderr)
        sys.exit(1)

    top_k = 5

    for i, a in enumerate(args[1:], 1):
        if a == '--top' and i + 1 < len(args):
            try:
                top_k = int(args[i + 1])
            except ValueError:
                print(f'Error: --top debe ser un número entero, no "{args[i + 1]}"', file=sys.stderr)
                sys.exit(1)

    candidates = query_index(query, top_k=top_k)
    if not candidates:
        print(json.dumps({'query': query, 'strategy': 'bm25', 'top_k': 0, 'candidates': []},
                         ensure_ascii=False, indent=2))
        return

    results = candidates[:top_k]
    output = {
        'query': query,
        'strategy': 'bm25',
        'top_k': len(results),
        'candidates': results
    }
    print(json.dumps(output, ensure_ascii=False, indent=2))


if __name__ == '__main__':
    main()
