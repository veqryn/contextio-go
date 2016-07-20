package cioutil

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

// RequestError is the error type returned by DoFormRequest and all cio api calls
type RequestError struct {
	error
	StatusCode int
	Payload    string
	Method     string
	URL        string
}

// String returns the same as Error()
func (e RequestError) String() string {
	return e.Error()
}

// Format prints out the error, any causes, a stacktrace, and the other fields in the struct
func (e RequestError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = fmt.Fprintf(s, "%+v\n", e.error)
			type temp struct {
				StatusCode int
				Payload    string
				Method     string
				URL        string
			}
			_, _ = fmt.Fprintf(s, "%+v", temp{StatusCode: e.StatusCode, Payload: e.Payload, Method: e.Method, URL: e.URL})
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, e.Error())
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", e.Error())
	}
}

// Cause returns the cause of any wrapped errors, or just the base error if no wrapped error
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
	return 0
}

// ErrorPayload returns the payload of the error, or an empty string
func ErrorPayload(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(RequestError); ok {
		return e.Payload
	}
	return ""
}

// ErrorMethod returns the method of the error, or an empty string
func ErrorMethod(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(RequestError); ok {
		return e.Method
	}
	return ""
}

// ErrorURL returns the URL of the error, or an empty string
func ErrorURL(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(RequestError); ok {
		return e.URL
	}
	return ""
}
