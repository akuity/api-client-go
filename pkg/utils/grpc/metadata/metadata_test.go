package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func TestRequestOriginRoundTrip(t *testing.T) {
	md := metadata.MD{}

	origin, ok := GetRequestOrigin(md)
	assert.False(t, ok)
	assert.Empty(t, origin)

	SetRequestOrigin(md, RequestOriginMCP)
	origin, ok = GetRequestOrigin(md)
	assert.True(t, ok)
	assert.Equal(t, RequestOriginMCP, origin)
}

func TestRequestOriginIsNotForwardedByGateway(t *testing.T) {
	_, ok := MatchIncomingMetadata(requestOriginMetadataKey)
	assert.False(t, ok)
	_, ok = MatchIncomingMetadata("X-Akuity-Request-Origin")
	assert.False(t, ok)
}
