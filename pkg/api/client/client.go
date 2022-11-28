package client

import (
	gwclient "github.com/akuity/api-client-go/pkg/api/gateway/client"
)

type GatewayClient interface {
	setGatewayClient(gwc gwclient.Client)
}

func newGatewayClient[T GatewayClient](client T, baseURL string, gwOpts ...gwclient.Option) T {
	client.setGatewayClient(gwclient.NewClient(baseURL, gwOpts...))
	return client
}
