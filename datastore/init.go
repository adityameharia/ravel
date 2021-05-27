package datastore

import (
	badger "github.com/dgraph-io/badger/v3"
)

var db *badger.DB

// Init initialises BadgerDB with the path provided.
func Init(path string) error {
	var err error
	options := badger.DefaultOptions(path)
	options.Logger = nil
	db, err = badger.Open(options)
	if err != nil {
		return err
	}
	return nil
}

// Close closes the connection to the BadgerDB instance
func Close() error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}
