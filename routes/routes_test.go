package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenRoutes(t *testing.T) {
	router := SetupRoutes()

	// Testando a rota "/"
	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "OK", resp.Body.String())

	// Testando a rota "/register"
	req, _ = http.NewRequest("POST", "/register", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusCreated, resp.Code)

	// Testando a rota "/login"
	req, _ = http.NewRequest("POST", "/login", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.JSONEq(t, `{"token": "jwt-token"}`, resp.Body.String())

	// Testando a rota "/password-recovery"
	req, _ = http.NewRequest("POST", "/password-recovery", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "Password recovery email sent", resp.Body.String())
}

func TestProtectedRoutes(t *testing.T) {
	router := SetupRoutes()

	// Testando rota protegida "/users" com token válido
	req, _ := http.NewRequest("GET", "/users", nil)
	req.Header.Set("Authorization", "Bearer valid-token")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.JSONEq(t, `[{"name": "John Doe"}]`, resp.Body.String())

	// Testando rota protegida "/users" com token inválido
	req, _ = http.NewRequest("GET", "/users", nil)
	req.Header.Set("Authorization", "Bearer invalid-token")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusUnauthorized, resp.Code)
}
