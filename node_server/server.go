package node_server

import (
	"context"
	"errors"

	"github.com/adityameharia/ravel/RavelNodePB"
	"github.com/adityameharia/ravel/node"
	"github.com/hashicorp/raft"
)

// Server implements the methods exposed via gRPC for a RavelNode
type Server struct {
	Node *node.RavelNode
}

// Join joins the passed in node to this node
func (s *Server) Join(ctx context.Context, req *RavelNodePB.Node) (*RavelNodePB.Void, error) {
	err := s.Node.Join(req.NodeId, req.RaftAddress)
	if err != nil {
		return nil, err
	}

	return &RavelNodePB.Void{}, nil
}

// Leave removes the passed in node from this leader
func (s *Server) Leave(ctx context.Context, req *RavelNodePB.Node) (*RavelNodePB.Void, error) {
	err := s.Node.Leave(req.NodeId)
	if err != nil {
		return nil, err
	}

	return &RavelNodePB.Void{}, nil
}

// IsLeader returns a boolean if this node is a leader or not
func (s *Server) IsLeader(ctx context.Context, v *RavelNodePB.Void) (*RavelNodePB.Boolean, error) {
	if s.Node.Raft.State() != raft.Leader {
		return &RavelNodePB.Boolean{Leader: false}, nil
	}
	return &RavelNodePB.Boolean{Leader: true}, nil

}

// Run executes the operation specified in "req", it can be {get, set, delete, getAndDelete}
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
	case "getAndDelete":
		val, err := s.Node.GetAndDelete(req.Key)
		if err != nil {
			return nil, err
		}
		return &RavelNodePB.Response{
			Msg:  "get and delete successful",
			Data: val}, nil
	default:
		return nil, errors.New("invalid operation")
	}
}
