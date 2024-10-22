package email

import "log"

type EmailService interface {
	SendRecoveryEmail(email, token string) error
}

type emailService struct {
	// You can add SMTP client or configurations here if needed
}

func NewEmailService() EmailService {
	return &emailService{}
}

// SendRecoveryEmail sends a password recovery email
func (s *emailService) SendRecoveryEmail(email, token string) error {
	// Here you would implement the logic to send the email.
	// This can include using an SMTP library or any email service provider.

	recoveryLink := "https://meusite.com/recovery?token=" + token
	message := "To recover your password, please click the link: " + recoveryLink

	// Replace with actual email sending logic
	log.Printf("Sending recovery email to %s: %s\n", email, message)

	return nil
}
