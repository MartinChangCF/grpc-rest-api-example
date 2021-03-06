// Code generated by protoc-gen-go. DO NOT EDIT.
// source: definition/auth.proto

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
	return fileDescriptor_73d0a0e567fd898b, []int{0}
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
	return fileDescriptor_73d0a0e567fd898b, []int{1}
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
	return fileDescriptor_73d0a0e567fd898b, []int{2}
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
	proto.RegisterType((*LoginRequest)(nil), "auth.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "auth.LoginReply")
	proto.RegisterType((*LogoutReply)(nil), "auth.LogoutReply")
}

func init() {
	proto.RegisterFile("definition/auth.proto", fileDescriptor_73d0a0e567fd898b)
}

var fileDescriptor_73d0a0e567fd898b = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x18, 0x85, 0x95, 0x0a, 0x4a, 0xf8, 0x81, 0x0a, 0x8c, 0xa8, 0x42, 0x60, 0x40, 0x16, 0x48, 0xa8,
	0x43, 0x22, 0x60, 0xeb, 0x06, 0x02, 0xb1, 0x54, 0x0c, 0x19, 0xd9, 0xdc, 0xd4, 0x4d, 0x2c, 0x05,
	0xff, 0x26, 0xf9, 0x0d, 0xca, 0xca, 0x15, 0x38, 0x05, 0xe7, 0xe1, 0x0a, 0x1c, 0x04, 0xc5, 0x49,
	0xaa, 0xb2, 0xb0, 0xf9, 0xf9, 0x3d, 0x7d, 0x7a, 0xff, 0x83, 0xa3, 0x85, 0x5c, 0x2a, 0xad, 0x48,
	0xa1, 0x8e, 0x85, 0xa5, 0x3c, 0x32, 0x25, 0x12, 0xb2, 0x8d, 0xe6, 0x1d, 0x9e, 0x66, 0x88, 0x59,
	0x21, 0x63, 0x61, 0x54, 0x2c, 0xb4, 0x46, 0x12, 0x4d, 0xae, 0x6a, 0x33, 0xe1, 0x49, 0xe7, 0x3a,
	0x35, 0xb7, 0xcb, 0x58, 0xbe, 0x18, 0xaa, 0x5b, 0x93, 0xdf, 0xc3, 0xee, 0x0c, 0x33, 0xa5, 0x13,
	0xf9, 0x6a, 0x65, 0x45, 0x2c, 0x80, 0x2d, 0x91, 0xa6, 0x68, 0x35, 0x05, 0xde, 0x99, 0x77, 0xb9,
	0x9d, 0xf4, 0x92, 0x85, 0xe0, 0x1b, 0x51, 0x55, 0xef, 0x58, 0x2e, 0x82, 0x81, 0xb3, 0x56, 0x9a,
	0x9f, 0x03, 0x74, 0x14, 0x53, 0xd4, 0x6c, 0x0c, 0xc3, 0x52, 0x56, 0xb6, 0x68, 0x11, 0x7e, 0xd2,
	0x29, 0x7e, 0x01, 0x3b, 0x33, 0xcc, 0xd0, 0xd2, 0xbf, 0xb1, 0xeb, 0x2f, 0x0f, 0x46, 0xb7, 0x96,
	0x72, 0xa9, 0x49, 0xa5, 0xee, 0x12, 0xf6, 0x08, 0x9b, 0x8e, 0xcf, 0x58, 0xe4, 0x8e, 0x5f, 0xaf,
	0x1c, 0xee, 0xff, 0xf9, 0x33, 0x45, 0xcd, 0x83, 0x8f, 0xef, 0x9f, 0xcf, 0x01, 0xe3, 0x7b, 0x6e,
	0x91, 0xb7, 0xab, 0xb8, 0x68, 0xbc, 0xa9, 0x37, 0x61, 0x4f, 0x30, 0x6c, 0x2b, 0xb0, 0x71, 0xd4,
	0xce, 0x12, 0xf5, 0xb3, 0x44, 0x0f, 0xcd, 0x2c, 0xe1, 0xc1, 0x8a, 0xd6, 0x17, 0xe5, 0xc7, 0x0e,
	0x77, 0xc8, 0x47, 0x6b, 0x38, 0xb4, 0x34, 0xf5, 0x26, 0x77, 0xf0, 0xec, 0xbb, 0xd1, 0x2d, 0xe5,
	0xf3, 0xa1, 0x23, 0xdd, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0x03, 0x27, 0x27, 0x96, 0xab, 0x01,
	0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/auth.Authentication/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) Logout(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*LogoutReply, error) {
	out := new(LogoutReply)
	err := c.cc.Invoke(ctx, "/auth.Authentication/Logout", in, out, opts...)
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
		FullMethod: "/auth.Authentication/Login",
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
		FullMethod: "/auth.Authentication/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Logout(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Authentication",
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
	Metadata: "definition/auth.proto",
}
