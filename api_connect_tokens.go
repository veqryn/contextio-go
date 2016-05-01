package ciolite

// Api functions that support: https://context.io/docs/lite/connect_tokens

import (
	"fmt"
)

// GetConnectTokenResponse data struct
// 	https://context.io/docs/lite/connect_tokens#get
// 	https://context.io/docs/lite/connect_tokens#id-get
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

	Created int `json:"created,omitempty"`
	Expires int `json:"expires,omitempty"`
	Used    int `json:"used,omitempty"`

	User GetConnectTokenUserResponse `json:"user,omitempty"`
}

// GetConnectTokenUserResponse data struct within GetConnectTokenResponse
// 	https://context.io/docs/lite/connect_tokens#get
// 	https://context.io/docs/lite/connect_tokens#id-get
type GetConnectTokenUserResponse struct {
	ID             string   `json:"id,omitempty"`
	EmailAddresses []string `json:"email_addresses,omitempty"`
	FirstName      string   `json:"first_name,omitempty"`
	LastName       string   `json:"last_name,omitempty"`
	Created        int      `json:"created,omitempty"`

	EmailAccounts []GetUsersEmailAccountsResponse `json:"email_accounts,omitempty"`
}

// CreateConnectTokenParams form values data struct.
// Requires CallbackURL, and optionally may have
// Email, FirstName, LastName, StatusCallbackURL.
// 	https://context.io/docs/lite/connect_tokens#post
// 	https://context.io/docs/lite/users/connect_tokens#post
type CreateConnectTokenParams struct {
	// Required:
	CallbackURL string `json:"callback_url"`

	// Optional:
	Email             string `json:"email,omitempty"`
	FirstName         string `json:"first_name,omitempty"`
	LastName          string `json:"last_name,omitempty"`
	StatusCallbackURL string `json:"status_callback_url,omitempty"`
}

// CreateConnectTokenResponse data struct
// 	https://context.io/docs/lite/connect_tokens#post
type CreateConnectTokenResponse struct {
	Success            bool   `json:"success,omitempty"`
	Token              string `json:"token,omitempty"`
	ResourceURL        string `json:"resource_url,omitempty"`
	BrowserRedirectURL string `json:"browser_redirect_url,omitempty"`
	AccessToken        string `json:"access_token,omitempty"`
	AccessTokenSecret  string `json:"access_token_secret,omitempty"`
}

// DeleteConnectTokenResponse data struct
// 	https://context.io/docs/lite/connect_tokens#id-delete
type DeleteConnectTokenResponse struct {
	Success bool `json:"success,omitempty"`
}

// GetConnectTokens get a list of connect tokens created with your API key.
// 	https://context.io/docs/lite/connect_tokens#get
func (cioLite CioLite) GetConnectTokens() ([]GetConnectTokenResponse, error) {

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
// 	https://context.io/docs/lite/connect_tokens#id-get
func (cioLite CioLite) GetConnectToken(token string) (GetConnectTokenResponse, error) {

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
// formValues requires CallbackURL, and optionally may have
// Email, FirstName, LastName, StatusCallbackURL
// 	https://context.io/docs/lite/connect_tokens#post
func (cioLite CioLite) CreateConnectToken(formValues CreateConnectTokenParams) (CreateConnectTokenResponse, error) {

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
func (cioLite CioLite) DeleteConnectToken(token string) (DeleteConnectTokenResponse, error) {

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

// EmailAccountMatching searches its EmailAccounts array for the provided email address,
// and returns the GetUsersEmailAccountsResponse Email Account that matches it.
func (user GetConnectTokenUserResponse) EmailAccountMatching(email string) (GetUsersEmailAccountsResponse, error) {
	return FindEmailAccountMatching(user.EmailAccounts, email)
}

// CheckConnectToken checks that the connect token was used and that CIO has access to the account
func (cioLite CioLite) CheckConnectToken(email string, contextioToken string) (GetConnectTokenResponse, error) {

	// Call Api
	response, err := cioLite.GetConnectToken(contextioToken)
	if err != nil {
		return response, err
	}

	// Confirm email matches
	if response.Email != email {
		return response, fmt.Errorf("Email does not match Context.io token")
	}

	// Confirm token was used (accepted/authorized)
	if response.Used == 0 {
		return response, fmt.Errorf("Context.io token not used yet")
	}

	// Confirm user exists
	if len(response.User.ID) == 0 {
		return response, fmt.Errorf("Context.io user not created yet")
	}

	// Confirm we have access
	account, err := response.User.EmailAccountMatching(email)
	if err != nil || account.Status != "OK" {
		return response, fmt.Errorf("Unable to access account using Context.io")
	}

	return response, nil
}
