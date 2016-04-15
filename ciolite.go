// Package ciolite is the Golang client library for the Lite Context.IO API
package ciolite

import "time"

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
	log            Logger
	Host           string
	RequestTimeout time.Duration
}

// NewCioLite returns a CIO Lite struct (without a logger) for accessing the CIO Lite API.
func NewCioLite(key string, secret string) CioLite {
	return NewCioLiteWithLogger(key, secret, nil)
}

// NewCioLiteWithLogger returns a CIO Lite struct (with a logger) for accessing the CIO Lite API.
func NewCioLiteWithLogger(key string, secret string, logger Logger) CioLite {
	return CioLite{
		apiKey:         key,
		apiSecret:      secret,
		log:            logger,
		Host:           DefaultHost,
		RequestTimeout: DefaultRequestTimeout,
	}
}

// Logger interface which *log.Logger uses. Allows injection of user specified loggers.
type Logger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
	Fatalf(format string, v ...interface{})
}
