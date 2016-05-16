// Package ciolite is the Golang client library for the Lite Context.IO API
package ciolite

import (
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/contextio/contextio-go/cioutil"
)

const (
	// DefaultHost is the default host of CIO Lite API
	DefaultHost = "https://api.context.io/lite"

	// DefaultRequestTimeout is the default timeout duration used on HTTP requests
	DefaultRequestTimeout = 120 * time.Second
)

// CioLite struct contains the api key and secret, along with an optional logger,
// and provides convenience functions for accessing all CIO Lite endpoints.
type CioLite struct {
	cioutil.Cio
}

// NewCioLite returns a CIO Lite struct (without a logger) for accessing the CIO Lite API.
func NewCioLite(key string, secret string) CioLite {
	return NewCioLiteWithLogger(key, secret, nil)
}

// NewCioLiteWithLogger returns a CIO Lite struct (with a logger) for accessing the CIO Lite API.
func NewCioLiteWithLogger(key string, secret string, logger cioutil.Logger) CioLite {
	return CioLite{Cio: cioutil.NewCio(key, secret, logger, DefaultHost, DefaultRequestTimeout)}
}

// NewTestCioLiteServer is a convenience function that returns a CioLite object
// and a *httptest.Server (which must be closed when done being used).
// The CioLite instance will hit the test server for all requests.
func NewTestCioLiteServer(key string, secret string, logger cioutil.Logger, handler http.Handler) (CioLite, *httptest.Server) {
	testServer := httptest.NewServer(handler)
	testCioLite := CioLite{Cio: cioutil.NewCio(key, secret, logger, testServer.URL, 5*time.Second)}
	return testCioLite, testServer
}
