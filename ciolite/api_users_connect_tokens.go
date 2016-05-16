package ciolite

// Api functions that support: https://context.io/docs/lite/users/connect_tokens

import (
	"fmt"

	"github.com/contextio/contextio-go/cioutil"
)

// GetUserConnectTokens gets a list of connect tokens created for a user.
// 	https://context.io/docs/lite/users/connect_tokens#get
func (cioLite CioLite) GetUserConnectTokens(userID string) ([]GetConnectTokenResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method: "GET",
		Path:   fmt.Sprintf("/users/%s/connect_tokens", userID),
	}

	// Make response
	var response []GetConnectTokenResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}

// GetUserConnectToken gets information about a given connect token for a specific user.
// 	https://context.io/docs/lite/users/connect_tokens#id-get
func (cioLite CioLite) GetUserConnectToken(userID string, token string) (GetConnectTokenResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method: "GET",
		Path:   fmt.Sprintf("/users/%s/connect_tokens/%s", userID, token),
	}

	// Make response
	var response GetConnectTokenResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}

// CreateUserConnectToken creates and obtains a new connect_token for a specific user.
// formValues requires CallbackURL, and may optionally have
// Email, FirstName, LastName, StatusCallbackURL
// 	https://context.io/docs/lite/users/connect_tokens#post
func (cioLite CioLite) CreateUserConnectToken(userID string, formValues CreateConnectTokenParams) (CreateConnectTokenResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method:     "POST",
		Path:       fmt.Sprintf("/users/%s/connect_tokens", userID),
		FormValues: formValues,
	}

	// Make response
	var response CreateConnectTokenResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}

// DeleteUserConnectToken removes a given connect token for a specific user.
// 	https://context.io/docs/lite/users/connect_tokens#id-delete
func (cioLite CioLite) DeleteUserConnectToken(userID string, token string) (DeleteConnectTokenResponse, error) {

	// Make request
	request := cioutil.ClientRequest{
		Method: "DELETE",
		Path:   fmt.Sprintf("/users/%s/connect_tokens/%s", userID, token),
	}

	// Make response
	var response DeleteConnectTokenResponse

	// Request
	err := cioLite.DoFormRequest(request, &response)

	return response, err
}
