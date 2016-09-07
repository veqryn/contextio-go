package cioutil

import (
	"net/url"
	"reflect"
	"testing"
)

// TestFormValues tests that the form values function returns the correct url.Values
func TestFormValuesAndQueryString(t *testing.T) {
	t.Parallel()

	params := struct {
		// Fields without omitempty should always be included
		BoolAlwaysInclude   bool   `json:"bool_always_include"`
		IntAlwaysInclude    int    `json:"int_always_include"`
		StringAlwaysInclude string `json:"string_always_include"`
		// Fields with omitempty that are zero values should not be included
		BoolTrue    bool   `json:"bool_true,omitempty"`
		BoolFalse   bool   `json:"bool_false,omitempty"`
		IntLarge    int    `json:"int_large,omitempty"`
		IntZero     int    `json:"int_zero,omitempty"`
		StringFull  string `json:"string_full,omitempty"`
		StringEmpty string `json:"string_empty,omitempty"`
	}{
		// Test values:
		BoolTrue:    true,
		BoolFalse:   false,
		IntLarge:    8194723,
		IntZero:     0,
		StringFull:  "hello world",
		StringEmpty: "",
	}

	expectedFormValues := url.Values{
		// Booleans get converted to 0 or 1
		"bool_always_include":   []string{"0"},
		"int_always_include":    []string{"0"},
		"string_always_include": []string{""},
		"bool_true":             []string{"1"},
		"int_large":             []string{"8194723"},
		"string_full":           []string{"hello world"},
	}

	formValues := FormValues(params)

	if !reflect.DeepEqual(formValues, expectedFormValues) {
		t.Error("Expected form values: ", expectedFormValues, "; Got: ", formValues)
	}

	expectedQueryString := "?bool_always_include=0&bool_true=1&int_always_include=0&int_large=8194723&string_always_include=&string_full=hello+world"

	queryString := QueryString(params)

	if queryString != expectedQueryString {
		t.Error("Expected query string: ", expectedQueryString, "; Got: ", queryString)
	}
}
