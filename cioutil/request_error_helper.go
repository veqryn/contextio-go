package cioutil

const (
	UnknownStatusCode = -1
	UnknownPayload    = "UNKNOWN"
	UnknownMethod     = "UNKNOWN"
	UnknownURL        = "UNKNOWN"
)

// ErrorStatusCode returns the StatusCode of the error, or 0
func ErrorStatusCode(err error) int {
	if err == nil {
		return 0
	}
	type ErrorStatusCoder interface {
		ErrorStatusCode() int
	}
	if e, ok := err.(ErrorStatusCoder); ok {
		return e.ErrorStatusCode()
	}
	return UnknownStatusCode
}

// ErrorPayload returns the payload of the error, or an empty string
func ErrorPayload(err error) string {
	if err == nil {
		return ""
	}
	type ErrorPayloader interface {
		ErrorPayload() string
	}
	if e, ok := err.(ErrorPayloader); ok {
		return e.ErrorPayload()
	}
	return UnknownPayload
}

// ErrorMethod returns the method of the error, or an empty string
func ErrorMethod(err error) string {
	if err == nil {
		return ""
	}
	type ErrorMethoder interface {
		ErrorMethod() string
	}
	if e, ok := err.(ErrorMethoder); ok {
		return e.ErrorMethod()
	}
	return UnknownMethod
}

// ErrorURL returns the URL of the error, or an empty string
func ErrorURL(err error) string {
	if err == nil {
		return ""
	}
	type ErrorURLer interface {
		ErrorURL() string
	}
	if e, ok := err.(ErrorURLer); ok {
		return e.ErrorURL()
	}
	return UnknownURL
}
