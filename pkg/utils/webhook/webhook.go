package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/pkg/errors"
)

const (
	SignaturePrefix = "sha256="
)

// ComputeSignature returns the signature of the given payload that can be
// used to validate Akuity webhook requests.
func ComputeSignature(secret string, payload []byte) (string, error) {
	signature, err := computeSignature(secret, payload)
	if err != nil {
		return "", errors.Wrap(err, "get signature from payload")
	}
	return SignaturePrefix + hex.EncodeToString(signature), nil
}

// VerifySignature verifies the given signature matches against the payload
func VerifySignature(signature, secret string, payload []byte) (bool, error) {
	actual, err := hex.DecodeString(strings.TrimPrefix(signature, SignaturePrefix))
	if err != nil {
		return false, errors.Wrap(err, "decode signature")
	}

	expected, err := computeSignature(secret, payload)
	if err != nil {
		return false, errors.Wrap(err, "get signature from payload")
	}
	return hmac.Equal(actual, expected), nil
}

func computeSignature(secret string, payload []byte) ([]byte, error) {
	sign := hmac.New(sha256.New, []byte(secret))
	if _, err := sign.Write(payload); err != nil {
		return nil, err
	}
	return sign.Sum(nil), nil
}
