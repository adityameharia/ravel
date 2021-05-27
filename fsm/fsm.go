package fsm

import (
	"log"
	"os"

	"github.com/adityameharia/ravel/db"
	"github.com/hashicorp/raft"
	"github.com/joho/godotenv"
)

type fsm struct {
	db *db.RavelDatabase
}

type data struct {
	Op    string `json:"op,omitempty"`
	Key   string `json:"key"`
	Value string `json:"value"`
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
}
