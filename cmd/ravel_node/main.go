package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/adityameharia/ravel/RavelClusterAdminPB"
	"github.com/adityameharia/ravel/RavelClusterPB"
	"github.com/adityameharia/ravel/node"
	"github.com/adityameharia/ravel/server"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type Config struct {
	storageDir       string
	gRPCAddr         string
	nodeID           string
	joinAddr         string
	raftInternalAddr string
	admingRPCAddr    string
	clusterID        int
}

var nodeConfig Config

var adminClient RavelClusterAdminPB.RavelClusterAdminClient

func init() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	id := uuid.New().String()

	flag.StringVar(&nodeConfig.storageDir, "storageDir", dirname, "Data Directory for Raft")
	flag.IntVar(&nodeConfig.clusterID, "clusterID", -1, "Unique id of the cluster the replica is supposed to join")
	flag.StringVar(&nodeConfig.gRPCAddr, "gRPCAddr", "", "Address (with port) at which gRPC server is started")
	flag.StringVar(&nodeConfig.nodeID, "nodeID", id, "Unique ID for the Node")
	flag.StringVar(&nodeConfig.joinAddr, "joinAddr", "", "Address of the leader node to which this node is supposed to join")
	flag.StringVar(&nodeConfig.raftInternalAddr, "raftAddr", "", "Raft internal communication address with port")
	flag.StringVar(&nodeConfig.admingRPCAddr, "admingRPCAddr", "", "GRPC address of the cluster admin")
}

func main() {
	flag.Parse()

	//setup gRPC client for admin cluster
	connAdmin, err := grpc.Dial(nodeConfig.admingRPCAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	adminClient = RavelClusterAdminPB.NewRavelClusterAdminClient(connAdmin)

	replica := &RavelClusterAdminPB.Cluster{ClusterID: int32(nodeConfig.clusterID)}

	//get address of leader node from admin cluster
	leader, err := adminClient.GetLeader(context.Background(), replica)
	if err != nil {
		log.Fatal("Unable to get the leader node from the server")
	}

	var newNode node.RavelNode

	newNode.Raft, newNode.Fsm, err = newNode.Open(leader.Data == "", nodeConfig.nodeID, nodeConfig.storageDir, nodeConfig.raftInternalAddr)
	if err != nil {
		log.Println(err)
	}

	if leader.Data != "" {
		if err := server.RequestJoinToLeader(nodeConfig.nodeID, leader.Data, nodeConfig.raftInternalAddr); err != nil {
			log.Fatalf("failed to join node at %s: %s", nodeConfig.joinAddr, err.Error())
		}
	}

	onSignalInterrupt()

	listener, err := net.Listen("tcp", nodeConfig.gRPCAddr)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	log.Printf("Starting TCP Server on %v for gRPC\n", nodeConfig.gRPCAddr)

	grpcServer := grpc.NewServer()
	RavelClusterPB.RegisterRavelClusterServer(grpcServer, &server.Server{Node: &newNode})
	err = grpcServer.Serve(listener)
}

func onSignalInterrupt() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	go func() {
		<-ch
		replica := &RavelClusterAdminPB.Cluster{ClusterID: int32(nodeConfig.clusterID)}
		l, err := adminClient.GetLeader(context.Background(), replica)
		if err != nil {
			log.Fatal("Unable to get the leader node from the server")
		}

		err = server.RequestLeaveToLeader(nodeConfig.nodeID, l.Data)
		if err != nil {
			log.Println(err)
		}

		os.Exit(1)
	}()
}
