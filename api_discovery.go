package ciolite

// Api functions that support: https://context.io/docs/lite/discovery

// GetDiscoveryResponse data struct
// 	https://context.io/docs/lite/discovery#get
type GetDiscoveryResponse struct {
	Email         string        `json:"email,omitempty"`
	Type          string        `json:"type,omitempty"`
	Documentation []interface{} `json:"documentation,omitempty"`

	Found bool `json:"found,omitempty"`

	IMAP struct {
		Server   string `json:"server,omitempty"`
		Username string `json:"username,omitempty"`

		UseSSL bool `json:"use_ssl,omitempty"`
		OAuth  bool `json:"oauth,omitempty"`

		Port int `json:"port,omitempty"`
	} `json:"imap,omitempty"`
}

// GetDiscoveryParams query values data struct
// 	https://context.io/docs/lite/discovery#get
type GetDiscoveryParams struct {
	// Required:
	SourceType string `json:"source_type,omitempty"`
	Email      string `json:"email,omitempty"`
}

// GetDiscovery attempts to discover connection settings for a given email address.
// queryValues requires SourceType and Email to be set.
// 	https://context.io/docs/lite/discovery#get
func (cioLite CioLite) GetDiscovery(queryValues GetDiscoveryParams) (GetDiscoveryResponse, error) {

	// Make request
	request := clientRequest{
		method:      "GET",
		path:        "/discovery",
		queryValues: queryValues,
	}

	// Make response
	var response GetDiscoveryResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
