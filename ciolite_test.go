package ciolite

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// TestNewCioLiteWithLogger tests the construction of CioLite
func TestNewCioLite(t *testing.T) {
	t.Parallel()
	NewTestCioLite(t)
}

// TestNewCioLiteWithLogger tests the construction of CioLite and *TestLogger objects
func TestNewCioLiteWithLogger(t *testing.T) {
	t.Parallel()
	NewTestCioLiteWithLogger(t)
}

// TestNewCioLiteWithLogger tests the construction of CioLite and *TestLogger objects
func TestNewTestCioLiteServer(t *testing.T) {
	t.Parallel()
	_, _, testServer, _ := NewTestCioLiteWithLoggerAndTestServer(t)
	defer testServer.Close()
}

// NewTestCioLite returns a new CioLite object
func NewTestCioLite(t *testing.T) CioLite {
	return NewCioLite(getEnv(t, "UNSUB_CIO_API_KEY"), getEnv(t, "UNSUB_CIO_API_SECRET"))
}

// NewTestCioLiteWithLogger returns a new CioLite object and *TestLogger object
func NewTestCioLiteWithLogger(t *testing.T) (CioLite, *TestLogger) {
	logger := &TestLogger{Buffer: &bytes.Buffer{}, t: t}
	cioLite := NewCioLiteWithLogger(getEnv(t, "UNSUB_CIO_API_KEY"), getEnv(t, "UNSUB_CIO_API_SECRET"), logger)
	return cioLite, logger
}

// NewTestCioLiteWithLoggerAndTestServer returns a new CioLite, *TestLogger, and *httptest.Server objects
func NewTestCioLiteWithLoggerAndTestServer(t *testing.T) (CioLite, *TestLogger, *httptest.Server, *http.ServeMux) {
	logger := &TestLogger{Buffer: &bytes.Buffer{}, t: t}
	mux := http.NewServeMux()
	cioLite, server := NewTestCioLiteServer(getEnv(t, "UNSUB_CIO_API_KEY"), getEnv(t, "UNSUB_CIO_API_SECRET"), logger, mux)
	return cioLite, logger, server, mux
}

// getEnv returns the named environment variable, or causes t.Fatal
func getEnv(t *testing.T, envVarName string) string {
	val := os.Getenv(envVarName)
	if len(val) == 0 {
		t.Fatal("Empty Environment Variable for: " + envVarName)
	}
	return val
}

// TestLogger is a *bytes.Buffer that implements the logging interface
type TestLogger struct {
	*bytes.Buffer
	t *testing.T
}

// Print prints the arguments to the buffer, using fmt.Sprint
func (l *TestLogger) Print(v ...interface{}) {
	_, err := l.Write([]byte(fmt.Sprint(v...)))
	if err != nil {
		l.t.Error("Error writing to logger: ", err)
	}
}

// Println prints the arguments to the buffer, using fmt.Sprint,appending a new line
func (l *TestLogger) Println(v ...interface{}) {
	_, err := l.Write([]byte(fmt.Sprint(v...) + "\n"))
	if err != nil {
		l.t.Error("Error writing to logger: ", err)
	}
}

// Printf prints the arguments to the buffer, using fmt.Sprintf
func (l *TestLogger) Printf(format string, v ...interface{}) {
	_, err := l.Write([]byte(fmt.Sprintf(format, v...)))
	if err != nil {
		l.t.Error("Error writing to logger: ", err)
	}
}

// Must panics if error is not nil
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
