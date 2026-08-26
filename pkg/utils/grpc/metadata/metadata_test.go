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

// TestRequestOriginIsNotForwardedByGateway locks the assumption
// internal/utils/grpc/auth.isMCPOrigin is built on: the MCP origin marker is
// set server-side by the MCP dispatch handler and must never be settable by a
// client. Forging it would both mark the caller as MCP-originated and swap in
// the MCP token validator, which accepts the MCP clients' audiences.
//
// The Grpc-Metadata- prefixed forms are the shape of #12244 (the agent API's
// cluster id): grpc-gateway's DefaultHeaderMatcher strips that prefix and
// forwards the rest, so any mux that switches to it would start forwarding
// this key. This allowlist matches exact keys and strips nothing, so both
// forms are refused.
func TestRequestOriginIsNotForwardedByGateway(t *testing.T) {
	for _, header := range []string{
		requestOriginMetadataKey,
		"X-Akuity-Request-Origin",
		"Grpc-Metadata-x-akuity-request-origin",
		"grpc-metadata-X-Akuity-Request-Origin",
	} {
		_, ok := MatchIncomingMetadata(header)
		assert.False(t, ok, "header %q must not be forwarded into gRPC metadata", header)
	}
}
