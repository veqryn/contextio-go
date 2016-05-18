package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts/folders/messages/attachments

import (
	"fmt"

	"github.com/contextio/contextio-go/cioutil"
)

// GetUserEmailAccountsFolderMessageAttachmentsResponse data struct
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/attachments#get
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/attachments#id-get
type GetUserEmailAccountsFolderMessageAttachmentsResponse struct {
	Type               int `json:"type,omitempty"`
	FileName           int `json:"file_name,omitempty"`
	BodySection        int `json:"body_section,omitempty"`
	ContentDisposition int `json:"content_disposition,omitempty"`
	EmailMessageID     int `json:"email_message_id,omitempty"`
	XAttachmentID      int `json:"x_attachment_id,omitempty"`

	Size         int `json:"size,omitempty"`
	AttachmentID int `json:"attachment_id,omitempty"`
}

// GetUserEmailAccountsFolderMessageAttachments gets listings of email attachments.
// queryValues may optionally contain Delimiter
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/attachments#get
func (cioLite CioLite) GetUserEmailAccountsFolderMessageAttachments(userID string, label string, folder string, messageID string, queryValues EmailAccountFolderDelimiterParam) ([]GetUserEmailAccountsFolderMessageAttachmentsResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method:      "GET",
		Path:        fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s/attachments", userID, label, folder, messageID),
		QueryValues: queryValues,
	}

	// Make response
	var response []GetUserEmailAccountsFolderMessageAttachmentsResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}

// GetUserEmailAccountsFolderMessageAttachment retrieves an email attachment.
// queryValues may optionally contain Delimiter
// 	https://context.io/docs/lite/users/email_accounts/folders/messages/attachments#id-get
func (cioLite CioLite) GetUserEmailAccountsFolderMessageAttachment(userID string, label string, folder string, messageID string, attachmentID string, queryValues EmailAccountFolderDelimiterParam) (GetUserEmailAccountsFolderMessageAttachmentsResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method:      "GET",
		Path:        fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s/attachments/%s", userID, label, folder, messageID, attachmentID),
		QueryValues: queryValues,
	}

	// Make response
	var response GetUserEmailAccountsFolderMessageAttachmentsResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}
