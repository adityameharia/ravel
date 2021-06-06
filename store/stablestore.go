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

	log.Println("StableStore: Initialised Stable Store")

	return &RavelStableStore{
		Db: &ravelDB,
	}, nil
}

// Set stores Key configuration in a stable manner.
func (s *RavelStableStore) Set(key []byte, val []byte) error {
	log.Println("StableStore: Set")
	return s.Db.Write(key, val)
}

// Get returns the value for the provided key
func (s *RavelStableStore) Get(key []byte) ([]byte, error) {
	log.Println("StableStore: Get")
	val, err := s.Db.Read(key)
	if err != nil {
		if err.Error() == badger.ErrKeyNotFound.Error() {
			val = []byte{}
			return val, nil
		} else {
			log.Fatalln("StableStore: Error retrieving key from db")
		}
	}

	return val, nil
}

// SetUint64 sets val as uint64 for the provided key
func (s *RavelStableStore) SetUint64(key []byte, val uint64) error {
	log.Println("StableStore: SetUint64")
	return s.Db.Write(key, uint64ToBytes(val))
}

// GetUint64 returns the value for the given key
func (s *RavelStableStore) GetUint64(key []byte) (uint64, error) {
	log.Println("StableStore: GetUint64")
	valBytes, err := s.Db.Read(key)

	var valUInt uint64
	if err != nil {
		if err.Error() == badger.ErrKeyNotFound.Error() {
			valUInt = 0
			return valUInt, nil
		} else {
			log.Fatalln("StableStore: Error retrieving key from db")
		}
	}

	valUInt = bytesToUint64(valBytes)
	return valUInt, nil
}
