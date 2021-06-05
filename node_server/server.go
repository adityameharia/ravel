package node_server

import (
	"context"
	"errors"

	"github.com/adityameharia/ravel/RavelNodePB"
	"github.com/adityameharia/ravel/node"
	"github.com/hashicorp/raft"
)

type Server struct {
	Node *node.RavelNode
}

func (s *Server) Join(ctx context.Context, req *RavelNodePB.Node) (*RavelNodePB.Void, error) {
	err := s.Node.Join(req.NodeId, req.RaftAddress)
	if err != nil {
		return nil, err
	}

	return &RavelNodePB.Void{}, nil
}

func (s *Server) Leave(ctx context.Context, req *RavelNodePB.Node) (*RavelNodePB.Void, error) {
	err := s.Node.Leave(req.NodeId)
	if err != nil {
		return nil, err
	}

	return &RavelNodePB.Void{}, nil
}

func (s *Server) IsLeader(ctx context.Context, v *RavelNodePB.Void) (*RavelNodePB.Boolean, error) {
	if s.Node.Raft.State() != raft.Leader {
		return &RavelNodePB.Boolean{Leader: false}, nil
	}
	return &RavelNodePB.Boolean{Leader: true}, nil

}

func (s *Server) Run(ctx context.Context, req *RavelNodePB.Command) (*RavelNodePB.Response, error) {
	switch req.Operation {
	case "get":
		val, err := s.Node.Get(req.Key)
		if err != nil {
			return nil, err
		}
		return &RavelNodePB.Response{
			Msg:  "get successful",
			Data: val}, nil
	case "set":
		err := s.Node.Set(req.Key, req.Value)
		if err != nil {
			return nil, err
		}
		return &RavelNodePB.Response{Msg: "set successful", Data: []byte{}}, nil
	case "delete":
		err := s.Node.Delete(req.Key)
		if err != nil {
			return nil, err
		}
		return &RavelNodePB.Response{Msg: "delete successful", Data: []byte{}}, nil
	case "getanddelete":
		val, err := s.Node.Get(req.Key)
		if err != nil {
			return nil, err
		}
		err = s.Node.Delete(req.Key)
		if err != nil {
			return nil, err
		}
		return &RavelNodePB.Response{
			Msg:  "get successful",
			Data: val}, nil
	default:
		return nil, errors.New("invalid operation")
	}
}
