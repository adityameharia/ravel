package datastore

import (
	badger "github.com/dgraph-io/badger/v3"
	"log"
)

var db *badger.DB

// Init initialises BadgerDB with the path provided.
func Init(path string) error {
	var err error

	options := badger.DefaultOptions(path)
	options.Logger = nil
	options.SyncWrites = true

	db, err = badger.Open(options)
	if err != nil {
		return err
	}
	return nil
}

// Close closes the connection to the BadgerDB instance
func Close() {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
