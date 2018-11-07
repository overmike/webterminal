// Code generated by protoc-gen-go. DO NOT EDIT.
// source: terminal.proto

/*
Package terminal is a generated protocol buffer package.

It is generated from these files:
	terminal.proto

It has these top-level messages:
	SessionRequest
	SessionResponse
*/
package terminal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SessionRequest struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *SessionRequest) Reset()                    { *m = SessionRequest{} }
func (m *SessionRequest) String() string            { return proto.CompactTextString(m) }
func (*SessionRequest) ProtoMessage()               {}
func (*SessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SessionRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SessionResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *SessionResponse) Reset()                    { *m = SessionResponse{} }
func (m *SessionResponse) String() string            { return proto.CompactTextString(m) }
func (*SessionResponse) ProtoMessage()               {}
func (*SessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SessionResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SessionRequest)(nil), "terminal.SessionRequest")
	proto.RegisterType((*SessionResponse)(nil), "terminal.SessionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Terminal service

type TerminalClient interface {
	Session(ctx context.Context, opts ...grpc.CallOption) (Terminal_SessionClient, error)
}

type terminalClient struct {
	cc *grpc.ClientConn
}

func NewTerminalClient(cc *grpc.ClientConn) TerminalClient {
	return &terminalClient{cc}
}

func (c *terminalClient) Session(ctx context.Context, opts ...grpc.CallOption) (Terminal_SessionClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Terminal_serviceDesc.Streams[0], c.cc, "/terminal.terminal/Session", opts...)
	if err != nil {
		return nil, err
	}
	x := &terminalSessionClient{stream}
	return x, nil
}

type Terminal_SessionClient interface {
	Send(*SessionRequest) error
	Recv() (*SessionResponse, error)
	grpc.ClientStream
}

type terminalSessionClient struct {
	grpc.ClientStream
}

func (x *terminalSessionClient) Send(m *SessionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *terminalSessionClient) Recv() (*SessionResponse, error) {
	m := new(SessionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Terminal service

type TerminalServer interface {
	Session(Terminal_SessionServer) error
}

func RegisterTerminalServer(s *grpc.Server, srv TerminalServer) {
	s.RegisterService(&_Terminal_serviceDesc, srv)
}

func _Terminal_Session_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TerminalServer).Session(&terminalSessionServer{stream})
}

type Terminal_SessionServer interface {
	Send(*SessionResponse) error
	Recv() (*SessionRequest, error)
	grpc.ServerStream
}

type terminalSessionServer struct {
	grpc.ServerStream
}

func (x *terminalSessionServer) Send(m *SessionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *terminalSessionServer) Recv() (*SessionRequest, error) {
	m := new(SessionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Terminal_serviceDesc = grpc.ServiceDesc{
	ServiceName: "terminal.terminal",
	HandlerType: (*TerminalServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Session",
			Handler:       _Terminal_Session_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "terminal.proto",
}

func init() { proto.RegisterFile("terminal.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x49, 0x2d, 0xca,
	0xcd, 0xcc, 0x4b, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0xb4,
	0xb8, 0xf8, 0x82, 0x53, 0x8b, 0x8b, 0x33, 0xf3, 0xf3, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b,
	0x84, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0x60, 0x5c, 0x25, 0x6d, 0x2e, 0x7e, 0xb8, 0xda, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54,
	0xdc, 0x8a, 0x8d, 0xfc, 0xb8, 0xe0, 0x96, 0x08, 0x39, 0x71, 0xb1, 0x43, 0x35, 0x0a, 0x49, 0xe8,
	0xc1, 0x9d, 0x82, 0x6a, 0xaf, 0x94, 0x24, 0x16, 0x19, 0x88, 0x2d, 0x1a, 0x8c, 0x06, 0x8c, 0x49,
	0x6c, 0x60, 0x97, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x69, 0xa6, 0xd5, 0x08, 0xcb, 0x00,
	0x00, 0x00,
}