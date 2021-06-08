package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/adityameharia/ravel/RavelNodePB"
	"github.com/adityameharia/ravel/node"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"

	"github.com/adityameharia/ravel/RavelClusterAdminPB"
	"github.com/adityameharia/ravel/node_server"
	"github.com/google/uuid"
)

// Config is the struct containing the configuration details of the node
type Config struct {
	// clusterID is ID of th cluster the node is a part of
	ClusterID string `yaml:"clusterid"`
	// nodeID is the nodes unique ID
	NodeID string `yaml:"nodeid"`
	// storageDir is the Data Directory for Raft
	StorageDir string `yaml:"storagedir"`
	// gRPCAddr is the Address (with port) at which gRPC server is started
	GRPCAddr string `yaml:"grpcaddr"`
	// raftInternalAddr is the Raft internal communication address with port
	RaftInternalAddr string `yaml:"raftaddr"`
	// adminGRPCAddr is the GRPC address of the cluster admin
	AdminGRPCAddr string `yaml:"adminrpcaddr"`
	// isLeader is a bool defining whether the node is a leader or not
	IsLeader bool `yaml:"leader"`
}

var yamlFile string
var nodeConfig Config
var adminClient RavelClusterAdminPB.RavelClusterAdminClient

func init() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	nodeConfig.NodeID = uuid.New().String()
	flag.StringVar(&yamlFile, "yaml", "", "Argument as yaml file or command line arguments")
	flag.StringVar(&nodeConfig.StorageDir, "storageDir", dirname, "Data Directory for Raft")
	flag.StringVar(&nodeConfig.GRPCAddr, "gRPCAddr", "", "Address (with port) at which gRPC server is started")
	flag.StringVar(&nodeConfig.RaftInternalAddr, "raftAddr", "", "Raft internal communication address with port")
	flag.StringVar(&nodeConfig.AdminGRPCAddr, "adminRPCAddr", "", "GRPC address of the cluster admin")
	flag.BoolVar(&nodeConfig.IsLeader, "leader", false, "Register this node as a new leader or not")
}

func main() {
	flag.Parse()

	if yamlFile != "" {
		err := readConf(yamlFile)
		if err != nil {
			log.Fatal("Unable to get the yaml file")
		}
	}

	adminConn, err := grpc.Dial(nodeConfig.AdminGRPCAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error in connecting to the Admin gRPC Server: ", err)
	}
	defer adminConn.Close()

	var ravelNode node.RavelNode
	adminClient = RavelClusterAdminPB.NewRavelClusterAdminClient(adminConn)

	if nodeConfig.IsLeader {
		ravelCluster, err := adminClient.JoinAsClusterLeader(context.TODO(), &RavelClusterAdminPB.Node{
			NodeId:      nodeConfig.NodeID,   // id of this node
			GrpcAddress: nodeConfig.GRPCAddr, // grpc address of this node
			RaftAddress: nodeConfig.RaftInternalAddr,
			ClusterId:   "", // cluster id is unknown thus empty
		})

		if err != nil {
			log.Fatal("Error in JoinAsClusterLeader: ", err)
		} else {
			nodeConfig.ClusterID = ravelCluster.ClusterId
			ravelNode.Raft, ravelNode.Fsm, err = ravelNode.Open(nodeConfig.IsLeader, nodeConfig.NodeID, nodeConfig.StorageDir, nodeConfig.RaftInternalAddr)
			if err != nil {
				log.Println(err)
			}

			// this node is the leader
		}
	} else {
		ravelCluster, err := adminClient.JoinExistingCluster(context.TODO(), &RavelClusterAdminPB.Node{
			NodeId:      nodeConfig.NodeID,
			GrpcAddress: nodeConfig.GRPCAddr,
			RaftAddress: nodeConfig.RaftInternalAddr,
			ClusterId:   "",
		})

		if err != nil {
			log.Fatal("Error in JoinExistingCluster: ", err)
		} else {
			log.Println("Cluster leader is: ", ravelCluster.LeaderGrpcAddress)
			nodeConfig.ClusterID = ravelCluster.ClusterId
			ravelNode.Raft, ravelNode.Fsm, err = ravelNode.Open(nodeConfig.IsLeader, nodeConfig.NodeID, nodeConfig.StorageDir, nodeConfig.RaftInternalAddr)
			if err != nil {
				log.Println(err)
			}

			log.Println("here")
			err = RequestJoinToClusterLeader(ravelCluster.LeaderGrpcAddress, &RavelNodePB.Node{
				NodeId:      nodeConfig.NodeID,
				ClusterId:   nodeConfig.ClusterID,
				GrpcAddress: nodeConfig.GRPCAddr,
				RaftAddress: nodeConfig.RaftInternalAddr,
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
			err := RequestLeaderUpdateToCluster(nodeConfig.AdminGRPCAddr, &RavelClusterAdminPB.Node{
				NodeId:      nodeConfig.NodeID,
				GrpcAddress: nodeConfig.GRPCAddr,
				RaftAddress: nodeConfig.RaftInternalAddr,
				ClusterId:   nodeConfig.ClusterID,
			})

			if err != nil {
				log.Println(err)
			}
		}
	}()

	//sends leave request to leader and admin on signal interrupt
	onSignalInterrupt()

	//starts the gRPC server
	listener, err := net.Listen("tcp", nodeConfig.GRPCAddr)
	if err != nil {
		log.Fatal("Error in starting TCP server: ", err)
	}
	log.Printf("Starting TCP Server on %v for gRPC\n", nodeConfig.GRPCAddr)

	grpcServer := grpc.NewServer()
	RavelNodePB.RegisterRavelNodeServer(grpcServer, &node_server.Server{
		Node: &ravelNode,
	})

	if nodeConfig.IsLeader {
		go initiateDataRelocation()
	}

	err = grpcServer.Serve(listener)
}

func readConf(path string) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(buf, &nodeConfig)
	if err != nil {
		return fmt.Errorf("in file %q: %v", path, err)
	}
	return nil
}

func initiateDataRelocation() {
	time.Sleep(5 * time.Second)
	resp, err := adminClient.InitiateDataRelocation(context.TODO(), &RavelClusterAdminPB.Cluster{
		ClusterId: nodeConfig.ClusterID,
	})

	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(resp.Data)
}

func onSignalInterrupt() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	go func() {
		<-ch
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

		resp, err := adminClient.LeaveCluster(context.TODO(), &RavelClusterAdminPB.Node{
			NodeId:      nodeConfig.NodeID,
			ClusterId:   nodeConfig.ClusterID,
			GrpcAddress: nodeConfig.GRPCAddr,
			RaftAddress: nodeConfig.RaftInternalAddr,
		})

		if err != nil {
			log.Println(err)
		}

		log.Println(resp)
		os.Exit(1)
	}()
}
