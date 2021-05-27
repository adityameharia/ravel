package db

import (
	"testing"
)

func TestClose(t *testing.T) {
	path := "/tmp/badger_test"
	var r RavelDatabase
	err := r.Init(path)
	defer r.Close()

	if err != nil {
		t.Error("Error in connecting to BadgerDB on Host Machine", err)
	}
}
