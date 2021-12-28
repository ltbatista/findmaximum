// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.17.3
// source: max/maxpb/max.proto

package maxpb

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

type Max struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Max) Reset() {
	*x = Max{}
	if protoimpl.UnsafeEnabled {
		mi := &file_max_maxpb_max_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Max) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Max) ProtoMessage() {}

func (x *Max) ProtoReflect() protoreflect.Message {
	mi := &file_max_maxpb_max_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Max.ProtoReflect.Descriptor instead.
func (*Max) Descriptor() ([]byte, []int) {
	return file_max_maxpb_max_proto_rawDescGZIP(), []int{0}
}

func (x *Max) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type MaxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Max *Max `protobuf:"bytes,1,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *MaxRequest) Reset() {
	*x = MaxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_max_maxpb_max_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaxRequest) ProtoMessage() {}

func (x *MaxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_max_maxpb_max_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaxRequest.ProtoReflect.Descriptor instead.
func (*MaxRequest) Descriptor() ([]byte, []int) {
	return file_max_maxpb_max_proto_rawDescGZIP(), []int{1}
}

func (x *MaxRequest) GetMax() *Max {
	if x != nil {
		return x.Max
	}
	return nil
}

type MaxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *MaxResponse) Reset() {
	*x = MaxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_max_maxpb_max_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaxResponse) ProtoMessage() {}

func (x *MaxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_max_maxpb_max_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaxResponse.ProtoReflect.Descriptor instead.
func (*MaxResponse) Descriptor() ([]byte, []int) {
	return file_max_maxpb_max_proto_rawDescGZIP(), []int{2}
}

func (x *MaxResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_max_maxpb_max_proto protoreflect.FileDescriptor

var file_max_maxpb_max_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x61, 0x78, 0x2f, 0x6d, 0x61, 0x78, 0x70, 0x62, 0x2f, 0x6d, 0x61, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x61, 0x78, 0x22, 0x1d, 0x0a, 0x03, 0x4d, 0x61,
	0x78, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x28, 0x0a, 0x0a, 0x4d, 0x61, 0x78,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x6d, 0x61, 0x78, 0x2e, 0x4d, 0x61, 0x78, 0x52, 0x03,
	0x6d, 0x61, 0x78, 0x22, 0x25, 0x0a, 0x0b, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x44, 0x0a, 0x0a, 0x4d, 0x61,
	0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64,
	0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x0f, 0x2e, 0x6d, 0x61, 0x78, 0x2e, 0x4d, 0x61,
	0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x6d, 0x61, 0x78, 0x2e, 0x4d,
	0x61, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x6d, 0x61, 0x78, 0x2f, 0x6d, 0x61, 0x78, 0x70, 0x62, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_max_maxpb_max_proto_rawDescOnce sync.Once
	file_max_maxpb_max_proto_rawDescData = file_max_maxpb_max_proto_rawDesc
)

func file_max_maxpb_max_proto_rawDescGZIP() []byte {
	file_max_maxpb_max_proto_rawDescOnce.Do(func() {
		file_max_maxpb_max_proto_rawDescData = protoimpl.X.CompressGZIP(file_max_maxpb_max_proto_rawDescData)
	})
	return file_max_maxpb_max_proto_rawDescData
}

var file_max_maxpb_max_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_max_maxpb_max_proto_goTypes = []interface{}{
	(*Max)(nil),         // 0: max.Max
	(*MaxRequest)(nil),  // 1: max.MaxRequest
	(*MaxResponse)(nil), // 2: max.MaxResponse
}
var file_max_maxpb_max_proto_depIdxs = []int32{
	0, // 0: max.MaxRequest.max:type_name -> max.Max
	1, // 1: max.MaxService.FindMaximum:input_type -> max.MaxRequest
	2, // 2: max.MaxService.FindMaximum:output_type -> max.MaxResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_max_maxpb_max_proto_init() }
func file_max_maxpb_max_proto_init() {
	if File_max_maxpb_max_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_max_maxpb_max_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Max); i {
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
		file_max_maxpb_max_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaxRequest); i {
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
		file_max_maxpb_max_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaxResponse); i {
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
			RawDescriptor: file_max_maxpb_max_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_max_maxpb_max_proto_goTypes,
		DependencyIndexes: file_max_maxpb_max_proto_depIdxs,
		MessageInfos:      file_max_maxpb_max_proto_msgTypes,
	}.Build()
	File_max_maxpb_max_proto = out.File
	file_max_maxpb_max_proto_rawDesc = nil
	file_max_maxpb_max_proto_goTypes = nil
	file_max_maxpb_max_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MaxServiceClient is the client API for MaxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MaxServiceClient interface {
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (MaxService_FindMaximumClient, error)
}

type maxServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMaxServiceClient(cc grpc.ClientConnInterface) MaxServiceClient {
	return &maxServiceClient{cc}
}

func (c *maxServiceClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (MaxService_FindMaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MaxService_serviceDesc.Streams[0], "/max.MaxService/FindMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &maxServiceFindMaximumClient{stream}
	return x, nil
}

type MaxService_FindMaximumClient interface {
	Send(*MaxRequest) error
	Recv() (*MaxResponse, error)
	grpc.ClientStream
}

type maxServiceFindMaximumClient struct {
	grpc.ClientStream
}

func (x *maxServiceFindMaximumClient) Send(m *MaxRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *maxServiceFindMaximumClient) Recv() (*MaxResponse, error) {
	m := new(MaxResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MaxServiceServer is the server API for MaxService service.
type MaxServiceServer interface {
	FindMaximum(MaxService_FindMaximumServer) error
}

// UnimplementedMaxServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMaxServiceServer struct {
}

func (*UnimplementedMaxServiceServer) FindMaximum(MaxService_FindMaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaximum not implemented")
}

func RegisterMaxServiceServer(s *grpc.Server, srv MaxServiceServer) {
	s.RegisterService(&_MaxService_serviceDesc, srv)
}

func _MaxService_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MaxServiceServer).FindMaximum(&maxServiceFindMaximumServer{stream})
}

type MaxService_FindMaximumServer interface {
	Send(*MaxResponse) error
	Recv() (*MaxRequest, error)
	grpc.ServerStream
}

type maxServiceFindMaximumServer struct {
	grpc.ServerStream
}

func (x *maxServiceFindMaximumServer) Send(m *MaxResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *maxServiceFindMaximumServer) Recv() (*MaxRequest, error) {
	m := new(MaxRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MaxService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "max.MaxService",
	HandlerType: (*MaxServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindMaximum",
			Handler:       _MaxService_FindMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "max/maxpb/max.proto",
}
