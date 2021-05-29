package fsm

import (
	"io"
	"log"

	"github.com/adityameharia/ravel/db"
	"github.com/hashicorp/raft"
	"github.com/vmihailenco/msgpack/v5"
)

type Fsm struct {
	Db *db.RavelDatabase
}

type LogData struct {
	Operation string `json:"op,omitempty"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

func NewFSM(path string) (*Fsm, error) {
	var r db.RavelDatabase
	err := r.Init(path)
	if err != nil {
		log.Fatal("FSM: Unable to Setup Database")
		return nil, err
	}

	log.Println("FSM: Initialised FSM")

	return &Fsm{
		Db: &r,
	}, nil
}

func (f *Fsm) Get(key string) (string, error) {
	log.Println("FSM: Getting Key")
	v, err := f.Db.Read([]byte(key))
	if err != nil {
		return "", err
	}
	return string(v), nil
}

//returns a FSMSnapshot object for future use by raft lib
func (f *Fsm) Snapshot() (raft.FSMSnapshot, error) {
	log.Println("FSM: Generate FSMSnapshot")
	return &FSMSnapshot{
		Db: f.Db,
	}, nil
}

func (f *Fsm) Apply(l *raft.Log) interface{} {
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

func (f *Fsm) Restore(r io.ReadCloser) error {
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

func (f *Fsm) Close() {
	log.Println("FSM: Close called")
	f.Db.Close()
}
