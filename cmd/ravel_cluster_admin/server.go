package main

import (
	"context"

	"github.com/adityameharia/ravel/RavelClusterAdminPB"
)

func (s *server) GetLeader(ctx context.Context, v *RavelClusterAdminPB.Cluster) (*RavelClusterAdminPB.Response, error) {
	return nil, nil
}

func (s *server) AddToReplicaMap(ctx context.Context, v *RavelClusterAdminPB.Node) (*RavelClusterAdminPB.Void, error) {
	return nil, nil
}
