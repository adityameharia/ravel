package db

import (
	"github.com/dgraph-io/badger/v3"
	"log"
)

// Read returns the value with the corresponding key in the db
func (r *RavelDatabase) Read(key []byte) ([]byte, error) {
	var value []byte

	err := r.Conn.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			value = append([]byte{}, val...)
			return nil
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return value, nil
}

// Write writes the key and value to BadgerDB
func (r *RavelDatabase) Write(key []byte, val []byte) error {
	err := r.Conn.Update(func(txn *badger.Txn) error {
		err := txn.Set(key, val)
		if err != nil {
			return err
		}
		return nil
	})

	return err
}

// Delete deletes the key value pair with the corresponding key from BadgerDB
func (r *RavelDatabase) Delete(key []byte) error {
	err := r.Conn.Update(func(txn *badger.Txn) error {
		err := txn.Delete(key)
		if err != nil {
			return err
		}
		return nil
	})

	return err
}
