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

// clusterInfo holds the information to represent a cluster
type clusterInfo struct {
	LeaderNode   *RavelClusterAdminPB.Node
	ReplicaCount uint64
}

// ClusterAdminGRPCServer is the entity that implements the gRPC server for the Cluster Admin
type ClusterAdminGRPCServer struct {
	mutex            sync.Mutex
	ClusterLeaderMap map[string]clusterInfo
	Server           *grpc.Server
}

// NewClusterAdminGRPCServer constructs and returns a ClusterAdminGRPCServer object
func NewClusterAdminGRPCServer() *ClusterAdminGRPCServer {
	var newServer ClusterAdminGRPCServer
	newServer.ClusterLeaderMap = make(map[string]clusterInfo)
	newServer.Server = grpc.NewServer()
	return &newServer
}

// JoinExistingCluster picks the cluster with the least number of replicas and returns information about that cluster
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

	return &RavelClusterAdminPB.Cluster{
		ClusterId:         minReplicaClusterID,
		LeaderGrpcAddress: s.ClusterLeaderMap[minReplicaClusterID].LeaderNode.GrpcAddress,
		LeaderRaftAddress: s.ClusterLeaderMap[minReplicaClusterID].LeaderNode.RaftAddress,
	}, nil
}

// JoinAsClusterLeader creates a new cluster adds "node" as the cluster leader,
// this also adds a member in the RavelConsistentHash entity
func (s *ClusterAdminGRPCServer) JoinAsClusterLeader(ctx context.Context, node *RavelClusterAdminPB.Node) (*RavelClusterAdminPB.Cluster, error) {
	log.Println("JoinAsClusterLeader: Request from", node.GrpcAddress)
	newClusterID := uuid.New().String()

	s.mutex.Lock()
	s.ClusterLeaderMap[newClusterID] = clusterInfo{node, 1}
	s.mutex.Unlock()

	log.Println("Adding", node.GrpcAddress, "as a new clusterID with ID:", newClusterID)

	return &RavelClusterAdminPB.Cluster{
		ClusterId:         newClusterID,
		LeaderGrpcAddress: node.GrpcAddress, // same as the node that sent the request
		LeaderRaftAddress: node.RaftAddress,
	}, nil
}

// UpdateClusterLeader updates "node" as the leader of it's cluster. This is called when a leader crashes and another
// leader is picked via the Leader Election in Raft.
func (s *ClusterAdminGRPCServer) UpdateClusterLeader(ctx context.Context, node *RavelClusterAdminPB.Node) (*RavelClusterAdminPB.Response, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if cInfo, exists := s.ClusterLeaderMap[node.ClusterId]; exists {
		s.ClusterLeaderMap[node.ClusterId] = clusterInfo{node, cInfo.ReplicaCount}
	} else {
		return nil, errors.New("invalid clusterID id")
	}

	log.Println(s.ClusterLeaderMap)

	return &RavelClusterAdminPB.Response{
		Data: "leader updated successfully",
	}, nil
}

// LeaveCluster decrements the replica count of the node's cluster
func (s *ClusterAdminGRPCServer) LeaveCluster(ctx context.Context, node *RavelClusterAdminPB.Node) (*RavelClusterAdminPB.Response, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	cInfo, exists := s.ClusterLeaderMap[node.ClusterId]
	if !exists {
		return nil, errors.New("invalid clusterID")
	}

	if len(s.ClusterLeaderMap) == 1 {
		// last remaining cluster in the system -> reset consistentHash -> delete info from ClusterLeaderMap
		log.Printf("Node: %v from Cluster: %v is the last standing Cluster Leader in the system\n", node.NodeId, node.ClusterId)
		log.Println("Resetting consistentHash, Removing", node.ClusterId, "from ClusterLeaderMap")

		consistentHash.Reset(271, 40, 1.2)
		err := consistentHash.BackupToDisk(RavelClusterAdminBackupPath)
		if err != nil {
			return nil, err
		}
		delete(s.ClusterLeaderMap, node.ClusterId)

		return &RavelClusterAdminPB.Response{
			Data: "Removing last standing cluster in the system",
		}, nil
	} else {
		if cInfo.ReplicaCount == 1 {
			// last remaining replica in the cluster -> remove cluster from ClusterLeaderMap -> remove cluster from consistentHash
			log.Printf("Node: %v from Cluster: %v is the last standing replica in the cluster\n", node.NodeId, node.ClusterId)
			log.Println("Deleting cluster from consistent hash and removing", node.ClusterId, "from ClusterLeaderMap")
			consistentHash.DeleteCluster(clusterID(node.ClusterId))
			delete(s.ClusterLeaderMap, node.ClusterId)

			return &RavelClusterAdminPB.Response{
				Data: "Deleting Cluster: " + node.ClusterId,
			}, nil
		} else {
			cInfo.ReplicaCount -= 1
			s.ClusterLeaderMap[node.ClusterId] = cInfo

			log.Println(s.ClusterLeaderMap)
			return &RavelClusterAdminPB.Response{
				Data: "replica count reduced",
			}, nil
		}
	}
}

// GetClusterLeader returns information about the leader node of the provided cluster
func (s *ClusterAdminGRPCServer) GetClusterLeader(ctx context.Context, cluster *RavelClusterAdminPB.Cluster) (*RavelClusterAdminPB.Node, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	cInfo, exists := s.ClusterLeaderMap[cluster.ClusterId]
	if !exists {
		return nil, errors.New("invalid clusterID id")
	}

	return cInfo.LeaderNode, nil
}

// InitiateDataRelocation adds the provided cluster as an owner to the consistent hashing setup,
// which in turns takes care of the movement of data
func (s *ClusterAdminGRPCServer) InitiateDataRelocation(ctx context.Context, cluster *RavelClusterAdminPB.Cluster) (*RavelClusterAdminPB.Response, error) {
	consistentHash.AddCluster(clusterID(cluster.ClusterId))
	return &RavelClusterAdminPB.Response{
		Data: "data relocation completed",
	}, nil
}

// WriteKeyValue writes the given key and value to the leader of the provided cluster.
// NOTE: this function is not exposed via gRPC
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

// ReadKey reads the value for the given key from the leader of the provided cluster.
// NOTE: this function is not exposed via gRPC
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

	return resp.Data, nil
}

// DeleteKey deletes the key and value on the server
// NOTE: this function is not exposed via gRPC
func (s *ClusterAdminGRPCServer) DeleteKey(key []byte, clusterID string) error {
	conn, err := grpc.Dial(s.ClusterLeaderMap[clusterID].LeaderNode.GrpcAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}

	client := RavelNodePB.NewRavelNodeClient(conn)
	resp, err := client.Run(context.TODO(), &RavelNodePB.Command{
		Operation: "delete",
		Key:       key,
	})

	if err != nil {
		return err
	}
	log.Println(resp.Msg)
	return nil
}

// ReadKeyAndDelete reads the key, value and then deletes it on the server
// NOTE: this function is not exposed via gRPC
func (s *ClusterAdminGRPCServer) ReadKeyAndDelete(key []byte, clusterID string) ([]byte, error) {
	conn, err := grpc.Dial(s.ClusterLeaderMap[clusterID].LeaderNode.GrpcAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := RavelNodePB.NewRavelNodeClient(conn)
	resp, err := client.Run(context.TODO(), &RavelNodePB.Command{
		Operation: "getAndDelete",
		Key:       key,
	})

	if err != nil {
		return nil, err
	}

	log.Println(resp.Msg)
	return resp.Data, nil
}
