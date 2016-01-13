// Package ciolite ...
package ciolite

// Api functions that support: https://context.io/docs/lite/connect_tokens

// Imports
import ()

// GetConnectTokensResponse ...
type GetConnectTokensResponse struct {
	Token              string `json:"token:omitempty"`
	Email              string `json:"email,omitempty"`
	CallbackURL        string `json:"callback_url,omitempty"`
	FirstName          string `json:"first_name,omitempty"`
	LastName           string `json:"last_name,omitempty"`
	ResourceURL        string `json:"resource_url,omitempty"`
	BrowserRedirectURL string `json:"browser_redirect_url,omitempty"`
	ServerLabel        string `json:"server_label,omitempty"`

	AccountLite bool `json:"account_lite:omitempty"`

	Created int `json:"created:omitempty"`
	Used    int `json:"used:omitempty"`
	Expires int `json:"expires:omitempty"`

	User struct {
		ID             string   `json:"id,omitempty"`
		EmailAddresses []string `json:"email_addresses,omitempty"`
		FirstName      string   `json:"first_name,omitempty"`
		LastName       string   `json:"last_name,omitempty"`

		EmailAccounts []GetUsersEmailAccountsResponse `json:"email_accounts,omitempty"`
	}
}

// CreateConnectTokenResponse ...
type CreateConnectTokenResponse struct {
	Success            string `json:"success:omitempty"`
	Token              string `json:"token:omitempty"`
	ResourceURL        string `json:"resource_url:omitempty"`
	BrowserRedirectURL string `json:"browser_redirect_url:omitempty"`
	AccessToken        string `json:"access_token:omitempty"`
	AccessTokenSecret  string `json:"access_token_secret:omitempty"`
}

// DeleteConnectTokenResponse ...
type DeleteConnectTokenResponse struct {
	Success string `json:"success:omitempty"`
}

// GetConnectTokens ...
// List of connect tokens created with your API key
// https://context.io/docs/lite/connect_tokens#get
func (cioLite *CioLite) GetConnectTokens() ([]GetConnectTokensResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   "/connect_tokens",
	}

	// Make response
	var response []GetConnectTokensResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}

// GetConnectToken ...
// Information about a given connect token
// https://context.io/docs/lite/connect_tokens#id-get
func (cioLite *CioLite) GetConnectToken(queryValues CioParams) (GetConnectTokensResponse, error) {

	// Make request
	request := clientRequest{
		method:      "GET",
		path:        "/connect_tokens",
		queryValues: queryValues,
	}

	// Make response
	var response GetConnectTokensResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}

// CreateConnectToken ...
// Obtain a new connect token
// https://context.io/docs/lite/connect_tokens#post
func (cioLite *CioLite) CreateConnectToken(formValues CioParams) (CreateConnectTokenResponse, error) {

	// Make request
	request := clientRequest{
		method:     "POST",
		path:       "/connect_tokens",
		formValues: formValues,
	}

	// Make response
	var response CreateConnectTokenResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}

// DeleteConnectToken ...
// Remove a given connect token
// https://context.io/docs/lite/connect_tokens#id-delete
func (cioLite *CioLite) DeleteConnectToken(queryValues CioParams) (DeleteConnectTokenResponse, error) {

	// Make request
	request := clientRequest{
		method:      "DELETE",
		path:        "/connect_tokens",
		queryValues: queryValues,
	}

	// Make response
	var response DeleteConnectTokenResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}
