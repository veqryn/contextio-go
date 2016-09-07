// Package ciolite is the Golang client library for the Lite Context.IO API
package ciolite

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"time"
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
	apiKey         string
	apiSecret      string
	Host           string
	RequestTimeout time.Duration

	// PreRequestHook is a function (mostly for logging) that will be executed
	// before the request is made, its arguments are the User ID, Account Label,
	// the Method (GET/POST/etc), the URL, and the redacted body values.
	PreRequestHook func(string, string, string, string, url.Values)

	// PostRequestShouldRetryHook is a function (mostly for logging) that will be
	// executed after each request is made, and will be called at least once.
	// Its arguments are the request Attempt # (starts at 1), User ID, Account Label,
	// the Method (GET/POST/etc), the URL, response Status Code, response Payload,
	// and any error received while attempting this request.
	// The returned boolean is whether this request should be retried or not, which
	// if false then this is the last call of this function, but if true means this
	// function will be called again.
	PostRequestShouldRetryHook func(int, string, string, string, string, int, string, error) bool

	// ResponseBodyCloseErrorHook is a function (purely for logging) that will
	// execute if there is an error closing the response body.
	ResponseBodyCloseErrorHook func(error)
}

// NewCioLite returns a CIO Lite struct (without a logger) for accessing the CIO Lite API.
func NewCioLite(key string, secret string) CioLite {
	return CioLite{
		apiKey:         key,
		apiSecret:      secret,
		Host:           DefaultHost,
		RequestTimeout: DefaultRequestTimeout,
	}
}

// NewTestCioLiteServer is a convenience function that returns a CioLite object
// and a *httptest.Server (which must be closed when done being used).
// The CioLite instance will hit the test server for all requests.
func NewTestCioLiteServer(key string, secret string, handler http.Handler) (CioLite, *httptest.Server) {
	testServer := httptest.NewServer(handler)
	testCioLite := CioLite{
		apiKey:         key,
		apiSecret:      secret,
		Host:           testServer.URL,
		RequestTimeout: 5 * time.Second,
	}
	return testCioLite, testServer
}

// ValidateCallback returns true if this Webhook Callback or User Account Status Callback authenticates
func (cio CioLite) ValidateCallback(token string, signature string, timestamp int) bool {
	// Hash timestamp and token with secret, compare to signature
	message := strconv.Itoa(timestamp) + token
	hash := hashHmac(sha256.New, message, cio.apiSecret)
	return len(hash) > 0 && signature == hash
}

// hashHmac returns the hash of a message hashed with the provided hash function, using the provided secret
func hashHmac(hashAlgorithm func() hash.Hash, message string, secret string) string {
	h := hmac.New(hashAlgorithm, []byte(secret))
	if _, err := h.Write([]byte(message)); err != nil {
		panic("hash.Hash unable to write message bytes, with error: " + err.Error())
	}
	return hex.EncodeToString(h.Sum(nil))
}
