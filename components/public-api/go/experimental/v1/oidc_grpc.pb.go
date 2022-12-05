// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: gitpod/experimental/v1/oidc.proto

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

// OIDCServiceClient is the client API for OIDCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OIDCServiceClient interface {
	// Creates a new OIDC client configuration.
	CreateClientConfig(ctx context.Context, in *CreateClientConfigRequest, opts ...grpc.CallOption) (*CreateClientConfigResponse, error)
	// Retrieves an OIDC client configuration by ID.
	GetClientConfig(ctx context.Context, in *GetClientConfigRequest, opts ...grpc.CallOption) (*GetClientConfigResponse, error)
	// Lists OIDC client configurations.
	ListClientConfigs(ctx context.Context, in *ListClientConfigsRequest, opts ...grpc.CallOption) (*ListClientConfigsResponse, error)
	// Updates modifiable properties of an existing OIDC client configuration.
	UpdateClientConfig(ctx context.Context, in *UpdateClientConfigRequest, opts ...grpc.CallOption) (*UpdateClientConfigResponse, error)
	// Removes a OIDC client configuration by ID.
	DeleteClientConfig(ctx context.Context, in *DeleteClientConfigRequest, opts ...grpc.CallOption) (*DeleteClientConfigResponse, error)
}

type oIDCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOIDCServiceClient(cc grpc.ClientConnInterface) OIDCServiceClient {
	return &oIDCServiceClient{cc}
}

func (c *oIDCServiceClient) CreateClientConfig(ctx context.Context, in *CreateClientConfigRequest, opts ...grpc.CallOption) (*CreateClientConfigResponse, error) {
	out := new(CreateClientConfigResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.OIDCService/CreateClientConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oIDCServiceClient) GetClientConfig(ctx context.Context, in *GetClientConfigRequest, opts ...grpc.CallOption) (*GetClientConfigResponse, error) {
	out := new(GetClientConfigResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.OIDCService/GetClientConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oIDCServiceClient) ListClientConfigs(ctx context.Context, in *ListClientConfigsRequest, opts ...grpc.CallOption) (*ListClientConfigsResponse, error) {
	out := new(ListClientConfigsResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.OIDCService/ListClientConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oIDCServiceClient) UpdateClientConfig(ctx context.Context, in *UpdateClientConfigRequest, opts ...grpc.CallOption) (*UpdateClientConfigResponse, error) {
	out := new(UpdateClientConfigResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.OIDCService/UpdateClientConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oIDCServiceClient) DeleteClientConfig(ctx context.Context, in *DeleteClientConfigRequest, opts ...grpc.CallOption) (*DeleteClientConfigResponse, error) {
	out := new(DeleteClientConfigResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.OIDCService/DeleteClientConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OIDCServiceServer is the server API for OIDCService service.
// All implementations must embed UnimplementedOIDCServiceServer
// for forward compatibility
type OIDCServiceServer interface {
	// Creates a new OIDC client configuration.
	CreateClientConfig(context.Context, *CreateClientConfigRequest) (*CreateClientConfigResponse, error)
	// Retrieves an OIDC client configuration by ID.
	GetClientConfig(context.Context, *GetClientConfigRequest) (*GetClientConfigResponse, error)
	// Lists OIDC client configurations.
	ListClientConfigs(context.Context, *ListClientConfigsRequest) (*ListClientConfigsResponse, error)
	// Updates modifiable properties of an existing OIDC client configuration.
	UpdateClientConfig(context.Context, *UpdateClientConfigRequest) (*UpdateClientConfigResponse, error)
	// Removes a OIDC client configuration by ID.
	DeleteClientConfig(context.Context, *DeleteClientConfigRequest) (*DeleteClientConfigResponse, error)
	mustEmbedUnimplementedOIDCServiceServer()
}

// UnimplementedOIDCServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOIDCServiceServer struct {
}

func (UnimplementedOIDCServiceServer) CreateClientConfig(context.Context, *CreateClientConfigRequest) (*CreateClientConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClientConfig not implemented")
}
func (UnimplementedOIDCServiceServer) GetClientConfig(context.Context, *GetClientConfigRequest) (*GetClientConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientConfig not implemented")
}
func (UnimplementedOIDCServiceServer) ListClientConfigs(context.Context, *ListClientConfigsRequest) (*ListClientConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClientConfigs not implemented")
}
func (UnimplementedOIDCServiceServer) UpdateClientConfig(context.Context, *UpdateClientConfigRequest) (*UpdateClientConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClientConfig not implemented")
}
func (UnimplementedOIDCServiceServer) DeleteClientConfig(context.Context, *DeleteClientConfigRequest) (*DeleteClientConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClientConfig not implemented")
}
func (UnimplementedOIDCServiceServer) mustEmbedUnimplementedOIDCServiceServer() {}

// UnsafeOIDCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OIDCServiceServer will
// result in compilation errors.
type UnsafeOIDCServiceServer interface {
	mustEmbedUnimplementedOIDCServiceServer()
}

func RegisterOIDCServiceServer(s grpc.ServiceRegistrar, srv OIDCServiceServer) {
	s.RegisterService(&OIDCService_ServiceDesc, srv)
}

func _OIDCService_CreateClientConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OIDCServiceServer).CreateClientConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.OIDCService/CreateClientConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OIDCServiceServer).CreateClientConfig(ctx, req.(*CreateClientConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OIDCService_GetClientConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClientConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OIDCServiceServer).GetClientConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.OIDCService/GetClientConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OIDCServiceServer).GetClientConfig(ctx, req.(*GetClientConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OIDCService_ListClientConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClientConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OIDCServiceServer).ListClientConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.OIDCService/ListClientConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OIDCServiceServer).ListClientConfigs(ctx, req.(*ListClientConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OIDCService_UpdateClientConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClientConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OIDCServiceServer).UpdateClientConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.OIDCService/UpdateClientConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OIDCServiceServer).UpdateClientConfig(ctx, req.(*UpdateClientConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OIDCService_DeleteClientConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClientConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OIDCServiceServer).DeleteClientConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.OIDCService/DeleteClientConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OIDCServiceServer).DeleteClientConfig(ctx, req.(*DeleteClientConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OIDCService_ServiceDesc is the grpc.ServiceDesc for OIDCService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OIDCService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitpod.experimental.v1.OIDCService",
	HandlerType: (*OIDCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClientConfig",
			Handler:    _OIDCService_CreateClientConfig_Handler,
		},
		{
			MethodName: "GetClientConfig",
			Handler:    _OIDCService_GetClientConfig_Handler,
		},
		{
			MethodName: "ListClientConfigs",
			Handler:    _OIDCService_ListClientConfigs_Handler,
		},
		{
			MethodName: "UpdateClientConfig",
			Handler:    _OIDCService_UpdateClientConfig_Handler,
		},
		{
			MethodName: "DeleteClientConfig",
			Handler:    _OIDCService_DeleteClientConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gitpod/experimental/v1/oidc.proto",
}
