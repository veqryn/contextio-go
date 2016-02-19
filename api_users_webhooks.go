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

// GetUsersWebHooksResponse ...
type GetUsersWebHooksResponse struct {
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

// CreateUserWebHookResponse ...
type CreateUserWebHookResponse struct {
	WebhookID   string `json:"webhook_id,omitempty"`
	ResourceURL string `json:"resource_url,omitempty"`

	Success bool `json:"success,omitempty"`
}

// ModifyWebHookResponse ...
type ModifyWebHookResponse struct {
	ResourceURL string `json:"resource_url,omitempty"`

	Success bool `json:"success,omitempty"`
}

// DeleteWebHookResponse ...
type DeleteWebHookResponse struct {
	Success bool `json:"success,omitempty"`
}

// WebHookCallback ...
type WebHookCallback struct {
	AccountID string `json:"account_id,omitempty"`
	WebhookID string `json:"webhook_id,omitempty"`
	Token     string `json:"token,omitempty" valid:"required"`
	Signature string `json:"signature,omitempty" valid:"required"`

	Timestamp int `json:"timestamp,omitempty" valid:"required"`

	// Data is an error message that gives more information about the cause of failure
	Data string `json:"data,omitempty"`

	MessageData WebHookMessageData `json:"message_data,omitempty"`
}

// WebHookMessageData ...
type WebHookMessageData struct {
	MessageID      string `json:"message_id,omitempty"`
	EmailMessageID string `json:"email_message_id,omitempty"`
	Subject        string `json:"subject,omitempty"`

	References []string `json:"references,omitempty"`
	Folders    []string `json:"folders,omitempty"`

	Date         int `json:"date,omitempty"`
	DateReceived int `json:"date_received,omitempty"`

	Addresses WebHookMessageDataAddresses `json:"addresses,omitempty"`

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

// WebHookMessageDataAddresses ...
type WebHookMessageDataAddresses struct {
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

// GetUserWebHooks gets listings of WebHooks configured for a user.
// 	https://context.io/docs/lite/users/webhooks#get
func (cioLite *CioLite) GetUserWebHooks(userID string) ([]GetUsersWebHooksResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   fmt.Sprintf("/users/%s/webhooks", userID),
	}

	// Make response
	var response []GetUsersWebHooksResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// GetUserWebHook gets the properties of a given WebHook.
// 	https://context.io/docs/lite/users/webhooks#id-get
func (cioLite *CioLite) GetUserWebHook(userID string, webhookID string) (GetUsersWebHooksResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   fmt.Sprintf("/users/%s/webhooks/%s", userID, webhookID),
	}

	// Make response
	var response GetUsersWebHooksResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// CreateUserWebHook creates a new WebHook on a user.
// formValues requires CioParams.CallbackURL, CioParams.FailureNotifUrl, and may optionally contain
// CioParams.FilterTo, CioParams.FilterFrom, CioParams.FilterCC, CioParams.FilterSubject,
// CioParams.FilterThread, CioParams.FilterNewImportant, CioParams.FilterFileName, CioParams.FilterFolderAdded,
// CioParams.FilterToDomain, CioParams.FilterFromDomain, CioParams.IncludeBody, CioParams.BodyType
// 	https://context.io/docs/lite/users/webhooks#post
func (cioLite *CioLite) CreateUserWebHook(userID string, formValues CioParams) (CreateUserWebHookResponse, error) {

	// Make request
	request := clientRequest{
		method:     "POST",
		path:       fmt.Sprintf("/users/%s/webhooks", userID),
		formValues: formValues,
	}

	// Make response
	var response CreateUserWebHookResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// ModifyUserWebHook changes the properties of a given WebHook.
// formValues requires CioParams.Active
// 	https://context.io/docs/lite/users/webhooks#id-post
func (cioLite *CioLite) ModifyUserWebHook(userID string, webhookID string, formValues CioParams) (ModifyWebHookResponse, error) {

	// Make request
	request := clientRequest{
		method:     "POST",
		path:       fmt.Sprintf("/users/%s/webhooks/%s", userID, webhookID),
		formValues: formValues,
	}

	// Make response
	var response ModifyWebHookResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// DeleteUserWebHookAccount cancels a WebHook.
// 	https://context.io/docs/lite/users/webhooks#id-delete
func (cioLite *CioLite) DeleteUserWebHookAccount(userID string, webhookID string) (DeleteWebHookResponse, error) {

	// Make request
	request := clientRequest{
		method: "DELETE",
		path:   fmt.Sprintf("/users/%s/webhooks/%s", userID, webhookID),
	}

	// Make response
	var response DeleteWebHookResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// Valid returns true if this WebHookCallback authenticates
func (whc WebHookCallback) Valid(cioLite *CioLite) bool {
	// Hash timestamp and token with secret, compare to signature
	message := strconv.Itoa(whc.Timestamp) + whc.Token
	hash := hashHmac(sha256.New, message, cioLite.apiSecret)
	return len(hash) > 0 && whc.Signature == hash
}

// hashHmac ...
func hashHmac(hashAlgorithm func() hash.Hash, message string, secret string) string {
	h := hmac.New(hashAlgorithm, []byte(secret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
