package db

import (
	"log"
	"testing"

	"github.com/dgraph-io/badger/v3"
)

func Setup() {
	path := "/tmp/badger_test"
	err := Init(path)
	if err != nil {
		log.Println("Error in starting connection with Badger")
		log.Println(err)
	}

	err = Db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("k1"), []byte("v1"))
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Println("Error in Setting up transaction_test.go")
		log.Println(err)
	}
}

func TestRead(t *testing.T) {
	Setup()
	defer Close()

	v, err := Read([]byte("k1"))
	if err != nil {
		t.Error("Error in Read", err)
	}

	if string(v) != "v1" {
		t.Error("Error in read value", err)
	}

}

func TestWrite(t *testing.T) {
	Setup()
	defer Close()

	err := Write([]byte("k2"), []byte("v2"))
	if err != nil {
		t.Error("Error in writing to Badger", err)
	}

}

func TestDelete(t *testing.T) {
	Setup()
	defer Close()

	err := Delete([]byte("k1"))
	if err != nil {
		t.Error("Error in deleting from Badger", err)
	}
}
