package context

import (
	"context"

	"github.com/gin-gonic/gin"
)

func Get[T any](ctx context.Context, k string) (T, bool) {
	v, ok := ctx.Value(k).(T)
	return v, ok
}

func Set[T any](ctx context.Context, k string, v T) context.Context {
	if gc, ok := ctx.(*gin.Context); ok {
		gc.Set(k, v)
		return ctx
	}
	//nolint:staticcheck
	return context.WithValue(ctx, k, v)
}
