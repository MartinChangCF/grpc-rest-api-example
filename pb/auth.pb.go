// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/auth.proto

package auth

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type LoginRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c43ab4a3390919ea, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginReply struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReply) Reset()         { *m = LoginReply{} }
func (m *LoginReply) String() string { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()    {}
func (*LoginReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c43ab4a3390919ea, []int{1}
}

func (m *LoginReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReply.Unmarshal(m, b)
}
func (m *LoginReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReply.Marshal(b, m, deterministic)
}
func (m *LoginReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReply.Merge(m, src)
}
func (m *LoginReply) XXX_Size() int {
	return xxx_messageInfo_LoginReply.Size(m)
}
func (m *LoginReply) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReply.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReply proto.InternalMessageInfo

func (m *LoginReply) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type LogoutReply struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutReply) Reset()         { *m = LogoutReply{} }
func (m *LogoutReply) String() string { return proto.CompactTextString(m) }
func (*LogoutReply) ProtoMessage()    {}
func (*LogoutReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c43ab4a3390919ea, []int{2}
}

func (m *LogoutReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutReply.Unmarshal(m, b)
}
func (m *LogoutReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutReply.Marshal(b, m, deterministic)
}
func (m *LogoutReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutReply.Merge(m, src)
}
func (m *LogoutReply) XXX_Size() int {
	return xxx_messageInfo_LogoutReply.Size(m)
}
func (m *LogoutReply) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutReply.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutReply proto.InternalMessageInfo

func (m *LogoutReply) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "LoginReply")
	proto.RegisterType((*LogoutReply)(nil), "LogoutReply")
}

func init() {
	proto.RegisterFile("pb/auth.proto", fileDescriptor_c43ab4a3390919ea)
}

var fileDescriptor_c43ab4a3390919ea = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x95, 0x4a, 0x84, 0x72, 0x6d, 0x3a, 0x18, 0xa9, 0x0a, 0x81, 0x01, 0x59, 0x20, 0xa1,
	0x0e, 0xb6, 0x80, 0xad, 0x12, 0x03, 0x12, 0x0c, 0x48, 0x9d, 0xf2, 0x06, 0x4e, 0x30, 0x69, 0xa4,
	0xe0, 0x33, 0xf1, 0x19, 0xd4, 0x95, 0x57, 0x60, 0xe3, 0xb5, 0x78, 0x05, 0x1e, 0x04, 0xd5, 0x49,
	0x21, 0x13, 0xe3, 0xaf, 0xcf, 0xfe, 0xee, 0xee, 0x87, 0xc4, 0x16, 0x52, 0x79, 0x5a, 0x0b, 0xdb,
	0x22, 0x61, 0x76, 0x52, 0x21, 0x56, 0x8d, 0x96, 0xca, 0xd6, 0x52, 0x19, 0x83, 0xa4, 0xa8, 0x46,
	0xe3, 0x7a, 0x7a, 0xdc, 0xd3, 0x90, 0x0a, 0xff, 0x24, 0xf5, 0xb3, 0xa5, 0x4d, 0x07, 0xf9, 0x1d,
	0x4c, 0x57, 0x58, 0xd5, 0x26, 0xd7, 0x2f, 0x5e, 0x3b, 0x62, 0x29, 0xec, 0xab, 0xb2, 0x44, 0x6f,
	0x28, 0x8d, 0x4e, 0xa3, 0x8b, 0x83, 0x7c, 0x17, 0x59, 0x06, 0x63, 0xab, 0x9c, 0x7b, 0xc3, 0xf6,
	0x31, 0x1d, 0x05, 0xf4, 0x9b, 0xf9, 0x19, 0x40, 0x6f, 0xb1, 0xcd, 0x86, 0xcd, 0x21, 0x6e, 0xb5,
	0xf3, 0x4d, 0xa7, 0x18, 0xe7, 0x7d, 0xe2, 0xe7, 0x30, 0x59, 0x61, 0x85, 0x9e, 0xfe, 0x7d, 0x76,
	0xf5, 0x19, 0xc1, 0xec, 0xd6, 0xd3, 0x5a, 0x1b, 0xaa, 0xcb, 0x70, 0x09, 0xbb, 0x81, 0xbd, 0xe0,
	0x67, 0x89, 0x18, 0x6e, 0x9b, 0x4d, 0xc4, 0xdf, 0x58, 0x9e, 0xbe, 0x7f, 0x7d, 0x7f, 0x8c, 0x18,
	0x4f, 0x42, 0x0f, 0xaf, 0x97, 0xb2, 0xd9, 0xb2, 0x65, 0xb4, 0x60, 0x0f, 0x10, 0x77, 0x83, 0xd9,
	0x5c, 0x74, 0x65, 0x88, 0x5d, 0x19, 0xe2, 0x7e, 0x5b, 0x46, 0x36, 0x15, 0x83, 0xcd, 0xf8, 0x51,
	0x30, 0x1d, 0xf2, 0xd9, 0xc0, 0x84, 0x9e, 0x96, 0xd1, 0xa2, 0x88, 0xc3, 0xc7, 0xeb, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xb8, 0x88, 0xb2, 0x27, 0x82, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthenticationClient is the client API for Authentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	Logout(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*LogoutReply, error)
}

type authenticationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationClient(cc grpc.ClientConnInterface) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/Authentication/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) Logout(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*LogoutReply, error) {
	out := new(LogoutReply)
	err := c.cc.Invoke(ctx, "/Authentication/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServer is the server API for Authentication service.
type AuthenticationServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Logout(context.Context, *empty.Empty) (*LogoutReply, error)
}

// UnimplementedAuthenticationServer can be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServer struct {
}

func (*UnimplementedAuthenticationServer) Login(ctx context.Context, req *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthenticationServer) Logout(ctx context.Context, req *empty.Empty) (*LogoutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}

func RegisterAuthenticationServer(s *grpc.Server, srv AuthenticationServer) {
	s.RegisterService(&_Authentication_serviceDesc, srv)
}

func _Authentication_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Authentication/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Authentication/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Logout(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Authentication_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Authentication_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/auth.proto",
}