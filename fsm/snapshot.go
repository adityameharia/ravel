package fsm

import (
	"log"

	"github.com/adityameharia/ravel/db"
	"github.com/hashicorp/raft"
)

// Snapshot implements the Snapshot interface
type Snapshot struct {
	Db *db.RavelDatabase
}

type KeyValue struct {
	Key   []byte `json:"key"`
	Value []byte `json:"value"`
}

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

func (f *Snapshot) Release() {
	log.Println("Snapshot: Snapshot finished")
}
