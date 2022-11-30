package client

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc"

	gwclient "github.com/akuity/api-client-go/pkg/api/gateway/client"
	argocdv1 "github.com/akuity/api-client-go/pkg/api/gen/argocd/v1"
)

type ArgoCDV1Client interface {
	GatewayClient
	argocdv1.ArgoCDServiceClient
}

type argoCDV1Client struct {
	gwc gwclient.Client
}

func NewArgoCDV1Client(baseURL string, opts ...gwclient.Option) ArgoCDV1Client {
	return newGatewayClient(&argoCDV1Client{}, baseURL, opts...)
}

func (c *argoCDV1Client) setGatewayClient(gwc gwclient.Client) {
	c.gwc = gwc
}

func (c *argoCDV1Client) ListInstanceVersions(
	ctx context.Context,
	req *argocdv1.ListInstanceVersionsRequest,
	_ ...grpc.CallOption,
) (*argocdv1.ListInstanceVersionsResponse, error) {
	gwReq := c.gwc.NewRequest(http.MethodGet, "/api/v1/argocd/versions")
	return gwclient.DoRequest[argocdv1.ListInstanceVersionsResponse](ctx, gwReq)
}

func (c *argoCDV1Client) ListOrganizationInstances(
	ctx context.Context,
	req *argocdv1.ListOrganizationInstancesRequest,
	_ ...grpc.CallOption,
) (*argocdv1.ListOrganizationInstancesResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances", req.GetOrganizationId())
	gwReq := c.gwc.NewRequest(http.MethodGet, path)
	return gwclient.DoRequest[argocdv1.ListOrganizationInstancesResponse](ctx, gwReq)
}

func (c *argoCDV1Client) CreateOrganizationInstance(
	ctx context.Context,
	req *argocdv1.CreateOrganizationInstanceRequest,
	_ ...grpc.CallOption,
) (*argocdv1.CreateOrganizationInstanceResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances", req.GetOrganizationId())
	gwReq := c.gwc.NewRequest(http.MethodPost, path).SetBody(req)
	return gwclient.DoRequest[argocdv1.CreateOrganizationInstanceResponse](ctx, gwReq)
}

func (c *argoCDV1Client) GetOrganizationInstance(
	ctx context.Context,
	req *argocdv1.GetOrganizationInstanceRequest,
	_ ...grpc.CallOption,
) (*argocdv1.GetOrganizationInstanceResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s", req.GetOrganizationId(), req.GetId())
	gwReq := c.gwc.NewRequest(http.MethodGet, path).
		SetQueryParam("id_type", req.GetIdType().String())
	return gwclient.DoRequest[argocdv1.GetOrganizationInstanceResponse](ctx, gwReq)
}

func (c *argoCDV1Client) PatchOrganizationInstance(
	ctx context.Context,
	req *argocdv1.PatchOrganizationInstanceRequest,
	_ ...grpc.CallOption,
) (*argocdv1.PatchOrganizationInstanceResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s", req.GetOrganizationId(), req.GetId())
	gwReq := c.gwc.NewRequest(http.MethodPatch, path).SetBody(req)
	return gwclient.DoRequest[argocdv1.PatchOrganizationInstanceResponse](ctx, gwReq)
}

func (c *argoCDV1Client) UpdateOrganizationInstance(
	ctx context.Context,
	req *argocdv1.UpdateOrganizationInstanceRequest,
	_ ...grpc.CallOption,
) (*argocdv1.UpdateOrganizationInstanceResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s", req.GetOrganizationId(), req.GetId())
	gwReq := c.gwc.NewRequest(http.MethodPut, path).SetBody(req)
	return gwclient.DoRequest[argocdv1.UpdateOrganizationInstanceResponse](ctx, gwReq)
}

func (c *argoCDV1Client) DeleteOrganizationInstance(
	ctx context.Context,
	req *argocdv1.DeleteOrganizationInstanceRequest,
	_ ...grpc.CallOption,
) (*argocdv1.DeleteOrganizationInstanceResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s", req.GetOrganizationId(), req.GetId())
	gwReq := c.gwc.NewRequest(http.MethodDelete, path)
	return gwclient.DoRequest[argocdv1.DeleteOrganizationInstanceResponse](ctx, gwReq)
}

func (c *argoCDV1Client) ListOrganizationInstanceClusters(
	ctx context.Context,
	req *argocdv1.ListOrganizationInstanceClustersRequest,
	_ ...grpc.CallOption,
) (*argocdv1.ListOrganizationInstanceClustersResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s/clusters", req.GetOrganizationId(), req.GetInstanceId())
	gwReq := c.gwc.NewRequest(http.MethodGet, path)
	return gwclient.DoRequest[argocdv1.ListOrganizationInstanceClustersResponse](ctx, gwReq)
}

func (c *argoCDV1Client) CreateOrganizationInstanceCluster(
	ctx context.Context,
	req *argocdv1.CreateOrganizationInstanceClusterRequest,
	_ ...grpc.CallOption,
) (*argocdv1.CreateOrganizationInstanceClusterResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s/clusters", req.GetOrganizationId(), req.GetInstanceId())
	gwReq := c.gwc.NewRequest(http.MethodPost, path).SetBody(req)
	return gwclient.DoRequest[argocdv1.CreateOrganizationInstanceClusterResponse](ctx, gwReq)
}

func (c *argoCDV1Client) GetOrganizationInstanceCluster(
	ctx context.Context,
	req *argocdv1.GetOrganizationInstanceClusterRequest,
	_ ...grpc.CallOption,
) (*argocdv1.GetOrganizationInstanceClusterResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s/clusters/%s", req.GetOrganizationId(), req.GetInstanceId(), req.GetId())
	gwReq := c.gwc.NewRequest(http.MethodGet, path).
		SetQueryParam("id_type", req.GetIdType().String())
	return gwclient.DoRequest[argocdv1.GetOrganizationInstanceClusterResponse](ctx, gwReq)
}

func (c *argoCDV1Client) GetOrganizationInstanceClusterManifests(
	ctx context.Context,
	req *argocdv1.GetOrganizationInstanceClusterManifestsRequest,
	_ ...grpc.CallOption,
) (*argocdv1.GetOrganizationInstanceClusterManifestsResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s/clusters/%s/manifests", req.GetOrganizationId(), req.GetInstanceId(), req.GetId())
	gwReq := c.gwc.NewRequest(http.MethodGet, path)
	return gwclient.DoRequest[argocdv1.GetOrganizationInstanceClusterManifestsResponse](ctx, gwReq)
}

func (c *argoCDV1Client) UpdateOrganizationInstanceCluster(
	ctx context.Context,
	req *argocdv1.UpdateOrganizationInstanceClusterRequest,
	_ ...grpc.CallOption,
) (*argocdv1.UpdateOrganizationInstanceClusterResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s/clusters/%s", req.GetOrganizationId(), req.GetInstanceId(), req.GetId())
	gwReq := c.gwc.NewRequest(http.MethodPut, path).SetBody(req)
	return gwclient.DoRequest[argocdv1.UpdateOrganizationInstanceClusterResponse](ctx, gwReq)
}

func (c *argoCDV1Client) DeleteOrganizationInstanceCluster(
	ctx context.Context,
	req *argocdv1.DeleteOrganizationInstanceClusterRequest,
	_ ...grpc.CallOption,
) (*argocdv1.DeleteOrganizationInstanceClusterResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s/clusters/%s", req.GetOrganizationId(), req.GetInstanceId(), req.GetId())
	gwReq := c.gwc.NewRequest(http.MethodDelete, path)
	return gwclient.DoRequest[argocdv1.DeleteOrganizationInstanceClusterResponse](ctx, gwReq)
}
