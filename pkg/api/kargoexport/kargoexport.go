// Package kargoexport holds helpers shared by the Kargo export API server and
// its clients (the akuity CLI and the Terraform provider) for translating
// between the streamed, one-resource-per-message ExportKargoInstanceStreamResponse
// and the aggregated ExportKargoInstanceResponse.
package kargoexport

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	kargov1 "github.com/akuity/api-client-go/pkg/api/gen/kargo/v1"
)

// AppendStreamResponse folds one ExportKargoInstanceStreamResponse into res:
// singular fields are assigned, repeated fields appended, and a message with no
// resource set is a no-op. The unary server builds its response with this same
// fold, so streaming and unary results cannot drift. Exhaustiveness over every
// oneof variant is enforced by Test_AppendStreamResponse_CoversEveryVariant.
func AppendStreamResponse(res *kargov1.ExportKargoInstanceResponse, msg *kargov1.ExportKargoInstanceStreamResponse) {
	switch r := msg.GetResource().(type) {
	case *kargov1.ExportKargoInstanceStreamResponse_Kargo:
		res.Kargo = r.Kargo
	case *kargov1.ExportKargoInstanceStreamResponse_Agent:
		res.Agents = append(res.Agents, r.Agent)
	case *kargov1.ExportKargoInstanceStreamResponse_KargoConfigmap:
		res.KargoConfigmap = r.KargoConfigmap
	case *kargov1.ExportKargoInstanceStreamResponse_Project:
		res.Projects = append(res.Projects, r.Project)
	case *kargov1.ExportKargoInstanceStreamResponse_Warehouse:
		res.Warehouses = append(res.Warehouses, r.Warehouse)
	case *kargov1.ExportKargoInstanceStreamResponse_Stage:
		res.Stages = append(res.Stages, r.Stage)
	case *kargov1.ExportKargoInstanceStreamResponse_AnalysisTemplate:
		res.AnalysisTemplates = append(res.AnalysisTemplates, r.AnalysisTemplate)
	case *kargov1.ExportKargoInstanceStreamResponse_PromotionTask:
		res.PromotionTasks = append(res.PromotionTasks, r.PromotionTask)
	case *kargov1.ExportKargoInstanceStreamResponse_ClusterPromotionTask:
		res.ClusterPromotionTasks = append(res.ClusterPromotionTasks, r.ClusterPromotionTask)
	case *kargov1.ExportKargoInstanceStreamResponse_ServiceAccount:
		res.ServiceAccounts = append(res.ServiceAccounts, r.ServiceAccount)
	case *kargov1.ExportKargoInstanceStreamResponse_Role:
		res.Roles = append(res.Roles, r.Role)
	case *kargov1.ExportKargoInstanceStreamResponse_RoleBinding:
		res.RoleBindings = append(res.RoleBindings, r.RoleBinding)
	case *kargov1.ExportKargoInstanceStreamResponse_Configmap:
		res.Configmaps = append(res.Configmaps, r.Configmap)
	case *kargov1.ExportKargoInstanceStreamResponse_ProjectConfig:
		res.ProjectConfigs = append(res.ProjectConfigs, r.ProjectConfig)
	case *kargov1.ExportKargoInstanceStreamResponse_MessageChannel:
		res.MessageChannels = append(res.MessageChannels, r.MessageChannel)
	case *kargov1.ExportKargoInstanceStreamResponse_ClusterMessageChannel:
		res.ClusterMessageChannels = append(res.ClusterMessageChannels, r.ClusterMessageChannel)
	case *kargov1.ExportKargoInstanceStreamResponse_EventRouter:
		res.EventRouters = append(res.EventRouters, r.EventRouter)
	case *kargov1.ExportKargoInstanceStreamResponse_CustomPromotionStep:
		res.CustomPromotionSteps = append(res.CustomPromotionSteps, r.CustomPromotionStep)
	case *kargov1.ExportKargoInstanceStreamResponse_ClusterConfig:
		res.ClusterConfigs = append(res.ClusterConfigs, r.ClusterConfig)
	}
}

// ExportClient is the subset of kargov1.KargoServiceGatewayClient needed to
// export an instance.
//
// TODO(export-stream): the unary method, the fallback in ExportKargoInstance,
// and streamEndpointUnsupported exist only for servers that predate the
// streaming endpoint. Delete them once the minimum supported server version
// serves streaming.
type ExportClient interface {
	ExportKargoInstance(context.Context, *kargov1.ExportKargoInstanceRequest) (*kargov1.ExportKargoInstanceResponse, error)
	ExportKargoInstanceStream(context.Context, *kargov1.ExportKargoInstanceStreamRequest) (<-chan *kargov1.ExportKargoInstanceStreamResponse, <-chan error, error)
}

// ExportKargoInstance calls the streaming export RPC and reassembles the chunks
// into a single ExportKargoInstanceResponse, giving callers (the CLI export/diff
// commands, the Terraform provider) the unary response shape without the unary
// RPC's message-size ceiling. It accepts the unary request type so callers keep
// building the request they always did.
//
// If the initial call fails in a way that suggests the server predates the
// streaming endpoint (see streamEndpointUnsupported), it falls back to the
// unary RPC. An error after the stream is established is returned as-is — never
// retried as a unary call that would only die on the message-size ceiling.
func ExportKargoInstance(ctx context.Context, c ExportClient, req *kargov1.ExportKargoInstanceRequest) (*kargov1.ExportKargoInstanceResponse, error) {
	respCh, errCh, err := c.ExportKargoInstanceStream(ctx, &kargov1.ExportKargoInstanceStreamRequest{
		OrganizationId: req.GetOrganizationId(),
		Id:             req.GetId(),
		WorkspaceId:    req.GetWorkspaceId(),
	})
	if err != nil {
		if streamEndpointUnsupported(err) {
			return c.ExportKargoInstance(ctx, req)
		}
		return nil, err
	}
	res := &kargov1.ExportKargoInstanceResponse{}
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

// streamEndpointUnsupported reports whether an initial ExportKargoInstanceStream
// error warrants falling back to the unary RPC. True for the codes a server
// lacking the endpoint produces: NotFound (grpc-gateway routing 404),
// Unimplemented (pure-gRPC missing method, routing 405), and Unknown
// (unparseable error body, e.g. a proxy's non-JSON 404 page). A genuine
// "instance not found" from a new server is indistinguishable from the routing
// 404 and also falls back — harmless, since the unary path fails fast with the
// same error before building any large response. Every other code is returned
// as-is: deterministic errors (PermissionDenied, ...) arrive typed from the
// gateway client, and a client-side cancellation or deadline must never be
// turned into a fresh unary export.
func streamEndpointUnsupported(err error) bool {
	switch status.Code(err) {
	case codes.NotFound, codes.Unknown, codes.Unimplemented:
		return true
	default:
		return false
	}
}
