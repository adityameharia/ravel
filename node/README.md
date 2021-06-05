## node

The node package implements some main functions the replica has to perform to get started and respond to requests.

- The `Open` function initialises structs that satisfy the
  - [StableStore](https://pkg.go.dev/github.com/hashicorp/raft#StableStore) interface
  - [LogStore](https://pkg.go.dev/github.com/hashicorp/raft#LogStore) interface
  - [FSM](https://pkg.go.dev/github.com/hashicorp/raft#FSM) interface
  - [SnapshotStore](https://pkg.go.dev/github.com/hashicorp/raft#SnapshotStore) interface and
  - [Transport](https://pkg.go.dev/github.com/hashicorp/raft#Transport) interface. 
  
It then uses these structs to initialise a new raft node.

- The `Get`, `Set` and `Delete` functions are called by the leader in the cluster on corresponding request from the
  admin
- The `Join` and `Leave` functions implements the logic for new replicas joining and leaving the replica.