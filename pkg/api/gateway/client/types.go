package client

import (
	"encoding/json"
)

type streamingResponse map[string]json.RawMessage

const (
	streamingResponseResultKey = "result"
	streamingResponseErrorKey  = "error"
)
