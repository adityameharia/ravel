package main

import (
	"context"
	"flag"
	"github.com/adityameharia/ravel/RavelNodePB"
	"github.com/adityameharia/ravel/node"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/adityameharia/ravel/RavelClusterAdminPB"
	"github.com/adityameharia/ravel/node_server"
	"github.com/google/uuid"
)

type Config struct {
	clusterID        string
	nodeID           string
	storageDir       string
	gRPCAddr         string
	raftInternalAddr string
	adminGRPCAddr    string
	isLeader         bool
}

var nodeConfig Config
var adminClient RavelClusterAdminPB.RavelClusterAdminClient

func init() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	nodeConfig.nodeID = uuid.New().String()
	flag.StringVar(&nodeConfig.storageDir, "storageDir", dirname, "Data Directory for Raft")
	flag.StringVar(&nodeConfig.gRPCAddr, "gRPCAddr", "", "Address (with port) at which gRPC server is started")
	flag.StringVar(&nodeConfig.raftInternalAddr, "raftAddr", "", "Raft internal communication address with port")
	flag.StringVar(&nodeConfig.adminGRPCAddr, "adminRPCAddr", "", "GRPC address of the cluster admin")
	flag.BoolVar(&nodeConfig.isLeader, "leader", false, "Register this node as a new leader or not")
}

//func main() {
//	flag.Parse()
//
//	//setup gRPC client for admin cluster
//	connAdmin, err := grpc.Dial(nodeConfig.admingRPCAddr, grpc.WithInsecure())
//	if err != nil {
//		log.Fatalf("can not connect with server %v", err)
//	}
//
//	adminClient = RavelClusterAdminPB.NewRavelClusterAdminClient(connAdmin)
//
//	replica := &RavelClusterAdminPB.Cluster{ClusterID: int32(nodeConfig.clusterID)}
//
//	//get address of leader node from admin cluster
//	leader, err := adminClient.GetLeader(context.Background(), replica)
//	if err != nil {
//		log.Fatal("Unable to get the leader node from the server")
//	}
//
//	var newNode node.RavelNode
//
//	newNode.Raft, newNode.Fsm, err = newNode.Open(leader.Data == "", nodeConfig.nodeID, nodeConfig.storageDir, nodeConfig.raftInternalAddr)
//	if err != nil {
//		log.Println(err)
//	}
//
//	if leader.Data != "" {
//		if err := server.RequestJoinToLeader(nodeConfig.nodeID, leader.Data, nodeConfig.raftInternalAddr); err != nil {
//			log.Fatalf("failed to join node at %s: %s", nodeConfig.joinAddr, err.Error())
//		}
//	}
//
//	onSignalInterrupt()
//
//	listener, err := net.Listen("tcp", nodeConfig.gRPCAddr)
//	if err != nil {
//		log.Fatalf("Error %v", err)
//	}
//	log.Printf("Starting TCP Server on %v for gRPC\n", nodeConfig.gRPCAddr)
//
//	grpcServer := grpc.NewServer()
//	RavelClusterPB.RegisterRavelClusterServer(grpcServer, &server.Server{Node: &newNode})
//	err = grpcServer.Serve(listener)
//}


func main() {
	flag.Parse()

	adminConn, err := grpc.Dial(nodeConfig.adminGRPCAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error in connecting to the Admin gRPC Server: ", err)
	}

	var ravelNode node.RavelNode
	adminClient = RavelClusterAdminPB.NewRavelClusterAdminClient(adminConn)

	if nodeConfig.isLeader {
		ravelCluster, err := adminClient.JoinAsClusterLeader(context.TODO(), &RavelClusterAdminPB.Node{
			NodeId: nodeConfig.nodeID, // id of this node
			GrpcAddress: nodeConfig.gRPCAddr, // grpc address of this node
			ClusterId: "", // cluster id is unknown thus empty
		})

		if err != nil {
			log.Fatal("Error in JoinAsClusterLeader: ", err)
		} else {
			nodeConfig.clusterID = ravelCluster.ClusterId
			ravelNode.Raft, ravelNode.Fsm, err = ravelNode.Open(nodeConfig.isLeader, nodeConfig.nodeID, nodeConfig.storageDir, nodeConfig.raftInternalAddr)
			if err != nil {
				log.Println(err)
			}

			// this node is the leader
		}
	} else {
		ravelCluster, err := adminClient.JoinExistingCluster(context.TODO(), &RavelClusterAdminPB.Node{
			NodeId: nodeConfig.nodeID,
			GrpcAddress: nodeConfig.gRPCAddr,
			ClusterId: "",
		})
		if err != nil {
			log.Fatal("Error in JoinExistingCluster: ", err)
		} else {
			log.Println("Cluster leader is: ", ravelCluster.LeaderGrpcAddress)
			nodeConfig.clusterID = ravelCluster.ClusterId
			ravelNode.Raft, ravelNode.Fsm, err = ravelNode.Open(nodeConfig.isLeader, nodeConfig.nodeID, nodeConfig.storageDir, nodeConfig.raftInternalAddr)
			if err != nil {
				log.Println(err)
			}

			log.Println("here")
			err = RequestJoinToClusterLeader(ravelCluster.LeaderGrpcAddress, &RavelNodePB.Node{
				NodeId: nodeConfig.nodeID,
				ClusterId: nodeConfig.clusterID,
				GrpcAddress: nodeConfig.gRPCAddr,
			})
			if err != nil {
				log.Println(err)
			}
		}
	}

	onSignalInterrupt()
	listener, err := net.Listen("tcp", nodeConfig.gRPCAddr)
	if err != nil {
		log.Fatal("Error in starting TCP server: ", err)
	}
	log.Printf("Starting TCP Server on %v for gRPC\n", nodeConfig.gRPCAddr)

	grpcServer := grpc.NewServer()
	RavelNodePB.RegisterRavelNodeServer(grpcServer, &node_server.Server{
		Node: &ravelNode,
	})
	err = grpcServer.Serve(listener)
}

func onSignalInterrupt() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	go func() {
		<-ch
		cluster := &RavelClusterAdminPB.Cluster{ClusterId: nodeConfig.clusterID}
		leaderNode, err := adminClient.GetClusterLeader(context.Background(), cluster)
		if err != nil {
			log.Fatal(err)
		}

		err = RequestLeaveToClusterLeader(leaderNode.GrpcAddress, &RavelNodePB.Node{
			NodeId: nodeConfig.nodeID,
			ClusterId: nodeConfig.clusterID,
			GrpcAddress: nodeConfig.gRPCAddr,
		})

		if err != nil {
			log.Println(err)
		}

		os.Exit(1)
	}()
}
