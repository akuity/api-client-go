package option

import (
	"bytes"
	"fmt"
	"io"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	publichttp "github.com/akuity/api-client-go/pkg/utils/http"
	"github.com/akuity/api-client-go/pkg/utils/proto"
)

var (
	ErrUnsupportedContentTypeEventStream = fmt.Errorf("unsupported content type: %v", publichttp.MIMETextEventStream)
)

var (
	Marshaller = &runtime.JSONPb{
		MarshalOptions:   proto.MarshalOptions,
		UnmarshalOptions: proto.UnmarshalOptions,
	}
)

func NewEventStreamMarshaller() runtime.Marshaler {
	return &eventStreamMarshaller{}
}

type eventStreamMarshaller struct{}

func (e *eventStreamMarshaller) encode(v interface{}) ([]byte, error) {
	data, err := Marshaller.Marshal(v)
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer
	b.WriteString("data: ")
	b.Write(data)
	b.WriteRune('\n')
	return b.Bytes(), nil
}

func (e *eventStreamMarshaller) Marshal(v interface{}) ([]byte, error) {
	return e.encode(v)
}

func (e *eventStreamMarshaller) Unmarshal(data []byte, v interface{}) error {
	return ErrUnsupportedContentTypeEventStream
}

func (e *eventStreamMarshaller) Decode(v interface{}) error {
	return ErrUnsupportedContentTypeEventStream
}

func (e *eventStreamMarshaller) NewDecoder(r io.Reader) runtime.Decoder {
	return e
}

func (e *eventStreamMarshaller) NewEncoder(w io.Writer) runtime.Encoder {
	return runtime.EncoderFunc(func(v interface{}) error {
		data, err := e.encode(v)
		if err != nil {
			return err
		}
		_, err = w.Write(data)
		return err
	})
}

func (e *eventStreamMarshaller) ContentType(_ interface{}) string {
	return publichttp.MIMETextEventStream
}
