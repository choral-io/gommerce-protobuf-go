// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: gommerce/iam/v1beta/realms.proto

package iam_v1beta

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
	RealmsService_CreateRealm_FullMethodName = "/gommerce.iam.v1beta.RealmsService/CreateRealm"
	RealmsService_GetRealm_FullMethodName    = "/gommerce.iam.v1beta.RealmsService/GetRealm"
	RealmsService_ListRealms_FullMethodName  = "/gommerce.iam.v1beta.RealmsService/ListRealms"
	RealmsService_UpdateRealm_FullMethodName = "/gommerce.iam.v1beta.RealmsService/UpdateRealm"
	RealmsService_DeleteRealm_FullMethodName = "/gommerce.iam.v1beta.RealmsService/DeleteRealm"
)

// RealmsServiceClient is the client API for RealmsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RealmsServiceClient interface {
	CreateRealm(ctx context.Context, in *CreateRealmRequest, opts ...grpc.CallOption) (*CreateRealmResponse, error)
	GetRealm(ctx context.Context, in *GetRealmRequest, opts ...grpc.CallOption) (*GetRealmResponse, error)
	ListRealms(ctx context.Context, in *ListRealmsRequest, opts ...grpc.CallOption) (*ListRealmsResponse, error)
	UpdateRealm(ctx context.Context, in *UpdateRealmRequest, opts ...grpc.CallOption) (*UpdateRealmResponse, error)
	DeleteRealm(ctx context.Context, in *DeleteRealmRequest, opts ...grpc.CallOption) (*DeleteRealmResponse, error)
}

type realmsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRealmsServiceClient(cc grpc.ClientConnInterface) RealmsServiceClient {
	return &realmsServiceClient{cc}
}

func (c *realmsServiceClient) CreateRealm(ctx context.Context, in *CreateRealmRequest, opts ...grpc.CallOption) (*CreateRealmResponse, error) {
	out := new(CreateRealmResponse)
	err := c.cc.Invoke(ctx, RealmsService_CreateRealm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realmsServiceClient) GetRealm(ctx context.Context, in *GetRealmRequest, opts ...grpc.CallOption) (*GetRealmResponse, error) {
	out := new(GetRealmResponse)
	err := c.cc.Invoke(ctx, RealmsService_GetRealm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realmsServiceClient) ListRealms(ctx context.Context, in *ListRealmsRequest, opts ...grpc.CallOption) (*ListRealmsResponse, error) {
	out := new(ListRealmsResponse)
	err := c.cc.Invoke(ctx, RealmsService_ListRealms_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realmsServiceClient) UpdateRealm(ctx context.Context, in *UpdateRealmRequest, opts ...grpc.CallOption) (*UpdateRealmResponse, error) {
	out := new(UpdateRealmResponse)
	err := c.cc.Invoke(ctx, RealmsService_UpdateRealm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realmsServiceClient) DeleteRealm(ctx context.Context, in *DeleteRealmRequest, opts ...grpc.CallOption) (*DeleteRealmResponse, error) {
	out := new(DeleteRealmResponse)
	err := c.cc.Invoke(ctx, RealmsService_DeleteRealm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RealmsServiceServer is the server API for RealmsService service.
// All implementations must embed UnimplementedRealmsServiceServer
// for forward compatibility
type RealmsServiceServer interface {
	CreateRealm(context.Context, *CreateRealmRequest) (*CreateRealmResponse, error)
	GetRealm(context.Context, *GetRealmRequest) (*GetRealmResponse, error)
	ListRealms(context.Context, *ListRealmsRequest) (*ListRealmsResponse, error)
	UpdateRealm(context.Context, *UpdateRealmRequest) (*UpdateRealmResponse, error)
	DeleteRealm(context.Context, *DeleteRealmRequest) (*DeleteRealmResponse, error)
	mustEmbedUnimplementedRealmsServiceServer()
}

// UnimplementedRealmsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRealmsServiceServer struct {
}

func (UnimplementedRealmsServiceServer) CreateRealm(context.Context, *CreateRealmRequest) (*CreateRealmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRealm not implemented")
}
func (UnimplementedRealmsServiceServer) GetRealm(context.Context, *GetRealmRequest) (*GetRealmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRealm not implemented")
}
func (UnimplementedRealmsServiceServer) ListRealms(context.Context, *ListRealmsRequest) (*ListRealmsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRealms not implemented")
}
func (UnimplementedRealmsServiceServer) UpdateRealm(context.Context, *UpdateRealmRequest) (*UpdateRealmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRealm not implemented")
}
func (UnimplementedRealmsServiceServer) DeleteRealm(context.Context, *DeleteRealmRequest) (*DeleteRealmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRealm not implemented")
}
func (UnimplementedRealmsServiceServer) mustEmbedUnimplementedRealmsServiceServer() {}

// UnsafeRealmsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RealmsServiceServer will
// result in compilation errors.
type UnsafeRealmsServiceServer interface {
	mustEmbedUnimplementedRealmsServiceServer()
}

func RegisterRealmsServiceServer(s grpc.ServiceRegistrar, srv RealmsServiceServer) {
	s.RegisterService(&RealmsService_ServiceDesc, srv)
}

func _RealmsService_CreateRealm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRealmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealmsServiceServer).CreateRealm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealmsService_CreateRealm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealmsServiceServer).CreateRealm(ctx, req.(*CreateRealmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealmsService_GetRealm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRealmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealmsServiceServer).GetRealm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealmsService_GetRealm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealmsServiceServer).GetRealm(ctx, req.(*GetRealmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealmsService_ListRealms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRealmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealmsServiceServer).ListRealms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealmsService_ListRealms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealmsServiceServer).ListRealms(ctx, req.(*ListRealmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealmsService_UpdateRealm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRealmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealmsServiceServer).UpdateRealm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealmsService_UpdateRealm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealmsServiceServer).UpdateRealm(ctx, req.(*UpdateRealmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RealmsService_DeleteRealm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRealmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealmsServiceServer).DeleteRealm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RealmsService_DeleteRealm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealmsServiceServer).DeleteRealm(ctx, req.(*DeleteRealmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RealmsService_ServiceDesc is the grpc.ServiceDesc for RealmsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RealmsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gommerce.iam.v1beta.RealmsService",
	HandlerType: (*RealmsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRealm",
			Handler:    _RealmsService_CreateRealm_Handler,
		},
		{
			MethodName: "GetRealm",
			Handler:    _RealmsService_GetRealm_Handler,
		},
		{
			MethodName: "ListRealms",
			Handler:    _RealmsService_ListRealms_Handler,
		},
		{
			MethodName: "UpdateRealm",
			Handler:    _RealmsService_UpdateRealm_Handler,
		},
		{
			MethodName: "DeleteRealm",
			Handler:    _RealmsService_DeleteRealm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gommerce/iam/v1beta/realms.proto",
}
