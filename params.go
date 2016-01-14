// Package ciolite ...
package ciolite

// Imports
import (
	"fmt"
	"net/url"
)

// CioParams ...
type CioParams struct {
	BodyType             string `json:"body_type:omitempty"`
	CallbackURL          string `json:"callback_url:omitempty"`
	Delimiter            string `json:"delimiter:omitempty"`
	Email                string `json:"email,omitempty"`
	FilterNotifURL       string `json:"failure_notif_url,omitempty"`
	IncludeHeaders       string `json:"include_headers:omitempty"`
	MigrateAccountID     string `json:"migrate_account_id,omitempty"`
	FirstName            string `json:"first_name,omitempty"`
	FilterTo             string `json:"filter_to:omitempty"`
	FilterFrom           string `json:"filter_from:omitempty"`
	FilterCC             string `json:"filter_cc:omitempty"`
	FilterSubject        string `json:"filter_subject:omitempty"`
	FilterThread         string `json:"filter_thread:omitempty"`
	FilterNewImportant   string `json:"filter_new_important:omitempty"`
	FilterFileName       string `json:"filter_file_name:omitempty"`
	FilterFolderAdded    string `json:"filter_folder_added:omitempty"`
	FilterToDomain       string `json:"filter_to_domain:omitempty"`
	FilterFromDomain     string `json:"filter_from_domain:omitempty"`
	LastName             string `json:"last_name,omitempty"`
	Password             string `json:"password,omitempty"`
	ProviderRefreshToken string `json:"provider_refresh_token,omitempty"`
	ProviderToken        string `json:"provider_token,omitempty"`
	ProviderTokenSecret  string `json:"provider_token_secret,omitempty"`
	ProviderConsumerKey  string `json:"provider_consumer_key,omitempty"`
	Server               string `json:"server,omitempty"`
	SourceType           string `json:"source_type,omitempty"`
	Status               string `json:"status:omitempty"`
	StatusCallbackURL    string `json:"status_callback_url:omitempty"`
	StatusOK             string `json:"status_ok:omitempty"`
	Type                 string `json:"type,omitempty"`
	Username             string `json:"username,omitempty"`

	Active            bool `json:"active:omitempty"`
	ForceStatusCheck  bool `json:"force_status_check:omitempty"`
	IncludeBody       bool `json:"include_body:omitempty"`
	IncludeFlags      bool `json:"include_flags:omitempty"`
	IncludeNamesOnly  bool `json:"include_names_only:omitempty"`
	Raw               bool `json:"raw:omitempty"`
	RawFileList       bool `json:"raw_file_list,omitempty"`
	SourceRawFileList bool `json:"source_raw_file_list"`
	UseSSL            bool `json:"use_ssl,omitempty"`

	Limit  int `json:"limit:omitempty"`
	Offset int `json:"offset:omitempty"`
	Port   int `json:"port,omitempty"`
}

// FormValues ...
// Make form values
func (cioParams CioParams) FormValues() url.Values {

	// Values
	values := url.Values{}

	// Strings
	if cioParams.BodyType != "" {
		values.Set("body_type", cioParams.BodyType)
	}
	if cioParams.CallbackURL != "" {
		values.Set("callback_url", cioParams.CallbackURL)
	}
	if cioParams.Delimiter != "" {
		values.Set("delimiter", cioParams.Delimiter)
	}
	if cioParams.Email != "" {
		values.Set("email", cioParams.Email)
	}
	if cioParams.FirstName != "" {
		values.Set("first_name", cioParams.FirstName)
	}
	if cioParams.LastName != "" {
		values.Set("last_name", cioParams.LastName)
	}
	if cioParams.Password != "" {
		values.Set("password", cioParams.Password)
	}
	if cioParams.ProviderRefreshToken != "" {
		values.Set("provider_refresh_token", cioParams.ProviderRefreshToken)
	}
	if cioParams.ProviderToken != "" {
		values.Set("provider_token", cioParams.ProviderToken)
	}
	if cioParams.ProviderTokenSecret != "" {
		values.Set("provider_token_secret", cioParams.ProviderTokenSecret)
	}
	if cioParams.ProviderConsumerKey != "" {
		values.Set("provider_consumer_key", cioParams.ProviderConsumerKey)
	}
	if cioParams.IncludeHeaders != "" {
		values.Set("include_headers", cioParams.IncludeHeaders)
	}
	if cioParams.Server != "" {
		values.Set("server", cioParams.Server)
	}
	if cioParams.Status != "" {
		values.Set("status", cioParams.Status)
	}
	if cioParams.StatusOK != "" {
		values.Set("status_ok", cioParams.StatusOK)
	}
	if cioParams.Type != "" {
		values.Set("type", cioParams.Type)
	}
	if cioParams.Username != "" {
		values.Set("username", cioParams.Username)
	}

	// Booleans
	if cioParams.ForceStatusCheck {
		values.Set("force_status_check", "1")
	}
	if cioParams.IncludeBody {
		values.Set("include_body", "1")
	}
	if cioParams.IncludeFlags {
		values.Set("include_flags", "1")
	}
	if cioParams.IncludeNamesOnly {
		values.Set("include_names_only", "1")
	}
	if cioParams.Raw {
		values.Set("raw", "1")
	}
	if cioParams.RawFileList {
		values.Set("raw_file_list", "1")
	}
	if cioParams.SourceRawFileList {
		values.Set("source_raw_file_list", "1")
	}
	if cioParams.UseSSL {
		values.Set("use_ssl", "1")
	}

	// Integers
	if cioParams.Limit != 0 {
		values.Set("limit", fmt.Sprintf("%d", cioParams.Limit))
	}
	if cioParams.Offset != 0 {
		values.Set("offset", fmt.Sprintf("%d", cioParams.Offset))
	}
	if cioParams.Port != 0 {
		values.Set("port", fmt.Sprintf("%d", cioParams.Port))
	}

	return values
}

// QueryString ...
// Make a query string
func (cioParams CioParams) QueryString() string {

	// Encode parameters
	encoded := cioParams.FormValues().Encode()
	if encoded == "" {
		return encoded
	}

	// Format
	return fmt.Sprintf("?%v", encoded)
}
