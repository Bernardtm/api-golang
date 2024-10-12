package controllers

import (
	"encoding/json"
	"net/http"
)

// HealthcheckController handles healthcheck-related HTTP requests
type HealthcheckController struct {
}

// NewHealthcheckController creates a new HealthcheckController
func NewHealthcheckController() *HealthcheckController {
	return &HealthcheckController{}
}

// Healthcheck returns the health status of the service
func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "UP"})
}
