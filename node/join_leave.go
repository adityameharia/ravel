package node

import (
	"encoding/json"
	"errors"
	"github.com/hashicorp/raft"
	"log"
)

func (n *RavelNode) Join(nodeID, addr string) error {
	log.Printf("received join request for remote node %s, addr %s\n", nodeID, addr)
	config := n.Raft.GetConfiguration()

	if err := config.Error(); err != nil {
		log.Printf("failed to get raft configuration\n")
		return err
	}

	if n.Raft.State() != raft.Leader {
		type response struct {
			Error string `json:"error"`
			Leader string `json:"leader"`
		}

		respBytes, err := json.Marshal(&response{
			Error: "node is not leader",
			Leader: string(n.Raft.Leader()),
		})

		if err != nil {
			log.Println("RavelNode.Join: Error converting response to JSON")
			return err
		}

		return errors.New(string(respBytes))
	}

	for _, server := range config.Configuration().Servers {
		if server.ID == raft.ServerID(nodeID) {
			log.Printf("node %s already joined raft cluster\n", nodeID)
			return nil
		}
	}

	f := n.Raft.AddVoter(raft.ServerID(nodeID), raft.ServerAddress(addr), 0, 0)
	if err := f.Error(); err != nil {
		return err
	}

	log.Printf("node %s at %s joined successfully\n", nodeID, addr)
	return nil
}

func (n *RavelNode) Leave(nodeID string) error {
	log.Printf("received leave request for remote node %s", nodeID)

	config := n.Raft.GetConfiguration()

	if err := config.Error(); err != nil {
		log.Printf("failed to get raft configuration")
		return err
	}

	if n.Raft.State() != raft.Leader {
		type response struct {
			Error string `json:"error"`
			Leader string `json:"leader"`
		}

		respBytes, err := json.Marshal(&response{
			Error: "node is not leader",
			Leader: string(n.Raft.Leader()),
		})

		if err != nil {
			log.Println("RavelNode.Join: Error converting response to JSON")
			return err
		}

		return errors.New(string(respBytes))
	}

	for _, server := range config.Configuration().Servers {
		if server.ID == raft.ServerID(nodeID) {
			f := n.Raft.RemoveServer(server.ID, 0, 0)
			if err := f.Error(); err != nil {
				log.Printf("failed to remove server %s\n", nodeID)
				return err
			}

			log.Printf("node %s leaved successfully\n", nodeID)
			return nil
		}
	}

	log.Printf("node %s not exists in raft group\n", nodeID)
	return nil
}
