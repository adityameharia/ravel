The fsm package implements the [FSM interface](https://pkg.go.dev/github.com/socketplane/socketplane/Godeps/_workspace/src/github.com/hashicorp/raft?utm_source=godoc#FSM) and the [FSMSnapshot interface](https://pkg.go.dev/github.com/socketplane/socketplane/Godeps/_workspace/src/github.com/hashicorp/raft?utm_source=godoc#FSMSnapshot) which are responsible for actually "applying" the log entries to our BadgerDB instance in a persistance manner.

<br/>

## FSM Interface 

<br/>
This interface makes use of the replicated log to "apply" logs,take snapshots and "restore" from snapshots.

<br/>
<br/>

- The Apply function is invoked once a log entry is commited and it is responsible for storing the data to the BadgerDB instace

- The Snapshot function is used for log compaction.Its returns and FSMSnapshot(a struct which implementsthe FSMSnapshot interface) which is used to save a snapshot of the FSM at that point in time i.e. its takes the state of the DB and creates a copy of the state so that previous logs can be deleted.

- The Restore function is used to restore an FSM from a snapshot i.e. restore the state of the DB to when the snapshot was taken thereby discarding all previous state.

## FSMSnapshot Interface

<br/>
FSMSnapshot is returned by an FSM in response to a Snapshot.This interface is responsible for dumping the current state of the DB i.e. snapshot of the FSM to the WriteCloser 'sink' which stored by the raft lib.

<br/>
<br/>

- The Persist function is the main function in this package which dumps the snapshot to the sink.

- The Release function is called when we are finished with the snapshot and all the data has been safely dumped.




