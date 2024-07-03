package option

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestUnmarshal(t *testing.T) {
	testCases := map[string]struct {
		input       string
		errExpected bool
		expected    proto.Message
	}{
		"unknown fields should return error": {
			input: `{
  "unknown_field": "unknown_value"
}`,
			errExpected: true,
			expected: &grpc_health_v1.HealthCheckRequest{
				Service: "fake_service",
			},
		},
		"omitted field should not return error": {
			input:       `{}`,
			errExpected: false,
			expected:    &grpc_health_v1.HealthCheckRequest{},
		},
		"structpb.Struct allows arbitrary fields": {
			input: `{
  "arbitrary_string": "arbitrary_value",
  "arbitrary_bool": true,
  "arbitrary_int": 1
}`,
			errExpected: false,
			expected: mustNewStruct(map[string]any{
				"arbitrary_string": "arbitrary_value",
				"arbitrary_bool":   true,
				"arbitrary_int":    1,
			}),
		},
	}
	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Create new message to assert unmarshal result
			actual := proto.Clone(tc.expected)
			proto.Reset(actual)

			err := UnmarshalOptions.Unmarshal([]byte(tc.input), actual)
			if tc.errExpected {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			if !cmp.Equal(tc.expected, actual, protocmp.Transform()) {
				require.Fail(t, cmp.Diff(tc.expected, actual, protocmp.Transform()))
			}
		})
	}
}

func mustNewStruct(m map[string]any) proto.Message {
	s, err := structpb.NewStruct(m)
	if err != nil {
		panic(err)
	}
	return s
}
