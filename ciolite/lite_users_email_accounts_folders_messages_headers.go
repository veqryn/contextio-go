package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts/folders/messages/headers

import (
	"fmt"
)

// GetUserEmailAccountsFolderMessageHeadersParams query values data struct.
// Optional: Delimiter, Raw.
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/headers#get
type GetUserEmailAccountsFolderMessageHeadersParams struct {
	// Optional:
	Delimiter string `json:"delimiter,omitempty"`
	Raw       bool   `json:"raw,omitempty"`
}

// GetUserEmailAccountsFolderMessageHeadersResponse data struct
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/headers#get
type GetUserEmailAccountsFolderMessageHeadersResponse struct {
	ResourceURL string `json:"resource_url,omitempty"`

	Headers map[string][]string `json:"headers,omitempty"`
}

// GetUserEmailAccountsFolderMessageHeaders gets the complete headers of a given email message.
// queryValues may optionally contain Delimiter, Raw
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/headers#get
func (cioLite CioLite) GetUserEmailAccountsFolderMessageHeaders(userID string, label string, folder string, messageID string, queryValues GetUserEmailAccountsFolderMessageHeadersParams) (GetUserEmailAccountsFolderMessageHeadersResponse, error) {

	// Make request
	request := clientRequest{
		Method:       "GET",
		Path:         fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s/headers", userID, label, folder, messageID),
		QueryValues:  queryValues,
		UserID:       userID,
		AccountLabel: label,
	}

	// Make response
	var response GetUserEmailAccountsFolderMessageHeadersResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
