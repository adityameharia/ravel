package fsm

import (
	"log"

	"github.com/adityameharia/ravel/db"
	"github.com/hashicorp/raft"
)

// Snapshot implements the raft.Snapshot interface
type Snapshot struct {
	Db *db.RavelDatabase
}

type KeyValue struct {
	Key   []byte `json:"key"`
	Value []byte `json:"value"`
}

// Persist writes a backup of the db to the sink
func (f *Snapshot) Persist(sink raft.SnapshotSink) error {
	log.Println("Snapshot: Starting Snapshot")

	_, err := f.Db.Conn.Backup(sink, 0)
	if err != nil {
		log.Println("Snapshot: Unable to take Snapshot")
		return err
	}

	err = sink.Close()
	if err != nil {
		log.Println("Snapshot: Unable to close Sink")
		return err
	}

	return nil
}

// Release releases the snapshot
func (f *Snapshot) Release() {
	log.Println("Snapshot: Snapshot finished")
}
