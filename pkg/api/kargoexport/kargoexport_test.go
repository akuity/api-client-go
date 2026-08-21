package kargoexport

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/structpb"

	kargov1 "github.com/akuity/api-client-go/pkg/api/gen/kargo/v1"
)

func mkStruct(t *testing.T, name string) *structpb.Struct {
	t.Helper()
	s, err := structpb.NewStruct(map[string]any{"metadata": map[string]any{"name": name}})
	require.NoError(t, err)
	return s
}

func structName(s *structpb.Struct) string {
	if s == nil {
		return ""
	}
	return s.AsMap()["metadata"].(map[string]any)["name"].(string)
}

func names(structs []*structpb.Struct) []string {
	var out []string
	for _, s := range structs {
		out = append(out, structName(s))
	}
	return out
}

// chunks builds one ExportKargoInstanceStreamResponse per resource in the order the
// server emits them, with repeated kinds appearing more than once. Exhaustiveness
// over every oneof variant is covered separately by
// Test_AppendStreamResponse_CoversEveryVariant.
func chunks(t *testing.T) []*kargov1.ExportKargoInstanceStreamResponse {
	t.Helper()
	return []*kargov1.ExportKargoInstanceStreamResponse{
		{Resource: &kargov1.ExportKargoInstanceStreamResponse_Kargo{Kargo: mkStruct(t, "kargo")}},
		{Resource: &kargov1.ExportKargoInstanceStreamResponse_Agent{Agent: mkStruct(t, "agent-1")}},
		{Resource: &kargov1.ExportKargoInstanceStreamResponse_Agent{Agent: mkStruct(t, "agent-2")}},
		{Resource: &kargov1.ExportKargoInstanceStreamResponse_KargoConfigmap{KargoConfigmap: mkStruct(t, "kargo-cm")}},
		{Resource: &kargov1.ExportKargoInstanceStreamResponse_Project{Project: mkStruct(t, "project-1")}},
		{Resource: &kargov1.ExportKargoInstanceStreamResponse_Project{Project: mkStruct(t, "project-2")}},
		{Resource: &kargov1.ExportKargoInstanceStreamResponse_Warehouse{Warehouse: mkStruct(t, "warehouse-1")}},
		{Resource: &kargov1.ExportKargoInstanceStreamResponse_Stage{Stage: mkStruct(t, "stage-1")}},
		{Resource: &kargov1.ExportKargoInstanceStreamResponse_ServiceAccount{ServiceAccount: mkStruct(t, "sa-1")}},
		{Resource: &kargov1.ExportKargoInstanceStreamResponse_Configmap{Configmap: mkStruct(t, "cm-1")}},
	}
}

func assertReassembled(t *testing.T, res *kargov1.ExportKargoInstanceResponse) {
	t.Helper()
	// singular fields
	require.Equal(t, "kargo", structName(res.Kargo))
	require.Equal(t, "kargo-cm", structName(res.KargoConfigmap))
	// repeated fields preserve order and count
	require.Equal(t, []string{"agent-1", "agent-2"}, names(res.Agents))
	require.Equal(t, []string{"project-1", "project-2"}, names(res.Projects))
	require.Equal(t, []string{"warehouse-1"}, names(res.Warehouses))
	require.Equal(t, []string{"stage-1"}, names(res.Stages))
	require.Equal(t, []string{"sa-1"}, names(res.ServiceAccounts))
	require.Equal(t, []string{"cm-1"}, names(res.Configmaps))
}

func Test_AppendStreamResponse(t *testing.T) {
	res := &kargov1.ExportKargoInstanceResponse{}
	for _, c := range chunks(t) {
		AppendStreamResponse(res, c)
	}
	assertReassembled(t, res)
}

func Test_AppendStreamResponse_EmptyMessageIsNoop(t *testing.T) {
	res := &kargov1.ExportKargoInstanceResponse{}
	AppendStreamResponse(res, &kargov1.ExportKargoInstanceStreamResponse{})
	require.Nil(t, res.Kargo)
	require.Empty(t, res.Projects)
}

// Test_StreamResponseFieldParity is the drift guard between the unary and streaming
// response shapes: every ExportKargoInstanceResponse field must have a
// same-numbered google.protobuf.Struct variant in the ExportKargoInstanceStreamResponse
// `resource` oneof, and vice versa. A developer adding a resource to the unary
// response cannot pass CI without adding the streaming variant.
func Test_StreamResponseFieldParity(t *testing.T) {
	unary := (&kargov1.ExportKargoInstanceResponse{}).ProtoReflect().Descriptor()
	stream := (&kargov1.ExportKargoInstanceStreamResponse{}).ProtoReflect().Descriptor()
	oneof := stream.Oneofs().ByName("resource")
	require.NotNil(t, oneof)

	const structType = protoreflect.FullName("google.protobuf.Struct")

	for i := 0; i < unary.Fields().Len(); i++ {
		fd := unary.Fields().Get(i)
		require.Equal(t, protoreflect.MessageKind, fd.Kind(), "unary field %s must be a Struct", fd.Name())
		require.Equal(t, structType, fd.Message().FullName(), "unary field %s must be a Struct", fd.Name())
		variant := stream.Fields().ByNumber(fd.Number())
		require.NotNil(t, variant,
			"unary field %s (#%d) has no stream oneof variant — add it to ExportKargoInstanceStreamResponse and a case to AppendStreamResponse",
			fd.Name(), fd.Number())
		require.NotNil(t, variant.ContainingOneof(), "stream field %s must live in the resource oneof", variant.Name())
		require.Equal(t, protoreflect.Name("resource"), variant.ContainingOneof().Name())
		require.Equal(t, structType, variant.Message().FullName())
	}
	for i := 0; i < oneof.Fields().Len(); i++ {
		fd := oneof.Fields().Get(i)
		require.NotNil(t, unary.Fields().ByNumber(fd.Number()),
			"stream variant %s (#%d) has no matching unary field", fd.Name(), fd.Number())
	}
	require.Equal(t, unary.Fields().Len(), oneof.Fields().Len())
}

// Test_AppendStreamResponse_CoversEveryVariant is the fold drift guard: every oneof
// variant, set via reflection, must populate the same-numbered unary field. A
// developer adding a variant but forgetting the AppendStreamResponse case fails here.
func Test_AppendStreamResponse_CoversEveryVariant(t *testing.T) {
	desc := (&kargov1.ExportKargoInstanceStreamResponse{}).ProtoReflect().Descriptor()
	oneof := desc.Oneofs().ByName("resource")
	require.NotNil(t, oneof)

	for i := 0; i < oneof.Fields().Len(); i++ {
		fd := oneof.Fields().Get(i)
		t.Run(string(fd.Name()), func(t *testing.T) {
			msg := &kargov1.ExportKargoInstanceStreamResponse{}
			msg.ProtoReflect().Set(fd, protoreflect.ValueOfMessage(mkStruct(t, "sentinel").ProtoReflect()))

			res := &kargov1.ExportKargoInstanceResponse{}
			AppendStreamResponse(res, msg)

			unaryFd := res.ProtoReflect().Descriptor().Fields().ByNumber(fd.Number())
			require.NotNil(t, unaryFd, "no unary field #%d — field parity is broken", fd.Number())
			require.True(t, res.ProtoReflect().Has(unaryFd),
				"AppendStreamResponse dropped variant %s — add its case to the switch", fd.Name())
		})
	}
}

// fakeStreamClient mimics the gateway client contract: ExportKargoInstanceStream
// closes the response channel on success (EOF) and, if err is set, delivers a single
// error on the error channel without closing the response channel — exactly how
// grpc-gateway-client's DoStreamingRequest behaves. If initErr is set it is returned
// as the initial error (as an older server lacking the endpoint would), and
// ExportKargoInstance returns the unary response so the fallback can be exercised.
type fakeStreamClient struct {
	msgs    []*kargov1.ExportKargoInstanceStreamResponse
	err     error
	initErr error

	unaryResp   *kargov1.ExportKargoInstanceResponse
	unaryErr    error
	unaryCalled bool
}

func (f *fakeStreamClient) ExportKargoInstance(context.Context, *kargov1.ExportKargoInstanceRequest) (*kargov1.ExportKargoInstanceResponse, error) {
	f.unaryCalled = true
	return f.unaryResp, f.unaryErr
}

func (f *fakeStreamClient) ExportKargoInstanceStream(ctx context.Context, _ *kargov1.ExportKargoInstanceStreamRequest) (<-chan *kargov1.ExportKargoInstanceStreamResponse, <-chan error, error) {
	if f.initErr != nil {
		return nil, nil, f.initErr
	}
	respCh := make(chan *kargov1.ExportKargoInstanceStreamResponse)
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

func Test_ExportKargoInstance_ReassemblesStream(t *testing.T) {
	c := &fakeStreamClient{msgs: chunks(t)}
	res, err := ExportKargoInstance(context.Background(), c, &kargov1.ExportKargoInstanceRequest{})
	require.NoError(t, err)
	assertReassembled(t, res)
	require.False(t, c.unaryCalled, "must not fall back to unary when the stream works")
}

func Test_ExportKargoInstance_ReturnsStreamError(t *testing.T) {
	wantErr := errors.New("boom")
	c := &fakeStreamClient{
		msgs: []*kargov1.ExportKargoInstanceStreamResponse{
			{Resource: &kargov1.ExportKargoInstanceStreamResponse_Project{Project: mkStruct(t, "project-1")}},
		},
		err: wantErr,
	}
	res, err := ExportKargoInstance(context.Background(), c, &kargov1.ExportKargoInstanceRequest{})
	require.ErrorIs(t, err, wantErr)
	require.Nil(t, res)
	require.False(t, c.unaryCalled, "a mid-stream error must NOT trigger the unary fallback")
}

func Test_ExportKargoInstance_FallsBackToUnaryOnInitError(t *testing.T) {
	unary := &kargov1.ExportKargoInstanceResponse{Kargo: mkStruct(t, "kargo")}
	c := &fakeStreamClient{
		initErr:   errors.New("unmarshal raw response: invalid character 'd'"),
		unaryResp: unary,
	}
	res, err := ExportKargoInstance(context.Background(), c, &kargov1.ExportKargoInstanceRequest{})
	require.NoError(t, err)
	require.True(t, c.unaryCalled, "must fall back to unary when the stream can't be established")
	require.Equal(t, "kargo", structName(res.Kargo))
}

func Test_ExportKargoInstance_FallsBackToUnaryOnNotFound(t *testing.T) {
	unary := &kargov1.ExportKargoInstanceResponse{Kargo: mkStruct(t, "kargo")}
	c := &fakeStreamClient{
		initErr:   status.Error(codes.NotFound, "Not Found"),
		unaryResp: unary,
	}
	res, err := ExportKargoInstance(context.Background(), c, &kargov1.ExportKargoInstanceRequest{})
	require.NoError(t, err)
	require.True(t, c.unaryCalled, "a NotFound initial error (old server's routing 404) must fall back to unary")
	require.Equal(t, "kargo", structName(res.Kargo))
}

func Test_ExportKargoInstance_DoesNotFallBackOnDistinctCode(t *testing.T) {
	for _, code := range []codes.Code{codes.Canceled, codes.DeadlineExceeded, codes.PermissionDenied, codes.InvalidArgument} {
		t.Run(code.String(), func(t *testing.T) {
			c := &fakeStreamClient{initErr: status.Error(code, "initial error")}
			res, err := ExportKargoInstance(context.Background(), c, &kargov1.ExportKargoInstanceRequest{})
			require.Equal(t, code, status.Code(err))
			require.Nil(t, res)
			require.False(t, c.unaryCalled, "an error carrying a distinct gRPC code must not trigger unary fallback")
		})
	}
}

func Test_ExportKargoInstance_FallbackPropagatesUnaryError(t *testing.T) {
	unaryErr := errors.New("not found")
	c := &fakeStreamClient{
		initErr:  errors.New("stream unavailable"),
		unaryErr: unaryErr,
	}
	res, err := ExportKargoInstance(context.Background(), c, &kargov1.ExportKargoInstanceRequest{})
	require.ErrorIs(t, err, unaryErr)
	require.Nil(t, res)
	require.True(t, c.unaryCalled)
}
