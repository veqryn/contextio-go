// Package ciolite ...
package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts/folders

// Imports
import (
	"fmt"
)

// GetUsersEmailAccountFoldersResponse ...
type GetUsersEmailAccountFoldersResponse struct {
	Name             string `json:"name,omitempty"`
	SymbolicName     string `json:"symbolic_name,omitempty"`
	NbMessages       string `json:"nb_messages,omitempty"`
	NbUnseenMessages string `json:"nb_unseen_messages,omitempty"`
	Delimiter        string `json:"delimiter,omitempty"`
	ResourceURL      string `json:"resource_url,omitempty"`
}

// CreateEmailAccountFolderResponse ...
type CreateEmailAccountFolderResponse struct {
	Status      string `json:"stats,omitempty"`
	Delimeter   string `json:"delimeter,omitempty"`
	ResourceURL string `json:"resource_url,omitempty"`
}

// GetUserEmailAccountsFolders ...
// List folders in an email account
// https://context.io/docs/lite/users/email_accounts/folders#get
func (cioLite *CioLite) GetUserEmailAccountsFolders(userID string, label string) ([]GetUsersEmailAccountFoldersResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   fmt.Sprintf("/users/%s/email_accounts/%s/folders", userID, label),
	}

	// Make response
	var response []GetUsersEmailAccountFoldersResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}

// GetUserEmailAccountFolder ...
// Returns information about a given folder
// https://context.io/docs/lite/users/email_accounts/folders#id-get
func (cioLite *CioLite) GetUserEmailAccountFolder(userID string, label string, folder string) (GetUsersEmailAccountFoldersResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s", userID, label, folder),
	}

	// Make response
	var response GetUsersEmailAccountFoldersResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}

// CreateUserEmailAccountFolder ...
// Create a folder on an email account
// https://context.io/docs/lite/users/email_accounts/folders#id-post
func (cioLite *CioLite) CreateUserEmailAccountFolder(userID string, label string, formValues CioParams) (CreateEmailAccountFolderResponse, error) {

	// Make request
	request := clientRequest{
		method:     "POST",
		path:       fmt.Sprintf("/users/%s/email_accounts/%s/folders/folder", userID, label),
		formValues: formValues,
	}

	// Make response
	var response CreateEmailAccountFolderResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}
