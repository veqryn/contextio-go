package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts/folders/messages/read

import (
	"fmt"
	"net/url"

	"github.com/contextio/contextio-go/cioutil"
)

// UserEmailAccountsFolderMessageReadResponse data struct
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/read#post
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/read#delete
type UserEmailAccountsFolderMessageReadResponse struct {
	Success bool `json:"success,omitempty"`
}

// MarkUserEmailAccountsFolderMessageRead marks the message as read.
// formValues may optionally contain Delimiter
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/read#post
func (cioLite CioLite) MarkUserEmailAccountsFolderMessageRead(userID string, label string, folder string, messageID string, formValues EmailAccountFolderDelimiterParam) (UserEmailAccountsFolderMessageReadResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method:     "POST",
		Path:       fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s/read", userID, label, url.QueryEscape(folder), url.QueryEscape(messageID)),
		FormValues: formValues,
	}

	// Make response
	var response UserEmailAccountsFolderMessageReadResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}

// MarkUserEmailAccountsFolderMessageUnRead marks the message as unread.
// formValues may optionally contain Delimiter
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/read#delete
func (cioLite CioLite) MarkUserEmailAccountsFolderMessageUnRead(userID string, label string, folder string, messageID string, formValues EmailAccountFolderDelimiterParam) (UserEmailAccountsFolderMessageReadResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method:     "DELETE",
		Path:       fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s/read", userID, label, url.QueryEscape(folder), url.QueryEscape(messageID)),
		FormValues: formValues,
	}

	// Make response
	var response UserEmailAccountsFolderMessageReadResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}
