package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts/folders/messages/headers

// Imports
import (
	"fmt"
)

// GetUserEmailAccountsFolderMessageHeadersResponse ...
type GetUserEmailAccountsFolderMessageHeadersResponse struct {
	ResourceURL int `json:"resource_url,omitempty"`

	Headers map[string]interface{} `json:"headers,omitempty"`
}

// GetUserEmailAccountsFolderMessageHeaders gets the complete headers of a given email message.
// queryValues may optionally contain CioParams.Delimiter, CioParams.Raw
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/headers#get
func (cioLite CioLite) GetUserEmailAccountsFolderMessageHeaders(userID string, label string, folder string, messageID string, queryValues CioParams) (GetUserEmailAccountsFolderMessageHeadersResponse, error) {

	// Make request
	request := clientRequest{
		method:      "GET",
		path:        fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s/headers", userID, label, folder, messageID),
		queryValues: queryValues,
	}

	// Make response
	var response GetUserEmailAccountsFolderMessageHeadersResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
