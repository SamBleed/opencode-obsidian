#!/usr/bin/env python3
"""retrieve.py — Hybrid retrieval orchestrator para Bunker OS.

Pipeline:
  query → BM25 (sparse candidates) → ollama rerank (dense cosine) → ranked results

Uso:
  python3 scripts/retrieve.py "texto de búsqueda" --top 5
  python3 scripts/retrieve.py "texto" --top 10 --no-rerank

Requiere: numpy, ollama con nomic-embed-text
"""

import json
import os
import subprocess
import sys

SCRIPTS = os.path.dirname(os.path.abspath(__file__))
VAULT = os.path.dirname(SCRIPTS)
sys.path.insert(0, SCRIPTS)
import importlib
bm25_mod = importlib.import_module('bm25-index')
query_index = bm25_mod.query_index
build_index = bm25_mod.build_index


def embed(texts):
    """Obtiene embeddings via ollama API."""
    if isinstance(texts, str):
        texts = [texts]

    try:
        import urllib.request
        payload = json.dumps({"model": "nomic-embed-text", "input": texts}).encode()
        req = urllib.request.Request(
            "http://localhost:11434/api/embed",
            data=payload,
            headers={"Content-Type": "application/json"}
        )
        with urllib.request.urlopen(req, timeout=30) as resp:
            data = json.loads(resp.read())
            return data.get("embeddings", [])
    except Exception as e:
        print(f'⚠️  Error en embed: {e}', file=sys.stderr)
        return None


def cosine_similarity(a, b):
    """Producto punto normalizado."""
    import numpy as np
    a = np.array(a)
    b = np.array(b)
    return float(np.dot(a, b) / (np.linalg.norm(a) * np.linalg.norm(b) + 1e-10))


def rerank(query, candidates, top_k=5):
    """Re-ordena candidatos por cosine similarity vía ollama."""
    if not candidates:
        return candidates

    texts = [c['preview'] for c in candidates]
    query_emb = embed(query)
    doc_embs = embed(texts)

    if query_emb is None or doc_embs is None:
        print('⚠️  Rerank falló (ollama no disponible), usando scores BM25', file=sys.stderr)
        return candidates[:top_k]

    query_vec = query_emb[0] if isinstance(query_emb[0], list) else query_emb
    doc_vecs = doc_embs

    for i, c in enumerate(candidates):
        if i < len(doc_vecs):
            c['similarity'] = round(cosine_similarity(query_vec, doc_vecs[i]), 4)
            c['combined'] = round(0.3 * c['score'] + 0.7 * c['similarity'], 4)
        else:
            c['similarity'] = 0
            c['combined'] = c['score']

    reranked = sorted(candidates, key=lambda x: -x.get('combined', x['score']))
    return reranked[:top_k]


def main():
    args = sys.argv[1:]
    if not args or args[0] in ('-h', '--help'):
        print('Uso: python3 scripts/retrieve.py "query" [--top N] [--no-rerank]')
        print('      python3 scripts/retrieve.py build  (construir/actualizar índice)')
        sys.exit(0)

    if args[0] == 'build':
        build_index()
        return
    elif args[0] == 'status':
        bm25_mod.main()
        return

    query = args[0]
    top_k = 5
    use_rerank = True

    for i, a in enumerate(args[1:], 1):
        if a == '--top' and i + 1 < len(args):
            top_k = int(args[i + 1])
        elif a == '--no-rerank':
            use_rerank = False

    # Paso 1: BM25 candidates
    candidates = query_index(query, top_k=top_k * 2)
    if not candidates:
        print(f'⚠️  Sin resultados para: {query}')
        return

    # Paso 2: Rerank (opcional)
    if use_rerank:
        results = rerank(query, candidates, top_k=top_k)
        strategy = 'bm25+rerank:cosine:nomic-embed-text'
    else:
        results = candidates[:top_k]
        strategy = 'bm25+no-rerank'

    # Output
    output = {
        'query': query,
        'strategy': strategy,
        'top_k': len(results),
        'candidates': results
    }
    print(json.dumps(output, ensure_ascii=False, indent=2))


if __name__ == '__main__':
    main()
