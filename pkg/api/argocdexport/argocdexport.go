// Package argocdexport holds helpers shared by the ArgoCD export API server and
// its clients (the akuity CLI and the Terraform provider) for translating between
// the streamed, one-resource-per-message ExportInstanceStreamResponse and the
// aggregated ExportInstanceResponse. See
// https://github.com/akuityio/akuity-platform/issues/3754.
package argocdexport

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	argocdv1 "github.com/akuity/api-client-go/pkg/api/gen/argocd/v1"
)

// AppendStreamResponse folds one ExportInstanceStreamResponse into res: singular
// fields are assigned, repeated fields appended, and a message with no resource
// set is a no-op. The unary server builds its response with this same fold, so
// streaming and unary results cannot drift.
func AppendStreamResponse(res *argocdv1.ExportInstanceResponse, msg *argocdv1.ExportInstanceStreamResponse) {
	switch r := msg.GetResource().(type) {
	case *argocdv1.ExportInstanceStreamResponse_Argocd:
		res.Argocd = r.Argocd
	case *argocdv1.ExportInstanceStreamResponse_ArgocdConfigmap:
		res.ArgocdConfigmap = r.ArgocdConfigmap
	case *argocdv1.ExportInstanceStreamResponse_ArgocdRbacConfigmap:
		res.ArgocdRbacConfigmap = r.ArgocdRbacConfigmap
	case *argocdv1.ExportInstanceStreamResponse_NotificationsConfigmap:
		res.NotificationsConfigmap = r.NotificationsConfigmap
	case *argocdv1.ExportInstanceStreamResponse_ImageUpdaterConfigmap:
		res.ImageUpdaterConfigmap = r.ImageUpdaterConfigmap
	case *argocdv1.ExportInstanceStreamResponse_ImageUpdaterSshConfigmap:
		res.ImageUpdaterSshConfigmap = r.ImageUpdaterSshConfigmap
	case *argocdv1.ExportInstanceStreamResponse_Cluster:
		res.Clusters = append(res.Clusters, r.Cluster)
	case *argocdv1.ExportInstanceStreamResponse_ArgocdKnownHostsConfigmap:
		res.ArgocdKnownHostsConfigmap = r.ArgocdKnownHostsConfigmap
	case *argocdv1.ExportInstanceStreamResponse_ArgocdTlsCertsConfigmap:
		res.ArgocdTlsCertsConfigmap = r.ArgocdTlsCertsConfigmap
	case *argocdv1.ExportInstanceStreamResponse_ConfigManagementPlugin:
		res.ConfigManagementPlugins = append(res.ConfigManagementPlugins, r.ConfigManagementPlugin)
	case *argocdv1.ExportInstanceStreamResponse_Application:
		res.Applications = append(res.Applications, r.Application)
	case *argocdv1.ExportInstanceStreamResponse_ApplicationSet:
		res.ApplicationSets = append(res.ApplicationSets, r.ApplicationSet)
	case *argocdv1.ExportInstanceStreamResponse_AppProject:
		res.AppProjects = append(res.AppProjects, r.AppProject)
	case *argocdv1.ExportInstanceStreamResponse_ImageUpdater:
		res.ImageUpdaters = append(res.ImageUpdaters, r.ImageUpdater)
	}
}

// ExportClient is the subset of argocdv1.ArgoCDServiceGatewayClient needed to
// export an instance.
//
// TODO(export-stream): the unary method, the fallback in ExportInstance, and
// streamEndpointUnsupported exist only for servers that predate the streaming
// endpoint. Delete them once the minimum supported server version serves
// streaming (https://github.com/akuityio/akuity-platform/issues/3754).
type ExportClient interface {
	ExportInstance(context.Context, *argocdv1.ExportInstanceRequest) (*argocdv1.ExportInstanceResponse, error)
	ExportInstanceStream(context.Context, *argocdv1.ExportInstanceStreamRequest) (<-chan *argocdv1.ExportInstanceStreamResponse, <-chan error, error)
}

// ExportInstance calls the streaming export RPC and reassembles the chunks into
// a single ExportInstanceResponse, giving callers (the CLI export/diff commands,
// the Terraform provider) the unary response shape without its message-size
// ceiling (20MB on the server's internal gateway-to-gRPC hop, see
// https://github.com/akuityio/akuity-platform/issues/3639). It accepts the unary
// request type so callers keep building the request they always did.
//
// If the initial call fails in a way that suggests the server predates the
// streaming endpoint (see streamEndpointUnsupported), it falls back to the unary
// RPC. An error after the stream is established is returned as-is — never
// retried as a unary call that would only die on the 20MB ceiling.
func ExportInstance(ctx context.Context, c ExportClient, req *argocdv1.ExportInstanceRequest) (*argocdv1.ExportInstanceResponse, error) {
	respCh, errCh, err := c.ExportInstanceStream(ctx, &argocdv1.ExportInstanceStreamRequest{
		OrganizationId: req.GetOrganizationId(),
		IdType:         req.GetIdType(),
		Id:             req.GetId(),
		WorkspaceId:    req.GetWorkspaceId(),
	})
	if err != nil {
		if streamEndpointUnsupported(err) {
			return c.ExportInstance(ctx, req)
		}
		return nil, err
	}
	res := &argocdv1.ExportInstanceResponse{}
	// The streaming client closes respCh on success and delivers exactly one
	// error (context cancellation included) on errCh, which is never closed.
	for {
		select {
		case err := <-errCh:
			return nil, err
		case msg, ok := <-respCh:
			if !ok {
				return res, nil
			}
			AppendStreamResponse(res, msg)
		}
	}
}

// streamEndpointUnsupported reports whether an initial ExportInstanceStream error
// warrants falling back to the unary RPC. True for the codes a server lacking the
// endpoint produces: NotFound (grpc-gateway routing 404), Unimplemented (pure-gRPC
// missing method, routing 405), and Unknown (unparseable error body, e.g. a
// proxy's non-JSON 404 page). A genuine "instance not found" from a new server is
// indistinguishable from the routing 404 and also falls back — harmless, since
// the unary path fails fast with the same error before building any large
// response. Every other code is returned as-is: deterministic errors
// (PermissionDenied, ...) arrive typed from the gateway client, and a client-side
// cancellation or deadline must never be turned into a fresh unary export.
func streamEndpointUnsupported(err error) bool {
	switch status.Code(err) {
	case codes.NotFound, codes.Unknown, codes.Unimplemented:
		return true
	default:
		return false
	}
}
