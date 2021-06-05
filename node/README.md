# Node

The node package implements some of the main functions the replica has to perform to get started and respond to requests.

- The Open function initialises a [StableStore interface](https://pkg.go.dev/github.com/socketplane/socketplane/Godeps/_workspace/src/github.com/hashicorp/raft?utm_source=godoc#StableStore), [LogStore interface](https://pkg.go.dev/github.com/socketplane/socketplane/Godeps/_workspace/src/github.com/hashicorp/raft?utm_source=godoc#LogStore), [FSM Interface](https://pkg.go.dev/github.com/socketplane/socketplane/Godeps/_workspace/src/github.com/hashicorp/raft?utm_source=godoc#FSM), [SnapshotStore Interface](https://pkg.go.dev/github.com/hashicorp/raft#SnapshotStore) and the [Transport Interface](https://pkg.go.dev/github.com/hashicorp/raft#Transport). It then uses these interfaces to initialize a new raft node.

- The Get,Set and Delete functions are called by the leader in the cluster on corresponding request from the admin

- The Join and Leave function implements the logic for new replicas joining and leaving the replica.