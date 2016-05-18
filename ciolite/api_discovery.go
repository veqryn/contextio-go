package ciolite

import "github.com/contextio/contextio-go/cioutil"

// Api functions that support: https://context.io/docs/lite/discovery

// GetDiscoveryParams query values data struct.
// Requires SourceType and Email.
// 	https://context.io/docs/lite/discovery#get
type GetDiscoveryParams struct {
	// Required:
	SourceType string `json:"source_type"`
	Email      string `json:"email"`
}

// GetDiscoveryResponse data struct
// 	https://context.io/docs/lite/discovery#get
type GetDiscoveryResponse struct {
	Email         string        `json:"email,omitempty"`
	Type          string        `json:"type,omitempty"`
	Documentation []interface{} `json:"documentation,omitempty"`

	Found bool `json:"found,omitempty"`

	IMAP GetDiscoveryIMAPResponse `json:"imap,omitempty"`
}

// GetDiscoveryIMAPResponse embedded data struct
// 	https://context.io/docs/lite/discovery#get
type GetDiscoveryIMAPResponse struct {
	Server   string `json:"server,omitempty"`
	Username string `json:"username,omitempty"`

	UseSSL bool `json:"use_ssl,omitempty"`
	OAuth  bool `json:"oauth,omitempty"`

	Port int `json:"port,omitempty"`
}

// GetDiscovery attempts to discover connection settings for a given email address.
// queryValues requires SourceType and Email to be set.
// 	https://context.io/docs/lite/discovery#get
func (cioLite CioLite) GetDiscovery(queryValues GetDiscoveryParams) (GetDiscoveryResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method:      "GET",
		Path:        "/discovery",
		QueryValues: queryValues,
	}

	// Make response
	var response GetDiscoveryResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}
