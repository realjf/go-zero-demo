// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transform.proto

package transform

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type LoginReq struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb4a498eeb2ba07d, []int{0}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReq) Reset()         { *m = RegisterReq{} }
func (m *RegisterReq) String() string { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()    {}
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb4a498eeb2ba07d, []int{1}
}

func (m *RegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReq.Unmarshal(m, b)
}
func (m *RegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReq.Marshal(b, m, deterministic)
}
func (m *RegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReq.Merge(m, src)
}
func (m *RegisterReq) XXX_Size() int {
	return xxx_messageInfo_RegisterReq.Size(m)
}
func (m *RegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReq proto.InternalMessageInfo

func (m *RegisterReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RegisterReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserInfoReq struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoReq) Reset()         { *m = UserInfoReq{} }
func (m *UserInfoReq) String() string { return proto.CompactTextString(m) }
func (*UserInfoReq) ProtoMessage()    {}
func (*UserInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb4a498eeb2ba07d, []int{2}
}

func (m *UserInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoReq.Unmarshal(m, b)
}
func (m *UserInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoReq.Marshal(b, m, deterministic)
}
func (m *UserInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoReq.Merge(m, src)
}
func (m *UserInfoReq) XXX_Size() int {
	return xxx_messageInfo_UserInfoReq.Size(m)
}
func (m *UserInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoReq proto.InternalMessageInfo

func (m *UserInfoReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type LoginResp struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResp) Reset()         { *m = LoginResp{} }
func (m *LoginResp) String() string { return proto.CompactTextString(m) }
func (*LoginResp) ProtoMessage()    {}
func (*LoginResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb4a498eeb2ba07d, []int{3}
}

func (m *LoginResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResp.Unmarshal(m, b)
}
func (m *LoginResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResp.Marshal(b, m, deterministic)
}
func (m *LoginResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResp.Merge(m, src)
}
func (m *LoginResp) XXX_Size() int {
	return xxx_messageInfo_LoginResp.Size(m)
}
func (m *LoginResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResp.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResp proto.InternalMessageInfo

func (m *LoginResp) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *LoginResp) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginResp) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type RegisterResp struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResp) Reset()         { *m = RegisterResp{} }
func (m *RegisterResp) String() string { return proto.CompactTextString(m) }
func (*RegisterResp) ProtoMessage()    {}
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb4a498eeb2ba07d, []int{4}
}

func (m *RegisterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResp.Unmarshal(m, b)
}
func (m *RegisterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResp.Marshal(b, m, deterministic)
}
func (m *RegisterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResp.Merge(m, src)
}
func (m *RegisterResp) XXX_Size() int {
	return xxx_messageInfo_RegisterResp.Size(m)
}
func (m *RegisterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResp.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResp proto.InternalMessageInfo

func (m *RegisterResp) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *RegisterResp) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RegisterResp) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UserInfoResp struct {
	Uid                  int64    `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoResp) Reset()         { *m = UserInfoResp{} }
func (m *UserInfoResp) String() string { return proto.CompactTextString(m) }
func (*UserInfoResp) ProtoMessage()    {}
func (*UserInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb4a498eeb2ba07d, []int{5}
}

func (m *UserInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoResp.Unmarshal(m, b)
}
func (m *UserInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoResp.Marshal(b, m, deterministic)
}
func (m *UserInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoResp.Merge(m, src)
}
func (m *UserInfoResp) XXX_Size() int {
	return xxx_messageInfo_UserInfoResp.Size(m)
}
func (m *UserInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoResp proto.InternalMessageInfo

func (m *UserInfoResp) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UserInfoResp) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserInfoResp) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginReq)(nil), "transform.LoginReq")
	proto.RegisterType((*RegisterReq)(nil), "transform.RegisterReq")
	proto.RegisterType((*UserInfoReq)(nil), "transform.UserInfoReq")
	proto.RegisterType((*LoginResp)(nil), "transform.LoginResp")
	proto.RegisterType((*RegisterResp)(nil), "transform.RegisterResp")
	proto.RegisterType((*UserInfoResp)(nil), "transform.UserInfoResp")
}

func init() { proto.RegisterFile("transform.proto", fileDescriptor_cb4a498eeb2ba07d) }

var fileDescriptor_cb4a498eeb2ba07d = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x42, 0x35, 0x99, 0x08, 0xca, 0x58, 0xb4, 0xe4, 0xa2, 0xe4, 0xe4, 0xa9, 0x87,
	0x7a, 0x53, 0x5f, 0x40, 0x10, 0x84, 0x85, 0x9e, 0x3c, 0x45, 0x3a, 0x0d, 0x81, 0x26, 0xbb, 0xce,
	0xa4, 0xf8, 0x6e, 0x3e, 0x9d, 0x24, 0xc6, 0xdd, 0x8d, 0xb5, 0xb7, 0xdc, 0x32, 0x33, 0x99, 0x8f,
	0xff, 0x9f, 0x7f, 0xe1, 0xbc, 0xe5, 0xa2, 0x91, 0xad, 0xe6, 0x7a, 0x69, 0x58, 0xb7, 0x1a, 0x13,
	0xdb, 0xc8, 0x9f, 0x20, 0x7e, 0xd1, 0x65, 0xd5, 0x28, 0xfa, 0xc0, 0x39, 0xcc, 0xa8, 0x2e, 0xaa,
	0xdd, 0x22, 0xb8, 0x0d, 0xee, 0x12, 0xf5, 0x53, 0x60, 0x06, 0xb1, 0x29, 0x44, 0x3e, 0x35, 0x6f,
	0x16, 0x61, 0x3f, 0xb0, 0x75, 0xfe, 0x06, 0xa9, 0xa2, 0xb2, 0x92, 0x96, 0xb8, 0x03, 0x64, 0x10,
	0xef, 0x85, 0xb8, 0x29, 0x6a, 0x1a, 0x18, 0xb6, 0x76, 0xf0, 0xf0, 0x18, 0x3c, 0xfa, 0x03, 0xbf,
	0x81, 0x74, 0x2d, 0xc4, 0xcf, 0xcd, 0x56, 0x77, 0xf0, 0x0b, 0x88, 0xf6, 0xd5, 0xa6, 0xe7, 0x46,
	0xaa, 0xfb, 0xcc, 0x5f, 0x21, 0x19, 0xb4, 0x8b, 0x39, 0x1c, 0x8f, 0xd4, 0x84, 0xc7, 0xd4, 0x44,
	0x9e, 0x9a, 0x5c, 0xc1, 0x99, 0xb3, 0x33, 0x1d, 0xd3, 0xb9, 0x98, 0x86, 0xb9, 0xfa, 0x0a, 0x20,
	0xb5, 0x11, 0x12, 0xe3, 0x0a, 0x66, 0xbb, 0xee, 0x10, 0x78, 0xb9, 0x74, 0x51, 0xff, 0xc6, 0x9a,
	0xcd, 0x0f, 0x9b, 0x62, 0xf0, 0x11, 0x62, 0x1e, 0xbc, 0xe2, 0x95, 0xf7, 0x87, 0x97, 0x67, 0x76,
	0xfd, 0x6f, 0x5f, 0x0c, 0x3e, 0xc0, 0x69, 0x49, 0xed, 0x9a, 0x64, 0xbc, 0xeb, 0xc5, 0x35, 0xda,
	0xf5, 0x0f, 0xf0, 0x7e, 0xd2, 0xbf, 0xc1, 0xfb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x27, 0x72,
	0x1a, 0xaa, 0x96, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransformerClient is the client API for Transformer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransformerClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
	GetUser(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error)
}

type transformerClient struct {
	cc *grpc.ClientConn
}

func NewTransformerClient(cc *grpc.ClientConn) TransformerClient {
	return &transformerClient{cc}
}

func (c *transformerClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/transform.transformer/login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transformerClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, "/transform.transformer/register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transformerClient) GetUser(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	out := new(UserInfoResp)
	err := c.cc.Invoke(ctx, "/transform.transformer/getUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransformerServer is the server API for Transformer service.
type TransformerServer interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	GetUser(context.Context, *UserInfoReq) (*UserInfoResp, error)
}

// UnimplementedTransformerServer can be embedded to have forward compatible implementations.
type UnimplementedTransformerServer struct {
}

func (*UnimplementedTransformerServer) Login(ctx context.Context, req *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedTransformerServer) Register(ctx context.Context, req *RegisterReq) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedTransformerServer) GetUser(ctx context.Context, req *UserInfoReq) (*UserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

func RegisterTransformerServer(s *grpc.Server, srv TransformerServer) {
	s.RegisterService(&_Transformer_serviceDesc, srv)
}

func _Transformer_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransformerServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transform.transformer/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransformerServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transformer_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransformerServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transform.transformer/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransformerServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transformer_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransformerServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transform.transformer/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransformerServer).GetUser(ctx, req.(*UserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transformer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transform.transformer",
	HandlerType: (*TransformerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "login",
			Handler:    _Transformer_Login_Handler,
		},
		{
			MethodName: "register",
			Handler:    _Transformer_Register_Handler,
		},
		{
			MethodName: "getUser",
			Handler:    _Transformer_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transform.proto",
}
