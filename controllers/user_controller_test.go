package controllers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"btmho/app/controllers"
	"btmho/app/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock implementation of UserService
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) ListUsers() ([]models.UserDTO, error) {
	args := m.Called()
	return args.Get(0).([]models.UserDTO), args.Error(1)
}

func TestListUsers(t *testing.T) {
	// Create an instance of the mock user service
	mockUserService := new(MockUserService)

	// Create a user controller with the mock service
	controller := controllers.NewUserController(mockUserService)

	// Prepare a response for the mock service
	expectedUsers := []models.UserDTO{
		{ID: "1", Email: "user1@example.com"},
		{ID: "2", Email: "user2@example.com"},
	}

	// Set up the mock to return expected users without error
	mockUserService.On("ListUsers").Return(expectedUsers, nil)

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the ListUsers handler
	handler := http.HandlerFunc(controller.ListUsers)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	var actualUsers []models.UserDTO
	err = json.NewDecoder(rr.Body).Decode(&actualUsers)
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, actualUsers)

	// Assert that the expectations were met
	mockUserService.AssertExpectations(t)
}

func TestListUsers_Error(t *testing.T) {
	// Create an instance of the mock user service
	mockUserService := new(MockUserService)

	// Create a user controller with the mock service
	controller := controllers.NewUserController(mockUserService)

	// Set up the mock to return an error
	mockUserService.On("ListUsers").Return(nil, errors.New("error fetching users"))

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the ListUsers handler
	handler := http.HandlerFunc(controller.ListUsers)
	handler.ServeHTTP(rr, req)

	// Check the status code is 500 Internal Server Error
	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	// Check the response body
	assert.Equal(t, "Error fetching users\n", rr.Body.String())

	// Assert that the expectations were met
	mockUserService.AssertExpectations(t)
}
