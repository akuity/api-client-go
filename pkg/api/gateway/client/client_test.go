package client

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	apikeyv1 "github.com/akuity/api-client-go/pkg/api/gen/apikey/v1"
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
	testSets := map[string]testSet[apikeyv1.GetAPIKeyResponse]{
		"use http": {
			newServerFn: httptest.NewServer,
			clientOptions: func(srv *httptest.Server) []Option {
				return []Option{
					WithHTTPClient(srv.Client()),
				}
			},
			req: &apikeyv1.GetAPIKeyRequest{Id: "abcdef"},
			expectedReq: `{
  "id":"abcdef"
}`,
			execErrExpected: false,
			res: `{
  "api_key": {
    "id": "abcdef"
  }
}`,
			expectedRes: &apikeyv1.GetAPIKeyResponse{
				ApiKey: &apikeyv1.APIKey{Id: "abcdef"},
			},
		},
		"use https": {
			newServerFn: httptest.NewTLSServer,
			clientOptions: func(srv *httptest.Server) []Option {
				return []Option{
					WithHTTPClient(srv.Client()),
				}
			},
			req: &apikeyv1.GetAPIKeyRequest{Id: "abcdef"},
			expectedReq: `{
  "id":"abcdef"
}`,
			execErrExpected: false,
			res: `{
  "api_key": {
    "id": "abcdef"
  }
}`,
			expectedRes: &apikeyv1.GetAPIKeyResponse{
				ApiKey: &apikeyv1.APIKey{Id: "abcdef"},
			},
		},
		"use https with skip-tls-verify": {
			newServerFn: httptest.NewTLSServer,
			clientOptions: func(srv *httptest.Server) []Option {
				return []Option{
					SkipTLSVerify(true),
				}
			},
			req: &apikeyv1.GetAPIKeyRequest{Id: "abcdef"},
			expectedReq: `{
  "id":"abcdef"
}`,
			execErrExpected: false,
			res: `{
  "api_key": {
    "id": "abcdef"
  }
}`,
			expectedRes: &apikeyv1.GetAPIKeyResponse{
				ApiKey: &apikeyv1.APIKey{Id: "abcdef"},
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
			res, err := DoRequest[apikeyv1.GetAPIKeyResponse](context.TODO(), req)
			if ts.execErrExpected {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.True(t, proto.Equal(res, ts.expectedRes))
		})
	}
}
