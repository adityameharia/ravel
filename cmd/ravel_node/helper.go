package main

import (
	"context"
	"log"

	"github.com/adityameharia/ravel/RavelClusterAdminPB"
	"github.com/adityameharia/ravel/RavelNodePB"
	"google.golang.org/grpc"
)

// RequestJoinToClusterLeader makes a new gRPC client and sends a join request to the leading replica in the cluster
func RequestJoinToClusterLeader(leaderGRPCAddr string, node *RavelNodePB.Node) error {
	conn, err := grpc.Dial(leaderGRPCAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
		return err
	}

	defer conn.Close()

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

// RequestLeaveToClusterLeader makes a new gRPC client and sends a leave request to the leading replica in the cluster
func RequestLeaveToClusterLeader(leaderGRPCAddr string, node *RavelNodePB.Node) error {
	conn, err := grpc.Dial(leaderGRPCAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error in RequestLeaveToClusterLeader: %v", err)
		return err
	}
	defer conn.Close()
	client := RavelNodePB.NewRavelNodeClient(conn)

	_, err = client.Leave(context.Background(), node)
	if err != nil {
		log.Printf("leave request failed: %v", err)
		return err
	}

	return nil
}

// RequestLeaderUpdateToCluster makes a new gRPC client and
// updates the admin server in case there is a change in leader in its cluster
func RequestLeaderUpdateToCluster(clusterAdminGRPCAddr string, node *RavelClusterAdminPB.Node) error {
	conn, err := grpc.Dial(clusterAdminGRPCAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error in RequestLeaderUpdateToCluster: %v", err)
		return err
	}
	defer conn.Close()
	client := RavelClusterAdminPB.NewRavelClusterAdminClient(conn)
	resp, err := client.UpdateClusterLeader(context.TODO(), node)
	if err != nil {
		log.Fatalf("Error in RequestLeaderUpdateToCluster: %v", err)
		return err
	}

	log.Println(resp.Data)
	return nil
}
