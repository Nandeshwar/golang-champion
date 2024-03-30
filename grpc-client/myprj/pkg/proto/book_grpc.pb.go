// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: book.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Login_Token_FullMethodName = "/proto.Login/token"
)

// LoginClient is the client API for Login service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginClient interface {
	Token(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*GetTokenResponse, error)
}

type loginClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginClient(cc grpc.ClientConnInterface) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) Token(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*GetTokenResponse, error) {
	out := new(GetTokenResponse)
	err := c.cc.Invoke(ctx, Login_Token_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServer is the server API for Login service.
// All implementations must embed UnimplementedLoginServer
// for forward compatibility
type LoginServer interface {
	Token(context.Context, *GetTokenReq) (*GetTokenResponse, error)
	mustEmbedUnimplementedLoginServer()
}

// UnimplementedLoginServer must be embedded to have forward compatible implementations.
type UnimplementedLoginServer struct {
}

func (UnimplementedLoginServer) Token(context.Context, *GetTokenReq) (*GetTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Token not implemented")
}
func (UnimplementedLoginServer) mustEmbedUnimplementedLoginServer() {}

// UnsafeLoginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginServer will
// result in compilation errors.
type UnsafeLoginServer interface {
	mustEmbedUnimplementedLoginServer()
}

func RegisterLoginServer(s grpc.ServiceRegistrar, srv LoginServer) {
	s.RegisterService(&Login_ServiceDesc, srv)
}

func _Login_Token_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).Token(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Login_Token_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).Token(ctx, req.(*GetTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Login_ServiceDesc is the grpc.ServiceDesc for Login service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Login_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Login",
	HandlerType: (*LoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "token",
			Handler:    _Login_Token_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book.proto",
}

const (
	Book_GetBookInfo_FullMethodName              = "/proto.Book/GetBookInfo"
	Book_GetBookInfoBidirectional_FullMethodName = "/proto.Book/GetBookInfoBidirectional"
)

// BookClient is the client API for Book service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookClient interface {
	GetBookInfo(ctx context.Context, in *GetBookInfoReq, opts ...grpc.CallOption) (*GetBookInfoResponse, error)
	GetBookInfoBidirectional(ctx context.Context, opts ...grpc.CallOption) (Book_GetBookInfoBidirectionalClient, error)
}

type bookClient struct {
	cc grpc.ClientConnInterface
}

func NewBookClient(cc grpc.ClientConnInterface) BookClient {
	return &bookClient{cc}
}

func (c *bookClient) GetBookInfo(ctx context.Context, in *GetBookInfoReq, opts ...grpc.CallOption) (*GetBookInfoResponse, error) {
	out := new(GetBookInfoResponse)
	err := c.cc.Invoke(ctx, Book_GetBookInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) GetBookInfoBidirectional(ctx context.Context, opts ...grpc.CallOption) (Book_GetBookInfoBidirectionalClient, error) {
	stream, err := c.cc.NewStream(ctx, &Book_ServiceDesc.Streams[0], Book_GetBookInfoBidirectional_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &bookGetBookInfoBidirectionalClient{stream}
	return x, nil
}

type Book_GetBookInfoBidirectionalClient interface {
	Send(*GetBookInfoReq) error
	Recv() (*GetBookInfoResponse, error)
	grpc.ClientStream
}

type bookGetBookInfoBidirectionalClient struct {
	grpc.ClientStream
}

func (x *bookGetBookInfoBidirectionalClient) Send(m *GetBookInfoReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bookGetBookInfoBidirectionalClient) Recv() (*GetBookInfoResponse, error) {
	m := new(GetBookInfoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BookServer is the server API for Book service.
// All implementations must embed UnimplementedBookServer
// for forward compatibility
type BookServer interface {
	GetBookInfo(context.Context, *GetBookInfoReq) (*GetBookInfoResponse, error)
	GetBookInfoBidirectional(Book_GetBookInfoBidirectionalServer) error
	mustEmbedUnimplementedBookServer()
}

// UnimplementedBookServer must be embedded to have forward compatible implementations.
type UnimplementedBookServer struct {
}

func (UnimplementedBookServer) GetBookInfo(context.Context, *GetBookInfoReq) (*GetBookInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookInfo not implemented")
}
func (UnimplementedBookServer) GetBookInfoBidirectional(Book_GetBookInfoBidirectionalServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBookInfoBidirectional not implemented")
}
func (UnimplementedBookServer) mustEmbedUnimplementedBookServer() {}

// UnsafeBookServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServer will
// result in compilation errors.
type UnsafeBookServer interface {
	mustEmbedUnimplementedBookServer()
}

func RegisterBookServer(s grpc.ServiceRegistrar, srv BookServer) {
	s.RegisterService(&Book_ServiceDesc, srv)
}

func _Book_GetBookInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).GetBookInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_GetBookInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).GetBookInfo(ctx, req.(*GetBookInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_GetBookInfoBidirectional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BookServer).GetBookInfoBidirectional(&bookGetBookInfoBidirectionalServer{stream})
}

type Book_GetBookInfoBidirectionalServer interface {
	Send(*GetBookInfoResponse) error
	Recv() (*GetBookInfoReq, error)
	grpc.ServerStream
}

type bookGetBookInfoBidirectionalServer struct {
	grpc.ServerStream
}

func (x *bookGetBookInfoBidirectionalServer) Send(m *GetBookInfoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bookGetBookInfoBidirectionalServer) Recv() (*GetBookInfoReq, error) {
	m := new(GetBookInfoReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Book_ServiceDesc is the grpc.ServiceDesc for Book service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Book_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Book",
	HandlerType: (*BookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBookInfo",
			Handler:    _Book_GetBookInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBookInfoBidirectional",
			Handler:       _Book_GetBookInfoBidirectional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "book.proto",
}
