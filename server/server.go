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

func (s *Server) Join(ctx context.Context, req *RavelClusterPB.Node) (*RavelClusterPB.Void, error) {
	err := s.Node.Join(req.NodeID, req.Address)
	if err != nil {
		return nil, err
	}

	return &RavelClusterPB.Void{}, nil
}

func (s *Server) Leave(ctx context.Context, req *RavelClusterPB.Node) (*RavelClusterPB.Void, error) {
	err := s.Node.Leave(req.NodeID)
	if err != nil {
		return nil, err
	}

	return &RavelClusterPB.Void{}, nil
}

func (s *Server) Run(ctx context.Context, req *RavelClusterPB.Command) (*RavelClusterPB.Response, error) {
	switch req.Operation {
	case "get":
		val, err := s.Node.Get(req.Key)
		if err != nil {
			return nil, err
		}
		return &RavelClusterPB.Response{Data: val}, nil
	case "set":
		err := s.Node.Set(req.Key, req.Value)
		if err != nil {
			return nil, err
		}
		return &RavelClusterPB.Response{Data: "set successful"}, nil
	case "delete":
		err := s.Node.Delete(req.Key)
		if err != nil {
			return nil, err
		}
		return &RavelClusterPB.Response{Data: "delete successful"}, nil
	default:
		return nil, errors.New("invalid operation")
	}
}
