package node

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"time"

	"github.com/adityameharia/ravel/fsm"
	"github.com/adityameharia/ravel/store"
	"github.com/hashicorp/raft"
)

// RavelNode represents a node inside the cluster.
type RavelNode struct {
	Fsm  *fsm.RavelFSM
	Raft *raft.Raft
}

// Open creates initialises a raft.Raft instance
func (n *RavelNode) Open(enableSingle bool, localID string, badgerPath string, BindAddr string) (*raft.Raft, *fsm.RavelFSM, error) {
	log.Println(enableSingle)
	log.Println("RavelNode: Opening node")
	var raftNode RavelNode

	//setting up Raft Config
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(localID)
	config.SnapshotThreshold = 5000

	log.Println(config)

	// Setup Raft communication
	addr, err := net.ResolveTCPAddr("tcp", BindAddr)
	if err != nil {
		log.Fatal("RavelNode: Unable to resolve TCP Bind Address")
		return nil, nil, err
	}
	transport, err := raft.NewTCPTransport(BindAddr, addr, 5, 2*time.Second, os.Stderr)
	if err != nil {
		log.Fatal("RavelNode: Unable to create NewTCPTransport")
		return nil, nil, err
	}

	// Create the snapshot store. This allows the Raft to truncate the log.
	snapshot, err := raft.NewFileSnapshotStore(badgerPath+"/snapshot", 1, os.Stderr)
	if err != nil {
		log.Fatal("RavelNode: Unable to create SnapShot store")
		return nil, nil, err
	}

	//creating log and stable store
	var logStore raft.LogStore
	var stableStore raft.StableStore

	logStore, err = store.NewRavelLogStore(badgerPath + "/logs")
	if err != nil {
		log.Fatal("RavelNode: Unable to create Log store")
		return nil, nil, err
	}

	f, err := fsm.NewFSM(badgerPath + "/fsm")
	if err != nil {
		log.Fatal("RavelNode: Unable to create FSM")
		return nil, nil, err
	}

	stableStore, err = store.NewRavelStableStore(badgerPath + "/stable")
	if err != nil {
		log.Fatal("RavelNode: Unable to create Stable store")
		return nil, nil, err
	}

	raftNode.Fsm = f

	r, err := raft.NewRaft(config, f, logStore, stableStore, snapshot, transport)
	if err != nil {
		log.Println(err)
		log.Fatal("RavelNode: Unable initialise raft node")

		return nil, nil, err
	}

	raftNode.Raft = r
	if enableSingle {
		configuration := raft.Configuration{
			Servers: []raft.Server{
				{
					ID:      config.LocalID,
					Address: transport.LocalAddr(),
				},
			},
		}
		r.BootstrapCluster(configuration)
	}

	return r, f, nil
}

// Get returns the value for the given key
func (n *RavelNode) Get(key []byte) (string, error) {
	if n.Raft.State() != raft.Leader {
		log.Println("RavelNode: Request sent to non leading replica")
		return "", raft.ErrNotLeader
	}
	return n.Fsm.Get(key)
}

// Set sets the key with the value
func (n *RavelNode) Set(key []byte, value []byte) error {
	if n.Raft.State() != raft.Leader {
		log.Println("RavelNode: Request sent to non leading replica")
		return raft.ErrNotLeader
	}

	var data fsm.LogData

	data.Operation = "set"
	data.Key = key
	data.Value = value

	dataBuffer, err := json.Marshal(data)
	if err != nil {
		log.Fatal("RavelNode: Unable to marhsal key value")
		return err
	}

	f := n.Raft.Apply(dataBuffer, 3*time.Second)

	return f.Error()
}

// Delete deletes the entry with given key
func (n *RavelNode) Delete(key []byte) error {
	if n.Raft.State() != raft.Leader {
		log.Println("RavelNode: Request sent to non leading replica")
		return raft.ErrNotLeader
	}

	var data fsm.LogData

	data.Operation = "delete"
	data.Key = key
	data.Value = []byte{}

	dataBuffer, err := json.Marshal(data)
	if err != nil {
		log.Fatal("RavelNode: Unable to marhsal key value")
		return err
	}

	f := n.Raft.Apply(dataBuffer, 3*time.Second)

	return f.Error()
}
