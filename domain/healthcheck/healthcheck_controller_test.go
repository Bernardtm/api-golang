package healthcheck_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"btmho/app/domain/healthcheck"

	"github.com/stretchr/testify/assert"
)

func TestHealthcheck(t *testing.T) {
	healthcheck.NewHealthcheckController()

	req, err := http.NewRequest(http.MethodGet, "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(healthcheck.Healthcheck)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := `{"status":"UP"}`
	assert.JSONEq(t, expected, rr.Body.String())
}
