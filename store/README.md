# Store

The store package implements the [StableStore](https://pkg.go.dev/github.com/hashicorp/raft#StableStore) interface and
the [LogStore](https://pkg.go.dev/github.com/hashicorp/raft#LogStore)
which are required for constructing a [new raft node](https://pkg.go.dev/github.com/hashicorp/raft#NewRaft)

These interfaces are used for storing and retrieving logs and other key configurations of the node.

## StableStore Interface

StableStore is used to provide stable storage of key configurations to ensure safety.The Set/Get and SetUint64/GetUint64
functions are used to set/get key-value pairs of type `[]byte` and `uint64` respectively.

## LogStore Interface

The LogStore interface stores and retrieves the logs in a persistent manner.

- The FirstIndex and LastIndex functions return the index property of the first and last log respectively. These
  functions are used to check whether the logs of the follower are consistent with that of the leader.

- The GetLog function gets the log with the given index and writes it to the pointer of
  type [Log](https://pkg.go.dev/github.com/hashicorp/raft#Log) passed to it.

- The StoreLog function is used to store a single Log to disk. It calls the StoreLogs function passing it an array of
  the given log.

- The StoreLogs function is perhaps the most important function in this interface. It takes in an array of logs and
  actually persists that data onto disk using [BadgerDB](https://github.com/dgraph-io/badger). It is used by the new
  nodes to store all the logs it receives from the leader or by the existing nodes to store individual logs.