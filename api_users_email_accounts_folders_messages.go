package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts/folders/messages

// Imports
import (
	"fmt"
)

// GetUsersEmailAccountFolderMessagesResponse ...
type GetUsersEmailAccountFolderMessagesResponse struct {
	EmailMessageID string `json:"email_message_id,omitempty"`
	Subject        string `json:"subject,omitempty"`
	MessageID      string `json:"message_id,omitempty"`
	InReplyTo      string `json:"in_reply_to,omitempty"`
	ResourceURL    string `json:"resource_url,omitempty"`

	Folders         []string `json:"folders,omitempty"`
	ListHeaders     []string `json:"list_headers,omitempty"`
	References      []string `json:"references,omitempty"`
	ReceivedHeaders []string `json:"received_headers,omitempty"`

	Addresses struct {
		From []struct {
			Email string `json:"email,omitempty"`
			Name  string `json:"name,omitempty"`
		} `json:"from,omitempty"`

		To []struct {
			Email string `json:"email,omitempty"`
			Name  string `json:"name,omitempty"`
		} `json:"to,omitempty"`

		Cc []struct {
			Email string `json:"email,omitempty"`
			Name  string `json:"name,omitempty"`
		} `json:"cc,omitempty"`

		Bcc []struct {
			Email string `json:"email,omitempty"`
			Name  string `json:"name,omitempty"`
		} `json:"bcc,omitempty"`

		Sender []struct {
			Email string `json:"email,omitempty"`
			Name  string `json:"name,omitempty"`
		} `json:"sender,omitempty"`

		ReplyTo []struct {
			Email string `json:"email,omitempty"`
			Name  string `json:"name,omitempty"`
		} `json:"reply_to,omitempty"`
	} `json:"addresses,omitempty"`

	PersonInfo map[string]interface {
	} `json:"person_info,omitempty"`

	Attachments []struct {
		Type               string `json:"type,omitempty"`
		FileName           string `json:"file_name,omitempty"`
		BodySection        string `json:"body_section,omitempty"`
		ContentDisposition string `json:"content_disposition,omitempty"`
		EmailMessageID     string `json:"email_message_id,omitempty"`
		XAttachmentID      string `json:"x_attachment_id,omitempty"`

		Size         int `json:"size,omitempty"`
		AttachmentID int `json:"attachment_id,omitempty"`
	} `json:"attachments,omitempty"`

	SentAt     int `json:"sent_at,omitempty"`
	ReceivedAt int `json:"received_at,omitempty"`
}

// MoveUserEmailAccountFolderMessageResponse ...
type MoveUserEmailAccountFolderMessageResponse struct {
	Success bool `json:"success,omitempty"`
}

// GetUserEmailAccountsFolderMessages gets listings of email messages for a user.
// queryValues may optionally contain CioParams.Delimiter, CioParams.IncludeBody, CioParams.BodyType,
// CioParams.IncludeHeaders, CioParams.IncludeFlags, CioParams.Limit, CioParams.Offset
// 	https://context.io/docs/lite/users/email_accounts/folders/messages#get
func (cioLite *CioLite) GetUserEmailAccountsFolderMessages(userID string, label string, folder string, queryValues CioParams) ([]GetUsersEmailAccountFolderMessagesResponse, error) {

	// Make request
	request := clientRequest{
		method:      "GET",
		path:        fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages", userID, label, folder),
		queryValues: queryValues,
	}

	// Make response
	var response []GetUsersEmailAccountFolderMessagesResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// GetUserEmailAccountFolderMessage gets file, contact and other information about a given email message.
// queryValues may optionally contain CioParams.Delimiter, CioParams.IncludeBody,
// CioParams.BodyType, CioParams.IncludeHeaders, CioParams.IncludeFlags
// 	https://context.io/docs/lite/users/email_accounts/folders/messages#id-get
func (cioLite *CioLite) GetUserEmailAccountFolderMessage(userID string, label string, folder string, messageID string, queryValues CioParams) (GetUsersEmailAccountFolderMessagesResponse, error) {

	// Make request
	request := clientRequest{
		method:      "GET",
		path:        fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s", userID, label, folder, messageID),
		queryValues: queryValues,
	}

	// Make response
	var response GetUsersEmailAccountFolderMessagesResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// MoveUserEmailAccountFolderMessage moves a message.
// formValues requires CioParams.NewFolderID, and may optionally contain CioParams.Delimiter
// 	https://context.io/docs/lite/users/email_accounts/folders/messages#id-put
func (cioLite *CioLite) MoveUserEmailAccountFolderMessage(userID string, label string, folder string, messageID string, formValues CioParams) (MoveUserEmailAccountFolderMessageResponse, error) {

	// Make request
	request := clientRequest{
		method:     "PUT",
		path:       fmt.Sprintf("/users/%s/email_accounts/%s/folders/%s/messages/%s", userID, label, folder, messageID),
		formValues: formValues,
	}

	// Make response
	var response MoveUserEmailAccountFolderMessageResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
