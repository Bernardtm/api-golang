package email

import (
	"bernardtm/backend/pkg/providers/emails"
	"errors"
)

// EmailService provides methods to send emails using any provider
type EmailService struct {
	provider emails.EmailProvider
}

// NewEmailService creates a new EmailService instance
func NewEmailService(provider emails.EmailProvider) *EmailService {
	return &EmailService{provider: provider}
}

// SendEmail sends an email using the configured provider
func (s *EmailService) SendEmail(email emails.EmailDto) error {
	if len(email.To) == 0 {
		return errors.New("email must have at least one recipient")
	}

	if email.Sender == "" {
		return errors.New("email must have a sender address")
	}

	// Delegate email sending to the configured provider
	return s.provider.Send(email)
}
