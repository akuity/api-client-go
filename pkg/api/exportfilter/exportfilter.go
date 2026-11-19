// Package exportfilter parses the resources entries of a filtered export request. The export
// servers and their clients share it, so a client can drop any entry the server would reject.
package exportfilter

import (
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/api/validation/path"
)

// Entry names one object in a filtered export.
type Entry struct {
	Group, Kind, Namespace, Name string
}

// String formats e the way Parse reads it.
func (e Entry) String() string {
	return strings.Join([]string{e.Group, e.Kind, e.Namespace, e.Name}, "/")
}

// Parse parses "<group>/<kind>/<namespace>/<name>", with an empty group for the core API group
// and an empty namespace for cluster-scoped kinds.
func Parse(s string) (Entry, error) {
	// Kubernetes names never contain "/", so a valid entry has exactly four parts.
	parts := strings.Split(s, "/")
	if len(parts) != 4 || parts[1] == "" || parts[3] == "" {
		return Entry{}, fmt.Errorf("invalid resources entry %q, want <group>/<kind>/<namespace>/<name>", s)
	}
	// client-go refuses these names before sending a request, which would fail the export.
	for _, seg := range parts[2:] {
		if msgs := path.IsValidPathSegmentName(seg); len(msgs) > 0 {
			return Entry{}, fmt.Errorf("invalid resources entry %q: %s", s, strings.Join(msgs, ", "))
		}
	}
	return Entry{Group: parts[0], Kind: parts[1], Namespace: parts[2], Name: parts[3]}, nil
}
