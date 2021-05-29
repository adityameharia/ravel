package fsm

import (
	"log"

	"github.com/adityameharia/ravel/db"
	"github.com/hashicorp/raft"
)

//implements the FSMSnapshot interface
type FSMSnapshot struct {
	Db *db.RavelDatabase
}

type KeyValue struct {
	Key   []byte `json:"key"`
	Value []byte `json:"value"`
}

func (f *FSMSnapshot) Persist(sink raft.SnapshotSink) error {

	log.Println("FSMSnapshot: Starting Snapshot")

	_, err := f.Db.Conn.Backup(sink, 0)
	if err != nil {
		log.Println("FSMSnapshot: Unable to take Snapshot")
		return err
	}

	err = sink.Close()
	if err != nil {
		log.Println("FSMSnapshot: Unable to close Sink")
		return err
	}

	return nil

	// c := make(chan KeyValue)

	// var kv KeyValue

	// //iterates over all the key in a seperate go routine and passes the values read into a channel
	// go f.db.Conn.View(func(txn *badger.Txn) error {

	// 	opts := badger.DefaultIteratorOptions
	// 	it := txn.NewIterator(opts)
	// 	defer it.Close()

	// 	for it.Rewind(); it.Valid(); it.Next() {
	// 		item := it.Item()
	// 		kv.Key = item.Key()

	// 		err := item.Value(func(v []byte) error {
	// 			kv.Value = v
	// 			return nil
	// 		})

	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		c <- kv

	// 	}
	// 	close(c)
	// 	return nil
	// })

	// var dataRead []KeyValue
	// var dr KeyValue
	// var ok bool

	// //reads the key values from the channel until it is closed and all the values have been read
	// for {
	// 	dr, ok = (<-c)
	// 	if !ok {
	// 		break
	// 	}

	// 	dataRead = append(dataRead, dr)
	// }

	// b, err := msgpack.Marshal(dataRead)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if _, err := sink.Write(b); err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("All keys have been persisted")

	// return nil
}

func (f *FSMSnapshot) Release() {
	log.Println("FSMSnapshot: Snapshot finised")
}
