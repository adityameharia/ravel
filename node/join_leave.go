package node

import (
	"errors"
	"log"

	"github.com/hashicorp/raft"
)

type Response struct {
	Error  string `json:"error"`
	Leader string `json:"leader"`
}

func (n *RavelNode) Join(nodeID, addr string) (string, error) {

	log.Printf("received join request for remote node %s, addr %s\n", nodeID, addr)
	if n.Raft.State() != raft.Leader {
		log.Println(string(n.Raft.Leader()))
		return string(n.Raft.Leader()), errors.New("node is not leader")
	}
	log.Println("1")
	config := n.Raft.GetConfiguration()
	log.Println("2")
	if err := config.Error(); err != nil {
		log.Printf("failed to get raft configuration\n")
		return "", err
	}
	log.Println("3")
	for _, server := range config.Configuration().Servers {
		if server.ID == raft.ServerID(nodeID) {
			log.Printf("node %s already joined raft cluster\n", nodeID)
			return "", errors.New("node already exists")
		}
	}
	log.Println("4")
	f := n.Raft.AddVoter(raft.ServerID(nodeID), raft.ServerAddress(addr), 0, 0)
	if err := f.Error(); err != nil {
		return "", err
	}
	log.Println("5")
	log.Printf("node %s at %s joined successfully\n", nodeID, addr)
	return "", nil
}

func (n *RavelNode) Leave(nodeID string) (string, error) {
	log.Printf("received leave request for remote node %s", nodeID)
	if n.Raft.State() != raft.Leader {
		return string(n.Raft.Leader()), errors.New("node is not leader")
	}

	config := n.Raft.GetConfiguration()

	if err := config.Error(); err != nil {
		log.Printf("failed to get raft configuration\n")
		return "", err
	}

	for _, server := range config.Configuration().Servers {
		if server.ID == raft.ServerID(nodeID) {
			f := n.Raft.RemoveServer(server.ID, 0, 0)
			if err := f.Error(); err != nil {
				log.Printf("failed to remove server %s\n", nodeID)
				return "", err
			}

			log.Printf("node %s left successfully\n", nodeID)
			return "", nil
		}
	}

	log.Printf("node %s does not exist in raft group\n", nodeID)
	return "", errors.New("node doesnt exists in the cluster")
}
