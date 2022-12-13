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

func (c *argoCDV1Client) WatchInstanceClusters(ctx context.Context, in *argocdv1.WatchInstanceClustersRequest, opts ...grpc.CallOption) (argocdv1.ArgoCDService_WatchInstanceClustersClient, error) {
	// TODO: implement watch instance cluster
	return nil, nil
}

func (c *argoCDV1Client) WatchInstances(ctx context.Context, in *argocdv1.WatchInstancesRequest, opts ...grpc.CallOption) (argocdv1.ArgoCDService_WatchInstancesClient, error) {
	// TODO: implement watch instance cluster
	return nil, nil
}

func (c *argoCDV1Client) ListInstanceVersions(
	ctx context.Context,
	req *argocdv1.ListInstanceVersionsRequest,
	_ ...grpc.CallOption,
) (*argocdv1.ListInstanceVersionsResponse, error) {
	gwReq := c.gwc.NewRequest(http.MethodGet, "/api/v1/argocd/versions")
	return gwclient.DoRequest[argocdv1.ListInstanceVersionsResponse](ctx, gwReq)
}

func (c *argoCDV1Client) ListInstances(
	ctx context.Context,
	req *argocdv1.ListInstancesRequest,
	_ ...grpc.CallOption,
) (*argocdv1.ListInstancesResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances", req.GetOrganizationId())
	gwReq := c.gwc.NewRequest(http.MethodGet, path)
	return gwclient.DoRequest[argocdv1.ListInstancesResponse](ctx, gwReq)
}

func (c *argoCDV1Client) CreateInstance(
	ctx context.Context,
	req *argocdv1.CreateInstanceRequest,
	_ ...grpc.CallOption,
) (*argocdv1.CreateInstanceResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances", req.GetOrganizationId())
	gwReq := c.gwc.NewRequest(http.MethodPost, path).SetBody(req)
	return gwclient.DoRequest[argocdv1.CreateInstanceResponse](ctx, gwReq)
}

func (c *argoCDV1Client) GetInstance(
	ctx context.Context,
	req *argocdv1.GetInstanceRequest,
	_ ...grpc.CallOption,
) (*argocdv1.GetInstanceResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s", req.GetOrganizationId(), req.GetId())
	gwReq := c.gwc.NewRequest(http.MethodGet, path).
		SetQueryParam("id_type", req.GetIdType().String())
	return gwclient.DoRequest[argocdv1.GetInstanceResponse](ctx, gwReq)
}

func (c *argoCDV1Client) PatchInstance(
	ctx context.Context,
	req *argocdv1.PatchInstanceRequest,
	_ ...grpc.CallOption,
) (*argocdv1.PatchInstanceResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s", req.GetOrganizationId(), req.GetId())
	gwReq := c.gwc.NewRequest(http.MethodPatch, path).SetBody(req)
	return gwclient.DoRequest[argocdv1.PatchInstanceResponse](ctx, gwReq)
}

func (c *argoCDV1Client) UpdateInstance(
	ctx context.Context,
	req *argocdv1.UpdateInstanceRequest,
	_ ...grpc.CallOption,
) (*argocdv1.UpdateInstanceResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s", req.GetOrganizationId(), req.GetId())
	gwReq := c.gwc.NewRequest(http.MethodPut, path).SetBody(req)
	return gwclient.DoRequest[argocdv1.UpdateInstanceResponse](ctx, gwReq)
}

func (c *argoCDV1Client) DeleteInstance(
	ctx context.Context,
	req *argocdv1.DeleteInstanceRequest,
	_ ...grpc.CallOption,
) (*argocdv1.DeleteInstanceResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s", req.GetOrganizationId(), req.GetId())
	gwReq := c.gwc.NewRequest(http.MethodDelete, path)
	return gwclient.DoRequest[argocdv1.DeleteInstanceResponse](ctx, gwReq)
}

func (c *argoCDV1Client) ListInstanceAccounts(
	ctx context.Context,
	req *argocdv1.ListInstanceAccountsRequest,
	_ ...grpc.CallOption,
) (*argocdv1.ListInstanceAccountsResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s/accounts", req.GetOrganizationId(), req.GetInstanceId())
	gwReq := c.gwc.NewRequest(http.MethodGet, path)
	return gwclient.DoRequest[argocdv1.ListInstanceAccountsResponse](ctx, gwReq)
}

func (c *argoCDV1Client) UpsertInstanceAccount(
	ctx context.Context,
	req *argocdv1.UpsertInstanceAccountRequest,
	_ ...grpc.CallOption,
) (*argocdv1.UpsertInstanceAccountResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s/accounts/%s", req.GetOrganizationId(), req.GetInstanceId(), req.GetName())
	gwReq := c.gwc.NewRequest(http.MethodPut, path).SetBody(req)
	return gwclient.DoRequest[argocdv1.UpsertInstanceAccountResponse](ctx, gwReq)
}

func (c *argoCDV1Client) UpdateInstanceAccountPassword(
	ctx context.Context,
	req *argocdv1.UpdateInstanceAccountPasswordRequest,
	_ ...grpc.CallOption,
) (*argocdv1.UpdateInstanceAccountPasswordResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s/accounts/%s/password", req.GetOrganizationId(), req.GetInstanceId(), req.GetName())
	gwReq := c.gwc.NewRequest(http.MethodPut, path).SetBody(req)
	return gwclient.DoRequest[argocdv1.UpdateInstanceAccountPasswordResponse](ctx, gwReq)
}

func (c *argoCDV1Client) RegenerateInstanceAccountPassword(
	ctx context.Context,
	req *argocdv1.RegenerateInstanceAccountPasswordRequest,
	_ ...grpc.CallOption,
) (*argocdv1.RegenerateInstanceAccountPasswordResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s/accounts/%s/regenerate-password", req.GetOrganizationId(), req.GetInstanceId(), req.GetName())
	gwReq := c.gwc.NewRequest(http.MethodPost, path).SetBody(req)
	return gwclient.DoRequest[argocdv1.RegenerateInstanceAccountPasswordResponse](ctx, gwReq)
}

func (c *argoCDV1Client) DeleteInstanceAccount(
	ctx context.Context,
	req *argocdv1.DeleteInstanceAccountRequest,
	_ ...grpc.CallOption,
) (*argocdv1.DeleteInstanceAccountResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s/accounts/%s", req.GetOrganizationId(), req.GetInstanceId(), req.GetName())
	gwReq := c.gwc.NewRequest(http.MethodDelete, path)
	return gwclient.DoRequest[argocdv1.DeleteInstanceAccountResponse](ctx, gwReq)
}

func (c *argoCDV1Client) ListInstanceClusters(
	ctx context.Context,
	req *argocdv1.ListInstanceClustersRequest,
	_ ...grpc.CallOption,
) (*argocdv1.ListInstanceClustersResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s/clusters", req.GetOrganizationId(), req.GetInstanceId())
	gwReq := c.gwc.NewRequest(http.MethodGet, path)
	return gwclient.DoRequest[argocdv1.ListInstanceClustersResponse](ctx, gwReq)
}

func (c *argoCDV1Client) CreateInstanceCluster(
	ctx context.Context,
	req *argocdv1.CreateInstanceClusterRequest,
	_ ...grpc.CallOption,
) (*argocdv1.CreateInstanceClusterResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s/clusters", req.GetOrganizationId(), req.GetInstanceId())
	gwReq := c.gwc.NewRequest(http.MethodPost, path).SetBody(req)
	return gwclient.DoRequest[argocdv1.CreateInstanceClusterResponse](ctx, gwReq)
}

func (c *argoCDV1Client) GetInstanceCluster(
	ctx context.Context,
	req *argocdv1.GetInstanceClusterRequest,
	_ ...grpc.CallOption,
) (*argocdv1.GetInstanceClusterResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s/clusters/%s", req.GetOrganizationId(), req.GetInstanceId(), req.GetId())
	gwReq := c.gwc.NewRequest(http.MethodGet, path).
		SetQueryParam("id_type", req.GetIdType().String())
	return gwclient.DoRequest[argocdv1.GetInstanceClusterResponse](ctx, gwReq)
}

func (c *argoCDV1Client) GetInstanceClusterManifests(
	ctx context.Context,
	req *argocdv1.GetInstanceClusterManifestsRequest,
	_ ...grpc.CallOption,
) (*argocdv1.GetInstanceClusterManifestsResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s/clusters/%s/manifests", req.GetOrganizationId(), req.GetInstanceId(), req.GetId())
	gwReq := c.gwc.NewRequest(http.MethodGet, path)
	return gwclient.DoRequest[argocdv1.GetInstanceClusterManifestsResponse](ctx, gwReq)
}

func (c *argoCDV1Client) UpdateInstanceCluster(
	ctx context.Context,
	req *argocdv1.UpdateInstanceClusterRequest,
	_ ...grpc.CallOption,
) (*argocdv1.UpdateInstanceClusterResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s/clusters/%s", req.GetOrganizationId(), req.GetInstanceId(), req.GetId())
	gwReq := c.gwc.NewRequest(http.MethodPut, path).SetBody(req)
	return gwclient.DoRequest[argocdv1.UpdateInstanceClusterResponse](ctx, gwReq)
}

func (c *argoCDV1Client) DeleteInstanceCluster(
	ctx context.Context,
	req *argocdv1.DeleteInstanceClusterRequest,
	_ ...grpc.CallOption,
) (*argocdv1.DeleteInstanceClusterResponse, error) {
	path := fmt.Sprintf("/api/v1/orgs/%s/argocd/instances/%s/clusters/%s", req.GetOrganizationId(), req.GetInstanceId(), req.GetId())
	gwReq := c.gwc.NewRequest(http.MethodDelete, path)
	return gwclient.DoRequest[argocdv1.DeleteInstanceClusterResponse](ctx, gwReq)
}
