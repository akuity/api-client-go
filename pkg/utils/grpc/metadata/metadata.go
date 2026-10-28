package metadata

import (
	"context"
	"strings"

	"google.golang.org/grpc/metadata"

	publichttp "github.com/akuity/api-client-go/pkg/utils/http"
)

const (
	platformMetadataKey     = "x-platform"
	apiKeyIDMetadataKey     = "x-akuity-api-key-id"
	apiKeySecretMetadataKey = "x-akuity-api-key-secret"
	userTokenMetadataKey    = "x-akuity-user-token"
	// serviceAccountTokenMetadataKey carries the short-lived credential a
	// service account receives from the OIDC token exchange.
	serviceAccountTokenMetadataKey = "x-akuity-service-account-token"
	argocdTokenMetadataKey         = "x-argocd-token"
	kargoTokenMetadataKey          = "x-kargo-token"
	requestIDMetadataKey           = "x-request-id"
	requestURLMetadataKey          = "x-request-url"
	requestMethodMetadataKey       = "x-request-method"
	refreshTokenMetadataKey        = "x-refresh-token"
	trustedPlatformHeader          = "x-trusted-platform-header"
	requestOriginMetadataKey       = "x-akuity-request-origin"
	forwardedForMetadataKey        = "x-forwarded-for"
	// forwardedHostMetadataKey is the host grpc-gateway records from the
	// request it serves. Server-owned, so it holds exactly that one value.
	forwardedHostMetadataKey = "x-forwarded-host"
	// forwardedPrefixMetadataKey carries the path prefix Traefik's stripPrefix
	// middleware removed (X-Forwarded-Prefix), which tells apart instances that
	// share a host under different basepaths. Server-owned.
	forwardedPrefixMetadataKey = "x-forwarded-prefix"
	// clientIPMetadataKey carries the caller address resolved at the HTTP
	// boundary. It is server-owned: see serverOwnedKeys for why no gateway may
	// let a client supply it.
	clientIPMetadataKey = "x-akuity-client-ip"
	// handoffTokenMetadataKey carries a shard handoff token the gateway read
	// from the redirect URL (internal/services/shardhandoff). Server-owned: a
	// client cannot supply it as a header, only as the URL parameter.
	handoffTokenMetadataKey = "x-akuity-handoff-token"
	// redirectMetadataKey is a response header a handler sets to answer an HTTP
	// request with a 307 to the given URL instead of a body; the gateway's
	// forwarder turns it into Location. It is never copied to the client as is.
	redirectMetadataKey = "x-akuity-redirect"

	UserTokenHeader = userTokenMetadataKey

	RequestOriginMCP = "mcp"
)

var allowedHeaders = map[string]bool{
	apiKeyIDMetadataKey:            true,
	apiKeySecretMetadataKey:        true,
	userTokenMetadataKey:           true,
	serviceAccountTokenMetadataKey: true,
	argocdTokenMetadataKey:         true,
	requestIDMetadataKey:           true,
	kargoTokenMetadataKey:          true,
}

// serverOwnedKeys are the metadata keys a gateway annotator derives from the
// request itself, so a client must never be able to supply them. grpc-gateway's
// default header matcher forwards any `Grpc-Metadata-<key>` header into
// metadata, and metadata.Join orders those ahead of the annotator's values, so
// md.Get would return the client's. The portal is covered by the allow-list
// above; gateways that need the permissive default must drop these explicitly
// (see gateway.MatchIncomingDroppingServerOwned).
var serverOwnedKeys = map[string]bool{
	requestURLMetadataKey:      true,
	requestMethodMetadataKey:   true,
	clientIPMetadataKey:        true,
	trustedPlatformHeader:      true,
	handoffTokenMetadataKey:    true,
	forwardedHostMetadataKey:   true,
	forwardedPrefixMetadataKey: true,
}

// GetForwardedHost returns the request host the gateway recorded. Anything
// but exactly one value is refused rather than guessed at.
func GetForwardedHost(md metadata.MD) (string, bool) {
	v := md.Get(forwardedHostMetadataKey)
	if len(v) != 1 || v[0] == "" {
		return "", false
	}
	return v[0], true
}

// SetForwardedPrefix records the prefixes the edge stripped from the request
// path (every X-Forwarded-Prefix value, so a repeated header stays visible to
// GetForwardedPrefix). Only gateway annotators may call it.
func SetForwardedPrefix(md metadata.MD, prefixes ...string) {
	md.Delete(forwardedPrefixMetadataKey)
	for _, p := range prefixes {
		if p != "" {
			md.Append(forwardedPrefixMetadataKey, p)
		}
	}
}

// GetForwardedPrefix returns the stripped request prefix, "" when the request
// carried none. ok is false when it carried more than one.
func GetForwardedPrefix(md metadata.MD) (prefix string, ok bool) {
	v := md.Get(forwardedPrefixMetadataKey)
	switch len(v) {
	case 0:
		return "", true
	case 1:
		return v[0], true
	default:
		return "", false
	}
}

func SetHandoffToken(md metadata.MD, token string) {
	md.Set(handoffTokenMetadataKey, token)
}

func GetHandoffToken(md metadata.MD) (string, bool) {
	v := first(md, handoffTokenMetadataKey)
	return v, v != ""
}

// RedirectHeader returns the response metadata that makes the HTTP gateway
// answer with a 307 to url. Unary handlers pass it to grpc.SetHeader, stream
// handlers to SendHeader, and return without a body.
func RedirectHeader(url string) metadata.MD {
	return metadata.Pairs(redirectMetadataKey, url)
}

// GetRedirect returns the redirect URL a handler set in its response headers.
func GetRedirect(md metadata.MD) (string, bool) {
	v := md.Get(redirectMetadataKey)
	if len(v) == 0 || v[0] == "" {
		return "", false
	}
	return v[0], true
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

func GetServiceAccountToken(md metadata.MD) (string, bool) {
	v := md.Get(serviceAccountTokenMetadataKey)
	if len(v) == 0 {
		return "", false
	}
	return v[0], true
}

func SetServiceAccountToken(md metadata.MD, token string) {
	md.Set(serviceAccountTokenMetadataKey, token)
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

// RequestURLFromContext returns the path and query the HTTP gateway recorded
// for this request, or "" when it did not pass through the gateway.
func RequestURLFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	u, _ := GetRequestURL(md)
	return u
}
