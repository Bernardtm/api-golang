package emails

import (
	"bernardtm/backend/configs"
	"encoding/base64"
	"fmt"
	"github.com/mailgun/mailgun-go"
)

type mailgunSMTPProvider struct {
	Domain string
	APIKey string
}

func NewMailgunProvider(config *configs.AppConfig) *mailgunSMTPProvider {
	return &mailgunSMTPProvider{
		Domain: config.MAILGUN_DOMAIN,
		APIKey: config.MAILGUN_API_KEY,
	}
}

func (m *mailgunSMTPProvider) Send(email EmailDto) error {
	mg := mailgun.NewMailgun(m.Domain, m.APIKey)

	message := mg.NewMessage(email.Sender, email.Subject, email.Body)

	for _, recipient := range email.To {
		err := message.AddRecipient(recipient)

		if err != nil {
			return err
		}
	}

	for _, cc := range email.CC {
		message.AddCC(cc)
	}

	for _, bcc := range email.BCC {
		message.AddBCC(bcc)
	}

	if email.IsHTML {
		message.SetHtml(email.Body)
	}

	for _, attachment := range email.Attachments {
		encodedContent := base64.StdEncoding.EncodeToString(attachment.Content)
		message.AddBufferAttachment(attachment.Filename, []byte(encodedContent))
	}

	_, _, err := mg.Send(message)
	if err != nil {
		return fmt.Errorf("failed to send email via Mailgun: %w", err)
	}

	fmt.Println("Email sent successfully via Mailgun!")
	return nil
}
