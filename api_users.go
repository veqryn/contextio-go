package ciolite

// Api functions that support: https://context.io/docs/lite/users

import (
	"fmt"
)

// GetUsersParams query values data struct.
// Optional: Email, Status, StatusOK, Limit, Offset.
// 	https://context.io/docs/lite/users#get
type GetUsersParams struct {
	// Optional:
	Email    string `json:"email,omitempty"`
	Status   string `json:"status,omitempty"`
	StatusOK string `json:"status_ok,omitempty"`
	Limit    int    `json:"limit,omitempty"`
	Offset   int    `json:"offset,omitempty"`
}

// GetUsersResponse data struct
// 	https://context.io/docs/lite/users#get
// 	https://context.io/docs/lite/users#id-get
type GetUsersResponse struct {
	ID             string   `json:"id,omitempty"`
	Username       string   `json:"username,omitempty"`
	EmailAddresses []string `json:"email_addresses,omitempty"`
	FirstName      string   `json:"first_name,omitempty"`
	LastName       string   `json:"last_name,omitempty"`
	ResourceURL    string   `json:"resource_url,omitempty"`

	EmailAccounts []GetUsersEmailAccountsResponse `json:"email_accounts,omitempty"`

	Created         int `json:"created,omitempty"`
	Suspended       int `json:"suspended,omitempty"`
	PasswordExpired int `json:"password_expired,omitempty"`
}

// CreateUserParams form values data struct.
// Can optionally be empty if just creating a user without any email accounts,
// but if creating a user and an email account at the same time then it is required to have:
// Email, Server, Username, UseSSL, Port, Type,
// and (if OAUTH) ProviderRefreshToken and ProviderConsumerKey,
// and (if not OAUTH) Password,
// and may optionally contain StatusCallbackURL,
// and (only for CreateUser) MigrateAccountID, FirstName, LastName.
// 	https://context.io/docs/lite/users#post
// 	https://context.io/docs/lite/users/email_accounts#post
type CreateUserParams struct {
	// Optional, but Required for creating an Email Account
	Email    string `json:"email,omitempty"`
	Server   string `json:"server,omitempty"`
	Username string `json:"username,omitempty"`
	Type     string `json:"type,omitempty"`
	UseSSL   bool   `json:"use_ssl,omitempty"`
	Port     int    `json:"port,omitempty"`

	// Optional, but Required for OAUTH:
	ProviderRefreshToken string `json:"provider_refresh_token,omitempty"`
	ProviderConsumerKey  string `json:"provider_consumer_key,omitempty"`

	// Optional, but Required for non-OAUTH:
	Password string `json:"password,omitempty"`

	// Optional:
	StatusCallbackURL string `json:"status_callback_url,omitempty"`

	// Optional for CreaseUser only (not used by CreateUserEmailAccount):
	MigrateAccountID  string `json:"migrate_account_id,omitempty"`
	FirstName         string `json:"first_name,omitempty"`
	LastName          string `json:"last_name,omitempty"`
}

// CreateUserResponse data struct
// 	https://context.io/docs/lite/users#post
type CreateUserResponse struct {
	Success bool   `json:"success,omitempty"`
	ID      string `json:"id,omitempty"`

	EmailAccount CreateEmailAccountResponse `json:"email_account,omitempty"`

	ResourceURL       string `json:"resource_url,omitempty"`
	AccessToken       string `json:"access_token,omitempty"`
	AccessTokenSecret string `json:"access_token_secret,omitempty"`
}

// ModifyUserParams form values data struct.
// Requires: FirstName, LastName.
// 	https://context.io/docs/lite/users#id-post
type ModifyUserParams struct {
	// Requires:
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// ModifyUserResponse data struct
// 	https://context.io/docs/lite/users#id-post
type ModifyUserResponse struct {
	Success     bool   `json:"success,omitempty"`
	ResourceURL string `json:"resource_url,omitempty"`
}

// DeleteUserResponse data struct
// 	https://context.io/docs/lite/users#id-delete
type DeleteUserResponse struct {
	Success     bool   `json:"success,omitempty"`
	ResourceURL string `json:"resource_url,omitempty"`
}

// GetUsers gets a list of users.
// queryValues may optionally contain Email, Status, StatusOK, Limit, Offset
// 	https://context.io/docs/lite/users#get
func (cioLite CioLite) GetUsers(queryValues GetUsersParams) ([]GetUsersResponse, error) {

	// Make request
	request := clientRequest{
		method:      "GET",
		path:        "/users",
		queryValues: queryValues,
	}

	// Make response
	var response []GetUsersResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// GetUser get details about a given user.
// 	https://context.io/docs/lite/users#id-get
func (cioLite CioLite) GetUser(userID string) (GetUsersResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   fmt.Sprintf("/users/%s", userID),
	}

	// Make response
	var response GetUsersResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// CreateUser create a new user.
// formValues can optionally be empty if just creating a user without any email accounts,
// but if creating a user and an email account at the same time then it is required to have:
// Email, Server, Username, UseSSL, Port, Type,
// and (if OAUTH) ProviderRefreshToken and ProviderConsumerKey,
// and (if not OAUTH) Password, and may optionally contain MigrateAccountID,
// FirstName, LastName, StatusCallbackURL
// 	https://context.io/docs/lite/users#post
func (cioLite CioLite) CreateUser(formValues CreateUserParams) (CreateUserResponse, error) {

	// Make request
	request := clientRequest{
		method:     "POST",
		path:       "/users",
		formValues: formValues,
	}

	// Make response
	var response CreateUserResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// ModifyUser modifies a given user.
// formValues requires FirstName, LastName
// 	https://context.io/docs/lite/users#id-post
func (cioLite CioLite) ModifyUser(userID string, formValues ModifyUserParams) (ModifyUserResponse, error) {

	// Make request
	request := clientRequest{
		method:     "POST",
		path:       fmt.Sprintf("/users/%s", userID),
		formValues: formValues,
	}

	// Make response
	var response ModifyUserResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// DeleteUser removes a given user.
// 	https://context.io/docs/lite/users#id-delete
func (cioLite CioLite) DeleteUser(userID string) (DeleteUserResponse, error) {

	// Make request
	request := clientRequest{
		method: "DELETE",
		path:   fmt.Sprintf("/users/%s", userID),
	}

	// Make response
	var response DeleteUserResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// EmailAccountMatching searches its EmailAccounts array for the provided email address,
// and returns the GetUsersEmailAccountsResponse Email Account that matches it.
func (user GetUsersResponse) EmailAccountMatching(email string) (GetUsersEmailAccountsResponse, error) {
	return FindEmailAccountMatching(user.EmailAccounts, email)
}
