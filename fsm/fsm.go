package fsm

import (
	"log"
	"os"

	"github.com/adityameharia/ravel/db"
	"github.com/dgraph-io/badger/v3"
	"github.com/joho/godotenv"
)

type fsm struct {
	db *badger.DB
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

	return &fsm{
		db: r.Conn,
	}, nil
}
