package store

import (
	"encoding/binary"
	"encoding/json"
	"log"

	"github.com/hashicorp/raft"
)

// Converts bytes to an integer
func bytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

// Converts a uint to a byte slice
func uint64ToBytes(u uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, u)
	return buf
}

// raftLogToBytes converts a raft.Log object to []bytes using msgpack serialization
func raftLogToBytes(l raft.Log) []byte {
	bytes, err := json.Marshal(l)
	if err != nil {
		log.Println(err)
	}

	return bytes
}

// bytesToRaftLog converts []byte to a raft.Log object using msgpack serialization
// and writes it on that pointer
func bytesToRaftLog(b []byte, raftLog *raft.Log) error {
	err := json.Unmarshal(b, raftLog)
	if err != nil {
		return err
	}

	return nil
}
