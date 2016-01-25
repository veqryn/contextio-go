// Package ciolite is the Golang client library for the Lite Context.IO API
package ciolite

// Imports
import (
	"github.com/Sirupsen/logrus"
)

// CioLite ...
type CioLite struct {
	apiKey    string
	apiSecret string

	log *logrus.Logger
}

// NewCioLite ...
func NewCioLite(key string, secret string, logger *logrus.Logger) *CioLite {

	if logger == nil {
		logger = logrus.New()
		logger.Level = logrus.ErrorLevel
	}

	return &CioLite{apiKey: key, apiSecret: secret, log: logger}
}
