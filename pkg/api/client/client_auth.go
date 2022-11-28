package client

import (
	"context"
	"net/http"

	"google.golang.org/grpc"

	gwclient "github.com/akuity/api-client-go/pkg/api/gateway/client"
	authv1 "github.com/akuity/api-client-go/pkg/api/gen/auth/v1"
)

type AuthV1Client interface {
	GatewayClient
	authv1.AuthServiceClient
}

type authV1Client struct {
	gwc gwclient.Client
}

func NewAuthV1Client(baseURL string, opts ...gwclient.Option) AuthV1Client {
	return newGatewayClient(&authV1Client{}, baseURL, opts...)
}

func (c *authV1Client) setGatewayClient(gwc gwclient.Client) {
	c.gwc = gwc
}

func (c *authV1Client) GetDeviceCode(ctx context.Context, _ *authv1.GetDeviceCodeRequest, _ ...grpc.CallOption) (*authv1.GetDeviceCodeResponse, error) {
	gwReq := c.gwc.NewRequest(http.MethodGet, "/api/v1/auth/device-code")
	return gwclient.DoRequest[authv1.GetDeviceCodeResponse](ctx, gwReq)
}

func (c *authV1Client) GetDeviceToken(ctx context.Context, req *authv1.GetDeviceTokenRequest, _ ...grpc.CallOption) (*authv1.GetDeviceTokenResponse, error) {
	gwReq := c.gwc.NewRequest(http.MethodPost, "/api/v1/auth/device-token").SetBody(req)
	return gwclient.DoRequest[authv1.GetDeviceTokenResponse](ctx, gwReq)
}
