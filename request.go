// Package ciolite ...
package ciolite

// Imports
import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/garyburd/go-oauth/oauth"
	"net/http"
	"time"
)

// clientRequest ...
// Defines information that can be used to make a request to Medium.
type clientRequest struct {
	method      string
	path        string
	formValues  CioParams
	queryValues CioParams
	format      string
}

const (
	// The default host of Medium's API
	host = "https://api.context.io/lite"

	// The default timeout duration used on HTTP requests
	defaultTimeout = 10 * time.Second
)

// doFormRequest ...
// Makes the actual request
func (cioLite *CioLite) doFormRequest(request clientRequest, result interface{}) error {

	// Construct the url
	url := host + request.path + request.queryValues.QueryString()

	// Construct the body
	var bodyReader *bytes.Reader
	if request.formValues.FormValues() != nil {
		bodyBytes := []byte(request.formValues.FormValues().Encode())
		bodyReader = bytes.NewReader(bodyBytes)
	} else {
		bodyReader = bytes.NewReader(make([]byte, 0, 0))
	}

	// Construct the request
	httpReq, err := http.NewRequest(request.method, url, bodyReader)
	if err != nil {
		return fmt.Errorf("Could not create request: %s", err)
	}

	// oAuth signature
	var client oauth.Client
	client.Credentials = oauth.Credentials{cioLite.apiKey, cioLite.apiSecret}

	// Add headers
	httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpReq.Header.Add("Accept", "application/json")
	httpReq.Header.Add("Accept-Charset", "utf-8")
	httpReq.Header.Add("Authorization", client.AuthorizationHeader(nil, request.method, httpReq.URL, request.formValues.FormValues()))

	// Create the HTTP client
	httpClient := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   defaultTimeout,
	}

	// Make the request
	res, err := httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("Failed to make request: %s", err)
	}

	// Determine status
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Invalid status code: %d", res.StatusCode)
	}

	// Parse the response
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(&result)
}
