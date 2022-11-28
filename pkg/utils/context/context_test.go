package context

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestSetGet(t *testing.T) {
	testSets := map[string]struct {
		ctxFn          func() context.Context
		newCtxExpected bool
	}{
		"gin context": {
			ctxFn: func() context.Context {
				ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
				return ctx
			},
			newCtxExpected: true,
		},
		"golang context": {
			ctxFn:          context.TODO,
			newCtxExpected: false,
		},
	}
	for name, testSet := range testSets {
		t.Run(name, func(t *testing.T) {
			testKey := "testKey"
			expectedValue := "testValue"
			ctx := testSet.ctxFn()
			newCtx := Set[string](ctx, testKey, expectedValue)
			if testSet.newCtxExpected {
				require.Equal(t, ctx, newCtx)
			} else {
				require.NotEqual(t, ctx, newCtx)
			}
			actualValue, ok := Get[string](newCtx, testKey)
			require.True(t, ok)
			require.Equal(t, expectedValue, actualValue)
		})
	}
}
