package main

import (
	"flag"
	"github.com/adityameharia/ravel/RavelClusterPB"
	"github.com/adityameharia/ravel/node"
	"github.com/adityameharia/ravel/server"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

type Config struct {
	storageDir       string
	gRPCAddr         string
	nodeID           string
	joinAddr         string
	raftInternalAddr string
}

var nodeConfig Config

func init() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	flag.StringVar(&nodeConfig.storageDir, "storageDir", dirname, "Data Directory for Raft")
	flag.StringVar(&nodeConfig.gRPCAddr, "gRPCAddr", "", "Address (with port) at which gRPC server is started")
	flag.StringVar(&nodeConfig.nodeID, "nodeID", "", "Unique ID for the Node")
	flag.StringVar(&nodeConfig.joinAddr, "joinAddr", "", "Address of the leader node to which this node is supposed to join")
	flag.StringVar(&nodeConfig.raftInternalAddr, "raftAddr", "", "Raft internal communication address with port")
}

func main() {
	flag.Parse()

	var err error
	var newNode node.RavelNode
	newNode.Raft, newNode.Fsm, err = newNode.Open(nodeConfig.joinAddr == "", nodeConfig.nodeID, nodeConfig.storageDir, nodeConfig.raftInternalAddr)
	if err != nil {
		log.Println(err)
	}

	if nodeConfig.joinAddr != "" {
		if err := server.RequestJoin(nodeConfig.nodeID, nodeConfig.joinAddr, nodeConfig.raftInternalAddr); err != nil {
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
		var temp string
		if nodeConfig.joinAddr == "" {
			temp = nodeConfig.gRPCAddr
		} else {
			temp = nodeConfig.joinAddr
		}

		err := server.RequestLeave(nodeConfig.nodeID, temp)
		if err != nil {
			log.Println(err)
		}

		os.Exit(1)
	}()
}
