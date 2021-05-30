package server

import (
	"context"
	"errors"
	"github.com/adityameharia/ravel/RavelClusterPB"
	"github.com/adityameharia/ravel/node"
)

type Server struct {
	Node *node.RavelNode
}

func (s *Server) Join(ctx context.Context, req *RavelClusterPB.Node) (*RavelClusterPB.Response, error) {
	joinResp := s.Node.Join(req.NodeID, req.Address)
	if joinResp.Error == "node node is not leader" {
		resp := RavelClusterPB.Response{Data: joinResp.Leader}
		return &resp, nil
	} else {
		resp := RavelClusterPB.Response{Data: ""}
		return &resp, errors.New(joinResp.Error)
	}
}

func (s *Server) Leave(ctx context.Context, req *RavelClusterPB.Node) (*RavelClusterPB.Response, error) {
	leaveResp := s.Node.Leave(req.NodeID)
	if leaveResp.Error == "node node is not leader" {
		resp := RavelClusterPB.Response{Data: leaveResp.Leader}
		return &resp, nil
	} else {
		resp := RavelClusterPB.Response{Data: ""}
		return &resp, errors.New(leaveResp.Error)
	}
}

func (s *Server) Run(ctx context.Context, req *RavelClusterPB.Command) (*RavelClusterPB.Response, error) {
	return nil, nil
}
