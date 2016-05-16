package ciolite

import (
	"net/url"
	"reflect"
	"testing"
)

// TestFormValues tests that the form values function returns the correct url.Values
func TestFormValuesAndQueryString(t *testing.T) {
	t.Parallel()

	params := struct {
		StringFull  string `json:"string_full"`
		StringEmpty string `json:"string_empty"`
		BoolTrue    bool   `json:"bool_true"`
		BoolFalse   bool   `json:"bool_false"`
		IntLarge    int    `json:"int_large"`
		IntZero     int    `json:"int_zero"`
	}{
		StringFull:  "hello world",
		StringEmpty: "",
		BoolTrue:    true,
		BoolFalse:   false,
		IntLarge:    8194723,
		IntZero:     0,
	}

	expectedFormValues := url.Values{
		"string_full": []string{"hello world"},
		"bool_true":   []string{"1"},
		"int_large":   []string{"8194723"},
	}

	formValues := FormValues(params)

	if !reflect.DeepEqual(formValues, expectedFormValues) {
		t.Error("Expected form values: ", expectedFormValues, "; Got: ", formValues)
	}

	expectedQueryString := "?bool_true=1&int_large=8194723&string_full=hello+world"

	queryString := QueryString(params)

	if queryString != expectedQueryString {
		t.Error("Expected query string: ", expectedQueryString, "; Got: ", queryString)
	}
}
