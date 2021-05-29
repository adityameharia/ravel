package store

import (
	"testing"
)

func TestSet(t *testing.T) {
	s, err := NewRavelStableStore("/tmp/badger/test/stable")
	if err != nil {
		t.Error("Error in newstablestore")
	}

	err = s.Set([]byte("testKey"), []byte("testValue"))
	if err != nil {
		t.Error("Error in Set stable store")
	}

}

func TestSetUint64(t *testing.T) {
	s, err := NewRavelStableStore("/tmp/badger/test/stable")
	if err != nil {
		t.Error("Error in newstablestore")
	}

	err = s.SetUint64([]byte("testKey"), 1)
	if err != nil {
		t.Error("Error in SetUint64 stable store")
	}
}

func TestGet(t *testing.T) {
	s, err := NewRavelStableStore("/tmp/badger/test/stable")
	if err != nil {
		t.Error("Error in newstablestore")
	}

	err = s.Set([]byte("testKey"), []byte("testValue"))
	if err != nil {
		t.Error("Error in Set stable store")
	}

	val, err := s.Get([]byte("testKey"))
	if err != nil {
		t.Error("Error in Get stable store")
	}
	t.Logf(string(val))
}

func TestGetUint64(t *testing.T) {
	s, err := NewRavelStableStore("/tmp/badger/test/stable")
	if err != nil {
		t.Error("Error in newstablestore")
	}

	err = s.SetUint64([]byte("testKey"), 234)
	if err != nil {
		t.Error("Error in Set stable store")
	}

	val, err := s.GetUint64([]byte("testKey"))
	if err != nil {
		t.Error("Error in Get stable store")
	}
	t.Log(val)
}
