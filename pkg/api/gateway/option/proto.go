package option

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	ServerMarshaller = &runtime.JSONPb{
		MarshalOptions:   ServerMarshalOptions,
		UnmarshalOptions: ServerUnmarshalOptions,
	}
	ClientMarshaller = &runtime.JSONPb{
		MarshalOptions:   ClientMarshalOptions,
		UnmarshalOptions: ClientUnmarshalOptions,
	}
)

var (
	ServerMarshalOptions = protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	ServerUnmarshalOptions = protojson.UnmarshalOptions{
		// Set DiscardUnknown as false to return error
		// if the request contains unknown fields.
		DiscardUnknown: false,
	}
	ClientMarshalOptions = protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	ClientUnmarshalOptions = protojson.UnmarshalOptions{
		// Client should support new response fields unless there are breaking changes,
		// set DiscardUnknown to true.
		DiscardUnknown: true,
	}
)
