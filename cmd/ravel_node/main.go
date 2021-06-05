package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/adityameharia/ravel/RavelNodePB"
	"github.com/adityameharia/ravel/node"
	"google.golang.org/grpc"

	"github.com/adityameharia/ravel/RavelClusterAdminPB"
	"github.com/adityameharia/ravel/node_server"
	"github.com/google/uuid"
)

// Config is the struct containing the configuration details of the node
type Config struct {
	// clusterID is ID of th cluster the node is a part of
	clusterID string
	// nodeID is the nodes unique ID
	nodeID string
	// storageDir is the Data Directory for Raft
	storageDir string
	// gRPCAddr is the Address (with port) at which gRPC server is started
	gRPCAddr string
	// raftInternalAddr is the Raft internal communication address with port
	raftInternalAddr string
	// adminGRPCAddr is the GRPC address of the cluster admin
	adminGRPCAddr string
	// isLeader is a bool defining whether the node is a leader or not
	isLeader bool
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
			NodeId:      nodeConfig.nodeID,   // id of this node
			GrpcAddress: nodeConfig.gRPCAddr, // grpc address of this node
			RaftAddress: nodeConfig.raftInternalAddr,
			ClusterId:   "", // cluster id is unknown thus empty
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
			NodeId:      nodeConfig.nodeID,
			GrpcAddress: nodeConfig.gRPCAddr,
			RaftAddress: nodeConfig.raftInternalAddr,
			ClusterId:   "",
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
				NodeId:      nodeConfig.nodeID,
				ClusterId:   nodeConfig.clusterID,
				GrpcAddress: nodeConfig.gRPCAddr,
				RaftAddress: nodeConfig.raftInternalAddr,
			})
			if err != nil {
				log.Println(err)
			}
		}
	}

	//updates the admin in case there is a change in leader
	go func() {
		leaderChange := <-ravelNode.Raft.LeaderCh()
		log.Println("Sending leader change req")
		if leaderChange {
			err := RequestLeaderUpdateToCluster(nodeConfig.adminGRPCAddr, &RavelClusterAdminPB.Node{
				NodeId:      nodeConfig.nodeID,
				GrpcAddress: nodeConfig.gRPCAddr,
				RaftAddress: nodeConfig.raftInternalAddr,
				ClusterId:   nodeConfig.clusterID,
			})

			if err != nil {
				log.Println(err)
			}
		}
	}()

	//sends leave request to leader and admin on signal interrupt
	onSignalInterrupt()

	//starts the gRPC server
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
			NodeId:      nodeConfig.nodeID,
			ClusterId:   nodeConfig.clusterID,
			GrpcAddress: nodeConfig.gRPCAddr,
		})

		if err != nil {
			log.Println(err)
		}

		resp, err := adminClient.LeaveCluster(context.TODO(), &RavelClusterAdminPB.Node{
			NodeId:      nodeConfig.nodeID,
			ClusterId:   nodeConfig.clusterID,
			GrpcAddress: nodeConfig.gRPCAddr,
			RaftAddress: nodeConfig.raftInternalAddr,
		})

		if err != nil {
			log.Println(err)
		}

		log.Println(resp)
		os.Exit(1)
	}()
}
