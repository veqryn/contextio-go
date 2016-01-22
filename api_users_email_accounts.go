// Package ciolite ...
package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts

// Imports
import (
	"fmt"
)

// GetUsersEmailAccountsResponse ...
type GetUsersEmailAccountsResponse struct {
	Status             string `json:"status,omitempty"`
	ResourceURL        string `json:"resource_url,omitempty"`
	Type               string `json:"type,omitempty"`
	AuthenticationType string `json:"authentication_type,omitempty"`
	Server             string `json:"server,omitempty"`
	Label              string `json:"label,omitempty"`
	Username           string `json:"username,omitempty"`

	UseSSL bool `json:"use_ssl,omitempty"`

	Port int `json:"port,omitempty"`
}

// CreateEmailAccountResponse ...
type CreateEmailAccountResponse struct {
	Status      string `json:"stats,omitempty"`
	Label       string `json:"label,omitempty"`
	ResourceURL string `json:"resource_url,omitempty"`
}

// ModifyEmailAccountResponse ...
type ModifyEmailAccountResponse struct {
	Success      string `json:"success,omitempty"`
	ResourceURL  string `json:"resource_url,omitempty"`
	FeedbackCode string `json:"feedback_code,omitempty"`
}

// DeleteEmailAccountResponse ...
type DeleteEmailAccountResponse struct {
	Success      string `json:"success,omitempty"`
	ResourceURL  string `json:"resource_url,omitempty"`
	FeedbackCode string `json:"feedback_code,omitempty"`
}

// GetUserEmailAccounts ...
// List of email accounts assigned to a user
// https://context.io/docs/lite/users/email_accounts#get
func (cioLite *CioLite) GetUserEmailAccounts(userID string) ([]GetUsersEmailAccountsResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   fmt.Sprintf("/users/%s/email_accounts", userID),
	}

	// Make response
	var response []GetUsersEmailAccountsResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}

// GetUserEmailAccount ...
// Parameters and status for an email account
// https://context.io/docs/lite/users/email_accounts#id-get
func (cioLite *CioLite) GetUserEmailAccount(userID string, label string) (GetUsersEmailAccountsResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   fmt.Sprintf("/users/%s/email_accounts/%s", userID, label),
	}

	// Make response
	var response GetUsersEmailAccountsResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}

// CreateUserEmailAccount ...
// Add a mailbox to a given user
// https://context.io/docs/lite/users/email_accounts#post
func (cioLite *CioLite) CreateUserEmailAccount(userID string, formValues CioParams) (CreateEmailAccountResponse, error) {

	// Make request
	request := clientRequest{
		method:     "POST",
		path:       fmt.Sprintf("/users/%s/email_accounts", userID),
		formValues: formValues,
	}

	// Make response
	var response CreateEmailAccountResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}

// ModifyUserEmailAccount ...
// Modify an email account on a given user
// https://context.io/docs/lite/users/email_accounts#id-post
func (cioLite *CioLite) ModifyUserEmailAccount(userID string, label string, formValues CioParams) (ModifyEmailAccountResponse, error) {

	// Make request
	request := clientRequest{
		method:     "POST",
		path:       fmt.Sprintf("/users/%s/email_accounts/%s", userID, label),
		formValues: formValues,
	}

	// Make response
	var response ModifyEmailAccountResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}

// DeleteUserEmailAccount ...
// Delete an email account of a user
// https://context.io/docs/lite/users/email_accounts#id-delete
func (cioLite *CioLite) DeleteUserEmailAccount(userID string, label string) (DeleteEmailAccountResponse, error) {

	// Make request
	request := clientRequest{
		method: "DELETE",
		path:   fmt.Sprintf("/users/%s/email_accounts/%s", userID, label),
	}

	// Make response
	var response DeleteEmailAccountResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}
