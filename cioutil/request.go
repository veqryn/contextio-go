package cioutil

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/Sirupsen/logrus"
	"github.com/garyburd/go-oauth/oauth"
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
		return httpReq, fmt.Errorf("Could not create request: %s", err)
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
		return fmt.Errorf("Failed to make request: %s", err)
	}

	// Parse the response
	defer func() {
		if closeErr := res.Body.Close(); closeErr != nil {
			logBodyCloseError(cio.Log, closeErr)
		}
	}()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Could not read response: %s", err)
	}
	resBodyString := string(resBody)

	// Unmarshal result
	err = json.Unmarshal(resBody, &result)

	// Log the response
	logResponse(cio.Log, httpReq.Method, cioURL, res.StatusCode, resBodyString, err)

	// Return special error if Status Code >= 400
	if res.StatusCode >= 400 {
		return fmt.Errorf("%d Status Code with Payload %s", res.StatusCode, resBodyString)
	}

	// Return Unmarshal error (if any) if Status Code is < 400
	return err
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
			redactedValues.Set("password", fmt.Sprintf("%x", md5.Sum([]byte(val))))
		}
		if val := redactedValues.Get("provider_refresh_token"); len(val) > 0 {
			redactedValues.Set("provider_refresh_token", fmt.Sprintf("%x", md5.Sum([]byte(val))))
		}
		if val := redactedValues.Get("provider_consumer_key"); len(val) > 0 {
			redactedValues.Set("provider_consumer_key", fmt.Sprintf("%x", md5.Sum([]byte(val))))
		}
		if val := redactedValues.Get("provider_consumer_secret"); len(val) > 0 {
			redactedValues.Set("provider_consumer_secret", fmt.Sprintf("%x", md5.Sum([]byte(val))))
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

		if logrusLogger, ok := log.(*logrus.Logger); ok {
			// If logrus, use structured fields
			logEntry := logrusLogger.WithFields(logrus.Fields{
				"httpMethod": method,
				"url":        cioURL,
				"statusCode": fmt.Sprintf("%d", statusCode),
				"payload":    responseBody})
			if unmarshalError != nil || statusCode >= 400 {
				logEntry.Warn("Received response from CIO")
			} else {
				logEntry.Debug("Received response from CIO")
			}

		} else {
			// Else just log with Println
			log.Println("Received response from " + method + " to: " + cioURL +
				" with status code: " + fmt.Sprintf("%d", statusCode) +
				" and payload: " + responseBody)
		}
	}
}
