package store

import (
	"log"

	"github.com/adityameharia/ravel/db"
	"github.com/dgraph-io/badger/v3"
	"github.com/hashicorp/raft"
)

// RavelLogStore implements raft.LogStore interface. The functions define the operations possible on the Logs which is
// maintained by every instance of raft.Raft
type RavelLogStore struct {
	Db *db.RavelDatabase
}

// NewRavelLogStore creates a new instance of RavelLogStore, logDBPath specifies the directory path to
// initialise the internal db.RavelDatabase instance. An entry in the Logs is of type raft.Log
func NewRavelLogStore(logDBPath string) (*RavelLogStore, error) {
	var ravelDB db.RavelDatabase
	err := ravelDB.Init(logDBPath)
	if err != nil {
		log.Fatalf("NewRavelLogStore: %v\n", err)
		return nil, err
	}

	log.Println("LogStore: Initialised Log Store")
	return &RavelLogStore{
		Db: &ravelDB,
	}, nil
}

// FirstIndex returns the Index property of the first entry in the Logs.
func (r *RavelLogStore) FirstIndex() (uint64, error) {
	log.Println("LogStore: FirstIndex")
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
		log.Fatalf("RavelLogStore.FirstIndex: %v\n", err)
		return 0, err
	}

	return key, nil
}

// LastIndex returns the Index property of the last entry in the Logs
func (r *RavelLogStore) LastIndex() (uint64, error) {
	log.Println("LogStore: LastIndex")
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
		log.Fatalf("RavelLogStore.LastIndex: %v\n", err)
		return 0, err
	}
	return key, nil
}

// GetLog writes the log on position "index" to the pointer "raftLog"
func (r *RavelLogStore) GetLog(index uint64, raftLog *raft.Log) error {
	key := uint64ToBytes(index)
	val, err := r.Db.Read(key)
	if err != nil {
		log.Printf("RavelLogStore.GetLog: %v\n", err)
		return err
	}

	return bytesToRaftLog(val, raftLog)
}

// StoreLog inserts a single raft.Log at the end of the Logs
func (r *RavelLogStore) StoreLog(l *raft.Log) error {
	log.Println("LogStore: StoreLog")
	return r.StoreLogs([]*raft.Log{l})
}

// StoreLogs inserts []raft.Log at the end of the Logs
func (r *RavelLogStore) StoreLogs(logs []*raft.Log) error {
	log.Println("LogStore: StoreLogs")
	for _, l := range logs {
		key := uint64ToBytes(l.Index)
		val := raftLogToBytes(*l)

		err := r.Db.Write(key, val)
		if err != nil {
			log.Fatalf("RavelLogStore.StoreLogs: %v\n", err)
			return err
		}
	}

	return nil
}

// DeleteRange deletes the entries from "min" to "max" position (both inclusive) in the Logs
func (r *RavelLogStore) DeleteRange(min uint64, max uint64) error {
	log.Println("LogStore: DeleteRange")
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
		log.Fatalf("RavelLogStore.DeleteRange: %v\n", err)
		return err
	}
	return nil
}
