package argocdv1

import (
	"errors"

	"google.golang.org/protobuf/proto"

	"github.com/akuityio/akuity-platform/pkg/api/gateway/http"
)

func init() {

	forward_ArgoCDService_WatchInstanceClusters_0 = http.NewStreamForwarder(func(message proto.Message) (string, error) {
		event, ok := message.(*WatchInstanceClustersResponse)
		if !ok {
			return "", errors.New("unexpected message type")
		}
		return event.Item.Id, nil
	})

	forward_ArgoCDService_WatchInstances_0 = http.NewStreamForwarder(func(message proto.Message) (string, error) {
		event, ok := message.(*WatchInstancesResponse)
		if !ok {
			return "", errors.New("unexpected message type")
		}
		return event.Item.Id, nil
	})

}
