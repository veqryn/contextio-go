package ciolite

// Api functions that support: https://context.io/docs/lite/connect_tokens

import (
	"fmt"
)

// GetConnectTokenResponse ...
type GetConnectTokenResponse struct {
	Token              string `json:"token,omitempty"`
	Email              string `json:"email,omitempty"`
	CallbackURL        string `json:"callback_url,omitempty"`
	FirstName          string `json:"first_name,omitempty"`
	LastName           string `json:"last_name,omitempty"`
	ResourceURL        string `json:"resource_url,omitempty"`
	BrowserRedirectURL string `json:"browser_redirect_url,omitempty"`
	ServerLabel        string `json:"server_label,omitempty"`

	AccountLite bool `json:"account_lite,omitempty"`
	Expires     bool `json:"expires,omitempty"`

	Created int `json:"created,omitempty"`
	Used    int `json:"used,omitempty"`

	User GetConnectTokenUserResponse `json:"user,omitempty"`
}

// GetConnectTokenUserResponse ...
type GetConnectTokenUserResponse struct {
	ID             string   `json:"id,omitempty"`
	EmailAddresses []string `json:"email_addresses,omitempty"`
	FirstName      string   `json:"first_name,omitempty"`
	LastName       string   `json:"last_name,omitempty"`
	Created        int      `json:"created,omitempty"`

	EmailAccounts []GetUsersEmailAccountsResponse `json:"email_accounts,omitempty"`
}

// CreateConnectTokenResponse ...
type CreateConnectTokenResponse struct {
	Success            bool   `json:"success,omitempty"`
	Token              string `json:"token,omitempty"`
	ResourceURL        string `json:"resource_url,omitempty"`
	BrowserRedirectURL string `json:"browser_redirect_url,omitempty"`
	AccessToken        string `json:"access_token,omitempty"`
	AccessTokenSecret  string `json:"access_token_secret,omitempty"`
}

// DeleteConnectTokenResponse ...
type DeleteConnectTokenResponse struct {
	Success bool `json:"success,omitempty"`
}

// GetConnectTokens get a list of connect tokens created with your API key.
// 	https://context.io/docs/lite/connect_tokens#get
func (cioLite *CioLite) GetConnectTokens() ([]GetConnectTokenResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   "/connect_tokens",
	}

	// Make response
	var response []GetConnectTokenResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// GetConnectToken gets information about a given connect token.
// https://context.io/docs/lite/connect_tokens#id-get
func (cioLite *CioLite) GetConnectToken(token string) (GetConnectTokenResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   fmt.Sprintf("/connect_tokens/%s", token),
	}

	// Make response
	var response GetConnectTokenResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// CreateConnectToken creates and obtains a new connect token.
// formValues requires CioParams.CallbackURL, and optionally may have
// CioParams.Email, CioParams.FirstName, CioParams.LastName, CioParams.StatusCallbackURL
// 	https://context.io/docs/lite/connect_tokens#post
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
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// DeleteConnectToken removes a given connect token
// 	https://context.io/docs/lite/connect_tokens#id-delete
func (cioLite *CioLite) DeleteConnectToken(token string) (DeleteConnectTokenResponse, error) {

	// Make request
	request := clientRequest{
		method: "DELETE",
		path:   fmt.Sprintf("/connect_tokens/%s", token),
	}

	// Make response
	var response DeleteConnectTokenResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// EmailAccountMatching ...
func (user GetConnectTokenUserResponse) EmailAccountMatching(email string) (GetUsersEmailAccountsResponse, error) {
	return FindEmailAccountMatching(user.EmailAccounts, email)
}
