package proto

import (
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	MarshalOptions = protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	UnmarshalOptions = protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
)
