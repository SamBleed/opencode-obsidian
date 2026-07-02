---
name: wiki-retrieve
description: >
  Hybrid retrieval para Bunker OS. BM25 + rerank semántico vía ollama
  (nomic-embed-text). Reemplaza la búsqueda plana con retrieval híbrido
  contextual. Triggers on: "/retrieve", "retrieve", "hybrid search",
  "búsqueda semántica", "find relevant", "buscar en el vault",
  "search chunks", "chunk search", "vault search".
---

# wiki-retrieve: Hybrid Retrieval para Bunker OS

Busca en el vault usando BM25 (búsqueda textual) + rerank semántico
(cosine similarity vía nomic-embed-text en ollama). Mejor que la
búsqueda plana porque entiende significado, no solo palabras exactas.

---

## Feature Detection

Antes de usar retrieve, verificá que el índice existe:

```bash
python3 scripts/retrieve.py build   # construir/actualizar índice
python3 scripts/retrieve.py status  # ver estado
```

Si el índice no existe, caé al flujo legacy (obsidian-vault_search_notes).

---

## Pipeline

```
query string
     │
     ▼
scripts/retrieve.py "<query>" --top 5
     │
     ├─ bm25-index.py query "<query>"    (candidatos sparse)
     ├─ rerank vía ollama nomic-embed-text (cosine similarity)
     │   embeddings cacheadas en .vault-meta/embed-cache.json
     └─ devuelve top-N con: path, score, similarity, preview
     │
     ▼
el skill lee las páginas y sintetiza respuesta
```

---

## Modos de Búsqueda

| Modo | Comando | Uso |
|------|---------|-----|
| **Quick** | `python3 scripts/retrieve.py "query" --top 3` | Búsqueda rápida, solo resultados |
| **Rerank** | `python3 scripts/retrieve.py "query" --top 5` | BM25 + rerank semántico (default) |
| **No rerank** | `python3 scripts/retrieve.py "query" --no-rerank` | Solo BM25 (más rápido) |
| **Build** | `python3 scripts/retrieve.py build` | Reconstruir índice |

---

## Integración con Búsqueda

Cuando recibas una consulta:

1. **Primero**: intentá `python3 scripts/retrieve.py status` para ver si el índice existe
2. **Si existe**: ejecutá `python3 scripts/retrieve.py "consulta" --top 5`
3. **Leé los resultados**: cada resultado tiene `path` (ruta del archivo) y `preview`
4. **Leé las páginas** más relevantes usando `obsidian-vault_read_note`
5. **Sintetizá** la respuesta con citas a las fuentes

Si el índice NO existe o el script falla:
- Usá `obsidian-vault_search_notes` con la consulta
- O leé `wiki/hot.md` + `wiki/index.md` + páginas relevantes (legacy)

---

## Mantenimiento del Índice

El índice hay que rebuildearlo cuando se agreguen/modifiquen páginas wiki:

```bash
# Después de ingestar fuentes nuevas:
python3 scripts/retrieve.py build

# Verificar estado:
python3 scripts/retrieve.py status
```

Idealmente correr `build` antes de cada sesión de búsqueda pesada.

---

## Tools Disponibles

- `python3 scripts/retrieve.py build` — construir índice BM25
- `python3 scripts/retrieve.py status` — estado del índice
- `python3 scripts/retrieve.py "query" --top N` — búsqueda híbrida
- `obsidian-vault_search_notes` — fallback FTS nativo
- `obsidian-vault_read_note` — leer páginas encontradas
