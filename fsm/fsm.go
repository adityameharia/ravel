package fsm

import (
	"encoding/json"
	"io"
	"log"

	"github.com/adityameharia/ravel/db"
	"github.com/hashicorp/raft"
)

// RavelFSM implements the raft.FSM interface. It represents the Finite State Machine in a RavelNode. The individual
// logs are "applied" on the FSM on receiving the commit RPC from the leader.
type RavelFSM struct {
	Db *db.RavelDatabase
}

// LogData represents the structure of individual commands on the Logs
type LogData struct {
	Operation string `json:"Operation"`
	Key       []byte `json:"Key"`
	Value     []byte `json:"Value"`
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
func (f *RavelFSM) Get(key []byte) ([]byte, error) {
	log.Println("FSM: Get")
	v, err := f.Db.Read(key)
	if err != nil {
		return []byte{}, err
	}
	return v, nil
}

// Snapshot returns an raft.FSMSnapshot which captures a snapshot of the data at that moment in time
func (f *RavelFSM) Snapshot() (raft.FSMSnapshot, error) {
	log.Println("FSM: Snapshot")
	return &Snapshot{
		Db: f.Db,
	}, nil
}

// Apply commits the given log to the database.
func (f *RavelFSM) Apply(l *raft.Log) interface{} {
	log.Println("FSM: Apply")
	var d LogData

	err := json.Unmarshal(l.Data, &d)
	if err != nil {
		log.Fatal("FSM: Unable to get data from log")
	}

	if d.Operation == "set" {
		return f.Db.Write(d.Key, d.Value)
	} else {
		return f.Db.Delete(d.Key)
	}

}

// Restore restores from the data from the last captured snapshot
func (f *RavelFSM) Restore(r io.ReadCloser) error {
	log.Println("FSM: Restore")
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
}

// Close will close the connection to the internal db.RavelDatabase instance
func (f *RavelFSM) Close() {
	log.Println("FSM: Close")
	f.Db.Close()
}
