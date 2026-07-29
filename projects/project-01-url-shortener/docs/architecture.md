Architecture Overview (starter)

- HTTP server (net/http)
- Handlers layer: routing and HTTP concerns
- Store interface: persistence abstraction
- FileStore: JSON-backed file persistence (data/urls.json)

Next steps for participants:
- Replace FileStore with SQLite-backed implementation
- Add input validation, rate limiting, auth
- Add metrics, logging, and health checks
- Add CI, Docker-compose, and deployment scripts
