package ciolite

// Api functions that support: https://context.io/docs/lite/users/email_accounts

// Imports
import (
	"fmt"
	"strings"
)

// GetUsersEmailAccountsResponse ...
type GetUsersEmailAccountsResponse struct {
	Status             string `json:"status,omitempty"`
	ResourceURL        string `json:"resource_url,omitempty"`
	Type               string `json:"type,omitempty"`
	AuthenticationType string `json:"authentication_type,omitempty"`
	Server             string `json:"server,omitempty"`
	Label              string `json:"label,omitempty"`
	Username           string `json:"username,omitempty"`

	UseSSL bool `json:"use_ssl,omitempty"`

	Port int `json:"port,omitempty"`
}

// CreateEmailAccountResponse ...
type CreateEmailAccountResponse struct {
	Status      string `json:"stats,omitempty"`
	Label       string `json:"label,omitempty"`
	ResourceURL string `json:"resource_url,omitempty"`
}

// ModifyEmailAccountResponse ...
type ModifyEmailAccountResponse struct {
	Success      bool   `json:"success,omitempty"`
	ResourceURL  string `json:"resource_url,omitempty"`
	FeedbackCode string `json:"feedback_code,omitempty"`
}

// DeleteEmailAccountResponse ...
type DeleteEmailAccountResponse struct {
	Success      bool   `json:"success,omitempty"`
	ResourceURL  string `json:"resource_url,omitempty"`
	FeedbackCode string `json:"feedback_code,omitempty"`
}

// GetUserEmailAccounts gets a list of email accounts assigned to a user.
// queryValues may optionally contain CioParams.Status, CioParams.StatusOK
// 	https://context.io/docs/lite/users/email_accounts#get
func (cioLite *CioLite) GetUserEmailAccounts(userID string, queryValues CioParams) ([]GetUsersEmailAccountsResponse, error) {

	// Make request
	request := clientRequest{
		method:      "GET",
		path:        fmt.Sprintf("/users/%s/email_accounts", userID),
		queryValues: queryValues,
	}

	// Make response
	var response []GetUsersEmailAccountsResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// GetUserEmailAccount gets the parameters and status for an email account.
// 	https://context.io/docs/lite/users/email_accounts#id-get
func (cioLite *CioLite) GetUserEmailAccount(userID string, label string) (GetUsersEmailAccountsResponse, error) {

	// Make request
	request := clientRequest{
		method: "GET",
		path:   fmt.Sprintf("/users/%s/email_accounts/%s", userID, label),
	}

	// Make response
	var response GetUsersEmailAccountsResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// CreateUserEmailAccount adds a mailbox to a given user.
// formValues requires CioParams.Email, CioParams.Server, CioParams.Username,
// CioParams.UseSSL, CioParams.Port, CioParams.Type,
// and (if OAUTH) CioParams.ProviderRefreshToken and CioParams.ProviderConsumerKey,
// and (if not OAUTH) CioParams.Password, and may optionally contain CioParams.StatusCallbackURL
// 	https://context.io/docs/lite/users/email_accounts#post
func (cioLite *CioLite) CreateUserEmailAccount(userID string, formValues CioParams) (CreateEmailAccountResponse, error) {

	// Make request
	request := clientRequest{
		method:     "POST",
		path:       fmt.Sprintf("/users/%s/email_accounts", userID),
		formValues: formValues,
	}

	// Make response
	var response CreateEmailAccountResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// ModifyUserEmailAccount modifies an email account on a given user.
// formValues optionally may contain CioParams.Status, CioParams.ForceStatusCheck, CioParams.Password,
// CioParams.ProviderRefreshToken, CioParams.ProviderConsumerKey, CioParams.StatusCallbackURL
// 	https://context.io/docs/lite/users/email_accounts#id-post
func (cioLite *CioLite) ModifyUserEmailAccount(userID string, label string, formValues CioParams) (ModifyEmailAccountResponse, error) {

	// Make request
	request := clientRequest{
		method:     "POST",
		path:       fmt.Sprintf("/users/%s/email_accounts/%s", userID, label),
		formValues: formValues,
	}

	// Make response
	var response ModifyEmailAccountResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// DeleteUserEmailAccount deletes an email account of a user.
// 	https://context.io/docs/lite/users/email_accounts#id-delete
func (cioLite *CioLite) DeleteUserEmailAccount(userID string, label string) (DeleteEmailAccountResponse, error) {

	// Make request
	request := clientRequest{
		method: "DELETE",
		path:   fmt.Sprintf("/users/%s/email_accounts/%s", userID, label),
	}

	// Make response
	var response DeleteEmailAccountResponse

	// Request
	err := cioLite.doFormRequest(request, &response)

	return response, err
}

// FindEmailAccountMatching ...
func FindEmailAccountMatching(emailAccounts []GetUsersEmailAccountsResponse, email string) (GetUsersEmailAccountsResponse, error) {

	if emailAccounts != nil {

		// Try to match against the username
		for _, emailAccount := range emailAccounts {
			if email == emailAccount.Username {
				return emailAccount, nil
			}
		}

		// Try to match the local part against the username or label
		localPart := upToSeparator(email, "@")

		for _, emailAccount := range emailAccounts {

			if localPart == upToSeparator(emailAccount.Username, "@") ||
				localPart == upToSeparator(emailAccount.Label, ":") {

				return emailAccount, nil
			}
		}
	}
	return GetUsersEmailAccountsResponse{}, fmt.Errorf("No email accounts match %s in %v", email, emailAccounts)
}

func upToSeparator(s string, sep string) string {
	idx := strings.Index(s, sep)
	if idx >= 0 {
		return s[:idx]
	}
	return s
}
