package ciolite

// Api functions that support: https://context.io/docs/lite/users/webhooks

// Imports
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
	Status      string `json:"stats,omitempty"`
	WebhookID   string `json:"webhook_id,omitempty"`
	ResourceURL string `json:"resource_url,omitempty"`
}

// ModifyWebHookResponse ...
type ModifyWebHookResponse struct {
	Success     string `json:"success,omitempty"`
	ResourceURL string `json:"resource_url,omitempty"`
}

// DeleteWebHookResponse ...
type DeleteWebHookResponse struct {
	Success string `json:"success,omitempty"`
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
