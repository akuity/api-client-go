package context

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/akuity/api-client-go/pkg/api/gateway/accesscontrol"
)

const (
	clientCredentialKey = "client_credential"
)

func GetClientCredential(ctx context.Context) accesscontrol.ClientCredential {
	v, ok := Get[accesscontrol.ClientCredential](ctx, clientCredentialKey)
	if !ok {
		return nil
	}
	return v
}

func SetClientCredential(ctx context.Context, cred accesscontrol.ClientCredential) context.Context {
	return Set[accesscontrol.ClientCredential](ctx, clientCredentialKey, cred)
}

func Get[T any](ctx context.Context, k string) (T, bool) {
	v, ok := ctx.Value(k).(T)
	return v, ok
}

func Set[T any](ctx context.Context, k string, v T) context.Context {
	if gc, ok := ctx.(*gin.Context); ok {
		gc.Set(k, v)
		return ctx
	}
	return context.WithValue(ctx, k, v)
}
