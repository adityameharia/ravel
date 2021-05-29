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
)

type Node struct {
	Mu   sync.Mutex
	Fsm  *fsm.Fsm
	Raft *raft.Raft
}

func (n *Node) Open(enableSingle bool, localID string, badgerPath string, raftdir string, BindAddr string) error {

	var raftNode Node

	opts := badger.DefaultOptions(badgerPath)

	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal("Unable to open DB")
		return err
	}

	raftNode.Fsm.Db.Conn = db

	//setting up Raft Config
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(localID)

	// Setup Raft communication
	addr, err := net.ResolveTCPAddr("tcp", BindAddr)
	if err != nil {
		return err
	}
	transport, err := raft.NewTCPTransport(BindAddr, addr, 5, 2*time.Second, os.Stderr)
	if err != nil {
		return err
	}

	// Create the snapshot store. This allows the Raft to truncate the log.
	snapshot, err := raft.NewFileSnapshotStore(raftdir, 1, os.Stderr)
	if err != nil {
		log.Fatal("Unable to create SnapShot store")
		return err
	}

	//creating log and stable store
	var logStore raft.LogStore
	var stableStore raft.StableStore
	var f raft.FSM

	logStore, err = store.NewRavelLogStore(raftdir + "/logs")
	if err != nil {
		log.Fatal("Unable to create Log store")
		return err
	}

	stableStore, err = store.NewRavelStableStore(raftdir + "/stable")
	if err != nil {
		log.Fatal("Unable to create Stable store")
		return err
	}

	f, err = fsm.NewFSM(raftdir)
	if err != nil {
		log.Fatal("Unable to create FSM")
		return err
	}

	r, err := raft.NewRaft(config, f, logStore, stableStore, snapshot, transport)
	if err != nil {
		log.Fatal("Unable initialise raft node")
		return err
	}

	raftNode.Raft = r

	return nil

}

func (n *Node) Get(key string) (string, error) {
	return n.Fsm.Get(key)
}

func (n *Node) Set(key string, value string) error {
	return nil
}
