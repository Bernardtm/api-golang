package email

import (
	"bernardtm/backend/pkg/providers/emails"
	"errors"
	"testing"
)

// MockEmailProvider is a mock implementation of the EmailProvider interface
type MockEmailProvider struct {
	SendFunc func(email emails.EmailDto) error // Function to simulate email sending
}

// Send calls the mock's SendFunc
func (m *MockEmailProvider) Send(email emails.EmailDto) error {
	return m.SendFunc(email)
}

// TestEmailService_SendEmail tests the SendEmail function of the EmailService
func TestEmailService_SendEmail(t *testing.T) {
	tests := []struct {
		name         string
		email        emails.EmailDto
		providerFunc func(email emails.EmailDto) error
		expectError  bool
	}{
		{
			name: "Valid email - no errors",
			email: emails.EmailDto{
				To:      []string{"recipient@example.com"},
				Sender:  "sender@example.com",
				Subject: "Test Subject",
				Body:    "Test Body",
			},
			providerFunc: func(email emails.EmailDto) error {
				// Simulate successful email sending
				return nil
			},
			expectError: false,
		},
		{
			name: "Missing recipient - should return error",
			email: emails.EmailDto{
				Sender:  "sender@example.com",
				Subject: "Test Subject",
				Body:    "Test Body",
			},
			providerFunc: func(email emails.EmailDto) error {
				return nil // This won't be called due to validation
			},
			expectError: true,
		},
		{
			name: "Provider fails to send email",
			email: emails.EmailDto{
				To:      []string{"recipient@example.com"},
				Sender:  "sender@example.com",
				Subject: "Test Subject",
				Body:    "Test Body",
			},
			providerFunc: func(email emails.EmailDto) error {
				return errors.New("failed to send email")
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a mock email provider with the test-specific behavior
			mockProvider := &MockEmailProvider{
				SendFunc: tt.providerFunc,
			}

			// Create the email service using the mock provider
			emailService := NewEmailService(mockProvider)

			// Attempt to send the email
			err := emailService.SendEmail(tt.email)

			// Validate the result
			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}
		})
	}
}
