package store

import (
	"github.com/adityameharia/ravel/db"
	"github.com/dgraph-io/badger/v3"
	"github.com/hashicorp/raft"
)

type RavelLogStore struct {
	Db *db.RavelDatabase
}

func NewRavelLogStore(logDBPath string) (*RavelLogStore, error) {
	var ravelDB db.RavelDatabase
	err := ravelDB.Init(logDBPath)
	if err != nil {
		return nil, err
	}

	return &RavelLogStore{
		Db: &ravelDB,
	}, nil
}

func (r *RavelLogStore) FirstIndex() (uint64, error) {
	var key uint64
	err := r.Db.Conn.View(func(txn *badger.Txn) error {
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
	var key uint64
	err := r.Db.Conn.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Reverse = true
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

func (r *RavelLogStore) GetLog(index uint64, raftLog *raft.Log) error {
	key := uint64ToBytes(index)
	val, err := r.Db.Read(key)
	if err != nil {
		return err
	}

	return bytesToRaftLog(val, raftLog)
}

func (r *RavelLogStore) StoreLog(log *raft.Log) error {
	return r.StoreLogs([]*raft.Log{log})
}

func (r *RavelLogStore) StoreLogs(logs []*raft.Log) error {
	for _, l := range logs {
		key := uint64ToBytes(l.Index)
		val := raftLogToBytes(*l)

		err := r.Db.Write(key, val)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *RavelLogStore) DeleteRange(min uint64, max uint64) error {
	minKey := uint64ToBytes(min)

	txn := r.Db.Conn.NewTransaction(true)
	defer txn.Discard()

	opts := badger.DefaultIteratorOptions
	it := txn.NewIterator(opts)
	it.Seek(minKey)

	for {
		key := it.Item().Key()

		if bytesToUint64(key) > max {
			break
		}

		err := r.Db.Delete(key)
		if err != nil {
			return err
		}

		it.Next()
	}

	if err := txn.Commit(); err != nil {
		return err
	}
	return nil
}
