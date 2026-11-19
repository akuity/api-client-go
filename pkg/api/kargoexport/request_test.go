package kargoexport

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	kargov1 "github.com/akuity/api-client-go/pkg/api/gen/kargo/v1"
)

type RequestRoundTripSuite struct {
	suite.Suite
}

func Test_RequestRoundTrip(t *testing.T) {
	suite.Run(t, new(RequestRoundTripSuite))
}

// fillEveryField sets each field of m to a non-zero value, so a conversion that drops a field
// shows up as a difference.
func (s *RequestRoundTripSuite) fillEveryField(m protoreflect.Message) {
	fields := m.Descriptor().Fields()
	for i := range fields.Len() {
		fd := fields.Get(i)
		val := "x-" + string(fd.Name())
		switch {
		case fd.IsList():
			s.Require().Equal(protoreflect.StringKind, fd.Kind(), "teach fillEveryField about list field %s", fd.Name())
			m.Mutable(fd).List().Append(protoreflect.ValueOfString(val))
		case fd.Kind() == protoreflect.StringKind:
			m.Set(fd, protoreflect.ValueOfString(val))
		case fd.Kind() == protoreflect.BoolKind:
			m.Set(fd, protoreflect.ValueOfBool(true))
		default:
			s.FailNow("teach fillEveryField about field " + string(fd.Name()))
		}
	}
}

func (s *RequestRoundTripSuite) TestUnaryRequest() {
	req := &kargov1.ExportKargoInstanceRequest{}
	s.fillEveryField(req.ProtoReflect())
	got := FromStreamRequest(ToStreamRequest(req))
	s.Require().True(proto.Equal(req, got), "a field was dropped converting to the stream request and back: %v", got)
}

func (s *RequestRoundTripSuite) TestStreamRequest() {
	req := &kargov1.ExportKargoInstanceStreamRequest{}
	s.fillEveryField(req.ProtoReflect())
	got := ToStreamRequest(FromStreamRequest(req))
	s.Require().True(proto.Equal(req, got), "a field was dropped converting to the unary request and back: %v", got)
}
