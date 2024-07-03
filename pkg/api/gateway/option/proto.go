package option

import (
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	MarshalOptions = protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	UnmarshalOptions = protojson.UnmarshalOptions{
		// Set DiscardUnknown as false to return error
		// if the request contains unknown fields.
		DiscardUnknown: false,
	}
)
