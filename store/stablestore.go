package store

import (
	"log"

	"github.com/adityameharia/ravel/db"
	"github.com/dgraph-io/badger/v3"
)

type RavelStableStore struct {
	Db *db.RavelDatabase
}

func NewRavelStableStore(logDBPath string) (*RavelStableStore, error) {
	var ravelDB db.RavelDatabase
	err := ravelDB.Init(logDBPath)
	if err != nil {
		log.Fatal("StableStore: Unable to setup new Stable Store")
		return nil, err
	}

	return &RavelStableStore{
		Db: &ravelDB,
	}, nil
}

func (s *RavelStableStore) Set(key []byte, val []byte) error {
	return s.Db.Write([]byte(key), []byte(val))
}

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

func (s *RavelStableStore) SetUint64(key []byte, val uint64) error {
	return s.Db.Write([]byte(key), uint64ToBytes(val))
}

func (s *RavelStableStore) GetUint64(key []byte) (uint64, error) {
	valBytes, err := s.Db.Read([]byte(key))
	valUInt := bytesToUint64(valBytes)
	if err == badger.ErrKeyNotFound {
		valUInt = 0
	}
	if err != nil && err != badger.ErrKeyNotFound {
		log.Println("StableStore: Error retrieving key from db")
	}
	return valUInt, nil
}
