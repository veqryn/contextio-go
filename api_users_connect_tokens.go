package ciolite

// Api functions that support: https://context.io/docs/lite/users/connect_tokens

import (
	"fmt"
)

// GetUserConnectTokens gets a list of connect tokens created for a user.
// 	https://context.io/docs/lite/users/connect_tokens#get
func (cioLite CioLite) GetUserConnectTokens(userID string) ([]GetConnectTokenResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   fmt.Sprintf("/users/%s/connect_tokens", userID),
	}

	// Make response
	var response []GetConnectTokenResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// GetUserConnectToken gets information about a given connect token for a specific user.
// 	https://context.io/docs/lite/users/connect_tokens#id-get
func (cioLite CioLite) GetUserConnectToken(userID string, token string) (GetConnectTokenResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   fmt.Sprintf("/users/%s/connect_tokens/%s", userID, token),
	}

	// Make response
	var response GetConnectTokenResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// CreateUserConnectToken creates and obtains a new connect_token for a specific user.
// formValues requires CioParams.CallbackURL, and may optionally have CioParams.Email,
// CioParams.FirstName, CioParams.LastName, CioParams.StatusCallbackURL
// 	https://context.io/docs/lite/users/connect_tokens#post
func (cioLite CioLite) CreateUserConnectToken(userID string, formValues CioParams) (CreateConnectTokenResponse, error) {

	// Make request
	request := clientRequest{
		method:     "POST",
		path:       fmt.Sprintf("/users/%s/connect_tokens", userID),
		formValues: formValues,
	}

	// Make response
	var response CreateConnectTokenResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// DeleteUserConnectToken removes a given connect token for a specific user.
// 	https://context.io/docs/lite/users/connect_tokens#id-delete
func (cioLite CioLite) DeleteUserConnectToken(userID string, token string) (DeleteConnectTokenResponse, error) {

	// Make request
	request := clientRequest{
		method: "DELETE",
		path:   fmt.Sprintf("/users/%s/connect_tokens/%s", userID, token),
	}

	// Make response
	var response DeleteConnectTokenResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}
