// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: apikey/v1/apikey.proto

package apikeyv1

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
	APIKeyService_GetAPIKey_FullMethodName                       = "/akuity.apikey.v1.APIKeyService/GetAPIKey"
	APIKeyService_DeleteAPIKey_FullMethodName                    = "/akuity.apikey.v1.APIKeyService/DeleteAPIKey"
	APIKeyService_RegenerateAPIKeySecret_FullMethodName          = "/akuity.apikey.v1.APIKeyService/RegenerateAPIKeySecret"
	APIKeyService_GetWorkspaceAPIKey_FullMethodName              = "/akuity.apikey.v1.APIKeyService/GetWorkspaceAPIKey"
	APIKeyService_DeleteWorkspaceAPIKey_FullMethodName           = "/akuity.apikey.v1.APIKeyService/DeleteWorkspaceAPIKey"
	APIKeyService_RegenerateWorkspaceAPIKeySecret_FullMethodName = "/akuity.apikey.v1.APIKeyService/RegenerateWorkspaceAPIKeySecret"
)

// APIKeyServiceClient is the client API for APIKeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIKeyServiceClient interface {
	GetAPIKey(ctx context.Context, in *GetAPIKeyRequest, opts ...grpc.CallOption) (*GetAPIKeyResponse, error)
	DeleteAPIKey(ctx context.Context, in *DeleteAPIKeyRequest, opts ...grpc.CallOption) (*DeleteAPIKeyResponse, error)
	RegenerateAPIKeySecret(ctx context.Context, in *RegenerateAPIKeySecretRequest, opts ...grpc.CallOption) (*RegenerateAPIKeySecretResponse, error)
	GetWorkspaceAPIKey(ctx context.Context, in *GetWorkspaceAPIKeyRequest, opts ...grpc.CallOption) (*GetWorkspaceAPIKeyResponse, error)
	DeleteWorkspaceAPIKey(ctx context.Context, in *DeleteWorkspaceAPIKeyRequest, opts ...grpc.CallOption) (*DeleteWorkspaceAPIKeyResponse, error)
	RegenerateWorkspaceAPIKeySecret(ctx context.Context, in *RegenerateWorkspaceAPIKeySecretRequest, opts ...grpc.CallOption) (*RegenerateWorkspaceAPIKeySecretResponse, error)
}

type aPIKeyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIKeyServiceClient(cc grpc.ClientConnInterface) APIKeyServiceClient {
	return &aPIKeyServiceClient{cc}
}

func (c *aPIKeyServiceClient) GetAPIKey(ctx context.Context, in *GetAPIKeyRequest, opts ...grpc.CallOption) (*GetAPIKeyResponse, error) {
	out := new(GetAPIKeyResponse)
	err := c.cc.Invoke(ctx, APIKeyService_GetAPIKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIKeyServiceClient) DeleteAPIKey(ctx context.Context, in *DeleteAPIKeyRequest, opts ...grpc.CallOption) (*DeleteAPIKeyResponse, error) {
	out := new(DeleteAPIKeyResponse)
	err := c.cc.Invoke(ctx, APIKeyService_DeleteAPIKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIKeyServiceClient) RegenerateAPIKeySecret(ctx context.Context, in *RegenerateAPIKeySecretRequest, opts ...grpc.CallOption) (*RegenerateAPIKeySecretResponse, error) {
	out := new(RegenerateAPIKeySecretResponse)
	err := c.cc.Invoke(ctx, APIKeyService_RegenerateAPIKeySecret_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIKeyServiceClient) GetWorkspaceAPIKey(ctx context.Context, in *GetWorkspaceAPIKeyRequest, opts ...grpc.CallOption) (*GetWorkspaceAPIKeyResponse, error) {
	out := new(GetWorkspaceAPIKeyResponse)
	err := c.cc.Invoke(ctx, APIKeyService_GetWorkspaceAPIKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIKeyServiceClient) DeleteWorkspaceAPIKey(ctx context.Context, in *DeleteWorkspaceAPIKeyRequest, opts ...grpc.CallOption) (*DeleteWorkspaceAPIKeyResponse, error) {
	out := new(DeleteWorkspaceAPIKeyResponse)
	err := c.cc.Invoke(ctx, APIKeyService_DeleteWorkspaceAPIKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIKeyServiceClient) RegenerateWorkspaceAPIKeySecret(ctx context.Context, in *RegenerateWorkspaceAPIKeySecretRequest, opts ...grpc.CallOption) (*RegenerateWorkspaceAPIKeySecretResponse, error) {
	out := new(RegenerateWorkspaceAPIKeySecretResponse)
	err := c.cc.Invoke(ctx, APIKeyService_RegenerateWorkspaceAPIKeySecret_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIKeyServiceServer is the server API for APIKeyService service.
// All implementations must embed UnimplementedAPIKeyServiceServer
// for forward compatibility
type APIKeyServiceServer interface {
	GetAPIKey(context.Context, *GetAPIKeyRequest) (*GetAPIKeyResponse, error)
	DeleteAPIKey(context.Context, *DeleteAPIKeyRequest) (*DeleteAPIKeyResponse, error)
	RegenerateAPIKeySecret(context.Context, *RegenerateAPIKeySecretRequest) (*RegenerateAPIKeySecretResponse, error)
	GetWorkspaceAPIKey(context.Context, *GetWorkspaceAPIKeyRequest) (*GetWorkspaceAPIKeyResponse, error)
	DeleteWorkspaceAPIKey(context.Context, *DeleteWorkspaceAPIKeyRequest) (*DeleteWorkspaceAPIKeyResponse, error)
	RegenerateWorkspaceAPIKeySecret(context.Context, *RegenerateWorkspaceAPIKeySecretRequest) (*RegenerateWorkspaceAPIKeySecretResponse, error)
	mustEmbedUnimplementedAPIKeyServiceServer()
}

// UnimplementedAPIKeyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAPIKeyServiceServer struct {
}

func (UnimplementedAPIKeyServiceServer) GetAPIKey(context.Context, *GetAPIKeyRequest) (*GetAPIKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAPIKey not implemented")
}
func (UnimplementedAPIKeyServiceServer) DeleteAPIKey(context.Context, *DeleteAPIKeyRequest) (*DeleteAPIKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAPIKey not implemented")
}
func (UnimplementedAPIKeyServiceServer) RegenerateAPIKeySecret(context.Context, *RegenerateAPIKeySecretRequest) (*RegenerateAPIKeySecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegenerateAPIKeySecret not implemented")
}
func (UnimplementedAPIKeyServiceServer) GetWorkspaceAPIKey(context.Context, *GetWorkspaceAPIKeyRequest) (*GetWorkspaceAPIKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkspaceAPIKey not implemented")
}
func (UnimplementedAPIKeyServiceServer) DeleteWorkspaceAPIKey(context.Context, *DeleteWorkspaceAPIKeyRequest) (*DeleteWorkspaceAPIKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkspaceAPIKey not implemented")
}
func (UnimplementedAPIKeyServiceServer) RegenerateWorkspaceAPIKeySecret(context.Context, *RegenerateWorkspaceAPIKeySecretRequest) (*RegenerateWorkspaceAPIKeySecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegenerateWorkspaceAPIKeySecret not implemented")
}
func (UnimplementedAPIKeyServiceServer) mustEmbedUnimplementedAPIKeyServiceServer() {}

// UnsafeAPIKeyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIKeyServiceServer will
// result in compilation errors.
type UnsafeAPIKeyServiceServer interface {
	mustEmbedUnimplementedAPIKeyServiceServer()
}

func RegisterAPIKeyServiceServer(s grpc.ServiceRegistrar, srv APIKeyServiceServer) {
	s.RegisterService(&APIKeyService_ServiceDesc, srv)
}

func _APIKeyService_GetAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIKeyServiceServer).GetAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: APIKeyService_GetAPIKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIKeyServiceServer).GetAPIKey(ctx, req.(*GetAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIKeyService_DeleteAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIKeyServiceServer).DeleteAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: APIKeyService_DeleteAPIKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIKeyServiceServer).DeleteAPIKey(ctx, req.(*DeleteAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIKeyService_RegenerateAPIKeySecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegenerateAPIKeySecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIKeyServiceServer).RegenerateAPIKeySecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: APIKeyService_RegenerateAPIKeySecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIKeyServiceServer).RegenerateAPIKeySecret(ctx, req.(*RegenerateAPIKeySecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIKeyService_GetWorkspaceAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkspaceAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIKeyServiceServer).GetWorkspaceAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: APIKeyService_GetWorkspaceAPIKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIKeyServiceServer).GetWorkspaceAPIKey(ctx, req.(*GetWorkspaceAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIKeyService_DeleteWorkspaceAPIKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkspaceAPIKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIKeyServiceServer).DeleteWorkspaceAPIKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: APIKeyService_DeleteWorkspaceAPIKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIKeyServiceServer).DeleteWorkspaceAPIKey(ctx, req.(*DeleteWorkspaceAPIKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIKeyService_RegenerateWorkspaceAPIKeySecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegenerateWorkspaceAPIKeySecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIKeyServiceServer).RegenerateWorkspaceAPIKeySecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: APIKeyService_RegenerateWorkspaceAPIKeySecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIKeyServiceServer).RegenerateWorkspaceAPIKeySecret(ctx, req.(*RegenerateWorkspaceAPIKeySecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// APIKeyService_ServiceDesc is the grpc.ServiceDesc for APIKeyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var APIKeyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "akuity.apikey.v1.APIKeyService",
	HandlerType: (*APIKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAPIKey",
			Handler:    _APIKeyService_GetAPIKey_Handler,
		},
		{
			MethodName: "DeleteAPIKey",
			Handler:    _APIKeyService_DeleteAPIKey_Handler,
		},
		{
			MethodName: "RegenerateAPIKeySecret",
			Handler:    _APIKeyService_RegenerateAPIKeySecret_Handler,
		},
		{
			MethodName: "GetWorkspaceAPIKey",
			Handler:    _APIKeyService_GetWorkspaceAPIKey_Handler,
		},
		{
			MethodName: "DeleteWorkspaceAPIKey",
			Handler:    _APIKeyService_DeleteWorkspaceAPIKey_Handler,
		},
		{
			MethodName: "RegenerateWorkspaceAPIKeySecret",
			Handler:    _APIKeyService_RegenerateWorkspaceAPIKeySecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apikey/v1/apikey.proto",
}
