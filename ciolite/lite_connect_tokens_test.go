package ciolite

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

// TestActualConnectTokenRequestToCioForGoogle tests actual CreateConnectToken,
// GetConnectToken, GetConnectTokens, and DeleteConnectToken requests to CIO.
// (internet connection required, real CIO key/secret required,
// gmail provider key setup previously required).
func TestActualConnectTokenRequestToCio(t *testing.T) {
	t.Parallel()

	cioLite, logger := NewTestCioLiteWithLogger(t)

	// create
	connectToken, err := cioLite.CreateConnectToken(CreateConnectTokenParams{
		CallbackURL: "https://bogusurl.com",
		Email:       "test@gmail.com",
	})

	if err != nil || !connectToken.Success || len(connectToken.BrowserRedirectURL) == 0 || len(connectToken.Token) == 0 {
		t.Error("Expected successful connect token; Got: ", connectToken, "; With Error: ", err, "; With Log: ", logger.String())
	}

	// get this one
	getConnectToken, err := cioLite.GetConnectToken(connectToken.Token)

	if err != nil ||
		getConnectToken.Email != "test@gmail.com" ||
		getConnectToken.CallbackURL != "https://bogusurl.com" ||
		!getConnectToken.AccountLite ||
		getConnectToken.Token != connectToken.Token ||
		// getConnectToken.BrowserRedirectURL != connectToken.BrowserRedirectURL ||
		getConnectToken.Used != 0 ||
		getConnectToken.User.ID != "" {

		t.Error("Expected GetConnectTokenResponse matching: ", connectToken, "; Got: ", getConnectToken, "; With Error: ", err, "; With Log: ", logger.String())
	}

	// get all
	getConnectTokens, err := cioLite.GetConnectTokens()

	found := false
	for _, getConnectTokenValue := range getConnectTokens {
		if reflect.DeepEqual(getConnectTokenValue, getConnectToken) {
			found = true
		}
	}

	if !found {
		t.Error("Expected to include: ", getConnectToken, "; Got: ", getConnectTokens, "; With Log: ", logger.String())
	}

	// check with bad email
	err = cioLite.CheckConnectToken(getConnectToken, "not_correct@gmail.com")
	expectedErrorText := "Email does not match Context.io token"
	if err.Error() != expectedErrorText {
		t.Error("Expected error: ", expectedErrorText, "; Got: ", err)
	}

	// check with good email
	err = cioLite.CheckConnectToken(getConnectToken, "test@gmail.com")
	expectedErrorText = "Context.io token not used yet"
	if err.Error() != expectedErrorText {
		t.Error("Expected error: ", expectedErrorText, "; Got: ", err)
	}

	// Check if accepted token's email accounts match (they shouldn't since its not accepted)
	_, err = getConnectToken.User.EmailAccountMatching("test@gmail.com")
	expectedErrorText = "No email accounts match"
	if !strings.HasPrefix(err.Error(), expectedErrorText) {
		t.Error("Expected error prefix: ", expectedErrorText, "; Got: ", err)
	}

	// delete
	deleteResponse, err := cioLite.DeleteConnectToken(connectToken.Token)

	if err != nil || !deleteResponse.Success {
		t.Error("Expected successful delete of connect token; Got: ", deleteResponse, "; With Error: ", err, "; With Log: ", logger.String())
	}

	if len(logger.String()) < 20 {
		t.Error("Expected some output from logger; Got: ", logger.String())
	}
}

// TestSimulatedGetConnectToken tests GetConnectToken with a simulated server
func TestSimulatedGetConnectToken(t *testing.T) {
	t.Parallel()

	tokenString := "axjogv7yipqnhj9c"
	responseString := `
{
  "token": "axjogv7yipqnhj9c",
  "email": "test@gmail.com",
  "account_lite": true,
  "created": 1462217246,
  "used": 1462217259,
  "status_callback_url": "https://yoursite.com/api/unsubscriber/v1/cio/account_status/callback",
  "callback_url": "https://yoursite.com",
  "user": {
    "id": "5727aa2a0be9af5d658b4568",
    "email_accounts": [
      {
        "status": "OK",
        "resource_url": "https://api.context.io/lite/users/5727aa2a0be9af5d658b4568/email_accounts/test%3A%3Agmail",
        "type": "imap",
        "authentication_type": "oauth2",
        "use_ssl": true,
        "server": "imap.googlemail.com",
        "label": "test::gmail",
        "username": "test@gmail.com",
        "port": 993
      }
    ],
    "email_addresses": [
      "test@gmail.com"
    ],
    "created": 1462217251,
    "first_name": "test@gmail.com",
    "last_name": "",
    "resource_url": "https://api.context.io/lite/users/5727aa2a0be9af5d658b4568"
  },
  "email_account_id": "test::gmail",
  "resource_url": "https://api.context.io/lite/connect_tokens/axjogv7yipqnhj9c",
  "expires": false
}`

	cioLite, logger, testServer, mux := NewTestCioLiteWithLoggerAndTestServer(t)
	defer testServer.Close()

	mux.HandleFunc("/connect_tokens/axjogv7yipqnhj9c", func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, responseString)
		Must(err)
	})

	expected := GetConnectTokenResponse{
		Token:             "axjogv7yipqnhj9c",
		Email:             "test@gmail.com",
		EmailAccountID:    "test::gmail",
		CallbackURL:       "https://yoursite.com",
		StatusCallbackURL: "https://yoursite.com/api/unsubscriber/v1/cio/account_status/callback",
		ResourceURL:       "https://api.context.io/lite/connect_tokens/axjogv7yipqnhj9c",
		// FirstName:      "test@gmail.com",
		// LastName:       "",
		// BrowserRedirectURL:  "", Not included yet for some reason...
		// ServerLabel:         "", Not included yet for some reason...
		AccountLite: true,
		Created:     1462217246,
		Used:        1462217259,
		Expires:     ExpiresMixed{Expires: nil},
		User: GetConnectTokenUserResponse{
			ID:             "5727aa2a0be9af5d658b4568",
			EmailAddresses: []string{"test@gmail.com"},
			FirstName:      "test@gmail.com",
			LastName:       "",
			Created:        1462217251,
			EmailAccounts: []GetUsersEmailAccountsResponse{{
				Status:             "OK",
				ResourceURL:        "https://api.context.io/lite/users/5727aa2a0be9af5d658b4568/email_accounts/test%3A%3Agmail",
				Type:               "imap",
				AuthenticationType: "oauth2",
				Server:             "imap.googlemail.com",
				Label:              "test::gmail",
				Username:           "test@gmail.com",
				UseSSL:             true,
				Port:               993,
			}},
		},
	}

	// Get Connect Token
	getConnectToken, err := cioLite.GetConnectToken(tokenString)

	if err != nil || !reflect.DeepEqual(getConnectToken, expected) {
		t.Error("Expected: ", expected, "; Got: ", getConnectToken, "; With Error: ", err, "; With Log: ", logger.String())
	}

	// Check Connect Token
	err = cioLite.CheckConnectToken(getConnectToken, "test@gmail.com")
	if err != nil {
		t.Error("Expected successful check connect token; Got: ", err)
	}

	// Check if accepted token's email accounts match (they shouldn't since its not accepted)
	emailAccount, err := getConnectToken.User.EmailAccountMatching("test@gmail.com")
	if err != nil || !reflect.DeepEqual(emailAccount, expected.User.EmailAccounts[0]) {
		t.Error("Expected: ", expected.User.EmailAccounts[0], "; Got: ", emailAccount, "; With Error: ", err, "; With Log: ", logger.String())
	}

	if len(logger.String()) < 20 {
		t.Error("Expected some output from logger; Got: ", logger.String())
	}
}

// TestExpiresMixed tests the mixed type json encoding/decoding of the "expires" field.
func TestExpiresMixed(t *testing.T) {
	t.Parallel()

	type ExpiresTest struct {
		MyField ExpiresMixed `json:"my_field"`
	}

	falseExpectedJSON := `{"my_field":false}`
	falseExpectedExpires := ExpiresTest{MyField: ExpiresMixed{}}

	timestamp := 1234509876
	timestampExpectedJSON := `{"my_field":1234509876}`
	timestampExpectedExpires := ExpiresTest{MyField: ExpiresMixed{Expires: &timestamp}}

	// Test FALSE
	var falseExpires ExpiresTest
	Must(json.Unmarshal([]byte(falseExpectedJSON), &falseExpires))

	if !reflect.DeepEqual(falseExpires, falseExpectedExpires) ||
		falseExpires.MyField.Unused() ||
		falseExpires.MyField.Timestamp() != -1 {
		t.Error(falseExpires)
	}

	falseJSON, err := json.Marshal(falseExpires)
	Must(err)

	if string(falseJSON) != falseExpectedJSON {
		t.Error(string(falseJSON))
	}

	// Test timestamp
	var timestampExpires ExpiresTest
	Must(json.Unmarshal([]byte(timestampExpectedJSON), &timestampExpires))

	if !reflect.DeepEqual(timestampExpires, timestampExpectedExpires) ||
		!timestampExpires.MyField.Unused() ||
		timestampExpires.MyField.Timestamp() != timestamp {
		t.Error(timestampExpires)
	}

	timestampJSON, err := json.Marshal(timestampExpires)
	Must(err)

	if string(timestampJSON) != timestampExpectedJSON {
		t.Error(string(timestampJSON))
	}
}
