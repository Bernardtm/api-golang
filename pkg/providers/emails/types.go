package emails

type AddressName struct {
	Email string `json:"Email"`
	Name  string `json:"Name"`
}

type EmailMailpitPayload struct {
	From        AddressName          `json:"From"`
	To          []AddressName        `json:"To"`
	CC          []AddressName        `json:"CC,omitempty"`
	BCC         []AddressName        `json:"BCC,omitempty"`
	Subject     string               `json:"Subject"`
	Text        string               `json:"Text"`
	ContentType string               `json:"Content_type"`
	Attachments []MailpitAttachtment `json:"Attachments"`
}

type MailpitAttachtment struct {
	Filename string `json:"Filename"`
	MIMEType string `json:"Mime_type"`
	Content  string `json:"Content"` // Base64-encoded content
}

type EmailMailgunPayload struct {
	From        AddressName         `json:"From"`
	To          []string            `json:"To"`
	CC          []string            `json:"CC,omitempty"`
	BCC         []string            `json:"BCC,omitempty"`
	Subject     string              `json:"Subject"`
	Text        string              `json:"Text"`
	ContentType string              `json:"Content_type"`
	Attachments []MailgunAttachment `json:"Attachments"`
}

type MailgunAttachment struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
}
