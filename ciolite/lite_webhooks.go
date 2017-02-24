package ciolite

// Api functions that support: https://context.io/docs/lite/webhooks

import (
	"fmt"
)

// GetWebhooks gets listings of Webhooks configured for the application.
// 	https://context.io/docs/lite/webhooks#get
func (cioLite CioLite) GetWebhooks() ([]GetUsersWebhooksResponse, error) {

	// Make request
	request := clientRequest{
		Method: "GET",
		Path:   "/lite/webhooks",
	}

	// Make response
	var response []GetUsersWebhooksResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// GetWebhook gets the properties of a given Webhook.
// 	https://context.io/docs/lite/webhooks#id-get
func (cioLite CioLite) GetWebhook(webhookID string) (GetUsersWebhooksResponse, error) {

	// Make request
	request := clientRequest{
		Method: "GET",
		Path:   fmt.Sprintf("/lite/webhooks/%s", webhookID),
	}

	// Make response
	var response GetUsersWebhooksResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// CreateWebhook creates a new Webhook for the application.
// formValues requires CallbackURL, FailureNotifUrl, and may optionally contain
// FilterTo, FilterFrom, FilterCC, FilterSubject, FilterThread,
// FilterNewImportant, FilterFileName, FilterFolderAdded, FilterToDomain,
// FilterFromDomain, IncludeBody, BodyType
// 	https://context.io/docs/lite/webhooks#post
func (cioLite CioLite) CreateWebhook(formValues CreateUserWebhookParams) (CreateUserWebhookResponse, error) {

	// Make request
	request := clientRequest{
		Method:     "POST",
		Path:       "/lite/webhooks",
		FormValues: formValues,
	}

	// Make response
	var response CreateUserWebhookResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// ModifyWebhook changes the properties of a given Webhook.
// formValues requires Active
// 	https://context.io/docs/lite/webhooks#id-post
func (cioLite CioLite) ModifyWebhook(webhookID string, formValues ModifyUserWebhookParams) (ModifyWebhookResponse, error) {

	// Make request
	request := clientRequest{
		Method:     "POST",
		Path:       fmt.Sprintf("/lite/webhooks/%s", webhookID),
		FormValues: formValues,
	}

	// Make response
	var response ModifyWebhookResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// DeleteWebhookAccount cancels a Webhook.
// 	https://context.io/docs/lite/webhooks#id-delete
func (cioLite CioLite) DeleteWebhookAccount(webhookID string) (DeleteWebhookResponse, error) {

	// Make request
	request := clientRequest{
		Method: "DELETE",
		Path:   fmt.Sprintf("/lite/webhooks/%s", webhookID),
	}

	// Make response
	var response DeleteWebhookResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
