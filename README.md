# Stratus Mesh

Stratus Mesh is a lightweight distributed-cloud orchestration API and simulation platform written in Go using Gin and GORM. It demonstrates patterns for node discovery, scheduling, replication, and observability for distributed cloud systems.

Features
- REST API for managing nodes and schedules
- GORM-based persistence (Postgres support)
- Simple in-memory service discovery and background replication worker
- Modular structure suitable for extension into real distributed systems
- Unit tests and Docker-compose setup for a local environment

Architecture

- cmd/stratus-mesh: entrypoint that wires configuration, DB, and services.
- internal/router: Gin router and HTTP route definitions.
- internal/handlers: HTTP handlers for nodes and schedules.
- internal/models: GORM models and migrations.
- internal/store: DB initialization and helper functions.
- internal/services: small packages simulating discovery and replication logic.
- pkg/config: configuration loader from env/file.

Flow
1. Startup loads configuration and initializes the DB (AutoMigrate).
2. Service discovery component maintains a simple registry for simulated nodes.
3. API accepts requests to register nodes and create schedules.
4. Background replicator tries to simulate replication to peer nodes (HTTP-based stub).

Quickstart (with Docker Compose)

Prerequisites: Docker, Docker Compose, make

1) Start dependencies

   docker-compose up -d

2) Build and run the app

   make run

3) API

- Register node: POST /api/v1/nodes {"name":"node-a","address":"http://127.0.0.1:8081"}
- List nodes: GET /api/v1/nodes
- Create schedule: POST /api/v1/schedules {"node_id":1,"task":"backup","cron":"@every 1m"}

Testing

- Unit tests use an in-memory sqlite database. Run:

   make test

Configuration

- Environment variables (see .env.example):
  - DATABASE_DSN: Postgres DSN (e.g. postgres://user:pass@db:5432/stratus?sslmode=disable)
  - PORT: HTTP port (default 8080)

Extending

- Swap the in-memory discovery for an actual gossip protocol (memberlist/Serf).
- Add leader election using Raft/etcd.
- Implement persistent job queue and worker orchestration.

License

MIT
