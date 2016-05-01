package ciolite

import (
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
	_, err = cioLite.CheckConnectToken("not_correct@gmail.com", connectToken.Token)
	expectedErrorText := "Email does not match Context.io token"
	if err.Error() != expectedErrorText {
		t.Error("Expected error: ", expectedErrorText, "; Got: ", err)
	}

	// check with good email
	_, err = cioLite.CheckConnectToken("test@gmail.com", connectToken.Token)
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
