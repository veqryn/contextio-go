package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts/folders/messages/body

import (
	"fmt"
)

// GetUserEmailAccountsFolderMessageBodyResponse data struct
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/body#get
type GetUserEmailAccountsFolderMessageBodyResponse struct {
	Type        string `json:"type,omitempty"`
	Charset     string `json:"charset,omitempty"`
	Content     string `json:"content,omitempty"`
	BodySection string `json:"body_section,omitempty"`
}

// GetUserEmailAccountsFolderMessageBody fetches the message body of a given email.
// queryValues may optionally contain CioParams.Delimiter, CioParams.Type
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/body#get
func (cioLite CioLite) GetUserEmailAccountsFolderMessageBody(userID string, label string, folder string, messageID string, queryValues CioParams) ([]GetUserEmailAccountsFolderMessageBodyResponse, error) {

	// Make request
	request := clientRequest{
		method:      "GET",
		path:        fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s/body", userID, label, folder, messageID),
		queryValues: queryValues,
	}

	// Make response
	var response []GetUserEmailAccountsFolderMessageBodyResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
