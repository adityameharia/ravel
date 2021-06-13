package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/adityameharia/ravel/RavelNodePB"
	"github.com/adityameharia/ravel/db"
	"github.com/adityameharia/ravel/node"
	"google.golang.org/grpc"

	"github.com/adityameharia/ravel/RavelClusterAdminPB"
	"github.com/adityameharia/ravel/node_server"
	"github.com/google/uuid"
)

// Config is the struct containing the configuration details of the node
type Config struct {
	ClusterID        string // ClusterID is ID of th cluster the node is a part of
	NodeID           string // NodeID is the nodes unique ID
	StorageDir       string // StorageDir is the Data Directory for Raft
	GRPCAddr         string // GRPCAddr is the Address (with port) at which gRPC server is started
	RaftInternalAddr string // RaftInternalAddr is the Raft internal communication address with port
	AdminGRPCAddr    string // AdminGRPCAddr is the address at which the cluster admin gRPC server is hosted
	IsLeader         bool   // IsLeader is a bool defining whether the node is a leader or not
}

var nodeConfig Config
var adminClient RavelClusterAdminPB.RavelClusterAdminClient

func init() {
	nodeConfig.NodeID = uuid.New().String()
	nodeConfig.GRPCAddr = "172.17.0.1:50000"
	nodeConfig.RaftInternalAddr = "172.17.0.1:60000"
	nodeConfig.StorageDir = "/ravel_node"

	//flag.StringVar(&nodeConfig.StorageDir, "storageDir", "", "Storage Dir")
	//flag.StringVar(&nodeConfig.GRPCAddr, "gRPCAddr", "", "GRPC Addr of this node")
	//flag.StringVar(&nodeConfig.RaftInternalAddr, "raftAddr", "", "Raft Internal address for this node")
	flag.StringVar(&nodeConfig.AdminGRPCAddr, "adminRPCAddr", "", "GRPC address of the cluster admin")
	flag.BoolVar(&nodeConfig.IsLeader, "leader", false, "Register this node as a new leader or not")
}

func main() {
	flag.Parse()

	var conf db.RavelDatabase
	err := conf.Init(nodeConfig.StorageDir + "/config")
	if err != nil {
		log.Fatal("FSM: Unable to Setup Database")
	}

	if nodeConfig.AdminGRPCAddr == "" {
		log.Fatal("adminRPCAddr has not been initialized")
	}

	adminConn, err := grpc.Dial(nodeConfig.AdminGRPCAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error in connecting to the Admin gRPC Server: ", err)
	}
	defer adminConn.Close()

	var ravelNode node.RavelNode
	adminClient = RavelClusterAdminPB.NewRavelClusterAdminClient(adminConn)

	cID, err := conf.Read([]byte("clusterID"))
	if err == nil {
		log.Println("texting")
		nodeConfig.ClusterID = string(cID)
		nID, err := conf.Read([]byte("nodeID"))
		if err != nil {
			log.Fatal("Cant get nodeID")
		}
		nodeConfig.NodeID = string(nID)
		ravelNode.Raft, ravelNode.Fsm, err = ravelNode.Open(false, nodeConfig.NodeID, nodeConfig.StorageDir, nodeConfig.RaftInternalAddr)
		if err != nil {
			log.Println(err)
		}

	} else {

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

				err = conf.Write([]byte("clusterID"), []byte(ravelCluster.ClusterId))
				if err != nil {
					log.Fatalf("Unable to clsuterID to disk")
				}
				err = conf.Write([]byte("nodeID"), []byte(nodeConfig.NodeID))
				if err != nil {
					log.Fatalf("Unable to clsuterID to disk")
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
				err = conf.Write([]byte("clusterID"), []byte(ravelCluster.ClusterId))
				if err != nil {
					log.Fatalf("Unable to clsuterID to disk")
				}
				err = conf.Write([]byte("nodeID"), []byte(nodeConfig.NodeID))
				if err != nil {
					log.Fatalf("Unable to clsuterID to disk")
				}
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

		err = os.RemoveAll(nodeConfig.StorageDir)
		if err != nil {
			log.Fatal(err)
		}

		os.Exit(1)
	}()
}
