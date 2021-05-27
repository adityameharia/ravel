package fsm

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/adityameharia/ravel/db"
	"github.com/dgraph-io/badger/v3"
	"github.com/hashicorp/raft"
)

type KeyValue struct {
	Key   []byte `json:"key"`
	Value []byte `json:"value"`
}

func Persist(sink raft.SnapshotSink) error {

	log.Println("Starting Snapshot")

	c := make(chan KeyValue)

	var kv KeyValue

	//iterates over all the key in a seperate go routine and passes the values read into a channel
	go db.KeyValueDB.View(func(txn *badger.Txn) error {

		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			kv.Key = item.Key()

			err := item.Value(func(v []byte) error {
				kv.Value = v
				return nil
			})

			if err != nil {
				log.Fatal(err)
			}

			c <- kv

		}
		close(c)
		return nil
	})

	KVBuffer := new(bytes.Buffer)
	var dataRead KeyValue
	var ok bool

	//reads the key values from the channel until it is closed and all the values have been read
	for {
		dataRead, ok = (<-c)
		if !ok {
			break
		}

		json.NewEncoder(KVBuffer).Encode(dataRead)

		if _, err := sink.Write(KVBuffer.Bytes()); err != nil {
			log.Fatal(err)
		}

	}

	log.Println("All keys have been persisted")

	return nil
}

func Release() {
	log.Println("Snapshot finised")
}
