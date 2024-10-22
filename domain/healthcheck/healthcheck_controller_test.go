package healthcheck_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"btmho/app/domain/healthcheck"

	"github.com/stretchr/testify/assert"
)

func TestHealthcheck(t *testing.T) {
	// Create a new HealthcheckController
	healthcheck.NewHealthcheckController()

	// Create a request to pass to our handler
	req, err := http.NewRequest(http.MethodGet, "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the Healthcheck handler
	handler := http.HandlerFunc(healthcheck.Healthcheck)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	expected := `{"status":"UP"}`
	assert.JSONEq(t, expected, rr.Body.String())
}
