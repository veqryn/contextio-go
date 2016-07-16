package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts/folders

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/contextio/contextio-go/cioutil"
)

// GetUserEmailAccountsFoldersParams query values data struct.
// Optional: IncludeNamesOnly.
// 	https://context.io/docs/lite/users/email_accounts/folders#get
type GetUserEmailAccountsFoldersParams struct {
	// Optional:
	IncludeNamesOnly bool `json:"include_names_only,omitempty"`
}

// GetUsersEmailAccountFoldersResponse data struct
// 	https://context.io/docs/lite/users/email_accounts/folders#get
// 	https://context.io/docs/lite/users/email_accounts/folders#id-get
type GetUsersEmailAccountFoldersResponse struct {
	Name             string `json:"name,omitempty"`
	SymbolicName     string `json:"symbolic_name,omitempty"`
	NbMessages       int    `json:"nb_messages,omitempty"`
	NbUnseenMessages int    `json:"nb_unseen_messages,omitempty"`
	Delimiter        string `json:"delimiter,omitempty"`
	ResourceURL      string `json:"resource_url,omitempty"`
}

// EmailAccountFolderDelimiterParam query values data struct.
// Optional: Delimiter
// 	https://context.io/docs/lite/users/email_accounts/folders#id-get
// 	https://context.io/docs/lite/users/email_accounts/folders#id-post
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/attachments#get
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/attachments#id-get
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/flags#get
//  https://context.io/docs/lite/users/email_accounts/folders/messages/raw#get
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/read#post
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/read#delete
type EmailAccountFolderDelimiterParam struct {
	// Optional:
	Delimiter string `json:"delimiter,omitempty"`
}

// CreateEmailAccountFolderResponse data struct
// 	https://context.io/docs/lite/users/email_accounts/folders#id-post
type CreateEmailAccountFolderResponse struct {
	Success bool `json:"success,omitempty"`
}

// GetUserEmailAccountsFolders gets a list of folders in an email account.
// queryValues may optionally contain IncludeNamesOnly
// 	https://context.io/docs/lite/users/email_accounts/folders#get
func (cioLite CioLite) GetUserEmailAccountsFolders(userID string, label string, queryValues GetUserEmailAccountsFoldersParams) ([]GetUsersEmailAccountFoldersResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method:      "GET",
		Path:        fmt.Sprintf("/users/%s/email_accounts/%s/folders", userID, label),
		QueryValues: queryValues,
	}

	// Make response
	var response []GetUsersEmailAccountFoldersResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}

// GetUserEmailAccountFolder gets information about a given folder.
// queryValues may optionally contain Delimiter
// 	https://context.io/docs/lite/users/email_accounts/folders#id-get
func (cioLite CioLite) GetUserEmailAccountFolder(userID string, label string, folder string, queryValues EmailAccountFolderDelimiterParam) (GetUsersEmailAccountFoldersResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method:      "GET",
		Path:        fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s", userID, label, url.QueryEscape(folder)),
		QueryValues: queryValues,
	}

	// Make response
	var response GetUsersEmailAccountFoldersResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}

// CreateUserEmailAccountFolder create a folder on an email account.
// This call will fail if the folder already exists.
// queryValues may optionally contain Delimiter
// 	https://context.io/docs/lite/users/email_accounts/folders#id-post
func (cioLite CioLite) CreateUserEmailAccountFolder(userID string, label string, folder string, formValues EmailAccountFolderDelimiterParam) (CreateEmailAccountFolderResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method:     "POST",
		Path:       fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s", userID, label, url.QueryEscape(folder)),
		FormValues: formValues,
	}

	// Make response
	var response CreateEmailAccountFolderResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}

// SafeCreateUserEmailAccountFolder will safely check if a folder exists, and create it if it does not.
// This function returns a bool representing whether it had to create a folder, and any errors it received.
// queryValues may optionally contain Delimiter
func (cioLite CioLite) SafeCreateUserEmailAccountFolder(userID string, label string, folder string, formValues EmailAccountFolderDelimiterParam) (bool, error) {

	existsResponse, err := cioLite.GetUserEmailAccountFolder(userID, label, folder, formValues)
	if err == nil && existsResponse.Name == folder {
		// It exists already, so return false and no error
		return false, nil
	}

	// CIO seems to have issues Getting a single specific folder, and Posting a new folder always gives an error if it already exists, so try getting the folder list and see if it is there already
	allFolders, err := cioLite.GetUserEmailAccountsFolders(userID, label, GetUserEmailAccountsFoldersParams{IncludeNamesOnly: true})
	if err == nil {
		for _, singleFolder := range allFolders {
			if singleFolder.Name == folder {
				return false, nil
			}
		}
	}

	createResponse, err := cioLite.CreateUserEmailAccountFolder(userID, label, folder, formValues)
	if err != nil {
		return true, err
	}
	if !createResponse.Success {
		return true, errors.New("Unable to create folder. CIO returned 200 but with Success=false")
	}
	// Created successfully
	return true, nil
}
