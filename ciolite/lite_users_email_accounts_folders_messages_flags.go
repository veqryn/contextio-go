package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts/folders/messages/flags

import (
	"fmt"
)

// GetUserEmailAccountsFolderMessageFlagsResponse data struct
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/flags#get
type GetUserEmailAccountsFolderMessageFlagsResponse struct {
	ResourceURL string `json:"resource_url,omitempty"`

	Flags UserEmailAccountsFolderMessageFlags `json:"flags,omitempty"`
}

// UserEmailAccountsFolderMessageFlags embedded data struct within GetUserEmailAccountsFolderMessageFlagsResponse
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/flags#get
type UserEmailAccountsFolderMessageFlags struct {
	Read     bool `json:"read,omitempty"`
	Answered bool `json:"answered,omitempty"`
	Flagged  bool `json:"flagged,omitempty"`
	Draft    bool `json:"draft,omitempty"`
}

// GetUserEmailAccountsFolderMessageFlags returns the message flags.
// queryValues may optionally contain Delimiter
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/flags#get
func (cioLite CioLite) GetUserEmailAccountsFolderMessageFlags(userID string, label string, folder string, messageID string, queryValues EmailAccountFolderDelimiterParam) (GetUserEmailAccountsFolderMessageFlagsResponse, error) {

	// Make request
	request := clientRequest{
		Method:       "GET",
		Path:         fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s/flags", userID, label, folder, messageID),
		QueryValues:  queryValues,
		UserID:       userID,
		AccountLabel: label,
	}

	// Make response
	var response GetUserEmailAccountsFolderMessageFlagsResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
