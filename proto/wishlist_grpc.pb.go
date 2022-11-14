// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// WishListServiceClient is the client API for WishListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WishListServiceClient interface {
	Create(ctx context.Context, in *CreateWishListReq, opts ...grpc.CallOption) (*CreateWishListResp, error)
	Add(ctx context.Context, in *AddItemReq, opts ...grpc.CallOption) (*AddItemResp, error)
	List(ctx context.Context, in *ListWishListReq, opts ...grpc.CallOption) (*ListWishListResp, error)
}

type wishListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWishListServiceClient(cc grpc.ClientConnInterface) WishListServiceClient {
	return &wishListServiceClient{cc}
}

func (c *wishListServiceClient) Create(ctx context.Context, in *CreateWishListReq, opts ...grpc.CallOption) (*CreateWishListResp, error) {
	out := new(CreateWishListResp)
	err := c.cc.Invoke(ctx, "/grpc.WishListService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wishListServiceClient) Add(ctx context.Context, in *AddItemReq, opts ...grpc.CallOption) (*AddItemResp, error) {
	out := new(AddItemResp)
	err := c.cc.Invoke(ctx, "/grpc.WishListService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wishListServiceClient) List(ctx context.Context, in *ListWishListReq, opts ...grpc.CallOption) (*ListWishListResp, error) {
	out := new(ListWishListResp)
	err := c.cc.Invoke(ctx, "/grpc.WishListService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WishListServiceServer is the server API for WishListService service.
// All implementations must embed UnimplementedWishListServiceServer
// for forward compatibility
type WishListServiceServer interface {
	Create(context.Context, *CreateWishListReq) (*CreateWishListResp, error)
	Add(context.Context, *AddItemReq) (*AddItemResp, error)
	List(context.Context, *ListWishListReq) (*ListWishListResp, error)
	mustEmbedUnimplementedWishListServiceServer()
}

// UnimplementedWishListServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWishListServiceServer struct {
}

func (UnimplementedWishListServiceServer) Create(context.Context, *CreateWishListReq) (*CreateWishListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedWishListServiceServer) Add(context.Context, *AddItemReq) (*AddItemResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedWishListServiceServer) List(context.Context, *ListWishListReq) (*ListWishListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedWishListServiceServer) mustEmbedUnimplementedWishListServiceServer() {}

// UnsafeWishListServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WishListServiceServer will
// result in compilation errors.
type UnsafeWishListServiceServer interface {
	mustEmbedUnimplementedWishListServiceServer()
}

func RegisterWishListServiceServer(s grpc.ServiceRegistrar, srv WishListServiceServer) {
	s.RegisterService(&WishListService_ServiceDesc, srv)
}

func _WishListService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWishListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WishListServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.WishListService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WishListServiceServer).Create(ctx, req.(*CreateWishListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WishListService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WishListServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.WishListService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WishListServiceServer).Add(ctx, req.(*AddItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WishListService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWishListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WishListServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.WishListService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WishListServiceServer).List(ctx, req.(*ListWishListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// WishListService_ServiceDesc is the grpc.ServiceDesc for WishListService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WishListService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.WishListService",
	HandlerType: (*WishListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _WishListService_Create_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _WishListService_Add_Handler,
		},
		{
			MethodName: "List",
			Handler:    _WishListService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/wishlist.proto",
}
