package fsm

import (
	"testing"
	"time"

	"github.com/hashicorp/raft"
	"github.com/vmihailenco/msgpack/v5"
)

func TestSnapshot(t *testing.T) {
	f, err := NewFSM("/tmp/badger/test/fsm")
	if err != nil {
		t.Error("Error in newFSM")
	}

	snapshot, err := f.Snapshot()
	if err != nil {
		t.Error("Error in snapshot")
	}
	t.Log(snapshot)

}

func TestApplyAndGet(t *testing.T) {
	f, err := NewFSM("/tmp/badger/test/fsm")
	if err != nil {
		t.Error("Error in newFSM")
	}

	com := &LogData{
		Operation: "set",
		Key:       []byte("testKey"),
		Value:     []byte("testValue"),
	}
	comByte, err := msgpack.Marshal(com)
	if err != nil {
		t.Error("Error marshalling logData")
	}

	l := raft.Log{
		Index:      1,
		Term:       0,
		Type:       raft.LogCommand,
		Data:       comByte,
		AppendedAt: time.Now(),
	}

	e := f.Apply(&l)
	if e != nil {
		t.Error("Error in apply fsm")
	}

	_, err = f.Get([]byte("testKey"))
	if e != nil {
		t.Error("Error in getting key which has been set")
	}
}

//not sure how to implement the restore test
// func TestRestore(t *testing.T){
// 	f, err := NewFSM("/tmp/badger/test/fsm")
// 	if err != nil {
// 		t.Error("Error in newFSM")
// 	}

// 	err=f.Restore()
// }
