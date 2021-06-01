package main

import (
	"context"
	"errors"
	"log"

	"github.com/adityameharia/ravel/RavelClusterAdminPB"
	"github.com/adityameharia/ravel/RavelClusterPB"
	"google.golang.org/grpc"
)

func (s *server) GetLeader(ctx context.Context, cluster *RavelClusterAdminPB.Cluster) (*RavelClusterAdminPB.Response, error) {

	if len(serverList[cluster.ClusterID]) == 0 {
		return &RavelClusterAdminPB.Response{Data: ""}, nil
	}

	conn, err := grpc.Dial(leader.gRPCAddress, grpc.WithInsecure())
	if err != nil {
		log.Printf("cannot connect with server %v", err)
	}

	v := &RavelClusterPB.Void{}
	client := RavelClusterPB.NewRavelClusterClient(conn)
	res, err := client.IsLeader(context.Background(), v)
	if err != nil {
		log.Printf("Is Leader request falied with server %v", err)
	}

	if res.Leader == true {
		return &RavelClusterAdminPB.Response{Data: leader.gRPCAddress}, nil
	} else {
		for _, rep := range serverList[cluster.ClusterID] {

			conn, err = grpc.Dial(rep.gRPCAddress, grpc.WithInsecure())
			if err != nil {
				log.Printf("cannot connect with server %v", err)
			}

			v := &RavelClusterPB.Void{}
			client = RavelClusterPB.NewRavelClusterClient(conn)
			res, err = client.IsLeader(context.Background(), v)
			if err != nil {
				log.Printf("Is Leader request falied with server %v", err)
			}

			if res.Leader == true {
				return &RavelClusterAdminPB.Response{Data: rep.gRPCAddress}, nil
			}
		}

		return nil, errors.New("No leader found")
	}
}

func (s *server) AddToReplicaMap(ctx context.Context, cluster *RavelClusterAdminPB.Node) (*RavelClusterAdminPB.Void, error) {
	rep := Replica{
		NodeID:      cluster.NodeID,
		gRPCAddress: cluster.GRPCaddress,
	}
	serverList[cluster.ClusterID] = append(serverList[cluster.ClusterID], rep)
	return &RavelClusterAdminPB.Void{}, nil
}
