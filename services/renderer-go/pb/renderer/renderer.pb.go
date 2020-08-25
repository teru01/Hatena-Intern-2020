// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: renderer.proto

package renderer

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RenderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Src string `protobuf:"bytes,1,opt,name=src,proto3" json:"src,omitempty"`
}

func (x *RenderRequest) Reset() {
	*x = RenderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_renderer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenderRequest) ProtoMessage() {}

func (x *RenderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_renderer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenderRequest.ProtoReflect.Descriptor instead.
func (*RenderRequest) Descriptor() ([]byte, []int) {
	return file_renderer_proto_rawDescGZIP(), []int{0}
}

func (x *RenderRequest) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

type RenderReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Html string `protobuf:"bytes,1,opt,name=html,proto3" json:"html,omitempty"`
}

func (x *RenderReply) Reset() {
	*x = RenderReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_renderer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenderReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenderReply) ProtoMessage() {}

func (x *RenderReply) ProtoReflect() protoreflect.Message {
	mi := &file_renderer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenderReply.ProtoReflect.Descriptor instead.
func (*RenderReply) Descriptor() ([]byte, []int) {
	return file_renderer_proto_rawDescGZIP(), []int{1}
}

func (x *RenderReply) GetHtml() string {
	if x != nil {
		return x.Html
	}
	return ""
}

var File_renderer_proto protoreflect.FileDescriptor

var file_renderer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x72, 0x22, 0x21, 0x0a, 0x0d, 0x52, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x72, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x72, 0x63, 0x22, 0x21, 0x0a,
	0x0b, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x74, 0x6d, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x74, 0x6d, 0x6c,
	0x32, 0x44, 0x0a, 0x08, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x06,
	0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x74, 0x65, 0x6e, 0x61, 0x2f, 0x48, 0x61, 0x74, 0x65,
	0x6e, 0x61, 0x2d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x2d, 0x32, 0x30, 0x32, 0x30, 0x2f, 0x70,
	0x62, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_renderer_proto_rawDescOnce sync.Once
	file_renderer_proto_rawDescData = file_renderer_proto_rawDesc
)

func file_renderer_proto_rawDescGZIP() []byte {
	file_renderer_proto_rawDescOnce.Do(func() {
		file_renderer_proto_rawDescData = protoimpl.X.CompressGZIP(file_renderer_proto_rawDescData)
	})
	return file_renderer_proto_rawDescData
}

var file_renderer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_renderer_proto_goTypes = []interface{}{
	(*RenderRequest)(nil), // 0: renderer.RenderRequest
	(*RenderReply)(nil),   // 1: renderer.RenderReply
}
var file_renderer_proto_depIdxs = []int32{
	0, // 0: renderer.Renderer.Render:input_type -> renderer.RenderRequest
	1, // 1: renderer.Renderer.Render:output_type -> renderer.RenderReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_renderer_proto_init() }
func file_renderer_proto_init() {
	if File_renderer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_renderer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenderRequest); i {
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
		file_renderer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenderReply); i {
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
			RawDescriptor: file_renderer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_renderer_proto_goTypes,
		DependencyIndexes: file_renderer_proto_depIdxs,
		MessageInfos:      file_renderer_proto_msgTypes,
	}.Build()
	File_renderer_proto = out.File
	file_renderer_proto_rawDesc = nil
	file_renderer_proto_goTypes = nil
	file_renderer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RendererClient is the client API for Renderer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RendererClient interface {
	Render(ctx context.Context, in *RenderRequest, opts ...grpc.CallOption) (*RenderReply, error)
}

type rendererClient struct {
	cc grpc.ClientConnInterface
}

func NewRendererClient(cc grpc.ClientConnInterface) RendererClient {
	return &rendererClient{cc}
}

func (c *rendererClient) Render(ctx context.Context, in *RenderRequest, opts ...grpc.CallOption) (*RenderReply, error) {
	out := new(RenderReply)
	err := c.cc.Invoke(ctx, "/renderer.Renderer/Render", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RendererServer is the server API for Renderer service.
type RendererServer interface {
	Render(context.Context, *RenderRequest) (*RenderReply, error)
}

// UnimplementedRendererServer can be embedded to have forward compatible implementations.
type UnimplementedRendererServer struct {
}

func (*UnimplementedRendererServer) Render(context.Context, *RenderRequest) (*RenderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Render not implemented")
}

func RegisterRendererServer(s *grpc.Server, srv RendererServer) {
	s.RegisterService(&_Renderer_serviceDesc, srv)
}

func _Renderer_Render_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RendererServer).Render(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/renderer.Renderer/Render",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RendererServer).Render(ctx, req.(*RenderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Renderer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "renderer.Renderer",
	HandlerType: (*RendererServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Render",
			Handler:    _Renderer_Render_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "renderer.proto",
}
