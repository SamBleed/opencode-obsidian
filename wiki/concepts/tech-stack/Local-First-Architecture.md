---
type: concept
title: "Local-First Architecture"
domain: tech-stack
created: 2026-05-13
updated: 2026-05-13
tags: [architecture, offline-first, sync, sqlite, wasm]
status: mature
---

# Local-First Architecture (v2026)

Local-First is an architectural pattern where the primary data source lives on the client device, not the server. It treats synchronization as a background process, similar to how Git manages distributed versions.

## 🏛️ The Mental Model: "Git for Data"
In a Local-First app, every user has a full (or partial) replica of the database.
- **Local Commits**: Mutations are applied instantly to the local DB (Zero Latency).
- **Rebase/Sync**: The sync engine pulls remote changes and "replays" local mutations on top of the new state.
- **Conflict Resolution**: Deterministic merging of concurrent changes.

## 🛠️ The Tech Stack (2026)

### 1. SQLite WASM + OPFS
The foundation of modern local-first. We use **wa-sqlite** or **PGlite** running in the browser's **Origin Private File System (OPFS)** for near-native disk performance.

### 2. Sync Engines
- **Replicache**: Uses a Git-like rebase model. Best for high-interactivity web apps.
- **PowerSync**: Syncs PostgreSQL to local SQLite. Best for production-grade mobile and web apps with complex filtering.
- **ElectricSQL**: Real-time sync for Postgres using a familiar SQL interface everywhere.

## 🔄 Conflict Resolution Strategies

| Strategy | Usage | Mechanism |
|----------|-------|-----------|
| **Last-Write-Wins (LWW)** | Simple CRUD | Most recent timestamp wins. |
| **Field-Level Merge** | Forms/Profiles | Merge changes to different fields of the same row. |
| **CRDTs (Conflict-free)** | Collaborative Docs | Mathematical guarantee of eventual consistency (e.g., Automerge, Yjs). |
| **Rebase/Replay** | Counters/Logic | Re-run the logic (e.g., `count + 1`) on the new server state. |

## 🚀 implementation Pattern (Replicache Example)

```typescript
const mutators = {
  async updateProfile(tx: WriteTransaction, { id, name }: { id: string; name: string }) {
    const profile = await tx.get(`profile/${id}`);
    await tx.set(`profile/${id}`, { ...profile, name });
  },
};

const rep = new Replicache({
  name: 'user-123',
  mutators,
  pushURL: '/api/sync/push',
  pullURL: '/api/sync/pull',
});
```

## ⚖️ Why Local-First?
- **User Experience**: 0ms interaction latency. UI never "loads."
- **Robustness**: Works in tunnels, airplanes, or when the server is down.
- **Privacy**: Personal data stays on the device until needed.
- **Cost**: Offloads query processing from the server to the client's CPU.

## Related
- [[React]]
- [[performance-optimization]]
- [[PostgreSQL]]
- [[Agentic-Workflows]] (Running agents locally)
