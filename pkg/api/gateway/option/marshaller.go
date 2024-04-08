package option

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

var Marshaller = &runtime.JSONPb{
	MarshalOptions:   MarshalOptions,
	UnmarshalOptions: UnmarshalOptions,
}
