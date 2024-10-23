package users_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"btmho/app/domain/users"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock implementation of UserService
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) ListUsers() ([]users.UserDTO, error) {
	args := m.Called()
	if user := args.Get(0); user != nil {
		return user.([]users.UserDTO), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestListUsers(t *testing.T) {
	mockUserService := new(MockUserService)

	controller := users.NewUserController(mockUserService)

	// Prepare a response for the mock service
	expectedUsers := []users.UserDTO{
		{ID: "1", Email: "user1@example.com"},
		{ID: "2", Email: "user2@example.com"},
	}

	mockUserService.On("ListUsers").Return(expectedUsers, nil)

	req, err := http.NewRequest(http.MethodGet, "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(controller.ListUsers)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var actualUsers []users.UserDTO
	err = json.NewDecoder(rr.Body).Decode(&actualUsers)
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, actualUsers)

	mockUserService.AssertExpectations(t)
}

func TestListUsersController_Error(t *testing.T) {
	mockUserService := new(MockUserService)

	controller := users.NewUserController(mockUserService)

	mockUserService.On("ListUsers").Return(nil, errors.New("error fetching users"))

	req, err := http.NewRequest(http.MethodGet, "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(controller.ListUsers)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	assert.Equal(t, "Error fetching users\n", rr.Body.String())

	mockUserService.AssertExpectations(t)
}
