package metadata

import (
	"strings"

	"google.golang.org/grpc/metadata"
)

const (
	apiKeyIDMetadataKey     = "x-akuity-api-key-id"
	apiKeySecretMetadataKey = "x-akuity-api-key-secret"
	userTokenMetadataKey    = "x-akuity-user-token"
	argocdTokenMetadataKey  = "x-argocd-token"
	kargoTokenMetadataKey   = "x-kargo-token"
	requestIDMetadataKey    = "x-request-id"
	refreshTokenMetadataKey = "x-refresh-token"
)

var allowedHeaders = map[string]bool{
	apiKeyIDMetadataKey:     true,
	apiKeySecretMetadataKey: true,
	userTokenMetadataKey:    true,
	argocdTokenMetadataKey:  true,
	requestIDMetadataKey:    true,
}

func MatchIncomingMetadata(header string) (string, bool) {
	if allowedHeaders[strings.ToLower(header)] {
		return header, true
	}

	return "", false
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

func GetClientIP(md metadata.MD) (string, bool) {
	if len(md.Get("x-forwarded-for")) == 0 {
		return "", false
	}
	forwardedFor := md.Get("x-forwarded-for")[0]
	if forwardedFor != "" {
		return strings.TrimSpace(strings.Split(forwardedFor, ",")[0]), true
	}
	return "", false
}
