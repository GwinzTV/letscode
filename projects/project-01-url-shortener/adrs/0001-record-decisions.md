# ADR 0001 — Persistence choice (starter)

Status: Accepted (updated)

Date: 2026-07-30

Context

- Beginner-friendly starter project; keep friction low.
- SQLite is recommended for production but requires extra dependencies and platform considerations.

Decision

Use a JSON-backed file store for the initial starter kit. This keeps the code simple and portable.

Update — Concurrency fix

During early testing a deadlock was discovered when the FileStore.persist method encoded the in-memory map while locks were held. The Save method acquires the write lock and called persist(), which attempted to acquire a read lock on the same sync.RWMutex — this causes a deadlock.

The implementation was updated to snapshot the map while holding the read lock, release the lock, and then write the snapshot to disk. This avoids holding locks during file I/O and prevents the deadlock.

Consequences

- Easy to understand and modify.
- Not suitable for high-concurrency or production workloads without replacement (e.g., SQLite/Postgres).
- The snapshot approach reduces risk of deadlock and keeps I/O off the critical path, but it does not provide transactional guarantees.
- Migration to SQLite or other DB should be straightforward via the Store interface.

Reference: internal/store/store.go — persist() now snapshots the map before encoding to disk.
