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

func killCluster(replicaCount int) {
	if replicaCount == 2 || replicaCount == -1 {

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
		leaderNode, err := adminClient.GetClusterLeader(context.Background(), cluster)
		if err != nil {
			log.Fatal(err)
		}

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

		log.Println("Server successfully removed")

		os.Exit(1)
	}
}
