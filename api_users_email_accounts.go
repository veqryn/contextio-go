// Package ciolite ...
package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts

// Imports
import ()

// GetUsersEmailAccountsResponse ...
type GetUsersEmailAccountsResponse struct {
	Status             string `json:"status:omitempty"`
	ResourceURL        string `json:"resource_url,omitempty"`
	Type               string `json:"type,omitempty"`
	AuthenticationType string `json:"authentication_type,omitempty"`
	Server             string `json:"server,omitempty"`
	Label              string `json:"label,omitempty"`
	Username           string `json:"username,omitempty"`

	UseSSL bool `json:"use_ssl:omitempty"`

	Port int `json:"port:omitempty"`
}
