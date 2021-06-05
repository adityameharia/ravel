# fsm

The fsm package contains the `RavelFSM` and `Snapshot` structs that implements
the [FSM](https://pkg.go.dev/github.com/hashicorp/raft#FSM) interface and
the [FSMSnapshot](https://pkg.go.dev/github.com/hashicorp/raft#FSMSnapshot) interface respectively which are required
for constructing a [new raft node](https://pkg.go.dev/github.com/hashicorp/raft#NewRaft) in
the [hashicorp/raft](https://pkg.go.dev/github.com/hashicorp/raft) library

These interfaces are responsible for actually "applying" the log entries to our BadgerDB instance in a persistent
manner.

## RavelFSM

This struct implements the [FSM](https://pkg.go.dev/github.com/hashicorp/raft#FSM) interface. This interface makes use
of the replicated log to "apply" logs, take snapshots and "restore" from snapshots.

- The `Apply` function is invoked once a log entry is committed and is responsible for storing the data to the BadgerDB
  instance. It checks the type of operation it is required to perform and then calls the corresponding function from
  the `db` package.

- The `Snapshot` function is used for log compaction. Its returns a `Snapshot` object (a struct which implements the
  FSMSnapshot interface) which is used to save a snapshot of the FSM at that point in time i.e. its takes the state of
  the DB and creates a copy of the state so that previous logs can be deleted.

- The `Restore` function is used to restore an FSM from a snapshot i.e. restore the state of the DB to when the snapshot
  was taken thereby discarding all the previous states. This can be done very easily using BadgerDB's inbuilt functions,
  first we drop all the current keys using the DropAll function and then call the Load function to restore the snapshot
  from backup.

## Snapshot

This struct implements the [FSMSnapshot](https://pkg.go.dev/github.com/hashicorp/raft#FSMSnapshot) interface. `Snapshot`
is returned by an FSM in response to a Snapshot call. This interface is responsible for dumping the current state of the
DB i.e. snapshot of the FSM to the [WriteCloser](https://pkg.go.dev/github.com/hashicorp/raft#SnapshotSink)
sink which stored by the raft lib.

- `Persist` is the main function which dumps the snapshot to the sink. We do this by simply taking a backup of our db
  and writing it into the sink.

- `Release` is called when we are finished with the snapshot and all the data has been safely dumped.


NOTE: FSM stands for Finite State Machine



