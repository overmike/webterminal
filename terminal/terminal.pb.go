// Code generated by protoc-gen-go. DO NOT EDIT.
// source: terminal.proto

package terminal

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SessionRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionRequest) Reset()         { *m = SessionRequest{} }
func (m *SessionRequest) String() string { return proto.CompactTextString(m) }
func (*SessionRequest) ProtoMessage()    {}
func (*SessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff8b8260c8ef16ad, []int{0}
}

func (m *SessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionRequest.Unmarshal(m, b)
}
func (m *SessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionRequest.Marshal(b, m, deterministic)
}
func (m *SessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionRequest.Merge(m, src)
}
func (m *SessionRequest) XXX_Size() int {
	return xxx_messageInfo_SessionRequest.Size(m)
}
func (m *SessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SessionRequest proto.InternalMessageInfo

func (m *SessionRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SessionResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionResponse) Reset()         { *m = SessionResponse{} }
func (m *SessionResponse) String() string { return proto.CompactTextString(m) }
func (*SessionResponse) ProtoMessage()    {}
func (*SessionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff8b8260c8ef16ad, []int{1}
}

func (m *SessionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionResponse.Unmarshal(m, b)
}
func (m *SessionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionResponse.Marshal(b, m, deterministic)
}
func (m *SessionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionResponse.Merge(m, src)
}
func (m *SessionResponse) XXX_Size() int {
	return xxx_messageInfo_SessionResponse.Size(m)
}
func (m *SessionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SessionResponse proto.InternalMessageInfo

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

func init() { proto.RegisterFile("terminal.proto", fileDescriptor_ff8b8260c8ef16ad) }

var fileDescriptor_ff8b8260c8ef16ad = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x49, 0x2d, 0xca,
	0xcd, 0xcc, 0x4b, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x64,
	0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b,
	0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0xea, 0x94, 0xb4, 0xb8, 0xf8, 0x82, 0x53, 0x8b, 0x8b,
	0x33, 0xf3, 0xf3, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0x73, 0x53,
	0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x25, 0x6d,
	0x2e, 0x7e, 0xb8, 0xda, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xdc, 0x8a, 0x8d, 0x12, 0xb9, 0xe0,
	0x4e, 0x10, 0x0a, 0xe5, 0x62, 0x87, 0x6a, 0x14, 0x92, 0xd0, 0x83, 0x3b, 0x14, 0xd5, 0x5e, 0x29,
	0x49, 0x2c, 0x32, 0x10, 0x5b, 0x94, 0x04, 0x9b, 0x2e, 0x3f, 0x99, 0xcc, 0xc4, 0x2d, 0xc4, 0xa9,
	0x0f, 0x53, 0xa2, 0xc1, 0x68, 0xc0, 0x98, 0xc4, 0x06, 0xf6, 0x82, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x8d, 0xde, 0xb2, 0x0f, 0xfc, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TerminalClient is the client API for Terminal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
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
	stream, err := c.cc.NewStream(ctx, &_Terminal_serviceDesc.Streams[0], "/terminal.terminal/Session", opts...)
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

// TerminalServer is the server API for Terminal service.
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
