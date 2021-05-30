package node

import (
	"github.com/hashicorp/raft"
	"log"
)

type Response struct {
	Error  string `json:"error"`
	Leader string `json:"leader"`
}

func (n *RavelNode) Join(nodeID, addr string) Response {
	log.Printf("received join request for remote node %s, addr %s\n", nodeID, addr)

	if n.Raft.State() != raft.Leader {
		respBytes := Response{
			Error:  "node is not leader",
			Leader: string(n.Raft.Leader()),
		}

		return respBytes
	}

	config := n.Raft.GetConfiguration()

	if err := config.Error(); err != nil {
		log.Printf("failed to get raft configuration\n")
		return Response{Error: err.Error(), Leader: ""}
	}

	for _, server := range config.Configuration().Servers {
		if server.ID == raft.ServerID(nodeID) {
			log.Printf("node %s already joined raft cluster\n", nodeID)
			return Response{Error: "node already exists", Leader: ""}
		}
	}

	f := n.Raft.AddVoter(raft.ServerID(nodeID), raft.ServerAddress(addr), 0, 0)
	if err := f.Error(); err != nil {
		return Response{Error: err.Error(), Leader: ""}
	}

	log.Printf("node %s at %s joined successfully\n", nodeID, addr)
	return Response{Error: "", Leader: ""}
}

func (n *RavelNode) Leave(nodeID string) Response {
	log.Printf("received leave request for remote node %s", nodeID)
	if n.Raft.State() != raft.Leader {
		respBytes := Response{
			Error:  "node is not leader",
			Leader: string(n.Raft.Leader()),
		}

		return respBytes
	}

	config := n.Raft.GetConfiguration()

	if err := config.Error(); err != nil {
		log.Printf("failed to get raft configuration")
		return Response{Error: err.Error(), Leader: ""}
	}

	for _, server := range config.Configuration().Servers {
		if server.ID == raft.ServerID(nodeID) {
			f := n.Raft.RemoveServer(server.ID, 0, 0)
			if err := f.Error(); err != nil {
				log.Printf("failed to remove server %s\n", nodeID)
				return Response{Error: err.Error(), Leader: ""}
			}

			log.Printf("node %s left successfully\n", nodeID)
			return Response{Error: "", Leader: ""}
		}
	}

	log.Printf("node %s not exist in raft group\n", nodeID)
	return Response{Error: "", Leader: ""}
}
