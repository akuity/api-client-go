// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: auth/v1/auth.proto

package authv1

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
	AuthService_GetDeviceCode_FullMethodName          = "/akuity.auth.v1.AuthService/GetDeviceCode"
	AuthService_GetDeviceToken_FullMethodName         = "/akuity.auth.v1.AuthService/GetDeviceToken"
	AuthService_RefreshAccessToken_FullMethodName     = "/akuity.auth.v1.AuthService/RefreshAccessToken"
	AuthService_GetOIDCProviderDetails_FullMethodName = "/akuity.auth.v1.AuthService/GetOIDCProviderDetails"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	GetDeviceCode(ctx context.Context, in *GetDeviceCodeRequest, opts ...grpc.CallOption) (*GetDeviceCodeResponse, error)
	GetDeviceToken(ctx context.Context, in *GetDeviceTokenRequest, opts ...grpc.CallOption) (*GetDeviceTokenResponse, error)
	RefreshAccessToken(ctx context.Context, in *RefreshAccessTokenRequest, opts ...grpc.CallOption) (*RefreshAccessTokenResponse, error)
	GetOIDCProviderDetails(ctx context.Context, in *GetOIDCProviderDetailsRequest, opts ...grpc.CallOption) (*GetOIDCProviderDetailsResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) GetDeviceCode(ctx context.Context, in *GetDeviceCodeRequest, opts ...grpc.CallOption) (*GetDeviceCodeResponse, error) {
	out := new(GetDeviceCodeResponse)
	err := c.cc.Invoke(ctx, AuthService_GetDeviceCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetDeviceToken(ctx context.Context, in *GetDeviceTokenRequest, opts ...grpc.CallOption) (*GetDeviceTokenResponse, error) {
	out := new(GetDeviceTokenResponse)
	err := c.cc.Invoke(ctx, AuthService_GetDeviceToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RefreshAccessToken(ctx context.Context, in *RefreshAccessTokenRequest, opts ...grpc.CallOption) (*RefreshAccessTokenResponse, error) {
	out := new(RefreshAccessTokenResponse)
	err := c.cc.Invoke(ctx, AuthService_RefreshAccessToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetOIDCProviderDetails(ctx context.Context, in *GetOIDCProviderDetailsRequest, opts ...grpc.CallOption) (*GetOIDCProviderDetailsResponse, error) {
	out := new(GetOIDCProviderDetailsResponse)
	err := c.cc.Invoke(ctx, AuthService_GetOIDCProviderDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	GetDeviceCode(context.Context, *GetDeviceCodeRequest) (*GetDeviceCodeResponse, error)
	GetDeviceToken(context.Context, *GetDeviceTokenRequest) (*GetDeviceTokenResponse, error)
	RefreshAccessToken(context.Context, *RefreshAccessTokenRequest) (*RefreshAccessTokenResponse, error)
	GetOIDCProviderDetails(context.Context, *GetOIDCProviderDetailsRequest) (*GetOIDCProviderDetailsResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) GetDeviceCode(context.Context, *GetDeviceCodeRequest) (*GetDeviceCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceCode not implemented")
}
func (UnimplementedAuthServiceServer) GetDeviceToken(context.Context, *GetDeviceTokenRequest) (*GetDeviceTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceToken not implemented")
}
func (UnimplementedAuthServiceServer) RefreshAccessToken(context.Context, *RefreshAccessTokenRequest) (*RefreshAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshAccessToken not implemented")
}
func (UnimplementedAuthServiceServer) GetOIDCProviderDetails(context.Context, *GetOIDCProviderDetailsRequest) (*GetOIDCProviderDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOIDCProviderDetails not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_GetDeviceCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetDeviceCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetDeviceCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetDeviceCode(ctx, req.(*GetDeviceCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetDeviceToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetDeviceToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetDeviceToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetDeviceToken(ctx, req.(*GetDeviceTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RefreshAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RefreshAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_RefreshAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RefreshAccessToken(ctx, req.(*RefreshAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetOIDCProviderDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOIDCProviderDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetOIDCProviderDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetOIDCProviderDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetOIDCProviderDetails(ctx, req.(*GetOIDCProviderDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "akuity.auth.v1.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDeviceCode",
			Handler:    _AuthService_GetDeviceCode_Handler,
		},
		{
			MethodName: "GetDeviceToken",
			Handler:    _AuthService_GetDeviceToken_Handler,
		},
		{
			MethodName: "RefreshAccessToken",
			Handler:    _AuthService_RefreshAccessToken_Handler,
		},
		{
			MethodName: "GetOIDCProviderDetails",
			Handler:    _AuthService_GetOIDCProviderDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/v1/auth.proto",
}
