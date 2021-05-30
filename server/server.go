package server

import (
	"context"

	"github.com/adityameharia/ravel/RavelClusterPB"
)

type Server struct{}

func (s *Server) AcceptJoin(ctx context.Context, req *RavelClusterPB.Join) (*RavelClusterPB.Response, error) {
	return nil, nil
}

func (s *Server) RequestJoin(ctx context.Context, req *RavelClusterPB.Join) (*RavelClusterPB.Response, error) {
	return nil, nil
}

func (s *Server) AcceptLeave(ctx context.Context, req *RavelClusterPB.Node) (*RavelClusterPB.Response, error) {
	return nil, nil
}

func (s *Server) RequestLeave(ctx context.Context, req *RavelClusterPB.Node) (*RavelClusterPB.Response, error) {
	return nil, nil
}

func (s *Server) Run(ctx context.Context, req *RavelClusterPB.Command) (*RavelClusterPB.Response, error) {
	return nil, nil
}
