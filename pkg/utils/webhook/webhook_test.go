package webhook

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testSecret           = "dummy-secret"
	testPayloadSignature = "sha256=b9b977ae4446dd7e1291ce3bc789cb2cf58e86a83ac0b1ce99feb04c37d0d480"
)

var testPayload = []byte(`{"event_time": "2024-06-26T15:04:05Z"}`)

func TestComputeSignature(t *testing.T) {
	actual, err := ComputeSignature(testSecret, testPayload)
	require.NoError(t, err)
	require.Equal(t, testPayloadSignature, actual)
}

func TestVerifySignature(t *testing.T) {
	testCases := map[string]struct {
		secret   string
		expected bool
	}{
		"secret match": {
			secret:   testSecret,
			expected: true,
		},
		"secret mismatch": {
			secret:   "wrong-secret",
			expected: false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual, err := VerifySignature(testPayloadSignature, tc.secret, testPayload)
			require.NoError(t, err)
			require.Equal(t, tc.expected, actual)
		})
	}
}
