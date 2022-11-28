package client

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc"

	gwclient "github.com/akuity/api-client-go/pkg/api/gateway/client"
	organizationv1 "github.com/akuity/api-client-go/pkg/api/gen/organization/v1"
)

type OrganizationV1Client interface {
	GatewayClient
	organizationv1.OrganizationServiceClient
}

type organizationV1Client struct {
	gwc gwclient.Client
}

func NewOrganizationV1Client(baseURL string, opts ...gwclient.Option) OrganizationV1Client {
	return newGatewayClient(&organizationV1Client{}, baseURL, opts...)
}

func (c *organizationV1Client) setGatewayClient(gwc gwclient.Client) {
	c.gwc = gwc
}

func (c *organizationV1Client) ListAuthenticatedUserOrganizations(
	ctx context.Context,
	req *organizationv1.ListAuthenticatedUserOrganizationsRequest,
	_ ...grpc.CallOption,
) (*organizationv1.ListAuthenticatedUserOrganizationsResponse, error) {
	gwReq := c.gwc.NewRequest(http.MethodGet, "/api/v1/organizations")
	return gwclient.DoRequest[organizationv1.ListAuthenticatedUserOrganizationsResponse](ctx, gwReq)
}

func (c *organizationV1Client) GetOrganization(
	ctx context.Context,
	req *organizationv1.GetOrganizationRequest,
	_ ...grpc.CallOption,
) (*organizationv1.GetOrganizationResponse, error) {
	path := fmt.Sprintf("/api/v1/organizations/%s", req.GetId())
	gwReq := c.gwc.NewRequest(http.MethodGet, path).
		SetQueryParam("id_type", req.GetIdType().String())
	return gwclient.DoRequest[organizationv1.GetOrganizationResponse](ctx, gwReq)
}
