package emails

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

// MailpitAPIProvider is an implementation of the EmailProvider interface using Mailpit API
type MailpitAPIProvider struct {
	APIBaseURL string // Base URL for the Mailpit API
	APIKey     string // Optional API key for authentication (if required)
}

// Send implements the EmailProvider interface for Mailpit API
func (m *MailpitAPIProvider) Send(email EmailDto) error {

	// Map email addresses to AddressName format
	to := MapAddresses(email.To)
	cc := MapAddresses(email.CC)
	bcc := MapAddresses(email.BCC)
	// Create the request payload
	payload := EmailMailpitPayload{
		From:    AddressName{Email: email.Sender, Name: ""},
		To:      to,
		CC:      cc,
		BCC:     bcc,
		Subject: email.Subject,
		Text:    email.Body,
	}

	if email.IsHTML {
		payload.ContentType = "text/html"
	}

	// Add attachments if any
	for _, attachment := range email.Attachments {
		payload.Attachments = append(payload.Attachments, struct {
			Filename string `json:"Filename"`
			MIMEType string `json:"Mime_type"`
			Content  string `json:"Content"`
		}{
			Filename: attachment.Filename,
			MIMEType: attachment.MIMEType,
			Content:  base64.StdEncoding.EncodeToString(attachment.Content),
		})
	}

	// Convert payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to serialize email payload: %w", err)
	}

	// Create the HTTP request
	url := fmt.Sprintf("%s/send", m.APIBaseURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// req.Header.Set("Content-Type", "application/json")
	// if m.APIKey != "" {
	// 	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", m.APIKey))
	// }

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request to Mailpit API: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("mailpit API returned error: %s", resp.Status)
	}

	fmt.Println("Email sent successfully via Mailpit API!")
	return nil
}
