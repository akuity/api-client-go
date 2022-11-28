package client

import (
	"net/http"
)

type Option func(*options)

type options struct {
	hc *http.Client

	skipTLSVerify bool
}

func WithHTTPClient(hc *http.Client) Option {
	return func(opts *options) {
		opts.hc = hc
	}
}

func SkipTLSVerify(skip bool) Option {
	return func(opts *options) {
		opts.skipTLSVerify = skip
	}
}
