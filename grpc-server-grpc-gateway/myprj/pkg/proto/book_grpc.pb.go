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
	Book_GetBookInfo_FullMethodName = "/proto.Book/GetBookInfo"
)

// BookClient is the client API for Book service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookClient interface {
	GetBookInfo(ctx context.Context, in *GetBookInfoReq, opts ...grpc.CallOption) (*GetBookInfoResponse, error)
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

// BookServer is the server API for Book service.
// All implementations must embed UnimplementedBookServer
// for forward compatibility
type BookServer interface {
	GetBookInfo(context.Context, *GetBookInfoReq) (*GetBookInfoResponse, error)
	mustEmbedUnimplementedBookServer()
}

// UnimplementedBookServer must be embedded to have forward compatible implementations.
type UnimplementedBookServer struct {
}

func (UnimplementedBookServer) GetBookInfo(context.Context, *GetBookInfoReq) (*GetBookInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookInfo not implemented")
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
	Streams:  []grpc.StreamDesc{},
	Metadata: "book.proto",
}
