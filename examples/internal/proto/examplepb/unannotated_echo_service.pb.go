// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/internal/proto/examplepb/unannotated_echo_service.proto

// Unannotated Echo Service
// Similar to echo_service.proto but without annotations. See
// unannotated_echo_service.yaml for the equivalent of the annotations in
// gRPC API configuration format.
//
// Echo Service API consists of a single service which returns
// a message.

package examplepb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	grpc "github.com/lfch/grpc"
	codes "github.com/lfch/grpc/codes"
	status "github.com/lfch/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// UnannotatedSimpleMessage represents a simple message sent to the unannotated Echo service.
type UnannotatedSimpleMessage struct {
	// Id represents the message identifier.
	Id                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Num                  int64              `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	Duration             *duration.Duration `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UnannotatedSimpleMessage) Reset()         { *m = UnannotatedSimpleMessage{} }
func (m *UnannotatedSimpleMessage) String() string { return proto.CompactTextString(m) }
func (*UnannotatedSimpleMessage) ProtoMessage()    {}
func (*UnannotatedSimpleMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab29419944ad9f3b, []int{0}
}

func (m *UnannotatedSimpleMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnannotatedSimpleMessage.Unmarshal(m, b)
}
func (m *UnannotatedSimpleMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnannotatedSimpleMessage.Marshal(b, m, deterministic)
}
func (m *UnannotatedSimpleMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnannotatedSimpleMessage.Merge(m, src)
}
func (m *UnannotatedSimpleMessage) XXX_Size() int {
	return xxx_messageInfo_UnannotatedSimpleMessage.Size(m)
}
func (m *UnannotatedSimpleMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_UnannotatedSimpleMessage.DiscardUnknown(m)
}

var xxx_messageInfo_UnannotatedSimpleMessage proto.InternalMessageInfo

func (m *UnannotatedSimpleMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UnannotatedSimpleMessage) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *UnannotatedSimpleMessage) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func init() {
	proto.RegisterType((*UnannotatedSimpleMessage)(nil), "grpc.gateway.examples.internal.examplepb.UnannotatedSimpleMessage")
}

func init() {
	proto.RegisterFile("examples/internal/proto/examplepb/unannotated_echo_service.proto", fileDescriptor_ab29419944ad9f3b)
}

var fileDescriptor_ab29419944ad9f3b = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x92, 0x31, 0x4b, 0xfb, 0x40,
	0x18, 0xc6, 0xb9, 0xe4, 0xcf, 0x9f, 0xf6, 0x2d, 0x88, 0xdc, 0x20, 0xb1, 0x83, 0x04, 0xa7, 0x4c,
	0x6f, 0xa0, 0xe2, 0x2e, 0xa1, 0x8e, 0x2e, 0x29, 0x2e, 0x2e, 0xe5, 0x92, 0xbc, 0xa6, 0x07, 0xe9,
	0x5d, 0xb8, 0xbb, 0xa8, 0xfd, 0x12, 0x0e, 0xe2, 0x57, 0xf3, 0xfb, 0xc8, 0x35, 0x4d, 0x74, 0x11,
	0x1c, 0x1c, 0xba, 0x25, 0xc7, 0x73, 0xbf, 0xdf, 0xf3, 0xc0, 0xc1, 0x0d, 0xbd, 0x88, 0x6d, 0xdb,
	0x90, 0x4d, 0xa5, 0x72, 0x64, 0x94, 0x68, 0xd2, 0xd6, 0x68, 0xa7, 0xd3, 0xc3, 0x79, 0x5b, 0xa4,
	0x9d, 0x12, 0x4a, 0x69, 0x27, 0x1c, 0x55, 0x6b, 0x2a, 0x37, 0x7a, 0x6d, 0xc9, 0x3c, 0xc9, 0x92,
	0x70, 0x1f, 0xe4, 0x49, 0x6d, 0xda, 0x12, 0x6b, 0xe1, 0xe8, 0x59, 0xec, 0x70, 0xc0, 0xe1, 0x80,
	0xc3, 0x11, 0x34, 0xbf, 0xa8, 0xb5, 0xae, 0x1b, 0xea, 0x05, 0x45, 0xf7, 0x98, 0x56, 0x9d, 0x11,
	0x4e, 0x6a, 0xd5, 0x93, 0x2e, 0x2d, 0x44, 0xf7, 0x5f, 0xae, 0x95, 0xf4, 0xd7, 0xee, 0xc8, 0x5a,
	0x51, 0x13, 0x3f, 0x81, 0x40, 0x56, 0x11, 0x8b, 0x59, 0x32, 0xcd, 0x03, 0x59, 0xf1, 0x53, 0x08,
	0x55, 0xb7, 0x8d, 0x82, 0x98, 0x25, 0x61, 0xee, 0x3f, 0xf9, 0x35, 0x4c, 0x06, 0x5e, 0x14, 0xc6,
	0x2c, 0x99, 0x2d, 0xce, 0xb1, 0x17, 0xe2, 0x20, 0xc4, 0xe5, 0x21, 0x90, 0x8f, 0xd1, 0xc5, 0x47,
	0x08, 0x67, 0xdf, 0xac, 0xb7, 0xe5, 0x46, 0xaf, 0xfa, 0x7d, 0xfc, 0x95, 0xc1, 0x3f, 0xff, 0xcf,
	0x33, 0xfc, 0xed, 0x46, 0xfc, 0x69, 0xc0, 0xfc, 0x0f, 0x18, 0xfc, 0x8d, 0xc1, 0xc4, 0x17, 0xca,
	0x74, 0xb5, 0x3b, 0x9a, 0x52, 0xef, 0x0c, 0xc0, 0x97, 0x5a, 0x52, 0x43, 0x8e, 0x8e, 0xa5, 0x56,
	0x36, 0x7b, 0x98, 0x8e, 0xa9, 0xe2, 0xff, 0xfe, 0x05, 0x5c, 0x7d, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xaa, 0x31, 0x7d, 0x8e, 0xee, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UnannotatedEchoServiceClient is the client API for UnannotatedEchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/github.com/lfch/grpc#ClientConn.NewStream.
type UnannotatedEchoServiceClient interface {
	// Echo method receives a simple message and returns it.
	//
	// The message posted as the id parameter will also be
	// returned.
	Echo(ctx context.Context, in *UnannotatedSimpleMessage, opts ...grpc.CallOption) (*UnannotatedSimpleMessage, error)
	// EchoBody method receives a simple message and returns it.
	EchoBody(ctx context.Context, in *UnannotatedSimpleMessage, opts ...grpc.CallOption) (*UnannotatedSimpleMessage, error)
	// EchoDelete method receives a simple message and returns it.
	EchoDelete(ctx context.Context, in *UnannotatedSimpleMessage, opts ...grpc.CallOption) (*UnannotatedSimpleMessage, error)
}

type unannotatedEchoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUnannotatedEchoServiceClient(cc grpc.ClientConnInterface) UnannotatedEchoServiceClient {
	return &unannotatedEchoServiceClient{cc}
}

func (c *unannotatedEchoServiceClient) Echo(ctx context.Context, in *UnannotatedSimpleMessage, opts ...grpc.CallOption) (*UnannotatedSimpleMessage, error) {
	out := new(UnannotatedSimpleMessage)
	err := c.cc.Invoke(ctx, "/grpc.gateway.examples.internal.examplepb.UnannotatedEchoService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unannotatedEchoServiceClient) EchoBody(ctx context.Context, in *UnannotatedSimpleMessage, opts ...grpc.CallOption) (*UnannotatedSimpleMessage, error) {
	out := new(UnannotatedSimpleMessage)
	err := c.cc.Invoke(ctx, "/grpc.gateway.examples.internal.examplepb.UnannotatedEchoService/EchoBody", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unannotatedEchoServiceClient) EchoDelete(ctx context.Context, in *UnannotatedSimpleMessage, opts ...grpc.CallOption) (*UnannotatedSimpleMessage, error) {
	out := new(UnannotatedSimpleMessage)
	err := c.cc.Invoke(ctx, "/grpc.gateway.examples.internal.examplepb.UnannotatedEchoService/EchoDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UnannotatedEchoServiceServer is the server API for UnannotatedEchoService service.
type UnannotatedEchoServiceServer interface {
	// Echo method receives a simple message and returns it.
	//
	// The message posted as the id parameter will also be
	// returned.
	Echo(context.Context, *UnannotatedSimpleMessage) (*UnannotatedSimpleMessage, error)
	// EchoBody method receives a simple message and returns it.
	EchoBody(context.Context, *UnannotatedSimpleMessage) (*UnannotatedSimpleMessage, error)
	// EchoDelete method receives a simple message and returns it.
	EchoDelete(context.Context, *UnannotatedSimpleMessage) (*UnannotatedSimpleMessage, error)
}

// UnimplementedUnannotatedEchoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUnannotatedEchoServiceServer struct {
}

func (*UnimplementedUnannotatedEchoServiceServer) Echo(ctx context.Context, req *UnannotatedSimpleMessage) (*UnannotatedSimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (*UnimplementedUnannotatedEchoServiceServer) EchoBody(ctx context.Context, req *UnannotatedSimpleMessage) (*UnannotatedSimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoBody not implemented")
}
func (*UnimplementedUnannotatedEchoServiceServer) EchoDelete(ctx context.Context, req *UnannotatedSimpleMessage) (*UnannotatedSimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoDelete not implemented")
}

func RegisterUnannotatedEchoServiceServer(s *grpc.Server, srv UnannotatedEchoServiceServer) {
	s.RegisterService(&_UnannotatedEchoService_serviceDesc, srv)
}

func _UnannotatedEchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnannotatedSimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnannotatedEchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.internal.examplepb.UnannotatedEchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnannotatedEchoServiceServer).Echo(ctx, req.(*UnannotatedSimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _UnannotatedEchoService_EchoBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnannotatedSimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnannotatedEchoServiceServer).EchoBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.internal.examplepb.UnannotatedEchoService/EchoBody",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnannotatedEchoServiceServer).EchoBody(ctx, req.(*UnannotatedSimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _UnannotatedEchoService_EchoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnannotatedSimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnannotatedEchoServiceServer).EchoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.internal.examplepb.UnannotatedEchoService/EchoDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnannotatedEchoServiceServer).EchoDelete(ctx, req.(*UnannotatedSimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _UnannotatedEchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.internal.examplepb.UnannotatedEchoService",
	HandlerType: (*UnannotatedEchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _UnannotatedEchoService_Echo_Handler,
		},
		{
			MethodName: "EchoBody",
			Handler:    _UnannotatedEchoService_EchoBody_Handler,
		},
		{
			MethodName: "EchoDelete",
			Handler:    _UnannotatedEchoService_EchoDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "examples/internal/proto/examplepb/unannotated_echo_service.proto",
}
