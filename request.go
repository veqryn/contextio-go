package ciolite

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/garyburd/go-oauth/oauth"
)

// clientRequest defines information that can be used to make a request
type clientRequest struct {
	method      string
	path        string
	formValues  interface{}
	queryValues interface{}
}

// doFormRequest makes the actual request
func (cioLite CioLite) doFormRequest(request clientRequest, result interface{}) error {

	// Construct the url
	url := cioLite.Host + request.path + QueryString(request.queryValues)

	// Construct the body
	var bodyReader io.Reader
	bodyValues := FormValues(request.formValues)
	bodyString := bodyValues.Encode()
	if len(bodyString) > 0 {
		bodyReader = bytes.NewReader([]byte(bodyString))
	}
	if cioLite.log != nil {
		cioLite.log.Println("Creating new request to: " + url + " with payload: " + bodyString)
	}

	// Construct the request
	httpReq, err := http.NewRequest(request.method, url, bodyReader)
	if err != nil {
		return fmt.Errorf("Could not create request: %s", err)
	}

	// oAuth signature
	var client oauth.Client
	client.Credentials = oauth.Credentials{Token: cioLite.apiKey, Secret: cioLite.apiSecret}

	// Add headers
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Accept-Charset", "utf-8")
	httpReq.Header.Set("User-Agent", "Golang CIOLite library v0.1")
	httpReq.Header.Set("Authorization", client.AuthorizationHeader(nil, request.method, httpReq.URL, bodyValues))

	// Create the HTTP client
	httpClient := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   cioLite.RequestTimeout,
	}

	// Make the request
	res, err := httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("Failed to make request: %s", err)
	}

	// Parse the response
	defer func() {
		if closeErr := res.Body.Close(); closeErr != nil {
			cioLite.log.Println("Unable to close response body, with error: " + closeErr.Error())
		}
	}()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Could not read response: %s", err)
	}
	if cioLite.log != nil {
		cioLite.log.Println("Received response from: " + url +
			" with status code: " + fmt.Sprintf("%d", res.StatusCode) +
			" and payload: " + string(resBody))
	}

	// Determine status
	if res.StatusCode >= 400 {
		return fmt.Errorf("Invalid status code: %d", res.StatusCode)
	}

	// Unmarshal result
	return json.Unmarshal(resBody, &result)
}
