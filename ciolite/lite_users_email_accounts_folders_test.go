package ciolite

import (
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

// TestSimulatedGetUserEmailAccountFolder tests GetUserEmailAccountFolder with a simulated server
func TestSimulatedGetUserEmailAccountFolder(t *testing.T) {
	t.Parallel()

	cioLite, logger, testServer, mux := NewTestCioLiteWithLoggerAndTestServer(t)
	defer testServer.Close()

	folderName := "ParentFolder/SubFolder With Spaces"
	expectedEncoding := "ParentFolder%2FSubFolder%20With%20Spaces"
	expectedResourceURL := "/users/123abc/email_accounts/0/folders/" + expectedEncoding

	cioLite.PreRequestHook = func(cioUserID string, cioLabel string, method string, requestURL string, bodyValues url.Values) {
		if !strings.HasSuffix(requestURL, expectedResourceURL) {
			t.Error("Expected requestURL suffix: ", expectedResourceURL, "; Got: ", requestURL)
		}
	}

	mux.HandleFunc("/users/123abc/email_accounts/0/folders/"+folderName, func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, `{
		 "name": "`+folderName+`",
		 "nb_messages": 5,
		 "nb_unseen_messages": 2,
		 "delimiter": "/",
		 "resource_url": "https://api.context.io/lite`+expectedResourceURL+`"
		 }`)
		Must(err)
	})

	expectedResponse := GetUsersEmailAccountFoldersResponse{
		Name:             folderName,
		NbMessages:       5,
		NbUnseenMessages: 2,
		Delimiter:        "/",
		ResourceURL:      "https://api.context.io/lite" + expectedResourceURL,
	}

	getFolderResp, err := cioLite.GetUserEmailAccountFolder("123abc", "0", folderName, EmailAccountFolderDelimiterParam{})

	if err != nil || !reflect.DeepEqual(getFolderResp, expectedResponse) {
		t.Error("Expected: ", expectedResponse, "; Got: ", getFolderResp, "; With Error: ", err, "; With Log: ", logger.String())
	}

	if len(logger.String()) < 20 {
		t.Error("Expected some output from logger; Got: ", logger.String())
	}
}
