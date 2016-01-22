// Package ciolite ...
package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts/folders/messages

// Imports
import (
	"fmt"
)

// GetUsersEmailAccountFoldersResponse ...
type GetUsersEmailAccountFolderMessagesResponse struct {
	EmailMessageID int `json:"email_message_id,omitempty"`
	Subject        int `json:"subject,omitempty"`
	MessageID      int `json:"message_id,omitempty"`
	InReplyTo      int `json:"in_reply_to,omitempty"`
	ResourceURL    int `json:"resource_url,omitempty"`

	Folders         []string `json:"folders,omitempty"`
	ListHeaders     []string `json:"list_headers,omitempty"`
	References      []string `json:"references,omitempty"`
	ReceivedHeaders []string `json:"received_headers,omitempty"`

	Addresses struct {
		From []struct {
			Email string `json:"email,omitempty"`
			Name  string `json:"name,omitempty"`
		} `json:"from,omitempty"`

		To []struct {
			Email string `json:"email,omitempty"`
			Name  string `json:"name,omitempty"`
		} `json:"to,omitempty"`

		Cc []struct {
			Email string `json:"email,omitempty"`
			Name  string `json:"cc,omitempty"`
		} `json:"to,omitempty"`
	} `json:"addresses,omitempty"`

	PersonInfo map[string]interface {
	} `json:"person_info,omitempty"`

	SentAt     int `json:"sent_at,omitempty"`
	ReceivedAt int `json:"received_at,omitempty"`
}

// MoveUserEmailAccountFolderMessageResponse ...
type MoveUserEmailAccountFolderMessageResponse struct {
	Success string `json:"success,omitempty"`
}

// GetUserEmailAccountsFolderMessages ...
// Listings of email messages for a user
// https://context.io/docs/lite/users/email_accounts/folders/messages#get
func (cioLite *CioLite) GetUserEmailAccountsFolderMessages(userID string, label string, folder string) ([]GetUsersEmailAccountFolderMessagesResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages", userID, label, folder),
	}

	// Make response
	var response []GetUsersEmailAccountFolderMessagesResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}

// GetUserEmailAccountFolderMessage ...
// File, contact and other information about a given email message
// https://context.io/docs/lite/users/email_accounts/folders/messages#id-get
func (cioLite *CioLite) GetUserEmailAccountFolderMessage(userID string, label string, folder string, messageID string) (GetUsersEmailAccountFolderMessagesResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s", userID, label, folder, messageID),
	}

	// Make response
	var response GetUsersEmailAccountFolderMessagesResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}

// MoveUserEmailAccountFolderMessage ...
// Move a message
// https://context.io/docs/lite/users/email_accounts/folders/messages#id-put
func (cioLite *CioLite) MoveUserEmailAccountFolderMessage(userID string, label string, folder string, messageID string) (MoveUserEmailAccountFolderMessageResponse, error) {

	// Make request
	request := clientRequest{
		method: "PUT",
		path:   fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s", userID, label, folder, messageID),
	}

	// Make response
	var response MoveUserEmailAccountFolderMessageResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}
