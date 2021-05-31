package server

import (
	"context"
	"errors"
	"log"

	"github.com/adityameharia/ravel/RavelClusterPB"
	"github.com/adityameharia/ravel/node"
)

type Server struct {
	Node *node.RavelNode
}

func (s *Server) Join(ctx context.Context, req *RavelClusterPB.Node) (*RavelClusterPB.Response, error) {

	log.Println(req.NodeID)
	log.Println(req.Address + "hi")
	log.Println(s.Node)
	leader, err := s.Node.Join(req.NodeID, req.Address)
	log.Println(err)
	if err != nil && err.Error() == "node is not leader" {
		resp := &RavelClusterPB.Response{Data: leader}
		log.Println(resp)
		return resp, err
	} else if err != nil {
		resp := &RavelClusterPB.Response{Data: ""}
		return resp, err
	}
	return &RavelClusterPB.Response{Data: ""}, nil
}

func (s *Server) Leave(ctx context.Context, req *RavelClusterPB.Node) (*RavelClusterPB.Response, error) {
	leader, err := s.Node.Leave(req.NodeID)
	if err != nil && err.Error() == "node is not leader" {
		resp := RavelClusterPB.Response{Data: leader}
		return &resp, err
	} else if err != nil {
		resp := RavelClusterPB.Response{Data: ""}
		return &resp, err
	}
	resp := RavelClusterPB.Response{Data: ""}
	return &resp, nil

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
