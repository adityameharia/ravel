package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/adityameharia/ravel/RavelClusterAdminPB"
	"github.com/adityameharia/ravel/RavelNodePB"
	"google.golang.org/grpc"
)

func killCluster() {

	config, err := conf.Read([]byte("config"))
	if err != nil {
		log.Fatal("Error reading config details from file")
	}

	err = json.Unmarshal(config, &nodeConfig)
	if err != nil {
		log.Fatal("Error reading config details from file")
	}

	adminConn, err := grpc.Dial(nodeConfig.AdminGRPCAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error in connecting to the Admin gRPC Server: ", err)
	}
	defer adminConn.Close()

	adminClient = RavelClusterAdminPB.NewRavelClusterAdminClient(adminConn)

	log.Println(nodeConfig)

	cluster := &RavelClusterAdminPB.Cluster{ClusterId: nodeConfig.ClusterID}
	log.Println("1")
	leaderNode, err := adminClient.GetClusterLeader(context.Background(), cluster)
	log.Println("2")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("3")
	err = RequestLeaveToClusterLeader(leaderNode.GrpcAddress, &RavelNodePB.Node{
		NodeId:      nodeConfig.NodeID,
		ClusterId:   nodeConfig.ClusterID,
		GrpcAddress: nodeConfig.GRPCAddr,
	})

	if err != nil {
		log.Println(err)
	}

	_, err = adminClient.LeaveCluster(context.TODO(), &RavelClusterAdminPB.Node{
		NodeId:      nodeConfig.NodeID,
		ClusterId:   nodeConfig.ClusterID,
		GrpcAddress: nodeConfig.GRPCAddr,
		RaftAddress: nodeConfig.RaftInternalAddr,
	})

	if err != nil {
		log.Println(err)
	}

	err = os.RemoveAll(nodeConfig.StorageDir)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(1)
}
