package e2e_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHealthcheckEndpoint is an end-to-end test for the /status endpoint
func TestAuthLoginEndpoint(t *testing.T) {
	firstToken, statusCode, err := AuthLoginEndpoint()
	if err != nil {
		t.Fatalf("Failed to login: %v", err)
	}

	// Check the status code
	assert.Equal(t, http.StatusOK, statusCode, "Expected status OK")
	assert.NotEmpty(t, firstToken, "Expected token to be non-empty")

	token, _, err := AuthLogin2StepEndpoint(firstToken)
	if err != nil {
		t.Fatalf("Failed to login: %v", err)
	}
	fmt.Println(token)
	assert.NotEmpty(t, token, "Expected token to be non-empty")
}

func AuthLoginEndpoint() (string, int, error) {
	// Create JSON payload
	payload := map[string]interface{}{
		"email":    "user@example.com",
		"password": "Password123#",
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", 0, fmt.Errorf("Failed to marshal payload: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, baseURL+contextPath+"/auth/login", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", 0, fmt.Errorf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Parse the response
	var responseData map[string]string
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		return "", 0, fmt.Errorf("Failed to decode response: %v", err)
	}

	return responseData["token"], resp.StatusCode, nil
}

// MailpitMessage represents a single email message from the Mailpit API
type MailpitMessage struct {
	ID        string    `json:"ID"`
	MessageID string    `json:"MessageID"`
	From      Address   `json:"From"`
	To        []Address `json:"To"`
	Subject   string    `json:"Subject"`
}
type MailpitResponse struct {
	Total    int              `json:"total"`
	Messages []MailpitMessage `json:"messages"`
}

type Address struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

func MailpitGetMessages() string {
	// MailpitResponse represents the structure of the Mailpit API response

	const (
		mailpitHost = "localhost"                             // Mailpit SMTP and API host
		mailpitPort = 1025                                    // Mailpit SMTP port
		mailpitAPI  = "http://localhost:8025/api/v1/messages" // Mailpit API endpoint to fetch emails
	)

	// Verify the email was received in Mailpit using its API
	resp, err := http.Get(mailpitAPI)
	if err != nil {
		fmt.Errorf("failed to connect to Mailpit API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("unexpected status code from Mailpit API: %d", resp.StatusCode)
	}

	// Parse the response
	var mailpitResp MailpitResponse
	fmt.Println(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&mailpitResp)
	if err != nil {
		fmt.Errorf("failed to parse Mailpit API response: %v", err)
	}

	// Validate the received email
	if mailpitResp.Total == 0 {
		fmt.Errorf("no emails found in Mailpit")
	}

	found := false
	for _, msg := range mailpitResp.Messages {
		if msg.Subject == "Código de Verificação" && len(msg.From.Address) > 0 && msg.From.Address == "no-reply@company.com" {
			found = true
			fmt.Printf("Email received: Subject=%s, From=%s, To=%v\n", msg.Subject, msg.From, msg.To)
			fmt.Println(found)
			return msg.ID
		}
	}
	return ""
}

type MailpitMessageResponse struct {
	ID        string    `json:"ID"`
	MessageID string    `json:"MessageID"`
	From      Address   `json:"From"`
	To        []Address `json:"To"`
	Subject   string    `json:"Subject"`
	Text      string    `json:"Text"`
}

func MailpitGetMessageText(messageID string) string {
	// MailpitResponse represents the structure of the Mailpit API response

	const (
		mailpitHost = "localhost" // Mailpit SMTP and API host
		mailpitPort = 1025        // Mailpit SMTP port
	)

	// Verify the email was received in Mailpit using its API
	url := "http://localhost:8025/api/v1/message/" + messageID
	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("failed to connect to Mailpit API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("unexpected status code from Mailpit API: %d", resp.StatusCode)
	}

	// Parse the response
	var mailpitResp MailpitMessageResponse
	fmt.Println(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&mailpitResp)
	if err != nil {
		fmt.Errorf("failed to parse Mailpit API response: %v", err)
	}

	// Validate the received email
	if len(mailpitResp.Text) > 0 {
		fmt.Errorf("no email text found in Mailpit")
	}

	return mailpitResp.Text
}

func AuthLogin2StepEndpoint(token string) (string, int, error) {
	id := MailpitGetMessages()
	messageText := MailpitGetMessageText(id)

	fmt.Println(messageText)
	code, err := extractVerificationCode(messageText)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Extracted code:", code)
	// Create JSON payload
	payload := map[string]interface{}{
		"otp": code,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", 0, fmt.Errorf("Failed to marshal payload: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, baseURL+contextPath+"/auth/login/verify", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", 0, fmt.Errorf("Failed to create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("unexpected status code from API: %d", resp.StatusCode)
	}

	// Parse the response
	var responseData map[string]string
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		return "", 0, fmt.Errorf("Failed to decode response: %v", err)
	}

	return responseData["token"], resp.StatusCode, nil
}

func extractVerificationCode(text string) (string, error) {
	// Define the regex pattern for a 6-digit number
	pattern := `\d{6}`
	re := regexp.MustCompile(pattern)

	// Find the first match in the text
	matches := re.FindString(text)
	if matches == "" {
		return "", fmt.Errorf("verification code not found")
	}
	return matches, nil
}
