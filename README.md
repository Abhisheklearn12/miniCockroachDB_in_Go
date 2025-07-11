

```md
# 🐓 mini-cockroachdb-go

**mini-cockroachdb-go** is a minimal, educational, distributed SQL database inspired by [CockroachDB](https://www.cockroachlabs.com/), implemented in Go with under 3000 lines of code.

It helps you deeply learn:

- ✅ Raft-based consensus
- ✅ MVCC key-value store
- ✅ Distributed transactions
- ✅ Lightweight SQL engine
- ✅ Gossip-based peer discovery
- ✅ Persistent storage via Pebble

> ⚠️ This project is intentionally minimal and focused on **learning system design and Go**, not production use.

---

## ⚙️ Architecture

```

```
    ┌────────────┐
    │   SQL      │  ← SQL Parser & Planner
    └────┬───────┘
         ↓
    ┌────────────┐
    │    KV      │  ← MVCC, Transactions
    └────┬───────┘
         ↓
    ┌────────────┐
    │   Raft     │  ← Distributed Consensus
    └────┬───────┘
         ↓
    ┌────────────┐
    │  Storage   │  ← Persistent engine (Pebble)
    └────────────┘
```

````

---

## 📁 Project Structure

```bash
mini-cockroachdb-go/
├── go.mod                  # Go module
├── cmd/
│   └── mini_cockroachdb/   # Entry point binary (this is your executable name)
│       └── main.go
├── api/                    # gRPC API and handlers
│   ├── service.proto
│   ├── service.pb.go
│   └── handler.go
├── raft/                   # Raft consensus
│   ├── node.go
│   ├── log.go
│   ├── transport.go
│   └── raft_test.go
├── kv/                     # MVCC key-value layer
│   ├── mvcc.go
│   ├── txn.go
│   └── kv_test.go
├── storage/                # Persistent Pebble store
│   ├── pebble.go
│   └── engine.go
├── sql/                    # SQL engine
│   ├── parser.go
│   ├── planner.go
│   └── executor.go
├── gossip/                 # Peer discovery via gossip
│   ├── gossip.go
│   └── info.go
├── node/                   # Node lifecycle
│   ├── node.go
│   └── server.go
├── util/                   # Shared utilities
│   ├── clock.go
│   ├── retry.go
│   └── log.go
└── README.md               # You're here
````

---

## 🚀 Getting Started

### 🔧 Prerequisites

* Go 1.22+
* Protobuf compiler + plugins:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

---

### 🛠 Setup

```bash
# Clone the project
git clone https://github.com/yourusername/mini-cockroachdb-go.git
cd mini-cockroachdb-go/mini_cockroachdb

# Install dependencies
go mod tidy

# Build the binary
go build -o mini-cockroachdb ./cmd/mini_cockroachdb
```

---

### 🧪 Run a Single Node

```bash
./mini_cockroachdb --node-id=1 --port=9001 --join=""
```

To simulate a cluster, start other nodes on different ports and pass `--join=host:port`.

---

## 🧠 What You’ll Learn

* Go concurrency (goroutines, channels, locking)
* Raft consensus (elections, log replication)
* MVCC and timestamp-based key-value versioning
* Distributed transactions and intents
* SQL parsing, planning, and execution
* Cluster coordination via gossip
* Persistent LSM-tree storage (via Pebble)

---

## 🧪 Run All Tests

```bash
go test ./...
```

---

## 📘 References

* [CockroachDB Architecture Docs](https://www.cockroachlabs.com/docs/stable/architecture/overview.html)
* [Raft Consensus Paper](https://raft.github.io/raft.pdf)
* [Google Spanner Paper](https://research.google/pubs/archive/45855.pdf)
* [Pebble Key-Value Store](https://github.com/cockroachdb/pebble)

---

## 🙋‍♂️ Author

**Abhishek Kumar**
---

## 🛡 License

MIT License – free to fork, learn, and build from.

---

## 🧭 Next Steps

Want help writing:

* `main.go` CLI setup?
* Raft `node.go` core loop?
* Basic SQL `parser.go` example?


