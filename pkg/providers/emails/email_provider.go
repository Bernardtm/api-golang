package emails

// EmailDto represents the structure of an email to be sent
type EmailDto struct {
	Sender      string            `json:"sender"`                // Sender email address
	To          []string          `json:"to"`                    // List of recipients
	CC          []string          `json:"cc,omitempty"`          // List of CC recipients (optional)
	BCC         []string          `json:"bcc,omitempty"`         // List of BCC recipients (optional)
	Subject     string            `json:"subject"`               // Email subject
	Body        string            `json:"body"`                  // Email body content
	IsHTML      bool              `json:"is_html"`               // Indicates if the body is in HTML format
	Attachments []Attachment      `json:"attachments,omitempty"` // Optional attachments
	Headers     map[string]string `json:"headers,omitempty"`     // Custom headers (optional)
}

// Attachment represents an email attachment
type Attachment struct {
	Filename string `json:"filename"`  // File name of the attachment
	Content  []byte `json:"content"`   // File content as bytes
	MIMEType string `json:"mime_type"` // MIME type of the file (e.g., "application/pdf")
}

// EmailProvider defines the interface for an email provider
type EmailProvider interface {
	Send(email EmailDto) error
}
