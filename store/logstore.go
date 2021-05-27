package store

import (
	"github.com/adityameharia/ravel/db"
	"github.com/dgraph-io/badger/v3"
	"github.com/hashicorp/raft"
	"log"
	"os"
)

type RavelLogStore struct {
	db *db.RavelDatabase
}

func NewRavelLogStore() (*RavelLogStore, error) {
	var ravelDb db.RavelDatabase
	path := os.Getenv("LOG_STORE_PATH")
	err := ravelDb.Init(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &RavelLogStore{
		db: &ravelDb,
	}, nil
}

func (r *RavelLogStore) FirstIndex() (uint64, error) {
	var key uint64
	err := r.db.Conn.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		it.Rewind()
		if it.Valid() {
			firstKey := it.Item().Key()
			key = bytesToUint64(firstKey)
		} else {
			key = 0
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return key, nil
}

func (r *RavelLogStore) LastIndex() (uint64, error) {
	return 0, nil
}

func (r *RavelLogStore) GetLog(index uint64, log *raft.Log) {}

func (r *RavelLogStore) StoreLog(log *raft.Log) error {
	return r.StoreLogs([]*raft.Log{log})
}

func (r *RavelLogStore) StoreLogs(logs []*raft.Log) error {
	for _, l := range logs {
		key := uint64ToBytes(l.Index)
		val := raftLogToBytes(*l)

		err := r.db.Write(key, val)
		return err
	}

	return nil
}

func (r *RavelLogStore) DeleteRange(min, max uint64) error {
	return nil
}
