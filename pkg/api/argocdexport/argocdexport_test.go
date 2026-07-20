package argocdexport

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"

	argocdv1 "github.com/akuity/api-client-go/pkg/api/gen/argocd/v1"
)

func mkStruct(t *testing.T, name string) *structpb.Struct {
	t.Helper()
	s, err := structpb.NewStruct(map[string]any{"metadata": map[string]any{"name": name}})
	require.NoError(t, err)
	return s
}

// chunks builds one ExportInstanceStreamResponse per resource covering every oneof variant, in the
// order the server emits them, with repeated kinds appearing more than once.
func chunks(t *testing.T) []*argocdv1.ExportInstanceStreamResponse {
	t.Helper()
	return []*argocdv1.ExportInstanceStreamResponse{
		{Resource: &argocdv1.ExportInstanceStreamResponse_Argocd{Argocd: mkStruct(t, "argocd")}},
		{Resource: &argocdv1.ExportInstanceStreamResponse_Cluster{Cluster: mkStruct(t, "cluster-1")}},
		{Resource: &argocdv1.ExportInstanceStreamResponse_Cluster{Cluster: mkStruct(t, "cluster-2")}},
		{Resource: &argocdv1.ExportInstanceStreamResponse_ConfigManagementPlugin{ConfigManagementPlugin: mkStruct(t, "cmp-1")}},
		{Resource: &argocdv1.ExportInstanceStreamResponse_ArgocdConfigmap{ArgocdConfigmap: mkStruct(t, "argocd-cm")}},
		{Resource: &argocdv1.ExportInstanceStreamResponse_NotificationsConfigmap{NotificationsConfigmap: mkStruct(t, "notifications-cm")}},
		{Resource: &argocdv1.ExportInstanceStreamResponse_ArgocdRbacConfigmap{ArgocdRbacConfigmap: mkStruct(t, "rbac-cm")}},
		{Resource: &argocdv1.ExportInstanceStreamResponse_ImageUpdaterConfigmap{ImageUpdaterConfigmap: mkStruct(t, "iu-cm")}},
		{Resource: &argocdv1.ExportInstanceStreamResponse_ImageUpdaterSshConfigmap{ImageUpdaterSshConfigmap: mkStruct(t, "iu-ssh-cm")}},
		{Resource: &argocdv1.ExportInstanceStreamResponse_ArgocdKnownHostsConfigmap{ArgocdKnownHostsConfigmap: mkStruct(t, "known-hosts-cm")}},
		{Resource: &argocdv1.ExportInstanceStreamResponse_ArgocdTlsCertsConfigmap{ArgocdTlsCertsConfigmap: mkStruct(t, "tls-certs-cm")}},
		{Resource: &argocdv1.ExportInstanceStreamResponse_Application{Application: mkStruct(t, "app-1")}},
		{Resource: &argocdv1.ExportInstanceStreamResponse_Application{Application: mkStruct(t, "app-2")}},
		{Resource: &argocdv1.ExportInstanceStreamResponse_ApplicationSet{ApplicationSet: mkStruct(t, "appset-1")}},
		{Resource: &argocdv1.ExportInstanceStreamResponse_AppProject{AppProject: mkStruct(t, "proj-1")}},
		{Resource: &argocdv1.ExportInstanceStreamResponse_ImageUpdater{ImageUpdater: mkStruct(t, "iu-1")}},
	}
}

func structName(s *structpb.Struct) string {
	if s == nil {
		return ""
	}
	return s.AsMap()["metadata"].(map[string]any)["name"].(string)
}

func assertReassembled(t *testing.T, res *argocdv1.ExportInstanceResponse) {
	t.Helper()
	// singular fields
	require.Equal(t, "argocd", structName(res.Argocd))
	require.Equal(t, "argocd-cm", structName(res.ArgocdConfigmap))
	require.Equal(t, "notifications-cm", structName(res.NotificationsConfigmap))
	require.Equal(t, "rbac-cm", structName(res.ArgocdRbacConfigmap))
	require.Equal(t, "iu-cm", structName(res.ImageUpdaterConfigmap))
	require.Equal(t, "iu-ssh-cm", structName(res.ImageUpdaterSshConfigmap))
	require.Equal(t, "known-hosts-cm", structName(res.ArgocdKnownHostsConfigmap))
	require.Equal(t, "tls-certs-cm", structName(res.ArgocdTlsCertsConfigmap))
	// repeated fields preserve order and count
	require.Equal(t, []string{"cluster-1", "cluster-2"}, names(res.Clusters))
	require.Equal(t, []string{"cmp-1"}, names(res.ConfigManagementPlugins))
	require.Equal(t, []string{"app-1", "app-2"}, names(res.Applications))
	require.Equal(t, []string{"appset-1"}, names(res.ApplicationSets))
	require.Equal(t, []string{"proj-1"}, names(res.AppProjects))
	require.Equal(t, []string{"iu-1"}, names(res.ImageUpdaters))
}

func names(structs []*structpb.Struct) []string {
	var out []string
	for _, s := range structs {
		out = append(out, structName(s))
	}
	return out
}

func Test_AppendStreamResponse(t *testing.T) {
	res := &argocdv1.ExportInstanceResponse{}
	for _, c := range chunks(t) {
		AppendStreamResponse(res, c)
	}
	assertReassembled(t, res)
}

func Test_AppendStreamResponse_EmptyMessageIsNoop(t *testing.T) {
	res := &argocdv1.ExportInstanceResponse{}
	AppendStreamResponse(res, &argocdv1.ExportInstanceStreamResponse{})
	require.Nil(t, res.Argocd)
	require.Empty(t, res.Applications)
}

// fakeStreamClient mimics the gateway client contract: ExportInstanceStream closes the response
// channel on success (EOF) and, if err is set, delivers a single error on the error channel without
// closing the response channel — exactly how grpc-gateway-client's DoStreamingRequest behaves. If
// initErr is set it is returned as the initial error (as an older server lacking the endpoint
// would), and ExportInstance returns the unary response so the fallback can be exercised.
type fakeStreamClient struct {
	msgs    []*argocdv1.ExportInstanceStreamResponse
	err     error
	initErr error

	unaryResp   *argocdv1.ExportInstanceResponse
	unaryErr    error
	unaryCalled bool
}

func (f *fakeStreamClient) ExportInstance(context.Context, *argocdv1.ExportInstanceRequest) (*argocdv1.ExportInstanceResponse, error) {
	f.unaryCalled = true
	return f.unaryResp, f.unaryErr
}

func (f *fakeStreamClient) ExportInstanceStream(ctx context.Context, _ *argocdv1.ExportInstanceStreamRequest) (<-chan *argocdv1.ExportInstanceStreamResponse, <-chan error, error) {
	if f.initErr != nil {
		return nil, nil, f.initErr
	}
	respCh := make(chan *argocdv1.ExportInstanceStreamResponse)
	// Unbuffered, matching grpc-gateway-client's DoStreamingRequest: it closes respCh on
	// EOF (success) or sends exactly one error on an unbuffered errCh without closing respCh.
	errCh := make(chan error)
	go func() {
		for _, m := range f.msgs {
			select {
			case <-ctx.Done():
				return
			case respCh <- m:
			}
		}
		if f.err != nil {
			errCh <- f.err
			return
		}
		close(respCh)
	}()
	return respCh, errCh, nil
}

func Test_ExportInstance_ReassemblesStream(t *testing.T) {
	c := &fakeStreamClient{msgs: chunks(t)}
	res, err := ExportInstance(context.Background(), c, &argocdv1.ExportInstanceRequest{})
	require.NoError(t, err)
	assertReassembled(t, res)
	require.False(t, c.unaryCalled, "must not fall back to unary when the stream works")
}

func Test_ExportInstance_ReturnsStreamError(t *testing.T) {
	wantErr := errors.New("boom")
	c := &fakeStreamClient{
		msgs: []*argocdv1.ExportInstanceStreamResponse{
			{Resource: &argocdv1.ExportInstanceStreamResponse_Application{Application: mkStruct(t, "app-1")}},
		},
		err: wantErr,
	}
	res, err := ExportInstance(context.Background(), c, &argocdv1.ExportInstanceRequest{})
	require.ErrorIs(t, err, wantErr)
	require.Nil(t, res)
	require.False(t, c.unaryCalled, "a mid-stream error must NOT trigger the unary fallback")
}

// Test_ExportInstance_FallsBackToUnaryOnInitError covers an initial error the gateway client could
// not parse into a status (codes.Unknown) — e.g. a proxy or load balancer answering the missing
// route on an old server with a non-JSON 404 page. The helper transparently falls back to the
// unary RPC.
func Test_ExportInstance_FallsBackToUnaryOnInitError(t *testing.T) {
	unary := &argocdv1.ExportInstanceResponse{Argocd: mkStruct(t, "argocd")}
	c := &fakeStreamClient{
		initErr:   errors.New("unmarshal raw response: invalid character 'd'"),
		unaryResp: unary,
	}
	res, err := ExportInstance(context.Background(), c, &argocdv1.ExportInstanceRequest{})
	require.NoError(t, err)
	require.True(t, c.unaryCalled, "must fall back to unary when the stream can't be established")
	require.Equal(t, "argocd", structName(res.Argocd))
}

// Test_ExportInstance_FallsBackToUnaryOnNotFound covers an older server that does not serve the
// streaming endpoint, reached through the current grpc-gateway-client: the server's grpc-gateway
// routing 404 ({"code":5,"message":"Not Found"}) parses cleanly now that the client strips the SSE
// "data:" framing, so the initial error arrives as a typed codes.NotFound rather than codes.Unknown.
// The helper must still fall back to unary.
func Test_ExportInstance_FallsBackToUnaryOnNotFound(t *testing.T) {
	unary := &argocdv1.ExportInstanceResponse{Argocd: mkStruct(t, "argocd")}
	c := &fakeStreamClient{
		initErr:   status.Error(codes.NotFound, "Not Found"),
		unaryResp: unary,
	}
	res, err := ExportInstance(context.Background(), c, &argocdv1.ExportInstanceRequest{})
	require.NoError(t, err)
	require.True(t, c.unaryCalled, "a NotFound initial error (old server's routing 404) must fall back to unary")
	require.Equal(t, "argocd", structName(res.Argocd))
}

// Test_ExportInstance_DoesNotFallBackOnDistinctCode covers initial errors carrying a distinct gRPC
// code: client-side cancellation/deadline, and deterministic server errors, which the gateway
// client delivers typed (it strips the SSE "data:" framing before parsing error bodies). None of
// these may be turned into a fresh unary export.
func Test_ExportInstance_DoesNotFallBackOnDistinctCode(t *testing.T) {
	for _, code := range []codes.Code{codes.Canceled, codes.DeadlineExceeded, codes.PermissionDenied, codes.InvalidArgument} {
		t.Run(code.String(), func(t *testing.T) {
			c := &fakeStreamClient{initErr: status.Error(code, "initial error")}
			res, err := ExportInstance(context.Background(), c, &argocdv1.ExportInstanceRequest{})
			require.Equal(t, code, status.Code(err))
			require.Nil(t, res)
			require.False(t, c.unaryCalled, "an error carrying a distinct gRPC code must not trigger unary fallback")
		})
	}
}

func Test_ExportInstance_FallbackPropagatesUnaryError(t *testing.T) {
	unaryErr := errors.New("not found")
	c := &fakeStreamClient{
		initErr:  errors.New("stream unavailable"),
		unaryErr: unaryErr,
	}
	res, err := ExportInstance(context.Background(), c, &argocdv1.ExportInstanceRequest{})
	require.ErrorIs(t, err, unaryErr)
	require.Nil(t, res)
	require.True(t, c.unaryCalled)
}
