package server

import (
	"context"
	"log"

	"github.com/adityameharia/ravel/RavelClusterPB"
	"google.golang.org/grpc"
)

func requestJoin(nodeID, joinAddr, raftAddr string) error {
	conn, err := grpc.Dial(joinAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	client := RavelClusterPB.NewRavelClusterClient(conn)

	node := &RavelClusterPB.Node{
		NodeID:  nodeID,
		Address: raftAddr,
	}

	res, err := client.Join(context.Background(), node)
	if res.Data=="node is not leader"

}
