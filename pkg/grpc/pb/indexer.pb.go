// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: github.com/dipdup-io/starknet-indexer/pkg/grpc/proto/indexer.proto

package pb

import (
	pb "github.com/dipdup-net/indexer-sdk/pkg/modules/grpc/pb"
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

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartBlock uint64     `protobuf:"varint,1,opt,name=start_block,json=startBlock,proto3" json:"start_block,omitempty"`
	Head       bool       `protobuf:"varint,2,opt,name=head,proto3" json:"head,omitempty"`
	Txs        *TxRequest `protobuf:"bytes,3,opt,name=txs,proto3" json:"txs,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribeRequest) GetStartBlock() uint64 {
	if x != nil {
		return x.StartBlock
	}
	return 0
}

func (x *SubscribeRequest) GetHead() bool {
	if x != nil {
		return x.Head
	}
	return false
}

func (x *SubscribeRequest) GetTxs() *TxRequest {
	if x != nil {
		return x.Txs
	}
	return nil
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response   *pb.SubscribeResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Block      *Block                `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	EndOfBlock bool                  `protobuf:"varint,3,opt,name=end_of_block,json=endOfBlock,proto3" json:"end_of_block,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDescGZIP(), []int{1}
}

func (x *Subscription) GetResponse() *pb.SubscribeResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *Subscription) GetBlock() *Block {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *Subscription) GetEndOfBlock() bool {
	if x != nil {
		return x.EndOfBlock
	}
	return false
}

var File_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto protoreflect.FileDescriptor

var file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x70,
	0x64, 0x75, 0x70, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x6e, 0x65, 0x74, 0x2d,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x46, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x70, 0x64, 0x75, 0x70, 0x2d, 0x6e,
	0x65, 0x74, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x69, 0x70, 0x64, 0x75, 0x70, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x6e,
	0x65, 0x74, 0x2d, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x43,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x70, 0x64, 0x75,
	0x70, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x6e, 0x65, 0x74, 0x2d, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64, 0x12, 0x22, 0x0a, 0x03,
	0x74, 0x78, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x03, 0x74, 0x78, 0x73,
	0x22, 0x8a, 0x01, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x34, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x20, 0x0a, 0x0c, 0x65,
	0x6e, 0x64, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x65, 0x6e, 0x64, 0x4f, 0x66, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x32, 0x93, 0x01,
	0x0a, 0x0e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3b, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x30, 0x01, 0x12, 0x44, 0x0a,
	0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x69, 0x70, 0x64, 0x75, 0x70, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x72,
	0x6b, 0x6e, 0x65, 0x74, 0x2d, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDescOnce sync.Once
	file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDescData = file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDesc
)

func file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDescGZIP() []byte {
	file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDescOnce.Do(func() {
		file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDescData)
	})
	return file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDescData
}

var file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_goTypes = []interface{}{
	(*SubscribeRequest)(nil),       // 0: proto.SubscribeRequest
	(*Subscription)(nil),           // 1: proto.Subscription
	(*TxRequest)(nil),              // 2: proto.TxRequest
	(*pb.SubscribeResponse)(nil),   // 3: proto.SubscribeResponse
	(*Block)(nil),                  // 4: proto.Block
	(*pb.UnsubscribeRequest)(nil),  // 5: proto.UnsubscribeRequest
	(*pb.UnsubscribeResponse)(nil), // 6: proto.UnsubscribeResponse
}
var file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_depIdxs = []int32{
	2, // 0: proto.SubscribeRequest.txs:type_name -> proto.TxRequest
	3, // 1: proto.Subscription.response:type_name -> proto.SubscribeResponse
	4, // 2: proto.Subscription.block:type_name -> proto.Block
	0, // 3: proto.IndexerService.Subscribe:input_type -> proto.SubscribeRequest
	5, // 4: proto.IndexerService.Unsubscribe:input_type -> proto.UnsubscribeRequest
	1, // 5: proto.IndexerService.Subscribe:output_type -> proto.Subscription
	6, // 6: proto.IndexerService.Unsubscribe:output_type -> proto.UnsubscribeResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_init() }
func file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_init() {
	if File_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto != nil {
		return
	}
	file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_entity_filters_proto_init()
	file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_response_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
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
			RawDescriptor: file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_goTypes,
		DependencyIndexes: file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_depIdxs,
		MessageInfos:      file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes,
	}.Build()
	File_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto = out.File
	file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDesc = nil
	file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_goTypes = nil
	file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_depIdxs = nil
}
