package db

import (
	"log"

	badger "github.com/dgraph-io/badger/v3"
)

type RavelDatabase struct {
	Conn *badger.DB
}

// Init initialises BadgerDB with the path provided.
func (r *RavelDatabase) Init(path string) (*badger.DB, error) {
	var err error

	options := badger.DefaultOptions(path)
	options.Logger = nil
	options.SyncWrites = true

	r.Conn, err = badger.Open(options)
	if err != nil {
		return nil, err
	}
	return r.Conn, nil
}

// Close closes the connection to the BadgerDB instance
func (r *RavelDatabase) Close() {
	err := r.Conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
