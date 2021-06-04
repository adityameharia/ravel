package main

import (
	"context"
	"errors"
	"github.com/adityameharia/ravel/RavelClusterAdminPB"
	"github.com/adityameharia/ravel/RavelNodePB"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"math"
	"sync"
)

type clusterInfo struct {
	LeaderNode   *RavelClusterAdminPB.Node
	ReplicaCount uint64
}

type ClusterAdminGRPCServer struct {
	mutex            sync.Mutex
	ClusterLeaderMap map[string]clusterInfo
	Server           *grpc.Server
}

func NewClusterAdminGRPCServer() *ClusterAdminGRPCServer {
	var newServer ClusterAdminGRPCServer
	newServer.ClusterLeaderMap = make(map[string]clusterInfo)
	newServer.Server = grpc.NewServer()
	return &newServer
}

func (s *ClusterAdminGRPCServer) JoinExistingCluster(ctx context.Context, node *RavelClusterAdminPB.Node) (*RavelClusterAdminPB.Cluster, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	log.Println("Join Existing Cluster: Request from", node.GrpcAddress)
	var minReplicaClusterID string = ""
	var minReplicaCount uint64 = math.MaxUint64
	for id, cInfo := range s.ClusterLeaderMap {
		if cInfo.ReplicaCount < minReplicaCount {
			minReplicaClusterID = id
			minReplicaCount = cInfo.ReplicaCount
		}
	}

	if minReplicaClusterID == "" {
		return nil, errors.New("no clusters found")
	}

	cInfo := s.ClusterLeaderMap[minReplicaClusterID]
	cInfo.ReplicaCount += 1
	s.ClusterLeaderMap[minReplicaClusterID] = cInfo

	log.Println(s.ClusterLeaderMap)

	return &RavelClusterAdminPB.Cluster{
		ClusterId: minReplicaClusterID,
		LeaderGrpcAddress: s.ClusterLeaderMap[minReplicaClusterID].LeaderNode.GrpcAddress,
		LeaderRaftAddress: s.ClusterLeaderMap[minReplicaClusterID].LeaderNode.RaftAddress,
	}, nil
}

func (s *ClusterAdminGRPCServer) JoinAsClusterLeader(ctx context.Context, node *RavelClusterAdminPB.Node) (*RavelClusterAdminPB.Cluster, error) {
	log.Println("JoinAsClusterLeader: Request from", node.GrpcAddress)
	newClusterID := uuid.New().String()
	s.mutex.Lock()
	s.ClusterLeaderMap[newClusterID] = clusterInfo{node, 1}
	s.mutex.Unlock()

	log.Println("Adding", node.GrpcAddress, "as a new cluster with ID:", newClusterID)
	return &RavelClusterAdminPB.Cluster{
		ClusterId: newClusterID,
		LeaderGrpcAddress: node.GrpcAddress, // same as the node that sent the request
		LeaderRaftAddress: node.RaftAddress,
	}, nil
}

func (s *ClusterAdminGRPCServer) UpdateClusterLeader(ctx context.Context, node *RavelClusterAdminPB.Node) (*RavelClusterAdminPB.Response, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if cInfo, exists := s.ClusterLeaderMap[node.ClusterId]; exists {
		s.ClusterLeaderMap[node.ClusterId] = clusterInfo{node, cInfo.ReplicaCount}
	} else {
		return nil, errors.New("invalid cluster id")
	}

	log.Println(s.ClusterLeaderMap)

	return &RavelClusterAdminPB.Response{
		Data: "leader updated successfully",
	}, nil
}

func (s *ClusterAdminGRPCServer) LeaveCluster(ctx context.Context, node *RavelClusterAdminPB.Node) (*RavelClusterAdminPB.Response, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	cInfo, exists := s.ClusterLeaderMap[node.ClusterId]
	if !exists {
		return nil, errors.New("invalid cluster id")
	}

	cInfo.ReplicaCount -= 1
	s.ClusterLeaderMap[node.ClusterId] = cInfo

	log.Println(s.ClusterLeaderMap)
	return &RavelClusterAdminPB.Response{
		Data: "replica count reduced",
	}, nil
}


func (s *ClusterAdminGRPCServer) GetClusterLeader(ctx context.Context, cluster *RavelClusterAdminPB.Cluster) (*RavelClusterAdminPB.Node, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	cInfo, exists := s.ClusterLeaderMap[cluster.ClusterId]
	if !exists {
		return nil, errors.New("invalid cluster id")
	}

	return cInfo.LeaderNode, nil
}

func (s *ClusterAdminGRPCServer) WriteKeyValue(key []byte, val []byte, clusterID string) error {
	conn, err := grpc.Dial(s.ClusterLeaderMap[clusterID].LeaderNode.GrpcAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}

	client := RavelNodePB.NewRavelNodeClient(conn)
	resp, err := client.Run(context.TODO(), &RavelNodePB.Command{
		Operation: "set",
		Key:       key,
		Value:     val,
	})

	if err != nil {
		return err
	}

	log.Println(resp.Data)
	return nil
}

func (s *ClusterAdminGRPCServer) ReadKey(key []byte, clusterID string) ([]byte, error) {
	conn, err := grpc.Dial(s.ClusterLeaderMap[clusterID].LeaderNode.GrpcAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := RavelNodePB.NewRavelNodeClient(conn)
	resp, err := client.Run(context.TODO(), &RavelNodePB.Command{
		Operation: "get",
		Key:       key,
	})

	if err != nil {
		return nil, err
	}

	log.Println(resp.Data)
	return nil
}
