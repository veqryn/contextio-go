package cioutil

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/garyburd/go-oauth/oauth"
	"github.com/pkg/errors"
)

// ClientRequest defines information that can be used to make a request
type ClientRequest struct {
	Method       string
	Path         string
	FormValues   interface{}
	QueryValues  interface{}
	UserID       string
	AccountLabel string
}

// DoFormRequest makes the actual request
func (cio Cio) DoFormRequest(request ClientRequest, result interface{}) error {

	// Construct the url
	cioURL := cio.Host + request.Path + QueryString(request.QueryValues)

	// Construct the body
	bodyValues := FormValues(request.FormValues)
	bodyString := bodyValues.Encode()

	// Before-Request Hook Function (logging)
	if cio.PreRequestHook != nil {
		cio.PreRequestHook(request.UserID, request.AccountLabel, request.Method, cioURL, redactBodyValues(bodyValues))
	}

	var (
		statusCode int
		resBody    string
		err        error
	)

	for i := 1; i <= 10; i++ {
		statusCode, resBody, err = cio.createAndSendRequest(request, cioURL, bodyString, bodyValues, result)
		// After-Request Hook Function (logging)
		if cio.PostRequestShouldRetryHook == nil || !cio.PostRequestShouldRetryHook(i, request.UserID, request.AccountLabel, request.Method, cioURL, statusCode, resBody, err) {
			break
		}
		time.Sleep(time.Duration(i*i*i) * time.Second) // Exponential backoff
	}

	return err
}

// createAndSendRequest creates the body io.Reader, the *http.Request, and sends the request, logging the response.
// Returns the status code, the response body, and any error
func (cio Cio) createAndSendRequest(request ClientRequest, cioURL string, bodyString string, bodyValues url.Values, result interface{}) (int, string, error) {

	var bodyReader io.Reader
	if len(bodyString) > 0 {
		bodyReader = bytes.NewReader([]byte(bodyString))
	}

	// Construct the request
	httpReq, err := cio.createRequest(request, cioURL, bodyReader, bodyValues)
	if err != nil {
		return 0, "", err
	}

	// Send the request
	return cio.sendRequest(httpReq, result, cioURL)
}

// createRequest creates the *http.Request object
func (cio Cio) createRequest(request ClientRequest, cioURL string, bodyReader io.Reader, bodyValues url.Values) (*http.Request, error) {
	// Construct the request
	httpReq, err := http.NewRequest(request.Method, cioURL, bodyReader)
	if err != nil {
		return httpReq, RequestError{errors.Wrap(err, "CIO: Failed to form request"), ErrorMetaData{Method: request.Method, URL: cioURL}}
	}

	// oAuth signature
	var client oauth.Client
	client.Credentials = oauth.Credentials{Token: cio.apiKey, Secret: cio.apiSecret}

	// Add headers
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Accept-Charset", "utf-8")
	httpReq.Header.Set("User-Agent", "Golang CIO Library")
	httpReq.Header.Set("Authorization", client.AuthorizationHeader(nil, request.Method, httpReq.URL, bodyValues))

	return httpReq, nil
}

// sendRequest sends the *http.Request, and returns the status code, the response body, and any error
func (cio Cio) sendRequest(httpReq *http.Request, result interface{}, cioURL string) (int, string, error) {
	// Create the HTTP client
	httpClient := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   cio.RequestTimeout,
	}

	// Make the request
	res, err := httpClient.Do(httpReq)
	if err != nil {
		return 0, "", RequestError{errors.Wrap(err, "CIO: Failed to make request"), ErrorMetaData{Method: httpReq.Method, URL: cioURL}}
	}

	// Parse the response
	defer func() {
		if closeErr := res.Body.Close(); closeErr != nil && cio.ResponseBodyCloseErrorHook != nil {
			cio.ResponseBodyCloseErrorHook(closeErr) // Logging
		}
	}()

	resBody, err := ioutil.ReadAll(res.Body)
	resBodyString := string(resBody)
	if err != nil {
		return res.StatusCode, resBodyString, RequestError{errors.Wrap(err, "CIO: Could not read response"), ErrorMetaData{Method: httpReq.Method, URL: cioURL, StatusCode: res.StatusCode, Payload: resBodyString}}
	}

	// Unmarshal result
	err = json.Unmarshal(resBody, &result)

	// Return own error if Status Code >= 400
	if res.StatusCode >= 400 {
		return res.StatusCode, resBodyString, RequestError{errors.New("CIO: Status Code >= 400"), ErrorMetaData{Method: httpReq.Method, URL: cioURL, StatusCode: res.StatusCode, Payload: resBodyString}}
	}

	// Return Unmarshal error (if any) if Status Code is < 400
	if err != nil {
		return res.StatusCode, resBodyString, RequestError{errors.Wrap(err, "CIO: Could not unmarshal payload"), ErrorMetaData{Method: httpReq.Method, URL: cioURL, StatusCode: res.StatusCode, Payload: resBodyString}}
	}
	return res.StatusCode, resBodyString, nil
}

// redactBodyValues returns a copy of the body values redacted
func redactBodyValues(bodyValues url.Values) url.Values {

	// Copy url.Values
	redactedValues := url.Values{}
	for k, v := range bodyValues {
		redactedValues[k] = v
	}

	// Redact sensitive information
	if val := redactedValues.Get("password"); len(val) > 0 {
		redactedValues.Set("password", "redacted")
	}
	if val := redactedValues.Get("provider_refresh_token"); len(val) > 0 {
		redactedValues.Set("provider_refresh_token", "redacted")
	}
	if val := redactedValues.Get("provider_consumer_key"); len(val) > 0 {
		redactedValues.Set("provider_consumer_key", "redacted")
	}
	if val := redactedValues.Get("provider_consumer_secret"); len(val) > 0 {
		redactedValues.Set("provider_consumer_secret", "redacted")
	}

	return redactedValues
}
