package ciolite

import (
	"io"
	"net/http"
	"reflect"
	"testing"
)

// TestSimulatedGetOauthProviders tests GetOauthProviders with a simulated server
func TestSimulatedGetOauthProviders(t *testing.T) {
	t.Parallel()

	cioLite, logger, testServer, mux := NewTestCioLiteWithLoggerAndTestServer(t)
	defer testServer.Close()

	mux.HandleFunc("/lite/oauth_providers", func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, `[ { "type": "GMAIL_OAUTH2",
      "provider_consumer_key": "123-abc.xzy.com",
      "provider_consumer_secret": "A2i[...]M9k",
      "resource_url": "https://api.context.io/lite/oauth_providers/123-abc.xzy.com" } ]`)
		Must(err)
	})

	expected := GetOAuthProvidersResponse{
		Type:                   "GMAIL_OAUTH2",
		ProviderConsumerKey:    "123-abc.xzy.com",
		ProviderConsumerSecret: "A2i[...]M9k",
		ResourceURL:            "https://api.context.io/lite/oauth_providers/123-abc.xzy.com",
	}

	oauthProviders, err := cioLite.GetOAuthProviders()

	if err != nil || !reflect.DeepEqual(oauthProviders, []GetOAuthProvidersResponse{expected}) {
		t.Error("Expected: ", []GetOAuthProvidersResponse{expected}, "; Got: ", oauthProviders, "; With Error: ", err, "; With Log: ", logger.String())
	}

	if len(logger.String()) < 20 {
		t.Error("Expected some output from logger; Got: ", logger.String())
	}
}

// TestSimulatedGetOauthProvider tests GetOauthProvider with a simulated server
func TestSimulatedGetOauthProvider(t *testing.T) {
	t.Parallel()

	cioLite, logger, testServer, mux := NewTestCioLiteWithLoggerAndTestServer(t)
	defer testServer.Close()

	mux.HandleFunc("/lite/oauth_providers/123-abc.xzy.com", func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, `{ "type": "GMAIL_OAUTH2",
      "provider_consumer_key": "123-abc.xzy.com",
      "provider_consumer_secret": "A2i[...]M9k",
      "resource_url": "https://api.context.io/lite/oauth_providers/123-abc.xzy.com" }`)
		Must(err)
	})

	expected := GetOAuthProvidersResponse{
		Type:                   "GMAIL_OAUTH2",
		ProviderConsumerKey:    "123-abc.xzy.com",
		ProviderConsumerSecret: "A2i[...]M9k",
		ResourceURL:            "https://api.context.io/lite/oauth_providers/123-abc.xzy.com",
	}

	oauthProvider, err := cioLite.GetOAuthProvider("123-abc.xzy.com")

	if err != nil || !reflect.DeepEqual(oauthProvider, expected) {
		t.Error("Expected: ", expected, "; Got: ", oauthProvider, "; With Error: ", err, "; With Log: ", logger.String())
	}

	if len(logger.String()) < 20 {
		t.Error("Expected some output from logger; Got: ", logger.String())
	}
}
