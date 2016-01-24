package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts/folders/messages/read

// Imports
import (
	"fmt"
)

// GetUserEmailAccountsFolderMessageHeadersResponse ...
type UserEmailAccountsFolderMessageReadResponse struct {
	Success int `json:"success,omitempty"`
}

// MarkUserEmailAccountsFolderMessageRead ...
// Mark the message as read
// https://context.io/docs/lite/users/email_accounts/folders/messages/read#post
func (cioLite *CioLite) MarkUserEmailAccountsFolderMessageRead(userID string, label string, folder string, messageID string) (UserEmailAccountsFolderMessageReadResponse, error) {

	// Make request
	request := clientRequest{
		method: "POST",
		path:   fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s/read", userID, label, folder, messageID),
	}

	// Make response
	var response UserEmailAccountsFolderMessageReadResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}

// MarkUserEmailAccountsFolderMessageRead ...
// Mark the message as unread
// https://context.io/docs/lite/users/email_accounts/folders/messages/read#delete
func (cioLite *CioLite) MarkUserEmailAccountsFolderMessageUnRead(userID string, label string, folder string, messageID string) (UserEmailAccountsFolderMessageReadResponse, error) {

	// Make request
	request := clientRequest{
		method: "DELETE",
		path:   fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s/read", userID, label, folder, messageID),
	}

	// Make response
	var response UserEmailAccountsFolderMessageReadResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}
