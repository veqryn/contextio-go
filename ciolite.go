// Package ciolite ...
package ciolite

// Imports
import ()

// CioLite ...
type CioLite struct {
	apiKey    string
	apiSecret string
}

// NewContextIOLite ...
func NewCioLite(key string, secret string) *CioLite {
	return &CioLite{apiKey: key, apiSecret: secret}
}
