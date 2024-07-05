package option

import (
	"net/http"

	"github.com/akuity/grpc-gateway-client/pkg/grpc/gateway"
	"github.com/akuity/grpc-gateway-client/pkg/http/roundtripper"
)

func NewClient(baseURL string, skipTLSVerify bool) gateway.Client {
	hc := &http.Client{}
	roundtripper.ApplyAuthorizationHeaderInjector(hc)
	return gateway.NewClient(baseURL,
		gateway.WithHTTPClient(hc),
		gateway.WithMarshaller(ClientMarshaller),
		gateway.SkipTLSVerify(skipTLSVerify),
	)
}
