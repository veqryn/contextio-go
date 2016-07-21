package cioutil

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

// RequestError is the error type returned by DoFormRequest and all cio api calls
type RequestError struct {
	error
	ErrorMetaData
}

// ErrorMetaData holds some meta-data about the error: StatusCode, Response Payload, Method used, and URL
type ErrorMetaData struct {
	StatusCode int
	Payload    string
	Method     string
	URL        string
}

const (
	UnknownStatusCode = -1
	UnknownPayload    = "UNKNOWN"
	UnknownMethod     = "UNKNOWN"
	UnknownURL        = "UNKNOWN"
)

// Cause returns the cause of any wrapped errors, or just the base error if no wrapped error.
// Can use with github.com/pkg/errors
func (e RequestError) Cause() error {
	return errors.Cause(e.error)
}

// ErrorStatusCode returns the StatusCode of the error, or 0
func ErrorStatusCode(err error) int {
	if err == nil {
		return 0
	}
	if e, ok := err.(RequestError); ok {
		return e.StatusCode
	}
	return UnknownStatusCode
}

// ErrorPayload returns the payload of the error, or an empty string
func ErrorPayload(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(RequestError); ok {
		return e.Payload
	}
	return UnknownPayload
}

// ErrorMethod returns the method of the error, or an empty string
func ErrorMethod(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(RequestError); ok {
		return e.Method
	}
	return UnknownMethod
}

// ErrorURL returns the URL of the error, or an empty string
func ErrorURL(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(RequestError); ok {
		return e.URL
	}
	return UnknownURL
}

// Format prints out the error, any causes, a stacktrace, and the other fields in the struct
func (e RequestError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = fmt.Fprintf(s, "%+v\n%+v", e.error, e.ErrorMetaData)
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, e.Error())
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", e.Error())
	}
}

// Error returns the Error string, any Wrapped Causes, and any StatusCode, Payload, Method, and URL that were set
func (e RequestError) Error() string {
	return fmt.Sprintf("%s; %+v", e.error, e.ErrorMetaData)
}

// String returns the same as Error()
func (e RequestError) String() string {
	return e.Error()
}
