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
	resp, err := client.Join(context.Background(), node)

	log.Println("test1")
	log.Println(err)
	log.Println(resp)
	log.Println("test2")
	if err != nil && err.Error() == "rpc error: code = Unknown desc = node already exists" {
		return nil
	} else if err != nil && err.Error() == "rpc error: code = Unknown desc = node is not leader" {
		log.Println("11")
		var e error
		conn, e = grpc.Dial(resp.Data, grpc.WithInsecure())
		if e != nil {
			log.Fatalf("can not connect with server %v", e)
		}
		log.Println("12")
		client := RavelClusterPB.NewRavelClusterClient(conn)
		node := &RavelClusterPB.Node{
			NodeID:  nodeID,
			Address: raftAddr,
		}
		log.Println("13")
		_, e = client.Join(context.Background(), node)
		if e != nil {
			log.Println(e)
			return err
		}
		log.Println("14")
	} else if err != nil {
		log.Println("15")
		log.Println(err.Error())
		log.Fatalf("join request falied with server %v", err)
		return err
	}
	log.Println("16")
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
	if err != nil && err.Error() == "rpc error: code = Unknown desc = node is not leader" {
		conn, err := grpc.Dial(res.Data, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("can not connect with server %v", err)
		}
		client := RavelClusterPB.NewRavelClusterClient(conn)

		node := &RavelClusterPB.Node{
			NodeID:  nodeID,
			Address: "",
		}
		_, err = client.Leave(context.Background(), node)
		if err != nil {
			log.Fatalf("leave request falied with server %v", err)
			return err
		}
	} else if err != nil {
		log.Fatalf("leave request falied with server %v", err)
		return err
	}

	return nil
}
