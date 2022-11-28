package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/status"

	gwoption "github.com/akuity/api-client-go/pkg/api/gateway/option"
)

type Client interface {
	NewRequest(method, url string) *resty.Request
	Raw() *resty.Client
}

var (
	_ Client = &client{}
)

type client struct {
	rc *resty.Client
}

func (c *client) NewRequest(method string, url string) *resty.Request {
	req := c.rc.NewRequest()
	req.Method = method
	req.URL = url
	return req
}

func (c *client) Raw() *resty.Client {
	return c.rc
}

func NewClient(baseURL string, opts ...Option) Client {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	if o.hc == nil {
		o.hc = &http.Client{}
	}
	rc := resty.NewWithClient(newGatewayHTTPClient(o.hc, o.skipTLSVerify))
	rc.SetBaseURL(baseURL)
	rc.JSONMarshal = gwoption.Marshaller.Marshal
	rc.JSONUnmarshal = gwoption.Marshaller.Unmarshal
	return &client{
		rc: rc,
	}
}

func DoRequest[T any](ctx context.Context, req *resty.Request) (*T, error) {
	var resBody T
	res, err := req.SetContext(ctx).
		SetResult(&resBody).
		SetError(&rpcstatus.Status{}).
		Send()
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}

	if res.IsError() {
		errRes, ok := res.Error().(*rpcstatus.Status)
		if !ok {
			return nil, fmt.Errorf("failed to cast error response: %s", res.String())
		}
		if err := status.ErrorProto(errRes); err != nil {
			return nil, err
		}
		return nil, status.Error(HTTPStatusToCode(res.StatusCode()), errRes.String())
	}

	data, ok := res.Result().(*T)
	if !ok {
		return nil, fmt.Errorf("failed to cast response: %s", res.String())
	}
	return data, nil
}
