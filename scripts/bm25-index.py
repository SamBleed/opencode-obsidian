#!/usr/bin/env python3
"""bm25-index.py — BM25 indexer for Bunker OS wiki.

Build:  python3 scripts/bm25-index.py build
Query:  python3 scripts/bm25-index.py query "tu busqueda"
Status: python3 scripts/bm25-index.py status

Index se guarda en .vault-meta/bm25/index.json
Chunks en .vault-meta/chunks/
"""

import json
import math
import os
import re
import sys
from collections import Counter
from glob import glob
from pathlib import Path

VAULT = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
INDEX_DIR = os.path.join(VAULT, '.vault-meta', 'bm25')
INDEX_FILE = os.path.join(INDEX_DIR, 'index.json')
CHUNKS_DIR = os.path.join(VAULT, '.vault-meta', 'chunks')


def tokenize(text):
    """Tokeniza texto a términos lowercase."""
    text = re.sub(r'[^a-záéíóúñüA-ZÁÉÍÓÚÑÜ0-9\s-]', ' ', text.lower())
    return [t for t in text.split() if len(t) > 1]


def chunk_page(filepath):
    """Divide una página MD en chunks (~500 tokens cada uno)."""
    with open(filepath, 'r', encoding='utf-8') as f:
        text = f.read()

    # Sacar frontmatter
    text = re.sub(r'^---.*?---\s*', '', text, flags=re.DOTALL)

    # Dividir por párrafos
    paragraphs = [p.strip() for p in text.split('\n\n') if p.strip()]
    chunks = []
    current = []
    current_len = 0

    for p in paragraphs:
        tokens = len(tokenize(p))
        if current_len + tokens > 500 and current:
            chunks.append('\n\n'.join(current))
            current = []
            current_len = 0
        current.append(p)
        current_len += tokens

    if current:
        chunks.append('\n\n'.join(current))

    return chunks if chunks else [text[:2000]]


def build_index():
    """Construye el índice BM25 de todas las páginas wiki."""
    os.makedirs(INDEX_DIR, exist_ok=True)
    os.makedirs(CHUNKS_DIR, exist_ok=True)

    files = glob(os.path.join(VAULT, 'wiki', '**', '*.md'), recursive=True)
    # Excluir meta
    files = [f for f in files if '/meta/' not in f and f.endswith('.md')]

    print(f'Indexando {len(files)} páginas...')

    # Cargar versión anterior si existe
    old_index = {}
    if os.path.exists(INDEX_FILE):
        try:
            old_index = json.load(open(INDEX_FILE))
        except Exception:
            pass

    doc_count = 0
    term_doc_freq = Counter()
    doc_terms = {}
    chunk_meta = []

    for filepath in files:
        rel = os.path.relpath(filepath, VAULT)
        chunks = chunk_page(filepath)

        for ci, chunk in enumerate(chunks):
            terms = tokenize(chunk)
            if not terms:
                continue

            doc_id = f'doc_{doc_count}'
            doc_terms[doc_id] = Counter(terms)
            chunk_meta.append({
                'doc_id': doc_id,
                'path': rel,
                'chunk_index': ci,
                'total_chunks': len(chunks),
                'preview': chunk[:200].replace('\n', ' ')
            })

            for t in set(terms):
                term_doc_freq[t] += 1

            doc_count += 1

    N = doc_count
    avgdl = sum(sum(tf.values()) for tf in doc_terms.values()) / max(N, 1)

    # BM25 params
    k1 = 1.5
    b = 0.75

    index = {
        'N': N,
        'avgdl': avgdl,
        'k1': k1,
        'b': b,
        'term_doc_freq': dict(term_doc_freq),
        'doc_terms': {k: dict(v) for k, v in doc_terms.items()},
        'chunks': chunk_meta,
        'stats': {
            'total_pages': len(files),
            'total_chunks': doc_count,
            'vocab_size': len(term_doc_freq),
        }
    }

    json.dump(index, open(INDEX_FILE, 'w'), ensure_ascii=False)
    print(f'✅ Índice guardado: {doc_count} chunks, {len(term_doc_freq)} términos')
    print(f'   {len(files)} páginas indexadas')


def query_index(query_str, top_k=10):
    """Busca en el índice BM25."""
    if not os.path.exists(INDEX_FILE):
        print('❌ Índice no encontrado. Ejecutá: python3 scripts/bm25-index.py build')
        return []

    index = json.load(open(INDEX_FILE))
    N = index['N']
    avgdl = index['avgdl']
    k1 = index['k1']
    b = index['b']
    term_doc_freq = index['term_doc_freq']
    doc_terms = index['doc_terms']
    chunks = index['chunks']

    query_terms = tokenize(query_str)
    if not query_terms:
        return []

    scores = {}
    for doc_id, tf in doc_terms.items():
        score = 0
        dl = sum(tf.values())
        for qt in query_terms:
            if qt in tf and qt in term_doc_freq:
                idf = math.log((N - term_doc_freq[qt] + 0.5) / (term_doc_freq[qt] + 0.5) + 1)
                tf_term = tf[qt]
                score += idf * (tf_term * (k1 + 1)) / (tf_term + k1 * (1 - b + b * dl / avgdl))
        if score > 0:
            scores[doc_id] = score

    ranked = sorted(scores.items(), key=lambda x: -x[1])

    results = []
    for doc_id, score in ranked[:top_k]:
        meta = next((c for c in chunks if c['doc_id'] == doc_id), {})
        results.append({
            'doc_id': doc_id,
            'score': round(score, 4),
            'path': meta.get('path', ''),
            'chunk': meta.get('chunk_index', 0),
            'preview': meta.get('preview', ''),
        })

    return results


def main():
    cmd = sys.argv[1] if len(sys.argv) > 1 else 'status'

    if cmd == 'build':
        build_index()
    elif cmd == 'query':
        q = ' '.join(sys.argv[2:])
        if not q:
            print('Uso: python3 scripts/bm25-index.py query "texto"')
            sys.exit(1)
        results = query_index(q)
        print(f'Resultados para: {q}\n')
        for r in results[:5]:
            print(f'  [{r["score"]:6.3f}] {r["path"]}')
            print(f'         {r["preview"][:100]}')
            print()
    elif cmd == 'status':
        if os.path.exists(INDEX_FILE):
            idx = json.load(open(INDEX_FILE))
            s = idx.get('stats', {})
            print(f'📊 BM25 Index Status')
            print(f'   Páginas: {s.get("total_pages", "?")}')
            print(f'   Chunks:  {s.get("total_chunks", "?")}')
            print(f'   Términos: {s.get("vocab_size", "?")}')
        else:
            print('📊 No hay índice. Ejecutá: python3 scripts/bm25-index.py build')
    else:
        print('Uso: python3 scripts/bm25-index.py build|query "texto"|status')


if __name__ == '__main__':
    main()
