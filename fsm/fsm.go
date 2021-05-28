package fsm

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/adityameharia/ravel/db"
	"github.com/hashicorp/raft"
	"github.com/joho/godotenv"
	"github.com/vmihailenco/msgpack/v5"
)

type fsm struct {
	db *db.RavelDatabase
}

type logData struct {
	Operation string `json:"op,omitempty"`
	Key       string `json:"key"`
	Value     string `json:"value"`
}

func NewFSM(path string) (*fsm, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	badgerPath := os.Getenv("BADGER_PATH")
	var r db.RavelDatabase
	err = r.Init(badgerPath)
	if err != nil {
		log.Fatal(("Unable to Setup Databas"))
		return nil, err
	}

	log.Println("Initialised FSM")

	return &fsm{
		db: &r,
	}, nil
}

func (f *fsm) Get(key string) (string, error) {
	v, err := f.db.Read([]byte(key))
	if err != nil {
		return "", err
	}
	return string(v), nil
}

//returns a FSMSnapshot object for future use by raft lib
func (f *fsm) Snapshot() (raft.FSMSnapshot, error) {
	log.Println("Generate FSMSnapshot")
	return &FSMSnapshot{
		db: f.db,
	}, nil
}

func (f *fsm) Apply(l *raft.Log) interface{} {
	var d logData

	err := msgpack.Unmarshal(l.Data, &d)
	if err != nil {
		log.Fatal("Unable to get data from log")
	}

	if d.Operation == "set" {
		return f.db.Write([]byte(d.Key), []byte(d.Value))
	} else {
		return f.db.Delete([]byte(d.Key))
	}

}

func (f *fsm) Restore(r io.ReadCloser) error {
	log.Println("Restoring from Snapshot")
	kvBuffer, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal("Unable to read Snapshot")
		return err
	}

	var KV []KeyValue

	err = msgpack.Unmarshal(kvBuffer, KV)
	if err != nil {
		log.Fatal(("Unable to unmarshal Snapshot"))
		return err
	}

	for _, kv := range KV {
		err = f.db.Write([]byte(kv.Key), []byte(kv.Value))
		if err != nil {
			log.Fatal(("Unable to write key"))
			return err
		}
	}

	log.Println("Snapshot restored")
	return nil
}

func (f *fsm) Close() {
	f.db.Close()
}
