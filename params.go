package ciolite

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)


type Params struct {
	BodyType               string `json:"body_type,omitempty"`
	CallbackURL            string `json:"callback_url,omitempty"`
	Delimiter              string `json:"delimiter,omitempty"`
	Email                  string `json:"email,omitempty"`
	FailureNotifURL        string `json:"failure_notif_url,omitempty"`
	FirstName              string `json:"first_name,omitempty"`
	MigrateAccountID       string `json:"migrate_account_id,omitempty"`
	NewFolderID            string `json:"new_folder_id,omitempty"`
	LastName               string `json:"last_name,omitempty"`
	Password               string `json:"password,omitempty"`
	ProviderRefreshToken   string `json:"provider_refresh_token,omitempty"`
	ProviderToken          string `json:"provider_token,omitempty"`
	ProviderTokenSecret    string `json:"provider_token_secret,omitempty"`
	ProviderConsumerKey    string `json:"provider_consumer_key,omitempty"`
	ProviderConsumerSecret string `json:"provider_consumer_secret,omitempty"`
	Server                 string `json:"server,omitempty"`
	SourceType             string `json:"source_type,omitempty"`
	Status                 string `json:"status,omitempty"`
	StatusCallbackURL      string `json:"status_callback_url,omitempty"`
	StatusOK               string `json:"status_ok,omitempty"`
	Type                   string `json:"type,omitempty"`
	Username               string `json:"username,omitempty"`

	Active            bool `json:"active,omitempty"`
	ForceStatusCheck  bool `json:"force_status_check,omitempty"`
	IncludeBody       bool `json:"include_body,omitempty"`
	IncludeHeaders    bool `json:"include_headers,omitempty"`
	IncludeFlags      bool `json:"include_flags,omitempty"`
	IncludeNamesOnly  bool `json:"include_names_only,omitempty"`
	Raw               bool `json:"raw,omitempty"`
	RawFileList       bool `json:"raw_file_list,omitempty"`
	SourceRawFileList bool `json:"source_raw_file_list,omitempty"`
	UseSSL            bool `json:"use_ssl,omitempty"`

	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
	Port   int `json:"port,omitempty"`
}

// FormValues returns valid FormValues for CIO Lite
func FormValues(cioParams interface{}) url.Values {

	// Values
	values := url.Values{}

	// dynamically iterate through struct fields
	refVal := reflect.ValueOf(cioParams)
	refType := reflect.TypeOf(cioParams)
	for i, numFields := 0, refVal.NumField(); i < numFields; i++ {
		fieldValue := refVal.Field(i)
		fieldType := refType.Field(i)

		// dynamically choose how to fill the values based on field type
		// and set the key to the json tag name
		switch fieldValue.Kind() {

		case reflect.String:
			v := fieldValue.String()
			if len(v) > 0 {
				values.Set(jsonName(fieldType), v)
			}

		case reflect.Bool:
			v := fieldValue.Bool()
			if v {
				values.Set(jsonName(fieldType), "1")
			}

		case reflect.Int:
			v := fieldValue.Int()
			if v != 0 {
				values.Set(jsonName(fieldType), fmt.Sprintf("%d", v))
			}

		default:
			panic("Unexpected CioParams type: " + fieldValue.Kind().String())
		}
	}

	return values
}

// QueryString returns a query string
func QueryString(cioParams interface{}) string {

	// Encode parameters
	encoded := FormValues(cioParams).Encode()
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
		panic(fmt.Sprintf("CioParam %s missing json name tag", sf.Name))
	}
	if indexComma >= 0 {
		return jsonTag[:indexComma]
	}
	return jsonTag
}
