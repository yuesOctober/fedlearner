// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pairing_service.proto

package proto

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

type Status struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_508234a4704e8446, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type RegisterRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Role                 string   `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Pairs                []*Pair  `protobuf:"bytes,3,rep,name=pairs,proto3" json:"pairs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_508234a4704e8446, []int{1}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *RegisterRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *RegisterRequest) GetPairs() []*Pair {
	if m != nil {
		return m.Pairs
	}
	return nil
}

type PairRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Pairs                []*Pair  `protobuf:"bytes,2,rep,name=pairs,proto3" json:"pairs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PairRequest) Reset()         { *m = PairRequest{} }
func (m *PairRequest) String() string { return proto.CompactTextString(m) }
func (*PairRequest) ProtoMessage()    {}
func (*PairRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_508234a4704e8446, []int{2}
}

func (m *PairRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PairRequest.Unmarshal(m, b)
}
func (m *PairRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PairRequest.Marshal(b, m, deterministic)
}
func (m *PairRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PairRequest.Merge(m, src)
}
func (m *PairRequest) XXX_Size() int {
	return xxx_messageInfo_PairRequest.Size(m)
}
func (m *PairRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PairRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PairRequest proto.InternalMessageInfo

func (m *PairRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *PairRequest) GetPairs() []*Pair {
	if m != nil {
		return m.Pairs
	}
	return nil
}

type Pair struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	LeaderIds            []string `protobuf:"bytes,2,rep,name=leader_ids,json=leaderIds,proto3" json:"leader_ids,omitempty"`
	FollowerIds          []string `protobuf:"bytes,3,rep,name=follower_ids,json=followerIds,proto3" json:"follower_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_508234a4704e8446, []int{3}
}

func (m *Pair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pair.Unmarshal(m, b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
}
func (m *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(m, src)
}
func (m *Pair) XXX_Size() int {
	return xxx_messageInfo_Pair.Size(m)
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Pair) GetLeaderIds() []string {
	if m != nil {
		return m.LeaderIds
	}
	return nil
}

func (m *Pair) GetFollowerIds() []string {
	if m != nil {
		return m.FollowerIds
	}
	return nil
}

type FinishRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Role                 string   `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FinishRequest) Reset()         { *m = FinishRequest{} }
func (m *FinishRequest) String() string { return proto.CompactTextString(m) }
func (*FinishRequest) ProtoMessage()    {}
func (*FinishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_508234a4704e8446, []int{4}
}

func (m *FinishRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FinishRequest.Unmarshal(m, b)
}
func (m *FinishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FinishRequest.Marshal(b, m, deterministic)
}
func (m *FinishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FinishRequest.Merge(m, src)
}
func (m *FinishRequest) XXX_Size() int {
	return xxx_messageInfo_FinishRequest.Size(m)
}
func (m *FinishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FinishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FinishRequest proto.InternalMessageInfo

func (m *FinishRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *FinishRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type ShutDownRequest struct {
	AppId                string   `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Role                 string   `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShutDownRequest) Reset()         { *m = ShutDownRequest{} }
func (m *ShutDownRequest) String() string { return proto.CompactTextString(m) }
func (*ShutDownRequest) ProtoMessage()    {}
func (*ShutDownRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_508234a4704e8446, []int{5}
}

func (m *ShutDownRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShutDownRequest.Unmarshal(m, b)
}
func (m *ShutDownRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShutDownRequest.Marshal(b, m, deterministic)
}
func (m *ShutDownRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShutDownRequest.Merge(m, src)
}
func (m *ShutDownRequest) XXX_Size() int {
	return xxx_messageInfo_ShutDownRequest.Size(m)
}
func (m *ShutDownRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShutDownRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShutDownRequest proto.InternalMessageInfo

func (m *ShutDownRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *ShutDownRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func init() {
	proto.RegisterType((*Status)(nil), "proto.Status")
	proto.RegisterType((*RegisterRequest)(nil), "proto.RegisterRequest")
	proto.RegisterType((*PairRequest)(nil), "proto.PairRequest")
	proto.RegisterType((*Pair)(nil), "proto.Pair")
	proto.RegisterType((*FinishRequest)(nil), "proto.FinishRequest")
	proto.RegisterType((*ShutDownRequest)(nil), "proto.ShutDownRequest")
}

func init() { proto.RegisterFile("pairing_service.proto", fileDescriptor_508234a4704e8446) }

var fileDescriptor_508234a4704e8446 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x41, 0xab, 0xd3, 0x40,
	0x10, 0x26, 0xaf, 0x2f, 0xe1, 0x65, 0xf2, 0xea, 0x83, 0xc5, 0x27, 0x45, 0x10, 0xfa, 0xe2, 0xa5,
	0x1e, 0x6c, 0x68, 0xbd, 0xa9, 0x17, 0x45, 0xaa, 0x3d, 0x08, 0x25, 0xbd, 0x89, 0x10, 0x36, 0xd9,
	0x69, 0xba, 0x98, 0x66, 0xd7, 0xdd, 0x8d, 0xa5, 0xff, 0xd2, 0x9f, 0x24, 0xd9, 0x4d, 0x8a, 0x06,
	0x15, 0x7a, 0xc9, 0xce, 0x7e, 0x33, 0xf3, 0x65, 0xe6, 0xfb, 0x16, 0xee, 0x25, 0xe5, 0x8a, 0xd7,
	0x65, 0xa6, 0x51, 0xfd, 0xe0, 0x05, 0xce, 0xa5, 0x12, 0x46, 0x10, 0xdf, 0x1e, 0xf1, 0x3b, 0x08,
	0xb6, 0x86, 0x9a, 0x46, 0x13, 0x02, 0xd7, 0x85, 0x60, 0x38, 0xf1, 0xa6, 0xde, 0xcc, 0x4f, 0x6d,
	0x4c, 0x9e, 0xc3, 0x18, 0x95, 0x12, 0x2a, 0x3b, 0xa0, 0xd6, 0xb4, 0xc4, 0xc9, 0xd5, 0xd4, 0x9b,
	0x85, 0xe9, 0xad, 0x05, 0x3f, 0x3b, 0x2c, 0xce, 0xe0, 0x2e, 0xc5, 0x92, 0x6b, 0x83, 0x2a, 0xc5,
	0xef, 0x0d, 0x6a, 0x43, 0xee, 0x21, 0xa0, 0x52, 0x66, 0x9c, 0x59, 0xb6, 0x30, 0xf5, 0xa9, 0x94,
	0x6b, 0xd6, 0xfe, 0x42, 0x89, 0xaa, 0x67, 0xb1, 0x31, 0x79, 0x00, 0xbf, 0x1d, 0x50, 0x4f, 0x46,
	0xd3, 0xd1, 0x2c, 0x5a, 0x46, 0x6e, 0xbc, 0xf9, 0x86, 0x72, 0x95, 0xba, 0x4c, 0xfc, 0x11, 0x22,
	0x7b, 0xfd, 0x3f, 0xf9, 0x99, 0xe8, 0xea, 0x9f, 0x44, 0x5f, 0xe1, 0xba, 0xbd, 0xb6, 0x73, 0x98,
	0x93, 0xc4, 0xae, 0xdf, 0xc6, 0xe4, 0x19, 0x40, 0x85, 0x94, 0xa1, 0xca, 0x38, 0x73, 0x1c, 0x61,
	0x1a, 0x3a, 0x64, 0xcd, 0x34, 0x79, 0x80, 0xdb, 0x9d, 0xa8, 0x2a, 0x71, 0xec, 0x0a, 0x46, 0xb6,
	0x20, 0xea, 0xb1, 0x35, 0xd3, 0xf1, 0x6b, 0x18, 0xaf, 0x78, 0xcd, 0xf5, 0xfe, 0x72, 0x15, 0xe2,
	0xb7, 0x70, 0xb7, 0xdd, 0x37, 0xe6, 0x83, 0x38, 0xd6, 0x97, 0x77, 0x2f, 0x7f, 0x7a, 0xf0, 0x68,
	0xe3, 0x5c, 0xde, 0x3a, 0x93, 0xc9, 0x02, 0x6e, 0x7a, 0x53, 0xc8, 0x93, 0x4e, 0x8a, 0x81, 0x4b,
	0x4f, 0xc7, 0x1d, 0xde, 0x3d, 0x80, 0x17, 0xbd, 0x3a, 0xbf, 0x2b, 0xf7, 0xf7, 0xd2, 0x97, 0x10,
	0xb8, 0x55, 0xc9, 0xe3, 0x2e, 0xf1, 0xc7, 0xe6, 0xc3, 0xf2, 0x05, 0xdc, 0xf4, 0xdb, 0x9d, 0x87,
	0x19, 0xac, 0x3b, 0x68, 0x79, 0xff, 0xe9, 0xcb, 0xaa, 0xe4, 0x66, 0xdf, 0xe4, 0xf3, 0x42, 0x1c,
	0x92, 0xfc, 0x64, 0x90, 0xd1, 0xba, 0xc0, 0x64, 0x87, 0xac, 0x42, 0xaa, 0x6a, 0x54, 0x09, 0x43,
	0x59, 0x89, 0x53, 0xf2, 0xad, 0xc9, 0x51, 0xd5, 0x68, 0x50, 0x67, 0x42, 0xa2, 0xa2, 0x46, 0xa8,
	0xc4, 0x52, 0xbd, 0xb1, 0xdf, 0x3c, 0xb0, 0xc7, 0xab, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7e,
	0x82, 0x5c, 0xb7, 0x08, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PairingServiceClient is the client API for PairingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PairingServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Status, error)
	Pair(ctx context.Context, in *PairRequest, opts ...grpc.CallOption) (*Status, error)
	Finish(ctx context.Context, in *FinishRequest, opts ...grpc.CallOption) (*Status, error)
	ShutDown(ctx context.Context, in *ShutDownRequest, opts ...grpc.CallOption) (*Status, error)
}

type pairingServiceClient struct {
	cc *grpc.ClientConn
}

func NewPairingServiceClient(cc *grpc.ClientConn) PairingServiceClient {
	return &pairingServiceClient{cc}
}

func (c *pairingServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/proto.PairingService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pairingServiceClient) Pair(ctx context.Context, in *PairRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/proto.PairingService/Pair", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pairingServiceClient) Finish(ctx context.Context, in *FinishRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/proto.PairingService/Finish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pairingServiceClient) ShutDown(ctx context.Context, in *ShutDownRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/proto.PairingService/ShutDown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PairingServiceServer is the server API for PairingService service.
type PairingServiceServer interface {
	Register(context.Context, *RegisterRequest) (*Status, error)
	Pair(context.Context, *PairRequest) (*Status, error)
	Finish(context.Context, *FinishRequest) (*Status, error)
	ShutDown(context.Context, *ShutDownRequest) (*Status, error)
}

// UnimplementedPairingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPairingServiceServer struct {
}

func (*UnimplementedPairingServiceServer) Register(ctx context.Context, req *RegisterRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedPairingServiceServer) Pair(ctx context.Context, req *PairRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pair not implemented")
}
func (*UnimplementedPairingServiceServer) Finish(ctx context.Context, req *FinishRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Finish not implemented")
}
func (*UnimplementedPairingServiceServer) ShutDown(ctx context.Context, req *ShutDownRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShutDown not implemented")
}

func RegisterPairingServiceServer(s *grpc.Server, srv PairingServiceServer) {
	s.RegisterService(&_PairingService_serviceDesc, srv)
}

func _PairingService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PairingServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PairingService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PairingServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PairingService_Pair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PairRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PairingServiceServer).Pair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PairingService/Pair",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PairingServiceServer).Pair(ctx, req.(*PairRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PairingService_Finish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PairingServiceServer).Finish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PairingService/Finish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PairingServiceServer).Finish(ctx, req.(*FinishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PairingService_ShutDown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutDownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PairingServiceServer).ShutDown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PairingService/ShutDown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PairingServiceServer).ShutDown(ctx, req.(*ShutDownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PairingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PairingService",
	HandlerType: (*PairingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _PairingService_Register_Handler,
		},
		{
			MethodName: "Pair",
			Handler:    _PairingService_Pair_Handler,
		},
		{
			MethodName: "Finish",
			Handler:    _PairingService_Finish_Handler,
		},
		{
			MethodName: "ShutDown",
			Handler:    _PairingService_ShutDown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pairing_service.proto",
}
