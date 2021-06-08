package main

import (
	"encoding/json"
	"github.com/adityameharia/ravel/db"
	"github.com/dgraph-io/badger/v3"
	"log"
)

// ReadPartitionOwnersFromDisk reads the RavelConsistentHash.PartitionOwners map from the disk
func ReadPartitionOwnersFromDisk(badgerPath string) (map[uint64]clusterID, error) {
	log.Printf("Reading PartitionOwners from path: %v\n", badgerPath)
	var backupDB db.RavelDatabase
	defer backupDB.Close()
	err := backupDB.Init(badgerPath + "/partition_owners")
	if err != nil {
		return nil, err
	}

	partitionOwnerMap := make(map[uint64]clusterID)
	err = backupDB.Conn.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()

			err = item.Value(func(val []byte) error {
				partitionOwnerMap[bytesToUint64(item.Key())] = clusterID(string(val))
				return nil
			})

			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return partitionOwnerMap, nil
}

// ReadPartitionKeyMapFromDisk reads the RavelConsistentHash.PartitionKeyMap from disk
func ReadPartitionKeyMapFromDisk(badgerPath string) (map[uint64]keySet, error) {
	log.Printf("Reading PartitionKeyMap from path: %v\n", badgerPath)
	var backupDB db.RavelDatabase
	defer backupDB.Close()
	err := backupDB.Init(badgerPath + "/partition_keymap")
	if err != nil {
		return nil, err
	}

	partitionKeyMap := make(map[uint64]keySet)
	err = backupDB.Conn.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		type kSetJSON struct {
			Keys []string `json:"keys"`
		}

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			err = item.Value(func(val []byte) error {
				var k kSetJSON
				err := json.Unmarshal(val, &k)
				if err != nil {
					return err
				}

				kSet := newKeySet()
				for _, key := range k.Keys {
					kSet.Insert([]byte(key))
				}

				partitionKeyMap[bytesToUint64(item.Key())] = kSet
				return nil
			})

			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return partitionKeyMap, nil
}
