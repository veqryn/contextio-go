package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts/folders/messages/flags

import (
	"fmt"

	"github.com/contextio/contextio-go/cioutil"
)

// GetUserEmailAccountsFolderMessageFlagsResponse data struct
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/flags#get
type GetUserEmailAccountsFolderMessageFlagsResponse struct {
	ResourceURL string `json:"resource_url,omitempty"`

	Flags struct {
		Read     bool `json:"read,omitempty"`
		Answered bool `json:"answered,omitempty"`
		Flagged  bool `json:"flagged,omitempty"`
		Draft    bool `json:"draft,omitempty"`
	} `json:"flags,omitempty"`
}

// GetUserEmailAccountsFolderMessageFlags returns the message flags.
// queryValues may optionally contain Delimiter
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/flags#get
func (cioLite CioLite) GetUserEmailAccountsFolderMessageFlags(userID string, label string, folder string, messageID string, queryValues EmailAccountFolderDelimiterParam) (GetUserEmailAccountsFolderMessageFlagsResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method:      "GET",
		Path:        fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s/flags", userID, label, folder, messageID),
		QueryValues: queryValues,
	}

	// Make response
	var response GetUserEmailAccountsFolderMessageFlagsResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}
