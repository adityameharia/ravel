package server

import (
	"context"
	"log"

	"github.com/adityameharia/ravel/RavelClusterPB"
	"google.golang.org/grpc"
)

func RequestJoin(nodeID, joinAddr, raftAddr string) error {

	conn, err := grpc.Dial(joinAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
		return err
	}

	client := RavelClusterPB.NewRavelClusterClient(conn)

	node := &RavelClusterPB.Node{
		NodeID:  nodeID,
		Address: raftAddr,
	}

	log.Println(node)
	res, err := client.Join(context.Background(), node)
	if err != nil && err.Error() == "rpc error: code = Unknown desc = node already exists" {
		return nil
	}
	if err != nil {

		log.Println(err.Error())
		log.Fatalf("join request falied with server %v", err)
		return err
	}
	if res.Data != "" {
		conn, err := grpc.Dial(res.Data, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("can not connect with server %v", err)
		}
		client := RavelClusterPB.NewRavelClusterClient(conn)

		node := &RavelClusterPB.Node{
			NodeID:  nodeID,
			Address: raftAddr,
		}
		_, err = client.Join(context.Background(), node)
		if err != nil {
			log.Fatalf("join request falied with server %v", err)
			return err
		}
	}
	return nil

}

func RequestLeave(nodeID, requestAddr string) error {
	conn, err := grpc.Dial(requestAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
		return err
	}
	client := RavelClusterPB.NewRavelClusterClient(conn)

	node := &RavelClusterPB.Node{
		NodeID:  nodeID,
		Address: "",
	}

	res, err := client.Leave(context.Background(), node)
	if err != nil {
		log.Fatalf("join request falied with server %v", err)
		return err
	}
	if res.Data != "" {
		conn, err := grpc.Dial(res.Data, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("can not connect with server %v", err)
		}
		client := RavelClusterPB.NewRavelClusterClient(conn)

		node := &RavelClusterPB.Node{
			NodeID:  nodeID,
			Address: "",
		}
		_, err = client.Join(context.Background(), node)
		if err != nil {
			log.Fatalf("join request falied with server %v", err)
			return err
		}
	}
	return nil
}
