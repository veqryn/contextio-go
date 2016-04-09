package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts/folders/messages/raw

import (
	"fmt"
)

// GetUserEmailAccountsFolderMessageRawResponse data struct
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/body#get
type GetUserEmailAccountsFolderMessageRawResponse string

// GetUserEmailAccountsFolderMessageRaw fetches the raw RFC-822 message text of a given email.
// queryValues may optionally contain CioParams.Delimiter
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/body#get
func (cioLite CioLite) GetUserEmailAccountsFolderMessageRaw(userID string, label string, folder string, messageID string, queryValues CioParams) (GetUserEmailAccountsFolderMessageRawResponse, error) {

	// Make request
	request := clientRequest{
		method:      "GET",
		path:        fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s/raw", userID, label, folder, messageID),
		queryValues: queryValues,
	}

	// Make response
	var response GetUserEmailAccountsFolderMessageRawResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
