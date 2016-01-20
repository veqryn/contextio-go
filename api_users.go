// Package ciolite ...
package ciolite

// Api functions that support: https://context.io/docs/lite/users

// Imports
import (
	"fmt"
)

// GetUsersResponse ...
type GetUsersResponse struct {
	ID             string   `json:"id,omitempty"`
	Username       string   `json:"username,omitempty"`
	EmailAddresses []string `json:"email_addresses,omitempty"`
	FirstName      string   `json:"first_name,omitempty"`
	LastName       string   `json:"last_name,omitempty"`
	ResourceURL    string   `json:"resource_url:omitempty"`

	EmailAccounts []GetUsersEmailAccountsResponse `json:"email_accounts,omitempty"`

	Created         int `json:"created:omitempty"`
	Suspended       int `json:"suspended:omitempty"`
	PasswordExpired int `json:"password_expired:omitempty"`
}

// CreateUserResponse ...
type CreateUserResponse struct {
	Success           string `json:"success:omitempty"`
	Id                string `json:"id:omitempty"`
	
	EmailAccount CreateEmailAccountResponse `json:"email_account,omitempty"`
	
	ResourceURL       string `json:"resource_url:omitempty"`
	AccessToken       string `json:"access_token:omitempty"`
	AccessTokenSecret string `json:"access_token_secret:omitempty"`
}

// ModifyUsersResponse ...
type ModifyUsersResponse struct {
	Success     string `json:"success:omitempty"`
	ResourceURL string `json:"resource_url:omitempty"`
}

// ModifyUserResponse ...
type ModifyUserResponse struct {
	Success     string `json:"success:omitempty"`
	ResourceURL string `json:"resource_url:omitempty"`
}

// DeleteUserResponse ...
type DeleteUserResponse struct {
	Success     string `json:"success:omitempty"`
	ResourceURL string `json:"resource_url:omitempty"`
}

// GetUsers ...
// List of users
// https://context.io/docs/lite/users#get
func (cioLite *CioLite) GetUsers() ([]GetUsersResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   "/users",
	}

	// Make response
	var response []GetUsersResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}

// GetUser ...
// Details about a given user
// https://context.io/docs/lite/users#id-get
func (cioLite *CioLite) GetUser(userID string) (GetUsersResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   fmt.Sprintf("/users/%s", userID),
	}

	// Make response
	var response GetUsersResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}

// CreateUser ...
// Add a new user
// https://context.io/docs/lite/users#post
func (cioLite *CioLite) CreateUser(formValues CioParams) (CreateUserResponse, error) {

	// Make request
	request := clientRequest{
		method:     "POST",
		path:       "/users",
		formValues: formValues,
	}

	// Make response
	var response CreateUserResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}

// ModifyUser ...
// Modify a given user
// https://context.io/docs/lite/users#id-post
func (cioLite *CioLite) ModifyUser(userID string, formValues CioParams) (ModifyUserResponse, error) {

	// Make request
	request := clientRequest{
		method:     "POST",
		path:       fmt.Sprintf("/users/%s", userID),
		formValues: formValues,
	}

	// Make response
	var response ModifyUserResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}

// DeleteUser ...
// Remove a given user
// https://context.io/docs/lite/users#id-delete
func (cioLite *CioLite) DeleteUser(userID string) (DeleteUserResponse, error) {

	// Make request
	request := clientRequest{
		method: "DELETE",
		path:   fmt.Sprintf("/users/%s", userID),
	}

	// Make response
	var response DeleteUserResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}
