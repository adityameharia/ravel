package db

import (
	"log"

	badger "github.com/dgraph-io/badger/v3"
)

type RavelDatabase struct {
	conn *badger.DB
}

// Init initialises BadgerDB with the path provided.
func (r *RavelDatabase) Init(path string) error {
	var err error

	options := badger.DefaultOptions(path)
	options.Logger = nil
	options.SyncWrites = true

	r.conn, err = badger.Open(options)
	if err != nil {
		return err
	}
	return nil
}

// Close closes the connection to the BadgerDB instance
func (r *RavelDatabase) Close() {
	err := r.conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
