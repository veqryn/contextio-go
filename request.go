package ciolite

// Imports
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/garyburd/go-oauth/oauth"
)

// clientRequest defines information that can be used to make a request
type clientRequest struct {
	method      string
	path        string
	formValues  CioParams
	queryValues CioParams
}

const (
	// The default host of API
	host = "https://api.context.io/lite"

	// The default timeout duration used on HTTP requests
	defaultTimeout = 20 * time.Second
)

// doFormRequest makes the actual request
func (cioLite *CioLite) doFormRequest(request clientRequest, result interface{}) error {

	// Construct the url
	url := host + request.path + request.queryValues.QueryString()

	// Construct the body
	var bodyReader io.Reader
	bodyValues := request.formValues.FormValues()
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
	httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpReq.Header.Add("Accept", "application/json")
	httpReq.Header.Add("Accept-Charset", "utf-8")
	httpReq.Header.Add("Authorization", client.AuthorizationHeader(nil, request.method, httpReq.URL, bodyValues))

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
