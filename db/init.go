package db

import (
	"log"

	badger "github.com/dgraph-io/badger/v3"
)

var Db *badger.DB

// Init initialises BadgerDB with the path provided.
func Init(path string) error {
	var err error

	options := badger.DefaultOptions(path)
	options.Logger = nil
	options.SyncWrites = true

	Db, err = badger.Open(options)
	if err != nil {
		return err
	}
	return nil
}

// Close closes the connection to the BadgerDB instance
func Close() {
	err := Db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
