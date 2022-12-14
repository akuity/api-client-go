// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	GetDeviceCode(ctx context.Context, in *GetDeviceCodeRequest, opts ...grpc.CallOption) (*GetDeviceCodeResponse, error)
	GetDeviceToken(ctx context.Context, in *GetDeviceTokenRequest, opts ...grpc.CallOption) (*GetDeviceTokenResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) GetDeviceCode(ctx context.Context, in *GetDeviceCodeRequest, opts ...grpc.CallOption) (*GetDeviceCodeResponse, error) {
	out := new(GetDeviceCodeResponse)
	err := c.cc.Invoke(ctx, "/akuity.auth.v1.AuthService/GetDeviceCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetDeviceToken(ctx context.Context, in *GetDeviceTokenRequest, opts ...grpc.CallOption) (*GetDeviceTokenResponse, error) {
	out := new(GetDeviceTokenResponse)
	err := c.cc.Invoke(ctx, "/akuity.auth.v1.AuthService/GetDeviceToken", in, out, opts...)
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
		FullMethod: "/akuity.auth.v1.AuthService/GetDeviceCode",
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
		FullMethod: "/akuity.auth.v1.AuthService/GetDeviceToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetDeviceToken(ctx, req.(*GetDeviceTokenRequest))
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/v1/auth.proto",
}
