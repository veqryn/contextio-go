package ciolite

// Api functions that support: https://context.io/docs/app/status_callback_url

// GetStatusCallbackURLResponse data struct
// 	https://context.io/docs/app/status_callback_url#get
type GetStatusCallbackURLResponse struct {
	StatusCallbackURL string `json:"status_callback_url,omitempty"`
	ResourceURL       string `json:"resource_url,omitempty"`
}

// CreateStatusCallbackURLParams form values data struct.
// Requires: StatusCallbackURL
// 	https://context.io/docs/app/status_callback_url#post
type CreateStatusCallbackURLParams struct {
	StatusCallbackURL string `json:"status_callback_url,omitempty"`
}

// CreateDeleteStatusCallbackURLResponse data struct
// 	https://context.io/docs/app/status_callback_url#post
// 	https://context.io/docs/app/status_callback_url#delete
type CreateDeleteStatusCallbackURLResponse struct {
	Success bool `json:"success,omitempty"`
}

// GetStatusCallbackURL gets a list of app status callback url's.
// 	https://context.io/docs/app/status_callback_url#get
func (cioLite CioLite) GetStatusCallbackURL() (GetStatusCallbackURLResponse, error) {

	// Make request
	request := clientRequest{
		Method: "GET",
		Path:   "/app/status_callback_url",
	}

	// Make response
	var response GetStatusCallbackURLResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// CreateStatusCallbackURL create an app status callback url.
// Requires: StatusCallbackURL
// 	https://context.io/docs/app/status_callback_url#post
func (cioLite CioLite) CreateStatusCallbackURL(formValues CreateStatusCallbackURLParams) (CreateDeleteStatusCallbackURLResponse, error) {

	// Make request
	request := clientRequest{
		Method:     "POST",
		Path:       "/app/status_callback_url",
		FormValues: formValues,
	}

	// Make response
	var response CreateDeleteStatusCallbackURLResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// DeleteStatusCallbackURL removes an app status callback url.
// 	https://context.io/docs/app/status_callback_url#delete
func (cioLite CioLite) DeleteStatusCallbackURL() (CreateDeleteStatusCallbackURLResponse, error) {

	// Make request
	request := clientRequest{
		Method: "DELETE",
		Path:   "/app/status_callback_url",
	}

	// Make response
	var response CreateDeleteStatusCallbackURLResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
