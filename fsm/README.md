# FSM

The fsm package implements the [FSM interface](https://pkg.go.dev/github.com/socketplane/socketplane/Godeps/_workspace/src/github.com/hashicorp/raft?utm_source=godoc#FSM) and the [FSMSnapshot interface](https://pkg.go.dev/github.com/socketplane/socketplane/Godeps/_workspace/src/github.com/hashicorp/raft?utm_source=godoc#FSMSnapshot) which are required for constructing a [new raft node](https://pkg.go.dev/github.com/socketplane/socketplane/Godeps/_workspace/src/github.com/hashicorp/raft?utm_source=godoc#NewRaft).These interfaces are responsible for actually "applying" the log entries to our BadgerDB instance in a persistance manner.

<br/>

## FSM Interface 

This interface makes use of the replicated log to "apply" logs,take snapshots and "restore" from snapshots.

- The Apply function is invoked once a log entry is commited and it is responsible for storing the data to the BadgerDB instace.It check the type of operation it is required to perform and then calls the corresponding function from the db package.

- The Snapshot function is used for log compaction.Its returns and FSMSnapshot(a struct which implementsthe FSMSnapshot interface) which is used to save a snapshot of the FSM at that point in time i.e. its takes the state of the DB and creates a copy of the state so that previous logs can be deleted.

- The Restore function is used to restore an FSM from a snapshot i.e. restore the state of the DB to when the snapshot was taken thereby discarding all previous state.This can be done very easily using BadgerDB's inbuilt functions,first we drop all the current keys using the DropAll function and then call the Load function to restore the snapshot from backup.

<br/>

## FSMSnapshot Interface

FSMSnapshot is returned by an FSM in response to a Snapshot.This interface is responsible for dumping the current state of the DB i.e. snapshot of the FSM to the [WriteCloser sink](https://pkg.go.dev/github.com/socketplane/socketplane/Godeps/_workspace/src/github.com/hashicorp/raft?utm_source=godoc#SnapshotSink) which stored by the raft lib.

- The Persist function is the main function in this package which dumps the snapshot to the sink.We do this by simply taking a backup of our db and writing it into the sink.

- The Release function is called when we are finished with the snapshot and all the data has been safely dumped.



