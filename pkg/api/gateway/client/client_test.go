package client

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/protobuf/proto"

	httputil "github.com/akuity/api-client-go/pkg/utils/http"
)

func TestDoRequest(t *testing.T) {
	type testSet[T any] struct {
		newServerFn     func(http.Handler) *httptest.Server
		clientOptions   func(*httptest.Server) []Option
		req             proto.Message
		expectedReq     string
		execErrExpected bool
		res             string
		expectedRes     *T
	}
	testSets := map[string]testSet[grpc_health_v1.HealthCheckResponse]{
		"use http": {
			newServerFn: httptest.NewServer,
			clientOptions: func(srv *httptest.Server) []Option {
				return []Option{
					WithHTTPClient(srv.Client()),
				}
			},
			req: &grpc_health_v1.HealthCheckRequest{Service: "test"},
			expectedReq: `{
  "service":"test"
}`,
			execErrExpected: false,
			res: `{
  "status": "SERVING"
}`,
			expectedRes: &grpc_health_v1.HealthCheckResponse{
				Status: grpc_health_v1.HealthCheckResponse_SERVING,
			},
		},
		"use https": {
			newServerFn: httptest.NewTLSServer,
			clientOptions: func(srv *httptest.Server) []Option {
				return []Option{
					WithHTTPClient(srv.Client()),
				}
			},
			req: &grpc_health_v1.HealthCheckRequest{Service: "test"},
			expectedReq: `{
  "service":"test"
}`,
			execErrExpected: false,
			res: `{
  "status": "SERVING"
}`,
			expectedRes: &grpc_health_v1.HealthCheckResponse{
				Status: grpc_health_v1.HealthCheckResponse_SERVING,
			},
		},
		"use https with skip-tls-verify": {
			newServerFn: httptest.NewTLSServer,
			clientOptions: func(srv *httptest.Server) []Option {
				return []Option{
					SkipTLSVerify(true),
				}
			},
			req: &grpc_health_v1.HealthCheckRequest{Service: "test"},
			expectedReq: `{
  "service":"test"
}`,
			execErrExpected: false,
			res: `{
  "status": "SERVING"
}`,
			expectedRes: &grpc_health_v1.HealthCheckResponse{
				Status: grpc_health_v1.HealthCheckResponse_SERVING,
			},
		},
	}

	for name, ts := range testSets {
		t.Run(name, func(t *testing.T) {
			srv := ts.newServerFn(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				require.Equal(t, http.MethodPost, r.Method)
				require.Equal(t, "/api/v1/some-endpoint", r.RequestURI)
				reqBody, err := io.ReadAll(r.Body)
				require.NoError(t, err)
				require.JSONEq(t, ts.expectedReq, string(reqBody))

				w.Header().Set("Content-Type", httputil.MIMEApplicationJSON)
				_, err = w.Write([]byte(ts.res))
				require.NoError(t, err)
			}))
			defer srv.Close()

			c := NewClient(srv.URL+"/api/v1", ts.clientOptions(srv)...)
			req := c.NewRequest(http.MethodPost, "/some-endpoint").SetBody(ts.req)
			res, err := DoRequest[grpc_health_v1.HealthCheckResponse](context.TODO(), req)
			if ts.execErrExpected {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.True(t, proto.Equal(res, ts.expectedRes))
		})
	}
}
