package client

import (
	"context"
	"net/http"
	"net/url"

	"google.golang.org/grpc"

	gwclient "github.com/akuity/api-client-go/pkg/api/gateway/client"
	systemv1 "github.com/akuity/api-client-go/pkg/api/gen/system/v1"
)

type SystemV1Client interface {
	GatewayClient
	systemv1.SystemServiceClient
}

type systemV1Client struct {
	gwc gwclient.Client
}

func NewSystemV1Client(baseURL string, opts ...gwclient.Option) SystemV1Client {
	return newGatewayClient(&systemV1Client{}, baseURL, opts...)
}

func (c *systemV1Client) setGatewayClient(gwc gwclient.Client) {
	c.gwc = gwc
}

func (c *systemV1Client) GetVersion(ctx context.Context, req *systemv1.GetVersionRequest, _ ...grpc.CallOption) (*systemv1.GetVersionResponse, error) {
	gwReq := c.gwc.NewRequest(http.MethodGet, "/api/v1/system/version")
	return gwclient.DoRequest[systemv1.GetVersionResponse](ctx, gwReq)
}

func (c *systemV1Client) GetSettings(ctx context.Context, req *systemv1.GetSettingsRequest, _ ...grpc.CallOption) (*systemv1.GetSettingsResponse, error) {
	gwReq := c.gwc.NewRequest(http.MethodGet, "/api/v1/system/settings")
	return gwclient.DoRequest[systemv1.GetSettingsResponse](ctx, gwReq)
}

func (c *systemV1Client) ListFeatures(ctx context.Context, req *systemv1.ListFeaturesRequest, _ ...grpc.CallOption) (*systemv1.ListFeaturesResponse, error) {
	gwReq := c.gwc.NewRequest(http.MethodGet, "/api/v1/system/features").
		SetQueryParamsFromValues(url.Values{
			"names": req.GetNames(),
		})
	return gwclient.DoRequest[systemv1.ListFeaturesResponse](ctx, gwReq)
}
