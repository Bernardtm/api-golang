package e2e_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHealthcheckEndpoint is an end-to-end test for the /status endpoint
func TestHealthcheckEndpoint(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, baseURL+"", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check the status code
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status OK")

	// Parse the response
	var responseData map[string]string
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	// Validate the response content
	assert.Equal(t, "ok", responseData["status"], "Status mismatch")
	assert.Equal(t, "1.0", responseData["version"], "Version mismatch")
}
