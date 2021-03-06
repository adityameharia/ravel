// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package RavelClusterAdminPB

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RavelClusterAdminClient is the client API for RavelClusterAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RavelClusterAdminClient interface {
	JoinExistingCluster(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Cluster, error)
	JoinAsClusterLeader(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Cluster, error)
	UpdateClusterLeader(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Response, error)
	LeaveCluster(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Response, error)
	GetClusterLeader(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Node, error)
	InitiateDataRelocation(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Response, error)
}

type ravelClusterAdminClient struct {
	cc grpc.ClientConnInterface
}

func NewRavelClusterAdminClient(cc grpc.ClientConnInterface) RavelClusterAdminClient {
	return &ravelClusterAdminClient{cc}
}

func (c *ravelClusterAdminClient) JoinExistingCluster(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, "/RavelClusterAdminPB.RavelClusterAdmin/JoinExistingCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ravelClusterAdminClient) JoinAsClusterLeader(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, "/RavelClusterAdminPB.RavelClusterAdmin/JoinAsClusterLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ravelClusterAdminClient) UpdateClusterLeader(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/RavelClusterAdminPB.RavelClusterAdmin/UpdateClusterLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ravelClusterAdminClient) LeaveCluster(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/RavelClusterAdminPB.RavelClusterAdmin/LeaveCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ravelClusterAdminClient) GetClusterLeader(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/RavelClusterAdminPB.RavelClusterAdmin/GetClusterLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ravelClusterAdminClient) InitiateDataRelocation(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/RavelClusterAdminPB.RavelClusterAdmin/InitiateDataRelocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RavelClusterAdminServer is the server API for RavelClusterAdmin service.
// All implementations should embed UnimplementedRavelClusterAdminServer
// for forward compatibility
type RavelClusterAdminServer interface {
	JoinExistingCluster(context.Context, *Node) (*Cluster, error)
	JoinAsClusterLeader(context.Context, *Node) (*Cluster, error)
	UpdateClusterLeader(context.Context, *Node) (*Response, error)
	LeaveCluster(context.Context, *Node) (*Response, error)
	GetClusterLeader(context.Context, *Cluster) (*Node, error)
	InitiateDataRelocation(context.Context, *Cluster) (*Response, error)
}

// UnimplementedRavelClusterAdminServer should be embedded to have forward compatible implementations.
type UnimplementedRavelClusterAdminServer struct {
}

func (UnimplementedRavelClusterAdminServer) JoinExistingCluster(context.Context, *Node) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinExistingCluster not implemented")
}
func (UnimplementedRavelClusterAdminServer) JoinAsClusterLeader(context.Context, *Node) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinAsClusterLeader not implemented")
}
func (UnimplementedRavelClusterAdminServer) UpdateClusterLeader(context.Context, *Node) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClusterLeader not implemented")
}
func (UnimplementedRavelClusterAdminServer) LeaveCluster(context.Context, *Node) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveCluster not implemented")
}
func (UnimplementedRavelClusterAdminServer) GetClusterLeader(context.Context, *Cluster) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterLeader not implemented")
}
func (UnimplementedRavelClusterAdminServer) InitiateDataRelocation(context.Context, *Cluster) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitiateDataRelocation not implemented")
}

// UnsafeRavelClusterAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RavelClusterAdminServer will
// result in compilation errors.
type UnsafeRavelClusterAdminServer interface {
	mustEmbedUnimplementedRavelClusterAdminServer()
}

func RegisterRavelClusterAdminServer(s grpc.ServiceRegistrar, srv RavelClusterAdminServer) {
	s.RegisterService(&RavelClusterAdmin_ServiceDesc, srv)
}

func _RavelClusterAdmin_JoinExistingCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RavelClusterAdminServer).JoinExistingCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RavelClusterAdminPB.RavelClusterAdmin/JoinExistingCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RavelClusterAdminServer).JoinExistingCluster(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _RavelClusterAdmin_JoinAsClusterLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RavelClusterAdminServer).JoinAsClusterLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RavelClusterAdminPB.RavelClusterAdmin/JoinAsClusterLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RavelClusterAdminServer).JoinAsClusterLeader(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _RavelClusterAdmin_UpdateClusterLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RavelClusterAdminServer).UpdateClusterLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RavelClusterAdminPB.RavelClusterAdmin/UpdateClusterLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RavelClusterAdminServer).UpdateClusterLeader(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _RavelClusterAdmin_LeaveCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RavelClusterAdminServer).LeaveCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RavelClusterAdminPB.RavelClusterAdmin/LeaveCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RavelClusterAdminServer).LeaveCluster(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _RavelClusterAdmin_GetClusterLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RavelClusterAdminServer).GetClusterLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RavelClusterAdminPB.RavelClusterAdmin/GetClusterLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RavelClusterAdminServer).GetClusterLeader(ctx, req.(*Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _RavelClusterAdmin_InitiateDataRelocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RavelClusterAdminServer).InitiateDataRelocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RavelClusterAdminPB.RavelClusterAdmin/InitiateDataRelocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RavelClusterAdminServer).InitiateDataRelocation(ctx, req.(*Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

// RavelClusterAdmin_ServiceDesc is the grpc.ServiceDesc for RavelClusterAdmin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RavelClusterAdmin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RavelClusterAdminPB.RavelClusterAdmin",
	HandlerType: (*RavelClusterAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinExistingCluster",
			Handler:    _RavelClusterAdmin_JoinExistingCluster_Handler,
		},
		{
			MethodName: "JoinAsClusterLeader",
			Handler:    _RavelClusterAdmin_JoinAsClusterLeader_Handler,
		},
		{
			MethodName: "UpdateClusterLeader",
			Handler:    _RavelClusterAdmin_UpdateClusterLeader_Handler,
		},
		{
			MethodName: "LeaveCluster",
			Handler:    _RavelClusterAdmin_LeaveCluster_Handler,
		},
		{
			MethodName: "GetClusterLeader",
			Handler:    _RavelClusterAdmin_GetClusterLeader_Handler,
		},
		{
			MethodName: "InitiateDataRelocation",
			Handler:    _RavelClusterAdmin_InitiateDataRelocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmd/ravel_cluster_admin/cluster_admin.proto",
}
