package node

import (
	"log"
	"net"
	"os"
	"sync"
	"time"

	"github.com/adityameharia/ravel/fsm"
	"github.com/adityameharia/ravel/store"
	"github.com/dgraph-io/badger/v3"
	"github.com/hashicorp/raft"
	"github.com/vmihailenco/msgpack/v5"
)

type Node struct {
	Mu   sync.Mutex
	Fsm  *fsm.Fsm
	Raft *raft.Raft
}

func (n *Node) Open(enableSingle bool, localID string, badgerPath string, raftdir string, BindAddr string) error {

	log.Println("Node: Opening node")

	var raftNode Node

	opts := badger.DefaultOptions(badgerPath)

	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal("Node: Unable to open DB")
		return err
	}

	raftNode.Fsm.Db.Conn = db

	//setting up Raft Config
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(localID)
	config.SnapshotThreshold = 5000

	// Setup Raft communication
	addr, err := net.ResolveTCPAddr("tcp", BindAddr)
	if err != nil {
		log.Fatal("Node: Unable to resolve TCP Bind Address")
		return err
	}
	transport, err := raft.NewTCPTransport(BindAddr, addr, 5, 2*time.Second, os.Stderr)
	if err != nil {
		log.Fatal("Node: Unable to create NewTCPTransport")
		return err
	}

	// Create the snapshot store. This allows the Raft to truncate the log.
	snapshot, err := raft.NewFileSnapshotStore(raftdir, 1, os.Stderr)
	if err != nil {
		log.Fatal("Node: Unable to create SnapShot store")
		return err
	}

	//creating log and stable store
	var logStore raft.LogStore
	var stableStore raft.StableStore
	var f raft.FSM

	logStore, err = store.NewRavelLogStore(raftdir + "/logs")
	if err != nil {
		log.Fatal("Node: Unable to create Log store")
		return err
	}

	stableStore, err = store.NewRavelStableStore(raftdir + "/stable")
	if err != nil {
		log.Fatal("Node: Unable to create Stable store")
		return err
	}

	f, err = fsm.NewFSM(raftdir)
	if err != nil {
		log.Fatal("Node: Unable to create FSM")
		return err
	}

	r, err := raft.NewRaft(config, f, logStore, stableStore, snapshot, transport)
	if err != nil {
		log.Fatal("Node: Unable initialise raft node")
		return err
	}

	raftNode.Raft = r

	return nil

}

func (n *Node) Get(key string) (string, error) {
	if n.Raft.State() != raft.Leader {
		log.Println("Node: Request sent to non leading replica")
		return "", raft.ErrNotLeader
	}
	return n.Fsm.Get(key)
}

func (n *Node) Set(key string, value string) error {
	if n.Raft.State() != raft.Leader {
		log.Println("Node: Request sent to non leading replica")
		return raft.ErrNotLeader
	}

	var data fsm.LogData

	data.Operation = "set"
	data.Key = key
	data.Value = value

	dataBuffer, err := msgpack.Marshal(data)
	if err != nil {
		log.Fatal("Node: Unable to marhsal key value")
		return err
	}

	f := n.Raft.Apply(dataBuffer, 3*time.Second)

	return f.Error()
}

func (n *Node) Delete(key string) error {
	if n.Raft.State() != raft.Leader {
		log.Println("Node: Request sent to non leading replica")
		return raft.ErrNotLeader
	}

	var data fsm.LogData

	data.Operation = "delete"
	data.Key = key
	data.Value = ""

	dataBuffer, err := msgpack.Marshal(data)
	if err != nil {
		log.Fatal("Node: Unable to marhsal key value")
		return err
	}

	f := n.Raft.Apply(dataBuffer, 3*time.Second)

	return f.Error()
}
