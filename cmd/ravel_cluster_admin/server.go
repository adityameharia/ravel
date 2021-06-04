package main

//import (
//	"context"
//	"errors"
//	"log"
//
//	"github.com/adityameharia/ravel/RavelClusterAdminPB"
//	"github.com/adityameharia/ravel/RavelClusterPB"
//	"google.golang.org/grpc"
//)
//
//func (s *server) GetLeader(ctx context.Context, clusterID *RavelClusterAdminPB.Cluster) (*RavelClusterAdminPB.Response, error) {
//	mu.Lock()
//	defer mu.Unlock()
//	if len(serverList[clusterID.ClusterID]) == 0 {
//		return &RavelClusterAdminPB.Response{Data: ""}, nil
//	}
//
//	conn, err := grpc.Dial(leader[clusterID.ClusterID].gRPCAddress, grpc.WithInsecure())
//	if err != nil {
//		log.Printf("cannot connect with server %v", err)
//	}
//
//	v := &RavelClusterPB.Void{}
//	client := RavelClusterPB.NewRavelClusterClient(conn)
//	res, err := client.IsLeader(context.Background(), v)
//	if err != nil {
//		log.Printf("Is Leader request falied with server %v", err)
//	}
//
//	if res.Leader == true {
//		return &RavelClusterAdminPB.Response{Data: leader[clusterID.ClusterID].gRPCAddress}, nil
//	} else {
//		for _, rep := range serverList[clusterID.ClusterID] {
//
//			conn, err = grpc.Dial(rep.gRPCAddress, grpc.WithInsecure())
//			if err != nil {
//				log.Printf("cannot connect with server %v", err)
//			}
//
//			v := &RavelClusterPB.Void{}
//			client = RavelClusterPB.NewRavelClusterClient(conn)
//			res, err = client.IsLeader(context.Background(), v)
//			if err != nil {
//				log.Printf("Is Leader request falied with server %v", err)
//			}
//
//			if res.Leader == true {
//				return &RavelClusterAdminPB.Response{Data: rep.gRPCAddress}, nil
//			}
//		}
//
//		return nil, errors.New("No leader found")
//	}
//}
//
//func (s *server) AddToReplicaMap(ctx context.Context, clusterID *RavelClusterAdminPB.Node) (*RavelClusterAdminPB.Void, error) {
//	rep := Replica{
//		NodeID:      clusterID.NodeID,
//		gRPCAddress: clusterID.GRPCaddress,
//	}
//	serverList[clusterID.ClusterID] = append(serverList[clusterID.ClusterID], rep)
//	return &RavelClusterAdminPB.Void{}, nil
//}
//
//func (s *server) RemoveReplicaFromMap(ctx context.Context, clusterID *RavelClusterAdminPB.Node) (*RavelClusterAdminPB.Void, error) {
//	for i, r := range serverList[clusterID.ClusterID] {
//		if r.NodeID == clusterID.NodeID {
//			serverList[clusterID.ClusterID] = RemoveIndex(serverList[clusterID.ClusterID], i)
//			return &RavelClusterAdminPB.Void{}, nil
//		}
//	}
//	return nil, errors.New("Replica not found in the server list")
//}
//
//func RemoveIndex(sl []Replica, index int) []Replica {
//	return append(sl[:index], sl[index+1:]...)
//}
