package option

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/akuity/api-client-go/pkg/utils/proto"
)

var (
	Marshaller = &runtime.JSONPb{
		MarshalOptions:   proto.MarshalOptions,
		UnmarshalOptions: proto.UnmarshalOptions,
	}
)
