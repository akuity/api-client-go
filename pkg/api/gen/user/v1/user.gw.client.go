// Code generated by protoc-gen-grpc-gateway-client. DO NOT EDIT.
// source: user/v1/user.proto

package userv1

import (
	context "context"
	gateway "github.com/akuity/grpc-gateway-client/pkg/grpc/gateway"
)

// UserServiceGatewayClient is the interface for UserService service client.
type UserServiceGatewayClient interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	UpdateUserUIPreferences(context.Context, *UpdateUserUIPreferencesRequest) (*UpdateUserUIPreferencesResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
}

func NewUserServiceGatewayClient(c gateway.Client) UserServiceGatewayClient {
	return &userServiceGatewayClient{
		gwc: c,
	}
}

type userServiceGatewayClient struct {
	gwc gateway.Client
}

func (c *userServiceGatewayClient) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	gwReq := c.gwc.NewRequest("GET", "/api/v1/users/me")
	return gateway.DoRequest[GetUserResponse](ctx, gwReq)
}

func (c *userServiceGatewayClient) UpdateUserUIPreferences(ctx context.Context, req *UpdateUserUIPreferencesRequest) (*UpdateUserUIPreferencesResponse, error) {
	gwReq := c.gwc.NewRequest("PUT", "/api/v1/users/me/ui-preferences")
	gwReq.SetBody(req)
	return gateway.DoRequest[UpdateUserUIPreferencesResponse](ctx, gwReq)
}

func (c *userServiceGatewayClient) DeleteUser(ctx context.Context, req *DeleteUserRequest) (*DeleteUserResponse, error) {
	gwReq := c.gwc.NewRequest("DELETE", "/api/v1/users/me")
	gwReq.SetBody(req)
	return gateway.DoRequest[DeleteUserResponse](ctx, gwReq)
}
