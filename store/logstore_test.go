package store

import (
	"github.com/hashicorp/raft"
	"log"
	"testing"
	"time"
)


func TestRavelLogStore_StoreLog(t *testing.T) {
	r, err := NewRavelLogStore("/tmp/badger/test")
	if err != nil {
		log.Println(err)
	}

	var logs []*raft.Log
	var i uint64
	for i = 0; i < 5; i++ {
		l := raft.Log{
			Index: i,
			Term: 0,
			Type: raft.LogCommand,
			Data: []byte("Test Log Data"),
			AppendedAt: time.Now(),
		}

		logs = append(logs, &l)
	}

	err = r.StoreLogs(logs)
	if err != nil {
		t.Error("Error in StoreLog", err)
	}

	var l raft.Log
	err = r.GetLog(2, &l)
	if err != nil {
		t.Error("Error in GetLog", err)
	}

	if l.Index != 2 {
		t.Error("Error in GetLog, expected l.Index to be 2 got ", l.Index)
	}

	var fi uint64
	fi, err = r.FirstIndex()
	if err != nil {
		t.Error("Error in FirstIndex", err)
	}

	if fi != 0 {
		t.Error("Error in FirstIndex, expected 0 got ", fi)
	}

	var li uint64
	li, err = r.LastIndex()
	if err != nil {
		t.Error("Error in LastIndex", err)
	}

	if li != 4 {
		t.Error("Error in LastIndex, expected 4 got ", li)
	}
}
