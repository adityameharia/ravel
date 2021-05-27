package datastore

import (
	"fmt"

	badger "github.com/dgraph-io/badger/v3"
)

var db *badger.DB

func OpenDb() error {
	var err error
	db, err = badger.Open(badger.DefaultOptions("/home/adi/badger"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func closeDb() {
	db.Close()
}
