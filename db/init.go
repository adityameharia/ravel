package db

import (
	"log"

	badger "github.com/dgraph-io/badger/v3"
)

// RavelDatabase represents an overlay on top of BadgerDB, it exposes Init, Close, Read, Write and Delete functions -
// these functions eventually perform mentioned operations on an instance of BadgerDB and these operations persist on disk
type RavelDatabase struct {
	Conn *badger.DB
}

// Init initialises BadgerDB with the path provided.
func (r *RavelDatabase) Init(path string) error {
	var err error

	options := badger.DefaultOptions(path)
	options.Logger = nil
	options.SyncWrites = true

	r.Conn, err = badger.Open(options)
	if err != nil {
		return err
	}
	return nil
}

// Close closes the connection to the BadgerDB instance
func (r *RavelDatabase) Close() {
	err := r.Conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
