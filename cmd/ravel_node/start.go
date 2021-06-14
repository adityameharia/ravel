package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"time"

	"github.com/adityameharia/ravel/RavelClusterAdminPB"
	"github.com/adityameharia/ravel/RavelNodePB"
	"github.com/adityameharia/ravel/node"
	"github.com/adityameharia/ravel/node_server"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
)

func startReplica() {
	if yamlFile != "" {
		err := readConf(yamlFile)
		if err != nil {
			log.Fatal("Unable to get the yaml file")
		}
	}

	if nodeConfig.AdminGRPCAddr == "" {
		log.Fatal("adminRPCAddr has not been initialized")
	}

	adminConn, err := grpc.Dial(nodeConfig.AdminGRPCAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error in connecting to the Admin gRPC Server: ", err)
	}
	defer adminConn.Close()

	adminClient = RavelClusterAdminPB.NewRavelClusterAdminClient(adminConn)

	var ravelNode node.RavelNode

	_, err = conf.Read([]byte("clusterID"))
	if err == nil {
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

	byteConfig, err := json.Marshal(nodeConfig)
	if err != nil {
		killCluster()
		log.Fatal("cannot write config to file")
	}

	err = conf.Write([]byte("config"), byteConfig)
	if err != nil {
		killCluster()
		log.Fatal("cannot write config to file")
	}

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
