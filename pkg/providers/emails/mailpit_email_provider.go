package emails

import (
	"bernardtm/backend/configs"
	"bytes"
	"encoding/base64"
	"fmt"
	"net/smtp"
	"strings"
)

// SMTPProvider is an implementation of the EmailProvider interface
type mailpitSMTPProvider struct {
	Host     string // SMTP server host (e.g., "localhost" for Mailpit)
	Port     int    // SMTP server port (e.g., 1025 for Mailpit)
	Username string // SMTP username (not used for Mailpit, but included for extensibility)
	Password string // SMTP password (not used for Mailpit, but included for extensibility)
}

func NewMailpitSMTPEmailProvider(config *configs.AppConfig) *mailpitSMTPProvider {
	return &mailpitSMTPProvider{
		Host: config.MAILPIT_HOST,
		Port: config.MAILPIT_PORT,
	}
}

// Send implements the EmailProvider interface for SMTP
func (s *mailpitSMTPProvider) Send(email EmailDto) error {
	// Connect to the SMTP server
	address := fmt.Sprintf("%s:%d", s.Host, s.Port)

	// Create the email headers
	var headers []string
	headers = append(headers, fmt.Sprintf("From: %s", email.Sender))
	headers = append(headers, fmt.Sprintf("To: %s", strings.Join(email.To, ", ")))
	if len(email.CC) > 0 {
		headers = append(headers, fmt.Sprintf("Cc: %s", strings.Join(email.CC, ", ")))
	}
	headers = append(headers, fmt.Sprintf("Subject: %s", email.Subject))
	headers = append(headers, "MIME-Version: 1.0")

	if email.IsHTML {
		headers = append(headers, "Content-Type: text/html; charset=\"utf-8\"")
	} else {
		headers = append(headers, "Content-Type: text/plain; charset=\"utf-8\"")
	}

	// Handle attachments if any
	var emailBody string
	if len(email.Attachments) > 0 {
		boundary := "my-boundary"
		headers = append(headers, fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s", boundary))

		var bodyBuffer bytes.Buffer
		bodyBuffer.WriteString(fmt.Sprintf("--%s\n", boundary))
		bodyBuffer.WriteString(fmt.Sprintf("Content-Type: %s; charset=\"utf-8\"\n\n", "text/plain"))
		bodyBuffer.WriteString(email.Body + "\n")

		// Add attachments
		for _, attachment := range email.Attachments {
			bodyBuffer.WriteString(fmt.Sprintf("--%s\n", boundary))
			bodyBuffer.WriteString(fmt.Sprintf("Content-Type: %s\n", attachment.MIMEType))
			bodyBuffer.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\n", attachment.Filename))
			bodyBuffer.WriteString("Content-Transfer-Encoding: base64\n\n")
			bodyBuffer.WriteString(base64.StdEncoding.EncodeToString(attachment.Content) + "\n")
		}
		bodyBuffer.WriteString(fmt.Sprintf("--%s--", boundary))
		emailBody = bodyBuffer.String()
	} else {
		emailBody = email.Body
	}

	// Combine headers and body
	message := strings.Join(headers, "\r\n") + "\r\n\r\n" + emailBody

	// Send the email
	err := smtp.SendMail(address, nil, email.Sender, append(email.To, append(email.CC, email.BCC...)...), []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Println("Email sent successfully!")
	return nil
}
