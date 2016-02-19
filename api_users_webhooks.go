package ciolite

// Api functions that support: https://context.io/docs/lite/users/webhooks

import (
	"fmt"
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
	Token     string `json:"token,omitempty"`
	Signature string `json:"signature,omitempty"`

	Timestamp int `json:"timestamp,omitempty"`

	MessageData WebHookMessageData `json:"message_data,omitempty"`
}

// WebHookMessageData ...
type WebHookMessageData struct {
	MessageID      string `json:"message_id,omitempty"`
	EmailMessageID string `json:"email_message_id,omitempty"`
	Subject        string `json:"subject,omitempty"`

	References []string `json:"references,omitempty"`
	Folders    []string `json:"folders,omitempty"`

	Date int `json:"date,omitempty"`

	Addresses  MessageAddresses `json:"addresses,omitempty"`
	PersonInfo PersonInfo       `json:"person_info,omitempty"`
}

// WebHookFailedCallback ...
type WebHookFailedCallback struct {
	AccountID string `json:"account_id,omitempty"`
	WebhookID string `json:"webhook_id,omitempty"`
	Data      string `json:"data,omitempty"`
	Token     string `json:"token,omitempty"`
	Signature string `json:"signature,omitempty"`

	Timestamp int `json:"timestamp,omitempty"`
}

// WebHookCallbackAuthentication ...
type WebHookCallbackAuthentication struct {
	Token     string `json:"token,omitempty" valid:"required"`
	Signature string `json:"signature,omitempty" valid:"required"`

	Timestamp int `json:"timestamp,omitempty" valid:"required"`
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
