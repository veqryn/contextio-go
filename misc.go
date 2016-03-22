package ciolite

// InboxName returns the name of the "Inbox" folder, based on the email provider (discovery type)
func InboxName(discoveryType string) string {
	switch discoveryType {
	case "gmail":
		return "INBOX"
	case "googleapps":
		return "INBOX"
	case "msliveconnect":
		return "Inbox"
	default:
		panic("Unknown discovery type: " + discoveryType)
	}
}
