package client

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"

	ctxutil "github.com/akuity/api-client-go/pkg/utils/context"
	httputil "github.com/akuity/api-client-go/pkg/utils/http"
)

var (
	_ http.RoundTripper = &gatewayTransport{}
)

type gatewayTransport struct {
	rt http.RoundTripper
}

func (g *gatewayTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	cred := ctxutil.GetClientCredential(ctx)
	if cred == nil {
		return g.rt.RoundTrip(req)
	}

	outReq := req.Clone(ctx)
	outReq.Header.Del(httputil.HeaderAuthorization)
	outReq.Header.Set(httputil.HeaderAuthorization, fmt.Sprintf("%s %s", cred.Scheme(), cred.Credential()))
	return g.rt.RoundTrip(outReq)
}

func newGatewayHTTPClient(baseClient *http.Client, skipTLSVerify bool) *http.Client {
	rt := baseClient.Transport
	if rt == nil {
		rt = http.DefaultTransport.(*http.Transport).Clone()
	}
	if ht, ok := rt.(*http.Transport); ok {
		if ht.TLSClientConfig == nil {
			ht.TLSClientConfig = &tls.Config{}
		}
		ht.TLSClientConfig.InsecureSkipVerify = skipTLSVerify
	}
	baseClient.Transport = &gatewayTransport{
		rt: rt,
	}
	return baseClient
}

// HTTPStatusToCode converts HTTP status code to gRPC error code
// https://github.com/grpc/grpc/blob/9d2a1a3d1aba56d94c56d2e01cf2511d1a082445/doc/http-grpc-status-mapping.md
func HTTPStatusToCode(code int) codes.Code {
	switch code {
	case http.StatusOK:
		return codes.OK
	case http.StatusBadRequest:
		return codes.Internal
	case http.StatusUnauthorized:
		return codes.Unauthenticated
	case http.StatusForbidden:
		return codes.PermissionDenied
	case http.StatusNotFound:
		return codes.Unimplemented
	case http.StatusTooManyRequests:
		return codes.Unavailable
	case http.StatusBadGateway:
		return codes.Unavailable
	case http.StatusServiceUnavailable:
		return codes.Unavailable
	case http.StatusGatewayTimeout:
		return codes.Unavailable
	default:
		return codes.Unknown
	}
}
