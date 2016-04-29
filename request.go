package ciolite

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"crypto/md5"

	"github.com/Sirupsen/logrus"
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
	cioURL := cioLite.Host + request.path + QueryString(request.queryValues)

	// Construct the body
	var bodyReader io.Reader
	bodyValues := FormValues(request.formValues)
	bodyString := bodyValues.Encode()
	if len(bodyString) > 0 {
		bodyReader = bytes.NewReader([]byte(bodyString))
	}
	logRequest(cioLite.Log, cioURL, bodyValues)

	// Construct the request
	httpReq, err := cioLite.createRequest(request, cioURL, bodyReader, bodyValues)
	if err != nil {
		return err
	}

	// Send the request
	return cioLite.sendRequest(httpReq, result, cioURL)
}

// createRequest creates the *http.Request object
func (cioLite CioLite) createRequest(request clientRequest, cioURL string, bodyReader io.Reader, bodyValues url.Values) (*http.Request, error) {
	// Construct the request
	httpReq, err := http.NewRequest(request.method, cioURL, bodyReader)
	if err != nil {
		return httpReq, fmt.Errorf("Could not create request: %s", err)
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

	return httpReq, nil
}

// sendRequest sends the *http.Request
func (cioLite CioLite) sendRequest(httpReq *http.Request, result interface{}, cioURL string) error {
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
			logBodyCloseError(cioLite.Log, closeErr)
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
	logResponse(cioLite.Log, cioURL, res.StatusCode, resBodyString, err)

	// Return special error if Status Code >= 400
	if res.StatusCode >= 400 {
		return fmt.Errorf("%d Status Code with Payload %s", res.StatusCode, resBodyString)
	}

	// Return Unmarshal error (if any) if Status Code is < 400
	return err
}

// logRequest logs the request about to be made to CIO, redacting sensitive information in the body
func logRequest(log Logger, cioURL string, bodyValues url.Values) {
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
			logrusLogger.WithFields(logrus.Fields{"url": cioURL, "payload": redactedValues.Encode()}).Debug("Creating new request to CIO")
		} else {
			// Else just log with Println
			log.Println("Creating new request to: " + cioURL + " with payload: " + redactedValues.Encode())
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
func logResponse(log Logger, cioURL string, statusCode int, responseBody string, unmarshalError error) {
	if log != nil {

		// TODO: redact access_token and access_token_secret before logging (only occurs with 3-legged oauth [rare])

		if logrusLogger, ok := log.(*logrus.Logger); ok {
			// If logrus, use structured fields
			logEntry := logrusLogger.WithFields(logrus.Fields{
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
			log.Println("Received response from: " + cioURL +
				" with status code: " + fmt.Sprintf("%d", statusCode) +
				" and payload: " + responseBody)
		}
	}
}
