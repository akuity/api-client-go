package metadata

import (
	"strings"

	"google.golang.org/grpc/metadata"
)

const (
	apiKeyIDMetadataKey     = "x-akuity-api-key-id"
	apiKeySecretMetadataKey = "x-akuity-api-key-secret"
	userTokenMetadataKey    = "x-akuity-user-token"
	requestIDMetadataKey    = "x-request-id"
)

var allowedHeaders = map[string]bool{
	apiKeyIDMetadataKey:     true,
	apiKeySecretMetadataKey: true,
	userTokenMetadataKey:    true,
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
