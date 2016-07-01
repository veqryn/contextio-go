package ciolite

import (
	"encoding/json"
	"testing"
)

// TestReceivingWebhook tests receiving and parsing WebhookCallback
func TestSimulatedReceivingWebhook(t *testing.T) {
	t.Parallel()

	emptyAddressesJSON := `
{
	"account_id": "abc4567XYZ",
	"webhook_id": "abc4567XYZ",
	"timestamp": 1467254577,
	"message_data": {
		"email_message_id": "<blah-1234@test.com>",
		"addresses": [],
		"references": [],
		"message_id": "<blah-1234@test.com>",
		"flags": {
			"flagged": false,
			"answered": false,
			"draft": false,
			"seen": false
		},
		"sources": [],
		"email_accounts": [],
		"files": [],
		"person_info": [],
		"date": 1464742490,
		"subject": "Test Subject",
		"folders": ["Inbox"]
	},
	"user_id": "abc4567XYZ",
	"token": "fake578Token",
	"signature": "fake123Signature"
}`

	fullAddressesJSON := `
{
	"account_id": "abc4567XYZ",
	"webhook_id": "abc4567XYZ",
	"token": "fake578Token",
	"signature": "fake123Signature",
	"timestamp": 1467254202,
	"message_data": {
		"message_id": "<blah-1234@test.com>",
		"email_message_id": "<blah-1234@test.com>",
		"subject": "Test Subject",
		"folders": ["Inbox"],
		"date": 1446919784,
		"date_received": 1446919783,
		"addresses": {
			"from": {
				"email": "from@test.com",
				"name": "John"
			},
			"to": [{
				"email": "to@test.com"
			}],
			"reply_to": [{
				"email": "reply@test.com"
			}]
		},
		"flags": {
			"seen": true
		},
		"email_accounts": []
	}
}`

	//type ExpiresTest struct {
	//	MyField ExpiresMixed `json:"my_field"`
	//}
	//
	//falseExpectedJSON := `{"my_field":false}`
	//falseExpectedExpires := ExpiresTest{MyField: ExpiresMixed{}}
	//
	//timestamp := 1234509876
	//timestampExpectedJSON := `{"my_field":1234509876}`
	//timestampExpectedExpires := ExpiresTest{MyField: ExpiresMixed{Expires: &timestamp}}

	// Test Empty Address Array
	var emptyAddresses WebhookCallback
	Must(json.Unmarshal([]byte(emptyAddressesJSON), &emptyAddresses))

	var fullAddresses WebhookCallback
	Must(json.Unmarshal([]byte(fullAddressesJSON), &fullAddresses))

	//if !reflect.DeepEqual(falseExpires, falseExpectedExpires) ||
	//falseExpires.MyField.Unused() ||
	//falseExpires.MyField.Timestamp() != -1 {
	//	t.Error(falseExpires)
	//}
	//
	//falseJSON, err := json.Marshal(falseExpires)
	//Must(err)
	//
	//if string(falseJSON) != falseExpectedJSON {
	//	t.Error(string(falseJSON))
	//}
	//
	//// Test timestamp
	//var timestampExpires ExpiresTest
	//Must(json.Unmarshal([]byte(timestampExpectedJSON), &timestampExpires))
	//
	//if !reflect.DeepEqual(timestampExpires, timestampExpectedExpires) ||
	//!timestampExpires.MyField.Unused() ||
	//timestampExpires.MyField.Timestamp() != timestamp {
	//	t.Error(timestampExpires)
	//}
	//
	//timestampJSON, err := json.Marshal(timestampExpires)
	//Must(err)
	//
	//if string(timestampJSON) != timestampExpectedJSON {
	//	t.Error(string(timestampJSON))
	//}
}
