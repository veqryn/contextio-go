// Package ciolite is the Golang client library for the Lite Context.IO API
package ciolite

// CioLite ...
type CioLite struct {
	apiKey    string
	apiSecret string
	log       Logger
}

// NewCioLite returns a CIO Lite struct (without a logger) for accessing the CIO Lite API.
func NewCioLite(key string, secret string) *CioLite {
	return NewCioLiteWithLogger(key, secret, nil)
}

// NewCioLiteWithLogger returns a CIO Lite struct (with a logger) for accessing the CIO Lite API.
func NewCioLiteWithLogger(key string, secret string, logger Logger) *CioLite {
	return &CioLite{apiKey: key, apiSecret: secret, log: logger}
}

// Logger interface which *log.Logger uses. Allows injection of user specified loggers.
type Logger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
	Fatalf(format string, v ...interface{})
}
