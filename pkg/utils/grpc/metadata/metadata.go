package metadata

import (
	"strings"

	"google.golang.org/grpc/metadata"

	publichttp "github.com/akuity/api-client-go/pkg/utils/http"
)

const (
	platformMetadataKey      = "x-platform"
	apiKeyIDMetadataKey      = "x-akuity-api-key-id"
	apiKeySecretMetadataKey  = "x-akuity-api-key-secret"
	userTokenMetadataKey     = "x-akuity-user-token"
	argocdTokenMetadataKey   = "x-argocd-token"
	kargoTokenMetadataKey    = "x-kargo-token"
	requestIDMetadataKey     = "x-request-id"
	requestURLMetadataKey    = "x-request-url"
	requestMethodMetadataKey = "x-request-method"
	refreshTokenMetadataKey  = "x-refresh-token"
	trustedPlatformHeader    = "x-trusted-platform-header"
	requestOriginMetadataKey = "x-akuity-request-origin"
	forwardedForMetadataKey  = "x-forwarded-for"
	// clientIPMetadataKey carries the caller address resolved at the HTTP
	// boundary. It is server-owned: see serverOwnedKeys for why no gateway may
	// let a client supply it.
	clientIPMetadataKey = "x-akuity-client-ip"

	UserTokenHeader = userTokenMetadataKey

	RequestOriginMCP = "mcp"
)

var allowedHeaders = map[string]bool{
	apiKeyIDMetadataKey:     true,
	apiKeySecretMetadataKey: true,
	userTokenMetadataKey:    true,
	argocdTokenMetadataKey:  true,
	requestIDMetadataKey:    true,
	kargoTokenMetadataKey:   true,
}

// serverOwnedKeys are the metadata keys a gateway annotator derives from the
// request itself, so a client must never be able to supply them. grpc-gateway's
// default header matcher forwards any `Grpc-Metadata-<key>` header into
// metadata, and metadata.Join orders those ahead of the annotator's values, so
// md.Get would return the client's. The portal is covered by the allow-list
// above; gateways that need the permissive default must drop these explicitly
// (see gateway.MatchIncomingDroppingServerOwned).
var serverOwnedKeys = map[string]bool{
	requestURLMetadataKey:    true,
	requestMethodMetadataKey: true,
	clientIPMetadataKey:      true,
	trustedPlatformHeader:    true,
}

// IsServerOwned reports whether the key is one the server derives from the
// request and a client may therefore not set.
func IsServerOwned(key string) bool {
	return serverOwnedKeys[strings.ToLower(key)]
}

type Platform string

const (
	PlatformAkuityPlatform Platform = "akuity-platform"
	PlatformArgoCD         Platform = "argocd"
	PlatformKargo          Platform = "kargo"
	PlatformAIMS           Platform = "aims"
)

func MatchIncomingMetadata(header string) (string, bool) {
	if allowedHeaders[strings.ToLower(header)] {
		return header, true
	}

	return "", false
}

func GetPlatform(md metadata.MD) (Platform, bool) {
	v := md.Get(platformMetadataKey)
	if len(v) == 0 {
		return "", false
	}
	return Platform(v[0]), true
}

func SetPlatform(md metadata.MD, platform string) {
	md.Set(platformMetadataKey, platform)
}

func SetTrustedPlatformHeader(md metadata.MD, platform string) {
	md.Set(trustedPlatformHeader, platform)
}

func GetAPIKey(md metadata.MD) (id, secret string, ok bool) {
	idv := md.Get(apiKeyIDMetadataKey)
	sv := md.Get(apiKeySecretMetadataKey)
	if len(idv) == 0 || len(sv) == 0 {
		return "", "", false
	}
	return idv[0], sv[0], true
}

func SetAPIKey(md metadata.MD, id, secret string) {
	md.Set(apiKeyIDMetadataKey, id)
	md.Set(apiKeySecretMetadataKey, secret)
}

func GetRefreshToken(md metadata.MD) (string, bool) {
	v := md.Get(refreshTokenMetadataKey)
	if len(v) == 0 {
		return "", false
	}
	return v[0], true
}

func SetRefreshToken(md metadata.MD, token string) {
	md.Set(refreshTokenMetadataKey, token)
}

func GetUserToken(md metadata.MD) (string, bool) {
	v := md.Get(userTokenMetadataKey)
	if len(v) == 0 {
		return "", false
	}
	return v[0], true
}

func SetUserToken(md metadata.MD, token string) {
	md.Set(userTokenMetadataKey, token)
}

func GetArgoCDToken(md metadata.MD) (string, bool) {
	v := md.Get(argocdTokenMetadataKey)
	if len(v) == 0 {
		return "", false
	}
	return v[0], true
}

func SetArgoCDToken(md metadata.MD, token string) {
	md.Set(argocdTokenMetadataKey, token)
}

func GetKargoToken(md metadata.MD) (string, bool) {
	v := md.Get(kargoTokenMetadataKey)
	if len(v) == 0 {
		return "", false
	}
	return v[0], true
}

func SetKargoToken(md metadata.MD, token string) {
	md.Set(kargoTokenMetadataKey, token)
}

func SetRequestOrigin(md metadata.MD, origin string) {
	md.Set(requestOriginMetadataKey, origin)
}

func GetRequestOrigin(md metadata.MD) (string, bool) {
	v := md.Get(requestOriginMetadataKey)
	if len(v) == 0 {
		return "", false
	}
	return v[0], true
}

func GetRequestID(md metadata.MD) (string, bool) {
	v := md.Get(requestIDMetadataKey)
	if len(v) == 0 {
		return "", false
	}
	return v[0], true
}

func SetRequestID(md metadata.MD, requestID string) {
	md.Set(requestIDMetadataKey, requestID)
}

func SetRequestURL(md metadata.MD, url string) {
	md.Set(requestURLMetadataKey, url)
}

func GetRequestURL(md metadata.MD) (string, bool) {
	v := md.Get(requestURLMetadataKey)
	if len(v) == 0 {
		return "", false
	}
	return v[0], true
}

func SetRequestMethod(md metadata.MD, method string) {
	md.Set(requestMethodMetadataKey, method)
}

func GetRequestMethod(md metadata.MD) (string, bool) {
	v := md.Get(requestMethodMetadataKey)
	if len(v) == 0 {
		return "", false
	}
	return v[0], true
}

// SetClientIP records an address already resolved from the inbound request.
// Callers that hold the *http.Request should use this rather than leaving the
// derivation to GetClientIP: grpc-gateway appends its own RemoteAddr to the
// x-forwarded-for it forwards, which on this deployment is the reverse proxy's
// pod address, so the chain in metadata has one more hop than the request had.
func SetClientIP(md metadata.MD, clientIP string) {
	md.Set(clientIPMetadataKey, clientIP)
}

// GetClientIP returns the caller's address: the one resolved at the HTTP
// boundary when it was recorded, and otherwise derived from the forwarded
// metadata. edge may be nil when the deployment declares no edge ranges; see
// ResolveClientIP for what that costs.
func GetClientIP(md metadata.MD, edge publichttp.EdgeChecker) string {
	if resolved := first(md, clientIPMetadataKey); resolved != "" {
		return resolved
	}
	return publichttp.ResolveClientIP(first(md, trustedPlatformHeader), first(md, forwardedForMetadataKey), edge)
}

func first(md metadata.MD, key string) string {
	v := md.Get(key)
	if len(v) == 0 {
		return ""
	}
	return v[0]
}
