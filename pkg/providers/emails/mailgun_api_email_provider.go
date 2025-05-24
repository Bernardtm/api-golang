package emails

import (
	"encoding/base64"
	"github.com/mailgun/mailgun-go"
)

type MailgunAPIProvider struct {
	APIBaseURL string
	APIKey     string
	Domain     string
}

func (m *MailgunAPIProvider) Send(email EmailDto) error {
	to := email.To
	cc := email.CC
	bcc := email.BCC

	payload := EmailMailgunPayload{
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

	for _, attachment := range email.Attachments {
		encodedContent := base64.StdEncoding.EncodeToString(attachment.Content)
		payload.Attachments = append(payload.Attachments, MailgunAttachment{
			Filename: attachment.Filename,
			Content:  encodedContent,
		})
	}

	mg := mailgun.NewMailgun(m.Domain, m.APIKey)

	message := mg.NewMessage(
		email.Sender, email.Subject, email.Body, to...)

	resp, _, err := mg.Send(message)

	if err != nil || (resp != "Queued" && resp != "Sent") {
		return err
	}

	return nil
}
