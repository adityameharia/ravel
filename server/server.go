package server

import (
	"context"

	"github.com/adityameharia/ravel/RavelClusterPB"
	"github.com/adityameharia/ravel/node"
)

type Server struct{
	Node *node.RavelNode
}

func (s *Server) Join(ctx context.Context, req *RavelClusterPB.Node) (*RavelClusterPB.Response, error) {
	err := s.Node.Join(req.NodeID, *(req.Address))
	if err != nil {

	}

	return nil, nil
}

func (s *Server) Leave(ctx context.Context, req *RavelClusterPB.Node) (*RavelClusterPB.Response, error) {
	return nil, nil
}

func (s *Server) Run(ctx context.Context, req *RavelClusterPB.Command) (*RavelClusterPB.Response, error) {
	return nil, nil
}
