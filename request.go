package ciolite

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
	if cioLite.Log != nil {
		if logrusLogger, ok := cioLite.Log.(*logrus.Logger); ok {
			logrusLogger.WithFields(logrus.Fields{"url": cioURL, "payload": bodyString}).Debug("Creating new request to CIO")
		} else {
			cioLite.Log.Println("Creating new request to: " + cioURL + " with payload: " + bodyString)
		}
	}

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
		if closeErr := res.Body.Close(); closeErr != nil && cioLite.Log != nil {
			if logrusLogger, ok := cioLite.Log.(*logrus.Logger); ok {
				logrusLogger.WithError(closeErr).Warn("Unable to close response body from CIO")
			} else {
				cioLite.Log.Println("Unable to close response body from CIO, with error: " + closeErr.Error())
			}
		}
	}()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Could not read response: %s", err)
	}
	if cioLite.Log != nil {
		if logrusLogger, ok := cioLite.Log.(*logrus.Logger); ok {
			logrusLogger.WithFields(logrus.Fields{
				"url":        cioURL,
				"statusCode": fmt.Sprintf("%d", res.StatusCode),
				"payload":    string(resBody),
			}).Debug("Received response from CIO")
		} else {
			cioLite.Log.Println("Received response from: " + cioURL +
				" with status code: " + fmt.Sprintf("%d", res.StatusCode) +
				" and payload: " + string(resBody))
		}
	}

	// Determine status
	if res.StatusCode >= 400 {
		return fmt.Errorf("Invalid status code: %d", res.StatusCode)
	}

	// Unmarshal result
	return json.Unmarshal(resBody, &result)
}
