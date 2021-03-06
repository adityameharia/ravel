// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.0
// source: ravel/cmd/ravel_node/ravel_node.proto

package RavelNodePB

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

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId      string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	GrpcAddress string `protobuf:"bytes,2,opt,name=grpc_address,json=grpcAddress,proto3" json:"grpc_address,omitempty"`
	RaftAddress string `protobuf:"bytes,3,opt,name=raft_address,json=raftAddress,proto3" json:"raft_address,omitempty"`
	ClusterId   string `protobuf:"bytes,4,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ravel_cmd_ravel_node_ravel_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_ravel_cmd_ravel_node_ravel_node_proto_msgTypes[0]
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
	return file_ravel_cmd_ravel_node_ravel_node_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *Node) GetGrpcAddress() string {
	if x != nil {
		return x.GrpcAddress
	}
	return ""
}

func (x *Node) GetRaftAddress() string {
	if x != nil {
		return x.RaftAddress
	}
	return ""
}

func (x *Node) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ravel_cmd_ravel_node_ravel_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_ravel_cmd_ravel_node_ravel_node_proto_msgTypes[1]
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
	return file_ravel_cmd_ravel_node_ravel_node_proto_rawDescGZIP(), []int{1}
}

type Boolean struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Leader bool `protobuf:"varint,1,opt,name=leader,proto3" json:"leader,omitempty"`
}

func (x *Boolean) Reset() {
	*x = Boolean{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ravel_cmd_ravel_node_ravel_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Boolean) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Boolean) ProtoMessage() {}

func (x *Boolean) ProtoReflect() protoreflect.Message {
	mi := &file_ravel_cmd_ravel_node_ravel_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Boolean.ProtoReflect.Descriptor instead.
func (*Boolean) Descriptor() ([]byte, []int) {
	return file_ravel_cmd_ravel_node_ravel_node_proto_rawDescGZIP(), []int{2}
}

func (x *Boolean) GetLeader() bool {
	if x != nil {
		return x.Leader
	}
	return false
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ravel_cmd_ravel_node_ravel_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_ravel_cmd_ravel_node_ravel_node_proto_msgTypes[3]
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
	return file_ravel_cmd_ravel_node_ravel_node_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Response) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation string `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	Key       []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value     []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ravel_cmd_ravel_node_ravel_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_ravel_cmd_ravel_node_ravel_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_ravel_cmd_ravel_node_ravel_node_proto_rawDescGZIP(), []int{4}
}

func (x *Command) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *Command) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Command) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_ravel_cmd_ravel_node_ravel_node_proto protoreflect.FileDescriptor

var file_ravel_cmd_ravel_node_ravel_node_proto_rawDesc = []byte{
	0x0a, 0x25, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x72, 0x61, 0x76, 0x65,
	0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x5f, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x52, 0x61, 0x76, 0x65, 0x6c, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x42, 0x22, 0x84, 0x01, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x67, 0x72, 0x70, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c,
	0x72, 0x61, 0x66, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x72, 0x61, 0x66, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x06,
	0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x30, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4f, 0x0a, 0x07, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xe9, 0x01, 0x0a,
	0x09, 0x52, 0x61, 0x76, 0x65, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x4a, 0x6f,
	0x69, 0x6e, 0x12, 0x14, 0x2e, 0x52, 0x61, 0x76, 0x65, 0x6c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x50, 0x42, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x14, 0x2e, 0x52, 0x61, 0x76, 0x65, 0x6c,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x42, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x33,
	0x0a, 0x05, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x14, 0x2e, 0x52, 0x61, 0x76, 0x65, 0x6c, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x42, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x1a, 0x14, 0x2e,
	0x52, 0x61, 0x76, 0x65, 0x6c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x42, 0x2e, 0x56,
	0x6f, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x17, 0x2e, 0x52, 0x61, 0x76,
	0x65, 0x6c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x42, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x1a, 0x18, 0x2e, 0x52, 0x61, 0x76, 0x65, 0x6c, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x50, 0x42, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a,
	0x08, 0x49, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x52, 0x61, 0x76, 0x65,
	0x6c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x42, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a,
	0x17, 0x2e, 0x52, 0x61, 0x76, 0x65, 0x6c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x42,
	0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x52, 0x61,
	0x76, 0x65, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x42, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ravel_cmd_ravel_node_ravel_node_proto_rawDescOnce sync.Once
	file_ravel_cmd_ravel_node_ravel_node_proto_rawDescData = file_ravel_cmd_ravel_node_ravel_node_proto_rawDesc
)

func file_ravel_cmd_ravel_node_ravel_node_proto_rawDescGZIP() []byte {
	file_ravel_cmd_ravel_node_ravel_node_proto_rawDescOnce.Do(func() {
		file_ravel_cmd_ravel_node_ravel_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_ravel_cmd_ravel_node_ravel_node_proto_rawDescData)
	})
	return file_ravel_cmd_ravel_node_ravel_node_proto_rawDescData
}

var file_ravel_cmd_ravel_node_ravel_node_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ravel_cmd_ravel_node_ravel_node_proto_goTypes = []interface{}{
	(*Node)(nil),     // 0: RavelClusterPB.Node
	(*Void)(nil),     // 1: RavelClusterPB.Void
	(*Boolean)(nil),  // 2: RavelClusterPB.Boolean
	(*Response)(nil), // 3: RavelClusterPB.Response
	(*Command)(nil),  // 4: RavelClusterPB.Command
}
var file_ravel_cmd_ravel_node_ravel_node_proto_depIdxs = []int32{
	0, // 0: RavelClusterPB.RavelNode.Join:input_type -> RavelClusterPB.Node
	0, // 1: RavelClusterPB.RavelNode.Leave:input_type -> RavelClusterPB.Node
	4, // 2: RavelClusterPB.RavelNode.Run:input_type -> RavelClusterPB.Command
	1, // 3: RavelClusterPB.RavelNode.IsLeader:input_type -> RavelClusterPB.Void
	1, // 4: RavelClusterPB.RavelNode.Join:output_type -> RavelClusterPB.Void
	1, // 5: RavelClusterPB.RavelNode.Leave:output_type -> RavelClusterPB.Void
	3, // 6: RavelClusterPB.RavelNode.Run:output_type -> RavelClusterPB.Response
	2, // 7: RavelClusterPB.RavelNode.IsLeader:output_type -> RavelClusterPB.Boolean
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ravel_cmd_ravel_node_ravel_node_proto_init() }
func file_ravel_cmd_ravel_node_ravel_node_proto_init() {
	if File_ravel_cmd_ravel_node_ravel_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ravel_cmd_ravel_node_ravel_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_ravel_cmd_ravel_node_ravel_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_ravel_cmd_ravel_node_ravel_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Boolean); i {
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
		file_ravel_cmd_ravel_node_ravel_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_ravel_cmd_ravel_node_ravel_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
			RawDescriptor: file_ravel_cmd_ravel_node_ravel_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ravel_cmd_ravel_node_ravel_node_proto_goTypes,
		DependencyIndexes: file_ravel_cmd_ravel_node_ravel_node_proto_depIdxs,
		MessageInfos:      file_ravel_cmd_ravel_node_ravel_node_proto_msgTypes,
	}.Build()
	File_ravel_cmd_ravel_node_ravel_node_proto = out.File
	file_ravel_cmd_ravel_node_ravel_node_proto_rawDesc = nil
	file_ravel_cmd_ravel_node_ravel_node_proto_goTypes = nil
	file_ravel_cmd_ravel_node_ravel_node_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RavelNodeClient is the client API for RavelNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RavelNodeClient interface {
	Join(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Void, error)
	Leave(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Void, error)
	Run(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Response, error)
	IsLeader(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Boolean, error)
}

type ravelNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewRavelNodeClient(cc grpc.ClientConnInterface) RavelNodeClient {
	return &ravelNodeClient{cc}
}

func (c *ravelNodeClient) Join(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/RavelClusterPB.RavelNode/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ravelNodeClient) Leave(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/RavelClusterPB.RavelNode/Leave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ravelNodeClient) Run(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/RavelClusterPB.RavelNode/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ravelNodeClient) IsLeader(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Boolean, error) {
	out := new(Boolean)
	err := c.cc.Invoke(ctx, "/RavelClusterPB.RavelNode/IsLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RavelNodeServer is the server API for RavelNode service.
type RavelNodeServer interface {
	Join(context.Context, *Node) (*Void, error)
	Leave(context.Context, *Node) (*Void, error)
	Run(context.Context, *Command) (*Response, error)
	IsLeader(context.Context, *Void) (*Boolean, error)
}

// UnimplementedRavelNodeServer can be embedded to have forward compatible implementations.
type UnimplementedRavelNodeServer struct {
}

func (*UnimplementedRavelNodeServer) Join(context.Context, *Node) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (*UnimplementedRavelNodeServer) Leave(context.Context, *Node) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Leave not implemented")
}
func (*UnimplementedRavelNodeServer) Run(context.Context, *Command) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (*UnimplementedRavelNodeServer) IsLeader(context.Context, *Void) (*Boolean, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsLeader not implemented")
}

func RegisterRavelNodeServer(s *grpc.Server, srv RavelNodeServer) {
	s.RegisterService(&_RavelNode_serviceDesc, srv)
}

func _RavelNode_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RavelNodeServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RavelClusterPB.RavelNode/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RavelNodeServer).Join(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _RavelNode_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RavelNodeServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RavelClusterPB.RavelNode/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RavelNodeServer).Leave(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _RavelNode_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RavelNodeServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RavelClusterPB.RavelNode/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RavelNodeServer).Run(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _RavelNode_IsLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RavelNodeServer).IsLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RavelClusterPB.RavelNode/IsLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RavelNodeServer).IsLeader(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _RavelNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RavelClusterPB.RavelNode",
	HandlerType: (*RavelNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _RavelNode_Join_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _RavelNode_Leave_Handler,
		},
		{
			MethodName: "Run",
			Handler:    _RavelNode_Run_Handler,
		},
		{
			MethodName: "IsLeader",
			Handler:    _RavelNode_IsLeader_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ravel/cmd/ravel_node/ravel_node.proto",
}
