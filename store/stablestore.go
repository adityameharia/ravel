package store

import (
	"log"

	"github.com/adityameharia/ravel/db"
	"github.com/dgraph-io/badger/v3"
)

// RavelStableStore implements the raft.StableStore interface. It stores the configuration for raft.Raft
type RavelStableStore struct {
	Db *db.RavelDatabase
}

// NewRavelStableStore creates a new instance of RavelStableStore
func NewRavelStableStore(stableStoreDBPath string) (*RavelStableStore, error) {
	var ravelDB db.RavelDatabase
	err := ravelDB.Init(stableStoreDBPath)
	if err != nil {
		log.Fatal("StableStore: Unable to setup new Stable Store")
		return nil, err
	}

	return &RavelStableStore{
		Db: &ravelDB,
	}, nil
}

// Set stores Key configuration in a stable manner.
func (s *RavelStableStore) Set(key []byte, val []byte) error {
	return s.Db.Write([]byte(key), []byte(val))
}

// Get returns the value for the provided key
func (s *RavelStableStore) Get(key []byte) ([]byte, error) {
	val, err := s.Db.Read([]byte(key))
	if err == badger.ErrKeyNotFound {
		val = []byte{}
	}
	if err != nil && err != badger.ErrKeyNotFound {
		log.Println("StableStore: Error retrieving key from db")
	}
	return val, nil
}

// SetUint64 sets val as uint64 for the provided key
func (s *RavelStableStore) SetUint64(key []byte, val uint64) error {
	return s.Db.Write(key, uint64ToBytes(val))
}

// GetUint64 returns the value for the given key
func (s *RavelStableStore) GetUint64(key []byte) (uint64, error) {
	valBytes, err := s.Db.Read(key)
	valUInt := bytesToUint64(valBytes)
	if err == badger.ErrKeyNotFound {
		valUInt = 0
	}
	if err != nil && err != badger.ErrKeyNotFound {
		log.Println("StableStore: Error retrieving key from db")
	}
	return valUInt, nil
}