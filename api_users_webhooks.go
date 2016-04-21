package ciolite

// Api functions that support: https://context.io/docs/lite/users/webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"strconv"
)

// GetUsersWebhooksResponse data struct
// 	https://context.io/docs/lite/users/webhooks#get
// 	https://context.io/docs/lite/users/webhooks#id-get
type GetUsersWebhooksResponse struct {
	CallbackURL        string `json:"callback_url,omitempty"`
	FailureNotifURL    string `json:"failure_notif_url,omitempty"`
	WebhookID          string `json:"webhook_id,omitempty"`
	FilterTo           string `json:"filter_to,omitempty"`
	FilterFrom         string `json:"filter_from,omitempty"`
	FilterCc           string `json:"filter_cc,omitempty"`
	FilterSubject      string `json:"filter_subject,omitempty"`
	FilterThread       string `json:"filter_thread,omitempty"`
	FilterNewImportant string `json:"filter_new_important,omitempty"`
	FilterFileName     string `json:"filter_file_name,omitempty"`
	FilterFolderAdded  string `json:"filter_folder_added,omitempty"`
	FilterToDomain     string `json:"filter_to_domain,omitempty"`
	FilterFromDomain   string `json:"filter_from_domain,omitempty"`
	BodyType           string `json:"body_type,omitempty"`
	ResourceURL        string `json:"resource_url,omitempty"`

	Active      bool `json:"active,omitempty"`
	Failure     bool `json:"failure,omitempty"`
	IncludeBody bool `json:"include_body,omitempty"`
}

// CreateUserWebhookParams form values data struct.
// Requires: CallbackURL, FailureNotifUrl, and may optionally contain
// FilterTo, FilterFrom, FilterCC, FilterSubject, FilterThread,
// FilterNewImportant, FilterFileName, FilterFolderAdded, FilterToDomain,
// FilterFromDomain, IncludeBody, BodyType
// 	https://context.io/docs/lite/users/webhooks#post
type CreateUserWebhookParams struct {
	// Requires:
	CallbackURL     string `json:"callback_url,omitempty"`
	FailureNotifURL string `json:"failure_notif_url,omitempty"`

	// Optional:
	FilterTo           string `json:"filter_to,omitempty"`
	FilterFrom         string `json:"filter_from,omitempty"`
	FilterCC           string `json:"filter_cc,omitempty"`
	FilterSubject      string `json:"filter_subject,omitempty"`
	FilterThread       string `json:"filter_thread,omitempty"`
	FilterNewImportant string `json:"filter_new_important,omitempty"`
	FilterFileName     string `json:"filter_file_name,omitempty"`
	FilterFolderAdded  string `json:"filter_folder_added,omitempty"`
	FilterToDomain     string `json:"filter_to_domain,omitempty"`
	FilterFromDomain   string `json:"filter_from_domain,omitempty"`
	BodyType           string `json:"body_type,omitempty"`
	IncludeBody        bool   `json:"include_body,omitempty"`
}

// CreateUserWebhookResponse data struct
// 	https://context.io/docs/lite/users/webhooks#post
type CreateUserWebhookResponse struct {
	WebhookID   string `json:"webhook_id,omitempty"`
	ResourceURL string `json:"resource_url,omitempty"`

	Success bool `json:"success,omitempty"`
}

// ModifyUserWebhookParams form values data struct.
// formValues requires Active
// 	https://context.io/docs/lite/users/webhooks#id-post
type ModifyUserWebhookParams struct {
	// Required:
	Active bool `json:"active,omitempty"`
}

// ModifyWebhookResponse data struct
// 	https://context.io/docs/lite/users/webhooks#id-post
type ModifyWebhookResponse struct {
	ResourceURL string `json:"resource_url,omitempty"`

	Success bool `json:"success,omitempty"`
}

// DeleteWebhookResponse data struct
// 	https://context.io/docs/lite/users/webhooks#id-delete
type DeleteWebhookResponse struct {
	Success bool `json:"success,omitempty"`
}

// WebhookCallback data struct that will be received from CIO
// 	https://context.io/docs/lite/users/webhooks#callbacks
type WebhookCallback struct {
	AccountID string `json:"account_id,omitempty"`
	WebhookID string `json:"webhook_id,omitempty"`
	Token     string `json:"token,omitempty" valid:"required"`
	Signature string `json:"signature,omitempty" valid:"required"`

	Timestamp int `json:"timestamp,omitempty" valid:"required"`

	// Data is an error message that gives more information about the cause of failure
	Data string `json:"data,omitempty"`

	MessageData WebhookMessageData `json:"message_data,omitempty"`
}

// WebhookMessageData data struct within WebhookCallback
// 	https://context.io/docs/lite/users/webhooks#callbacks
type WebhookMessageData struct {
	MessageID      string `json:"message_id,omitempty"`
	EmailMessageID string `json:"email_message_id,omitempty"`
	Subject        string `json:"subject,omitempty"`

	References []string `json:"references,omitempty"`
	Folders    []string `json:"folders,omitempty"`

	Date         int `json:"date,omitempty"`
	DateReceived int `json:"date_received,omitempty"`

	Addresses WebhookMessageDataAddresses `json:"addresses,omitempty"`

	PersonInfo PersonInfo `json:"person_info,omitempty"`

	Flags struct {
		Flagged  bool `json:"flagged,omitempty"`
		Answered bool `json:"answered,omitempty"`
		Draft    bool `json:"draft,omitempty"`
		Seen     bool `json:"seen,omitempty"`
	} `json:"flags,omitempty"`

	Sources []struct {
		Label  string `json:"label,omitempty"`
		Folder string `json:"folder,omitempty"`
		UID    int    `json:"uid,omitempty"`
	} `json:"sources,omitempty"`

	EmailAccounts []struct {
		Label  string `json:"label,omitempty"`
		Folder string `json:"folder,omitempty"`
		UID    int    `json:"uid,omitempty"`
	} `json:"email_accounts,omitempty"`

	Files []struct {
		ContentID          string `json:"content_id,omitempty"`
		Type               string `json:"type,omitempty"`
		XAttachmentID      string `json:"x_attachment_id,omitempty"`
		FileName           string `json:"file_name,omitempty"`
		BodySection        string `json:"body_section,omitempty"`
		ContentDisposition string `json:"content_disposition,omitempty"`
		MainFileName       string `json:"main_file_name,omitempty"`

		FileNameStructure [][]string `json:"file_name_structure,omitempty"`

		AttachmentID int `json:"attachment_id,omitempty"`
		Size         int `json:"size,omitempty"`

		IsEmbedded bool `json:"is_embedded,omitempty"`
	} `json:"files,omitempty"`
}

// WebhookMessageDataAddresses struct within WebhookMessageData
// 	https://context.io/docs/lite/users/webhooks#callbacks
type WebhookMessageDataAddresses struct {
	From struct {
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
}

// GetUserWebhooks gets listings of Webhooks configured for a user.
// 	https://context.io/docs/lite/users/webhooks#get
func (cioLite CioLite) GetUserWebhooks(userID string) ([]GetUsersWebhooksResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   fmt.Sprintf("/users/%s/webhooks", userID),
	}

	// Make response
	var response []GetUsersWebhooksResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// GetUserWebhook gets the properties of a given Webhook.
// 	https://context.io/docs/lite/users/webhooks#id-get
func (cioLite CioLite) GetUserWebhook(userID string, webhookID string) (GetUsersWebhooksResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   fmt.Sprintf("/users/%s/webhooks/%s", userID, webhookID),
	}

	// Make response
	var response GetUsersWebhooksResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// CreateUserWebhook creates a new Webhook on a user.
// formValues requires CallbackURL, FailureNotifUrl, and may optionally contain
// FilterTo, FilterFrom, FilterCC, FilterSubject, FilterThread,
// FilterNewImportant, FilterFileName, FilterFolderAdded, FilterToDomain,
// FilterFromDomain, IncludeBody, BodyType
// 	https://context.io/docs/lite/users/webhooks#post
func (cioLite CioLite) CreateUserWebhook(userID string, formValues CreateUserWebhookParams) (CreateUserWebhookResponse, error) {

	// Make request
	request := clientRequest{
		method:     "POST",
		path:       fmt.Sprintf("/users/%s/webhooks", userID),
		formValues: formValues,
	}

	// Make response
	var response CreateUserWebhookResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// ModifyUserWebhook changes the properties of a given Webhook.
// formValues requires Active
// 	https://context.io/docs/lite/users/webhooks#id-post
func (cioLite CioLite) ModifyUserWebhook(userID string, webhookID string, formValues ModifyUserWebhookParams) (ModifyWebhookResponse, error) {

	// Make request
	request := clientRequest{
		method:     "POST",
		path:       fmt.Sprintf("/users/%s/webhooks/%s", userID, webhookID),
		formValues: formValues,
	}

	// Make response
	var response ModifyWebhookResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// DeleteUserWebhookAccount cancels a Webhook.
// 	https://context.io/docs/lite/users/webhooks#id-delete
func (cioLite CioLite) DeleteUserWebhookAccount(userID string, webhookID string) (DeleteWebhookResponse, error) {

	// Make request
	request := clientRequest{
		method: "DELETE",
		path:   fmt.Sprintf("/users/%s/webhooks/%s", userID, webhookID),
	}

	// Make response
	var response DeleteWebhookResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// ValidateWebhookCallback returns true if this WebhookCallback authenticates
func (cioLite CioLite) ValidateWebhookCallback(whc WebhookCallback) bool {
	// Hash timestamp and token with secret, compare to signature
	message := strconv.Itoa(whc.Timestamp) + whc.Token
	hash := hashHmac(sha256.New, message, cioLite.apiSecret)
	return len(hash) > 0 && whc.Signature == hash
}

// hashHmac returns the hash of a message hashed with the provided hash function, using the provided secret
func hashHmac(hashAlgorithm func() hash.Hash, message string, secret string) string {
	h := hmac.New(hashAlgorithm, []byte(secret))
	if _, err := h.Write([]byte(message)); err != nil {
		panic("hash.Hash unable to write message bytes, with error: " + err.Error())
	}
	return hex.EncodeToString(h.Sum(nil))
}
