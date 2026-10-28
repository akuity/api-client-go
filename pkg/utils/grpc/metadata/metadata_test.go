package metadata

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

func TestRequestMethodRoundTrip(t *testing.T) {
	md := metadata.MD{}

	method, ok := GetRequestMethod(md)
	assert.False(t, ok)
	assert.Empty(t, method)

	SetRequestMethod(md, "GET")
	method, ok = GetRequestMethod(md)
	assert.True(t, ok)
	assert.Equal(t, "GET", method)
}

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

func TestServiceAccountTokenMetadata(t *testing.T) {
	md := metadata.MD{}
	_, ok := GetServiceAccountToken(md)
	require.False(t, ok)
	SetServiceAccountToken(md, "tok")
	got, ok := GetServiceAccountToken(md)
	require.True(t, ok)
	require.Equal(t, "tok", got)
	_, allowed := MatchIncomingMetadata("x-akuity-service-account-token")
	require.True(t, allowed)
}

func TestRedirectHeaderRoundTrip(t *testing.T) {
	_, ok := GetRedirect(metadata.MD{})
	require.False(t, ok)
	url, ok := GetRedirect(RedirectHeader("https://us1.example.com/api/v1/x"))
	require.True(t, ok)
	require.Equal(t, "https://us1.example.com/api/v1/x", url)
	// Never copied to the client as a raw header: only the forwarder reads it.
	_, exposed := MatchIncomingMetadata("x-akuity-redirect")
	require.False(t, exposed)
}

// The portal gateway installs MatchIncomingMetadata as its header matcher, so
// it alone keeps a client from supplying a server-owned key as a
// Grpc-Metadata-* header. A forged x-request-url would defeat the shard
// handoff token's binding to the request it was issued for.
func TestServerOwnedKeysAreNotForwarded(t *testing.T) {
	require.NotEmpty(t, serverOwnedKeys)
	for key := range serverOwnedKeys {
		for _, header := range []string{key, "Grpc-Metadata-" + key, strings.ToUpper(key)} {
			_, ok := MatchIncomingMetadata(header)
			assert.False(t, ok, "header %q must not be forwarded into gRPC metadata", header)
		}
	}
}
