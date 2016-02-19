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
	NbMessages       int `json:"nb_messages,omitempty"`
	NbUnseenMessages int `json:"nb_unseen_messages,omitempty"`
	Delimiter        string `json:"delimiter,omitempty"`
	ResourceURL      string `json:"resource_url,omitempty"`
}

// CreateEmailAccountFolderResponse ...
type CreateEmailAccountFolderResponse struct {
	Success bool `json:"success,omitempty"`
}

// GetUserEmailAccountsFolders gets a list of folders in an email account.
// queryValues may optionally contain CioParams.IncludeNamesOnly
// 	https://context.io/docs/lite/users/email_accounts/folders#get
func (cioLite *CioLite) GetUserEmailAccountsFolders(userID string, label string, queryValues CioParams) ([]GetUsersEmailAccountFoldersResponse, error) {

	// Make request
	request := clientRequest{
		method:      "GET",
		path:        fmt.Sprintf("/users/%s/email_accounts/%s/folders", userID, label),
		queryValues: queryValues,
	}

	// Make response
	var response []GetUsersEmailAccountFoldersResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// GetUserEmailAccountFolder gets information about a given folder.
// queryValues may optionally contain CioParams.Delimiter
// 	https://context.io/docs/lite/users/email_accounts/folders#id-get
func (cioLite *CioLite) GetUserEmailAccountFolder(userID string, label string, folder string, queryValues CioParams) (GetUsersEmailAccountFoldersResponse, error) {

	// Make request
	request := clientRequest{
		method:      "GET",
		path:        fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s", userID, label, folder),
		queryValues: queryValues,
	}

	// Make response
	var response GetUsersEmailAccountFoldersResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// CreateUserEmailAccountFolder create a folder on an email account.
// queryValues may optionally contain CioParams.Delimiter
// 	https://context.io/docs/lite/users/email_accounts/folders#id-post
func (cioLite *CioLite) CreateUserEmailAccountFolder(userID string, label string, folder string, formValues CioParams) (CreateEmailAccountFolderResponse, error) {

	// Make request
	request := clientRequest{
		method:     "POST",
		path:       fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s", userID, label, folder),
		formValues: formValues,
	}

	// Make response
	var response CreateEmailAccountFolderResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
