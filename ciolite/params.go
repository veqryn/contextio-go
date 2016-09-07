package cioutil

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

// FormValues returns valid FormValues for CIO
func FormValues(cioFormValueParams interface{}) url.Values {

	// Values
	values := url.Values{}

	// If uninitialized, return empty url.Values
	if cioFormValueParams == nil {
		return values
	}

	// dynamically iterate through struct fields
	refVal := reflect.ValueOf(cioFormValueParams)
	refType := reflect.TypeOf(cioFormValueParams)
	for i, numFields := 0, refVal.NumField(); i < numFields; i++ {
		fieldValue := refVal.Field(i)
		fieldType := refType.Field(i)

		// dynamically choose how to fill the values based on field type
		// and set the key to the json tag name
		switch fieldValue.Kind() {

		case reflect.String:
			v := fieldValue.String()
			if len(v) > 0 || !jsonOmitempty(fieldType) {
				values.Set(jsonName(fieldType), v)
			}

		case reflect.Bool:
			v := fieldValue.Bool()
			if v || !jsonOmitempty(fieldType) {
				// boolean values are set to 0 or 1
				if v {
					values.Set(jsonName(fieldType), "1")
				} else {
					values.Set(jsonName(fieldType), "0")
				}
			}

		case reflect.Int:
			v := fieldValue.Int()
			if v != 0 || !jsonOmitempty(fieldType) {
				values.Set(jsonName(fieldType), fmt.Sprintf("%d", v))
			}

		default:
			panic("Unexpected parameter type: " + fieldValue.Kind().String())
		}
	}

	return values
}

// QueryString returns a query string
func QueryString(cioQueryValueParams interface{}) string {

	// Encode parameters
	encoded := FormValues(cioQueryValueParams).Encode()
	if encoded == "" {
		return encoded
	}

	// Format
	return fmt.Sprintf("?%s", encoded)
}

// jsonName returns the json name based on the json tag of the struct field
func jsonName(sf reflect.StructField) string {
	jsonTag := sf.Tag.Get("json")
	indexComma := strings.Index(jsonTag, ",")
	if len(jsonTag) == 0 || indexComma == 0 {
		// Impossible situation: they all have a name tag
		panic(fmt.Sprintf("Parameter %s missing json name tag", sf.Name))
	}
	if indexComma > 0 {
		return jsonTag[:indexComma]
	}
	return jsonTag
}

// jsonOmitempty returns true if json tags of this field include "omitempty"
func jsonOmitempty(sf reflect.StructField) bool {
	jsonTag := sf.Tag.Get("json")
	indexComma := strings.Index(jsonTag, ",")
	if len(jsonTag) == 0 || indexComma == 0 {
		// Impossible situation: they all have a name tag
		panic(fmt.Sprintf("Parameter %s missing json name tag", sf.Name))
	}
	if indexComma > 0 {
		return jsonTagContains(jsonTag[indexComma+1:], "omitempty")
	}
	return false
}

// jsonTagContains returns true of the JSON tag options (after the name tag)
// contain the desired option (ex: omitempty).
func jsonTagContains(tagOptions string, optionName string) bool {
	s := tagOptions
	for s != "" {
		var next string
		i := strings.Index(s, ",")
		if i >= 0 {
			s, next = s[:i], s[i+1:]
		}
		if s == optionName {
			return true
		}
		s = next
	}
	return false
}
