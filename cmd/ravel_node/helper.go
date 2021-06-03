package main

import (
	"context"
	"github.com/adityameharia/ravel/RavelNodePB"
	"google.golang.org/grpc"
	"log"
)

func RequestJoinToClusterLeader(leaderGRPCAddr string, node *RavelNodePB.Node) error {
	conn, err := grpc.Dial(leaderGRPCAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
		return err
	}

	client := RavelNodePB.NewRavelNodeClient(conn)
	_, err = client.Join(context.Background(), node)

	if err != nil && err.Error() == "rpc error: code = Unknown desc = node already exists" {
		return nil
	} else if err != nil {
		log.Println(err.Error())
		log.Fatalf("join request falied with server %v", err)
		return err
	}

	return nil
}

func RequestLeaveToClusterLeader(leaderGRPCAddr string, node *RavelNodePB.Node) error {
	conn, err := grpc.Dial(leaderGRPCAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
		return err
	}
	client := RavelNodePB.NewRavelNodeClient(conn)

	_, err = client.Leave(context.Background(), node)
	if err != nil {
		log.Fatalf("join request falied with server %v", err)
		return err
	}

	return nil
}
