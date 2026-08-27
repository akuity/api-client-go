package http

import "strings"

// EdgeChecker reports whether an address belongs to the edge that fronts the
// deployment — Cloudflare's published ranges in the SaaS shape. It is an
// interface here so this package stays free of the matcher's dependencies;
// internal/utils/ip.NewEdgeChecker builds one from configuration.
type EdgeChecker interface {
	Contains(addr string) bool
}

// ResolveClientIP picks the caller's address from what the edge and the proxy
// in front of us reported: the value of the deployment's trusted platform
// header (e.g. CF-Connecting-IP) and the X-Forwarded-For chain.
//
// The right-most X-Forwarded-For entry is the one the reverse proxy wrote from
// the connection it accepted, so unlike every entry to its left it is not
// caller input — a client cannot forge the address it is connecting from.
// That makes it the thing to check the edge against: when it belongs to the
// edge, the request really did arrive through it and the trusted header the
// edge set is worth believing.
//
// When it does not belong to the edge, the request reached the origin some
// other way — straight at the load balancer, say — and any trusted header on
// it was written by the caller. The proxy-written entry is then the caller's
// real address, so it is returned instead of the header: still the right
// answer, and one nobody can spoof.
//
// With no edge configured this keeps the older behaviour: the trusted header
// when it is set, otherwise the left-most entry, which is what gin's ClientIP
// resolves to under a trust-all proxy configuration. That is spoofable, and it
// is why enforcing callers should run with the edge ranges configured.
//
// Both the HTTP and the gRPC metadata paths resolve through here, so one
// request cannot be judged by two different addresses.
func ResolveClientIP(trustedPlatformValue, forwardedFor string, edge EdgeChecker) string {
	trusted := strings.TrimSpace(trustedPlatformValue)

	if edge != nil && trusted != "" {
		if peer := rightmostForwardedFor(forwardedFor); peer != "" {
			if !edge.Contains(peer) {
				return peer
			}
			return trusted
		}
	}

	if trusted != "" {
		return trusted
	}
	return leftmostForwardedFor(forwardedFor)
}

// leftmostForwardedFor returns the first entry in the chain: the address the
// original caller claimed, which every hop after it merely passed along.
func leftmostForwardedFor(forwardedFor string) string {
	if forwardedFor == "" {
		return ""
	}
	leftmost, _, _ := strings.Cut(forwardedFor, ",")
	return strings.TrimSpace(leftmost)
}

// rightmostForwardedFor returns the last entry: the address the nearest proxy
// observed on its own connection.
func rightmostForwardedFor(forwardedFor string) string {
	if forwardedFor == "" {
		return ""
	}
	_, rightmost, found := lastCut(forwardedFor, ",")
	if !found {
		return strings.TrimSpace(forwardedFor)
	}
	return strings.TrimSpace(rightmost)
}

func lastCut(s, sep string) (before, after string, found bool) {
	if i := strings.LastIndex(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}
