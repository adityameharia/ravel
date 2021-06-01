// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.0
// source: ravel/cmd/ravel_cluster_admin/clusterAdmin.proto

package RavelClusterAdminPB

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_rawDescGZIP(), []int{0}
}

type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterID int32 `protobuf:"varint,1,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_rawDescGZIP(), []int{1}
}

func (x *Cluster) GetClusterID() int32 {
	if x != nil {
		return x.ClusterID
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeID      string `protobuf:"bytes,1,opt,name=NodeID,proto3" json:"NodeID,omitempty"`
	GRPCaddress string `protobuf:"bytes,2,opt,name=gRPCaddress,proto3" json:"gRPCaddress,omitempty"`
	ClusterID   int32  `protobuf:"varint,3,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_rawDescGZIP(), []int{3}
}

func (x *Node) GetNodeID() string {
	if x != nil {
		return x.NodeID
	}
	return ""
}

func (x *Node) GetGRPCaddress() string {
	if x != nil {
		return x.GRPCaddress
	}
	return ""
}

func (x *Node) GetClusterID() int32 {
	if x != nil {
		return x.ClusterID
	}
	return 0
}

var File_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto protoreflect.FileDescriptor

var file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_rawDesc = []byte{
	0x0a, 0x30, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x72, 0x61, 0x76, 0x65,
	0x6c, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x52, 0x61, 0x76, 0x65, 0x6c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x42, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22,
	0x27, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x22, 0x1e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5e, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x67, 0x52, 0x50, 0x43,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x67,
	0x52, 0x50, 0x43, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x32, 0xa6, 0x01, 0x0a, 0x11, 0x52, 0x61, 0x76,
	0x65, 0x6c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x48,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x52, 0x61,
	0x76, 0x65, 0x6c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50,
	0x42, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x1d, 0x2e, 0x52, 0x61, 0x76, 0x65,
	0x6c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x42, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x54,
	0x6f, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4d, 0x61, 0x70, 0x12, 0x19, 0x2e, 0x52, 0x61,
	0x76, 0x65, 0x6c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50,
	0x42, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x19, 0x2e, 0x52, 0x61, 0x76, 0x65, 0x6c, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x42, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x52, 0x61, 0x76, 0x65, 0x6c, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x42, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_rawDescOnce sync.Once
	file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_rawDescData = file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_rawDesc
)

func file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_rawDescGZIP() []byte {
	file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_rawDescOnce.Do(func() {
		file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_rawDescData = protoimpl.X.CompressGZIP(file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_rawDescData)
	})
	return file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_rawDescData
}

var file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_goTypes = []interface{}{
	(*Void)(nil),     // 0: RavelClusterAdminPB.Void
	(*Cluster)(nil),  // 1: RavelClusterAdminPB.cluster
	(*Response)(nil), // 2: RavelClusterAdminPB.Response
	(*Node)(nil),     // 3: RavelClusterAdminPB.Node
}
var file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_depIdxs = []int32{
	1, // 0: RavelClusterAdminPB.RavelClusterAdmin.GetLeader:input_type -> RavelClusterAdminPB.cluster
	3, // 1: RavelClusterAdminPB.RavelClusterAdmin.AddToReplicaMap:input_type -> RavelClusterAdminPB.Node
	2, // 2: RavelClusterAdminPB.RavelClusterAdmin.GetLeader:output_type -> RavelClusterAdminPB.Response
	0, // 3: RavelClusterAdminPB.RavelClusterAdmin.AddToReplicaMap:output_type -> RavelClusterAdminPB.Void
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_init() }
func file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_init() {
	if File_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_goTypes,
		DependencyIndexes: file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_depIdxs,
		MessageInfos:      file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_msgTypes,
	}.Build()
	File_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto = out.File
	file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_rawDesc = nil
	file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_goTypes = nil
	file_ravel_cmd_ravel_cluster_admin_clusterAdmin_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RavelClusterAdminClient is the client API for RavelClusterAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RavelClusterAdminClient interface {
	GetLeader(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Response, error)
	AddToReplicaMap(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Void, error)
}

type ravelClusterAdminClient struct {
	cc grpc.ClientConnInterface
}

func NewRavelClusterAdminClient(cc grpc.ClientConnInterface) RavelClusterAdminClient {
	return &ravelClusterAdminClient{cc}
}

func (c *ravelClusterAdminClient) GetLeader(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/RavelClusterAdminPB.RavelClusterAdmin/GetLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ravelClusterAdminClient) AddToReplicaMap(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/RavelClusterAdminPB.RavelClusterAdmin/AddToReplicaMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RavelClusterAdminServer is the server API for RavelClusterAdmin service.
type RavelClusterAdminServer interface {
	GetLeader(context.Context, *Cluster) (*Response, error)
	AddToReplicaMap(context.Context, *Node) (*Void, error)
}

// UnimplementedRavelClusterAdminServer can be embedded to have forward compatible implementations.
type UnimplementedRavelClusterAdminServer struct {
}

func (*UnimplementedRavelClusterAdminServer) GetLeader(context.Context, *Cluster) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeader not implemented")
}
func (*UnimplementedRavelClusterAdminServer) AddToReplicaMap(context.Context, *Node) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToReplicaMap not implemented")
}

func RegisterRavelClusterAdminServer(s *grpc.Server, srv RavelClusterAdminServer) {
	s.RegisterService(&_RavelClusterAdmin_serviceDesc, srv)
}

func _RavelClusterAdmin_GetLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RavelClusterAdminServer).GetLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RavelClusterAdminPB.RavelClusterAdmin/GetLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RavelClusterAdminServer).GetLeader(ctx, req.(*Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _RavelClusterAdmin_AddToReplicaMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RavelClusterAdminServer).AddToReplicaMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RavelClusterAdminPB.RavelClusterAdmin/AddToReplicaMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RavelClusterAdminServer).AddToReplicaMap(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

var _RavelClusterAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RavelClusterAdminPB.RavelClusterAdmin",
	HandlerType: (*RavelClusterAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLeader",
			Handler:    _RavelClusterAdmin_GetLeader_Handler,
		},
		{
			MethodName: "AddToReplicaMap",
			Handler:    _RavelClusterAdmin_AddToReplicaMap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ravel/cmd/ravel_cluster_admin/clusterAdmin.proto",
}
