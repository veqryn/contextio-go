package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts/folders/messages/body

// Imports
import (
	"fmt"
)

// GetUserEmailAccountsFolderMessageBodyResponse ...
type GetUserEmailAccountsFolderMessageBodyResponse struct {
	Type        int `json:"type,omitempty"`
	Charset     int `json:"charset,omitempty"`
	Content     int `json:"content,omitempty"`
	BodySection int `json:"body_section,omitempty"`
}

// GetUserEmailAccountsFolderMessageBody ...
// Fetch the message body of a given email
// https://context.io/docs/lite/users/email_accounts/folders/messages/body#get
func (cioLite *CioLite) GetUserEmailAccountsFolderMessageBody(userID string, label string, folder string, messageID string) ([]GetUserEmailAccountsFolderMessageBodyResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s/body", userID, label, folder, messageID),
	}

	// Make response
	var response []GetUserEmailAccountsFolderMessageBodyResponse

	// Request
	err := cioLite.doFormRequest(request, response)

	return response, err
}
