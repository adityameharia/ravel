package server

import (
	"context"

	"github.com/adityameharia/ravel/RavelClusterPB"
)

type Server struct{}

func (s *Server) Join(ctx context.Context, req *RavelClusterPB.Node) (*RavelClusterPB.Response, error) {
	return nil, nil
}

func (s *Server) Leave(ctx context.Context, req *RavelClusterPB.Node) (*RavelClusterPB.Response, error) {
	return nil, nil
}

func (s *Server) Run(ctx context.Context, req *RavelClusterPB.Command) (*RavelClusterPB.Response, error) {
	return nil, nil
}
