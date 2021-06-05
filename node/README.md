# Node

The node package implements some main functions the replica has to perform to get started and respond to requests.

- The Open function initialises a [StableStore interface](https://pkg.go.dev/github.com/hashicorp/raft#StableStore)
  , [LogStore interface](https://pkg.go.dev/github.com/hashicorp/raft#LogStore)
  , [FSM Interface](https://pkg.go.dev/github.com/hashicorp/raft#FSM)
  , [SnapshotStore Interface](https://pkg.go.dev/github.com/hashicorp/raft#SnapshotStore) and
  the [Transport Interface](https://pkg.go.dev/github.com/hashicorp/raft#Transport). It then uses these interfaces to
  initialise a new raft node.

- The `Get`, `Set` and `Delete` functions are called by the leader in the cluster on corresponding request from the admin
- The `Join` and `Leave` functions implements the logic for new replicas joining and leaving the replica.