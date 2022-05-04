// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: demo/api/bff/v1/bff.proto

package v1

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

// BffServiceClient is the client API for BffService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BffServiceClient interface {
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*Order, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
}

type bffServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBffServiceClient(cc grpc.ClientConnInterface) BffServiceClient {
	return &bffServiceClient{cc}
}

func (c *bffServiceClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := c.cc.Invoke(ctx, "/demo.api.bff.v1.BffService/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bffServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/demo.api.bff.v1.BffService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bffServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/demo.api.bff.v1.BffService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BffServiceServer is the server API for BffService service.
// All implementations must embed UnimplementedBffServiceServer
// for forward compatibility
type BffServiceServer interface {
	GetOrder(context.Context, *GetOrderRequest) (*Order, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	mustEmbedUnimplementedBffServiceServer()
}

// UnimplementedBffServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBffServiceServer struct {
}

func (UnimplementedBffServiceServer) GetOrder(context.Context, *GetOrderRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedBffServiceServer) GetUser(context.Context, *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedBffServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedBffServiceServer) mustEmbedUnimplementedBffServiceServer() {}

// UnsafeBffServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BffServiceServer will
// result in compilation errors.
type UnsafeBffServiceServer interface {
	mustEmbedUnimplementedBffServiceServer()
}

func RegisterBffServiceServer(s grpc.ServiceRegistrar, srv BffServiceServer) {
	s.RegisterService(&BffService_ServiceDesc, srv)
}

func _BffService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BffServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.api.bff.v1.BffService/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BffServiceServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BffService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BffServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.api.bff.v1.BffService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BffServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BffService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BffServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.api.bff.v1.BffService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BffServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BffService_ServiceDesc is the grpc.ServiceDesc for BffService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BffService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "demo.api.bff.v1.BffService",
	HandlerType: (*BffServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrder",
			Handler:    _BffService_GetOrder_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _BffService_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _BffService_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo/api/bff/v1/bff.proto",
}
