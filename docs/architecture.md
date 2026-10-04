# Stratus Mesh Architecture

This document explains the high-level architecture and where to extend the project.

Components

- API Layer (Gin): A REST interface exposing node and schedule operations.
- Persistence (GORM): Models persisted to a relational DB (Postgres recommended).
- Discovery (in-memory): A simple registry for discovered peers; swap out for gossip/consul/etcd.
- Replicator (background worker): Simulates replication/propagation across discovered peers.

Extensibility Points

- Replace Discovery with memberlist, Serf, or Consul for real peer discovery.
- Add leader election (etcd/raft) to coordinate scheduled job executions.
- Implement a reliable replication protocol and request deduplication.

Operational Notes

- Run database migrations on startup with AutoMigrate (development). For production, use a proper migration tool.
- Secrets and DSNs should be injected through environment variables or secret storage.
