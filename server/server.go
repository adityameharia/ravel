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
	log.Println("5")
	log.Println(req.NodeID)
	log.Println(req.Address + "hi")
	log.Println(s.Node)
	joinResp := s.Node.Join(req.NodeID, req.Address)
	log.Println("6")
	if joinResp.Error == "node is not leader" {
		resp := RavelClusterPB.Response{Data: joinResp.Leader}
		return &resp, nil
	} else if joinResp.Error == "" {
		resp := RavelClusterPB.Response{Data: ""}
		return &resp, nil
	} else {
		resp := RavelClusterPB.Response{Data: ""}
		return &resp, errors.New(joinResp.Error)
	}
}

func (s *Server) Leave(ctx context.Context, req *RavelClusterPB.Node) (*RavelClusterPB.Response, error) {
	leaveResp := s.Node.Leave(req.NodeID)
	if leaveResp.Error == "node is not leader" {
		resp := RavelClusterPB.Response{Data: leaveResp.Leader}
		return &resp, nil
	} else if leaveResp.Error == "" {
		resp := RavelClusterPB.Response{Data: ""}
		return &resp, nil
	} else {
		resp := RavelClusterPB.Response{Data: ""}
		return &resp, errors.New(leaveResp.Error)
	}
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
