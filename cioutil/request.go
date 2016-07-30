package cioutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
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
	logRequest(cio.Log, request.Method, cioURL, bodyValues)

	statusCode, resBody, err := cio.createAndSendRequest(request, cioURL, bodyString, bodyValues, result)

	// Retry if Status Code >= 500 and RetryServerErr is set to true
	if cio.RetryServerErr && shouldRetryOnce(statusCode, err) {
		time.Sleep(1 * time.Second)
		logResponse(cio.Log, true, request.Method, cioURL, statusCode, resBody, errors.Cause(err))
		statusCode, resBody, err = cio.createAndSendRequest(request, cioURL, bodyString, bodyValues, result)
	}

	// Log the response
	logResponse(cio.Log, false, request.Method, cioURL, statusCode, resBody, errors.Cause(err))

	return err
}

// shouldRetryOnce returns true if the request should be retried
func shouldRetryOnce(statusCode int, err error) bool {
	// Retry if a connection can not be made (network blip), and also on CIO Server errors,
	// and also if the nonce has been used (CIO seems to have issues with nonce collisions).
	return statusCode >= 500 || statusCode == 0 ||
		(statusCode == 401 && strings.Contains(strings.ToLower(ErrorPayload(err)), "nonce"))
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
		if closeErr := res.Body.Close(); closeErr != nil {
			logBodyCloseError(cio.Log, closeErr)
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
			log.Printf("Creating new %s request to: %s with payload: %s\n", method, cioURL, redactedValues.Encode())
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
			log.Printf("Unable to close response body from CIO, with error: %s\n", closeError.Error())
		}
	}
}

// logResponse logs the response from CIO, if any logger is set
func logResponse(log Logger, retry bool, method string, cioURL string, statusCode int, responseBody string, err error) {
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
			if !retry && (err != nil || statusCode >= 400) {
				if err != nil {
					logEntry = logEntry.WithError(err)
				}
				logEntry.Warn("Received response from CIO")
			} else {
				logEntry.Debug("Received response from CIO")
			}

		} else {
			// Else just log with Println
			if err != nil {
				log.Printf("Received response from %s to: %s with status code: %d and error: %s and payload snippet: %s\n", method, cioURL, statusCode, err, responseBody)
			} else {
				log.Printf("Received response from %s to: %s with status code: %d and payload snippet: %s\n", method, cioURL, statusCode, responseBody)
			}
		}
	}
}
