package cioutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/Sirupsen/logrus"
	"github.com/garyburd/go-oauth/oauth"
	"github.com/pkg/errors"
)

// ClientRequest defines information that can be used to make a request
type ClientRequest struct {
	Method      string
	Path        string
	FormValues  interface{}
	QueryValues interface{}
}

// DoFormRequest makes the actual request
func (cio Cio) DoFormRequest(request ClientRequest, result interface{}) error {

	// Construct the url
	cioURL := cio.Host + request.Path + QueryString(request.QueryValues)

	// Construct the body
	var bodyReader io.Reader
	bodyValues := FormValues(request.FormValues)
	bodyString := bodyValues.Encode()
	if len(bodyString) > 0 {
		bodyReader = bytes.NewReader([]byte(bodyString))
	}
	logRequest(cio.Log, request.Method, cioURL, bodyValues)

	// Construct the request
	httpReq, err := cio.createRequest(request, cioURL, bodyReader, bodyValues)
	if err != nil {
		return err
	}

	// Send the request
	return cio.sendRequest(httpReq, result, cioURL)
}

// createRequest creates the *http.Request object
func (cio Cio) createRequest(request ClientRequest, cioURL string, bodyReader io.Reader, bodyValues url.Values) (*http.Request, error) {
	// Construct the request
	httpReq, err := http.NewRequest(request.Method, cioURL, bodyReader)
	if err != nil {
		return httpReq, RequestError{error: errors.Wrap(err, "CIO: Failed to form request"), Method: request.Method, URL: cioURL}
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

// sendRequest sends the *http.Request
func (cio Cio) sendRequest(httpReq *http.Request, result interface{}, cioURL string) error {
	// Create the HTTP client
	httpClient := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   cio.RequestTimeout,
	}

	// Make the request
	res, err := httpClient.Do(httpReq)
	if err != nil {
		return RequestError{error: errors.Wrap(err, "CIO: Failed to make request"), Method: httpReq.Method, URL: cioURL}
	}

	// Parse the response
	defer func() {
		if closeErr := res.Body.Close(); closeErr != nil {
			logBodyCloseError(cio.Log, closeErr)
		}
	}()

	resBody, err := ioutil.ReadAll(res.Body)
	resBodyString := string(resBody)
	if err != nil {
		return RequestError{error: errors.Wrap(err, "CIO: Could not read response"), Method: httpReq.Method, URL: cioURL, StatusCode: res.StatusCode, Payload: resBodyString}
	}

	// Unmarshal result
	err = json.Unmarshal(resBody, &result)

	// Log the response
	logResponse(cio.Log, httpReq.Method, cioURL, res.StatusCode, resBodyString, err)

	// Return own error if Status Code >= 400
	if res.StatusCode >= 400 {
		return RequestError{error: errors.New("CIO: Status Code >= 400"), Method: httpReq.Method, URL: cioURL, StatusCode: res.StatusCode, Payload: resBodyString}
	}

	// Return Unmarshal error (if any) if Status Code is < 400
	if err != nil {
		return RequestError{error: errors.Wrap(err, "CIO: Could not unmarshal payload"), Method: httpReq.Method, URL: cioURL, StatusCode: res.StatusCode, Payload: resBodyString}
	}
	return nil
}

// logRequest logs the request about to be made to CIO, redacting sensitive information in the body
func logRequest(log Logger, method string, cioURL string, bodyValues url.Values) {
	if log != nil {

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

		// Actually log
		if logrusLogger, ok := log.(*logrus.Logger); ok {
			// If logrus, use structured fields
			logrusLogger.WithFields(logrus.Fields{"httpMethod": method, "url": cioURL, "payload": redactedValues.Encode()}).Debug("Creating new request to CIO")
		} else {
			// Else just log with Println
			log.Println("Creating new " + method + " request to: " + cioURL + " with payload: " + redactedValues.Encode())
		}
	}
}

// logBodyCloseError logs any error that happens when trying to close the *http.Response.Body
func logBodyCloseError(log Logger, closeError error) {
	if log != nil {
		if logrusLogger, ok := log.(*logrus.Logger); ok {
			// If logrus, use structured fields
			logrusLogger.WithError(closeError).Warn("Unable to close response body from CIO")
		} else {
			// Else just log with Println
			log.Println("Unable to close response body from CIO, with error: " + closeError.Error())
		}
	}
}

// logResponse logs the response from CIO, if any logger is set
func logResponse(log Logger, method string, cioURL string, statusCode int, responseBody string, unmarshalError error) {
	if log != nil {

		// TODO: redact access_token and access_token_secret before logging (only occurs with 3-legged oauth [rare])

		// Take only the first 2000 characters from the responseBody, which should be more than enough to debug anything
		if bodyLen := len(responseBody); bodyLen > 2000 {
			responseBody = responseBody[:2000]
		}

		if logrusLogger, ok := log.(*logrus.Logger); ok {
			// If logrus, use structured fields
			logEntry := logrusLogger.WithFields(logrus.Fields{
				"httpMethod":     method,
				"url":            cioURL,
				"statusCode":     fmt.Sprintf("%d", statusCode),
				"payloadSnippet": responseBody})
			if unmarshalError != nil || statusCode >= 400 {
				logEntry.Warn("Received response from CIO")
			} else {
				logEntry.Debug("Received response from CIO")
			}

		} else {
			// Else just log with Println
			log.Println("Received response from " + method + " to: " + cioURL +
				" with status code: " + fmt.Sprintf("%d", statusCode) +
				" and payload snippet: " + responseBody)
		}
	}
}
