package datastore

import (
	"testing"
)

func TestClose(t *testing.T) {
	path := "/tmp/badger_test"
	err := Init(path)
	if err != nil {
		t.Error("Error in connecting to BadgerDB on Host Machine", err)
	}

	err = Close()
	if err != nil {
		t.Error("Error in closing connection to BadgerDB on Host Machine", err)
	}
}
