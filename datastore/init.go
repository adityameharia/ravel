package datastore

import (
	badger "github.com/dgraph-io/badger/v3"
)

var db *badger.DB

// Init initialises BadgerDB with the path provided.
func Init(path string) error {
	var err error
	db, err = badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return err
	}
	return nil
}

func Close() error {
	err := db.Close()
	if err != nil {
		return err
	}

	return nil
}
