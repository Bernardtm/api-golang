package emails

import (
	"bernardtm/backend/configs"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	mailpitHost = "localhost"                             // Mailpit SMTP and API host
	mailpitPort = 1025                                    // Mailpit SMTP port
	mailpitAPI  = "http://localhost:8025/api/v1/messages" // Mailpit API endpoint to fetch emails
)

// MailpitMessage represents a single email message from the Mailpit API
type MailpitMessage struct {
	ID        string    `json:"ID"`
	MessageID string    `json:"MessageID"`
	From      Address   `json:"From"`
	To        []Address `json:"To"`
	Subject   string    `json:"Subject"`
}

type Address struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

// MailpitResponse represents the structure of the Mailpit API response
type MailpitResponse struct {
	Total    int              `json:"total"`
	Messages []MailpitMessage `json:"messages"`
}

// TestSMTPProviderSend tests the SMTPProvider with Mailpit
func TestSMTPProviderSend(t *testing.T) {
	// Setup SMTPProvider for Mailpit
	appConfig := &configs.AppConfig{
		MAILPIT_HOST: "localhost",
		MAILPIT_PORT: 1025,
	}
	mailpitSMTPProvider := NewMailpitSMTPEmailProvider(appConfig)

	// Create a test email
	email := EmailDto{
		To:      []string{"recipient@example.com"},
		Sender:  "sender@example.com",
		Subject: "Test Email Subject1",
		Body:    "This is a test email body.",
		IsHTML:  false,
	}

	// Send the email using SMTPProvider
	err := mailpitSMTPProvider.Send(email)
	if err != nil {
		t.Fatalf("failed to send email: %v", err)
	}

	// Allow some time for Mailpit to receive the email
	time.Sleep(1 * time.Second)

	// Verify the email was received in Mailpit using its API
	resp, err := http.Get(mailpitAPI)
	if err != nil {
		t.Fatalf("failed to connect to Mailpit API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code from Mailpit API: %d", resp.StatusCode)
	}

	// Parse the response
	var mailpitResp MailpitResponse
	fmt.Println(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&mailpitResp)
	if err != nil {
		t.Fatalf("failed to parse Mailpit API response: %v", err)
	}

	// Validate the received email
	if mailpitResp.Total == 0 {
		t.Fatal("no emails found in Mailpit")
	}

	found := false
	for _, msg := range mailpitResp.Messages {
		if msg.Subject == email.Subject && len(msg.From.Address) > 0 && msg.From.Address == email.Sender {
			found = true
			fmt.Printf("Email received: Subject=%s, From=%s, To=%v\n", msg.Subject, msg.From, msg.To)
			break
		}
	}

	// Assert that the email was found
	assert.True(t, found, "email with subject %q not found in Mailpit", email.Subject)
}
