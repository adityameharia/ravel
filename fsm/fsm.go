package fsm

import (
	"io"
	"log"

	"github.com/adityameharia/ravel/db"
	"github.com/hashicorp/raft"
	"github.com/vmihailenco/msgpack/v5"
)

// RavelFSM implements the raft.FSM interface. It represents the Finite State Machine in a RavelNode. The individual
// logs are "applied" on the FSM on receiving the commit RPC from the leader.
type RavelFSM struct {
	Db *db.RavelDatabase
}

// LogData represents the structure of individual commands on the Logs
type LogData struct {
	Operation string
	Key       string
	Value     string
}

// NewFSM creates an instance of RavelFSM
func NewFSM(path string) (*RavelFSM, error) {
	var r db.RavelDatabase
	err := r.Init(path)
	if err != nil {
		log.Fatal("FSM: Unable to Setup Database")
		return nil, err
	}

	log.Println("FSM: Initialised FSM")

	return &RavelFSM{
		Db: &r,
	}, nil
}

// Get returns the value for the provided key
func (f *RavelFSM) Get(key string) (string, error) {
	log.Println("FSM: Getting Key")
	v, err := f.Db.Read([]byte(key))
	if err != nil {
		return "", err
	}
	return string(v), nil
}

// Snapshot returns an raft.FSMSnapshot which captures a snapshot of the data at that moment in time
func (f *RavelFSM) Snapshot() (raft.FSMSnapshot, error) {
	log.Println("FSM: Generate FSMSnapshot")
	return &FSMSnapshot{
		Db: f.Db,
	}, nil
}

// Apply commits the given log to the database.
func (f *RavelFSM) Apply(l *raft.Log) interface{} {
	log.Println("FSM: Applying set/delete")
	var d LogData

	err := msgpack.Unmarshal(l.Data, &d)
	if err != nil {
		log.Fatal("FSM: Unable to get data from log")
	}

	if d.Operation == "set" {
		return f.Db.Write([]byte(d.Key), []byte(d.Value))
	} else {
		return f.Db.Delete([]byte(d.Key))
	}

}

// Restore restores from the data from the last captured snapshot
func (f *RavelFSM) Restore(r io.ReadCloser) error {
	log.Println("FSM: Restore called")
	err := f.Db.Conn.DropAll()
	if err != nil {
		log.Fatal("FSM: Unable to delete previous state")
		return err
	}
	err = f.Db.Conn.Load(r, 100)
	if err != nil {
		log.Fatal("FSM: Unable to restore Snapshot")
		return err
	}
	return nil

	// log.Println("Restoring from Snapshot")
	// kvBuffer, err := ioutil.ReadAll(r)
	// if err != nil {
	// 	log.Fatal("Unable to read Snapshot")
	// 	return err
	// }

	// var KV []KeyValue

	// err = msgpack.Unmarshal(kvBuffer, KV)
	// if err != nil {
	// 	log.Fatal(("Unable to unmarshal Snapshot"))
	// 	return err
	// }

	// for _, kv := range KV {
	// 	err = f.db.Write([]byte(kv.Key), []byte(kv.Value))
	// 	if err != nil {
	// 		log.Fatal(("Unable to write key"))
	// 		return err
	// 	}
	// }

	// log.Println("Snapshot restored")
}

// Close will close the connection to the internal db.RavelDatabase instance
func (f *RavelFSM) Close() {
	log.Println("FSM: Close called")
	f.Db.Close()
}
