package auth_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"btmho/app/domain/auth"
	"btmho/app/domain/users"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) RegisterUser(user *users.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockAuthService) Login(credentials auth.Credentials) (string, error) {
	args := m.Called(credentials)
	return args.String(0), args.Error(1)
}

func (m *MockAuthService) RecoverPassword(email string) error {
	args := m.Called(email)
	return args.Error(0)
}

func TestRegister_UserCreatedSuccessfully(t *testing.T) {
	mockAuthService := new(MockAuthService)
	controller := auth.NewAuthController(mockAuthService)

	user := users.User{
		FullName:        "John Doe",
		Email:           "john@example.com",
		Password:        "SecurePassword123!",
		ConfirmPassword: "SecurePassword123!",
	}

	mockAuthService.On("RegisterUser", &user).Return(nil)

	body, _ := json.Marshal(user)
	req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Register)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var response map[string]string
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "User created successfully", response["message"])

	mockAuthService.AssertExpectations(t)
}

func TestRegister_InvalidInput(t *testing.T) {
	mockAuthService := new(MockAuthService)
	controller := auth.NewAuthController(mockAuthService)

	req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte("invalid")))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Register)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "Invalid input\n", rr.Body.String())
}

func TestRegister_ErrorWhileRegisteringUser(t *testing.T) {
	mockAuthService := new(MockAuthService)
	controller := auth.NewAuthController(mockAuthService)

	user := users.User{
		FullName:        "John Doe",
		Email:           "john@example.com",
		Password:        "SecurePassword123!",
		ConfirmPassword: "SecurePassword123!",
	}

	mockAuthService.On("RegisterUser", &user).Return(errors.New("user already exists"))

	body, _ := json.Marshal(user)
	req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Register)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "user already exists\n", rr.Body.String())

	mockAuthService.AssertExpectations(t)
}

func TestLogin_UserLoggedInSuccessfully(t *testing.T) {
	mockAuthService := new(MockAuthService)
	controller := auth.NewAuthController(mockAuthService)

	credentials := auth.Credentials{
		Email:    "john@example.com",
		Password: "SecurePassword123!",
	}

	mockAuthService.On("Login", credentials).Return("jwt.token.here", nil)

	body, _ := json.Marshal(credentials)
	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Login)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]string
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "jwt.token.here", response["token"])

	mockAuthService.AssertExpectations(t)
}

func TestLogin_InvalidInput(t *testing.T) {
	mockAuthService := new(MockAuthService)
	controller := auth.NewAuthController(mockAuthService)

	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte("invalid")))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Login)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "Invalid input\n", rr.Body.String())
}

func TestLogin_ErrorWhileLoggingIn(t *testing.T) {
	mockAuthService := new(MockAuthService)
	controller := auth.NewAuthController(mockAuthService)

	credentials := auth.Credentials{
		Email:    "john@example.com",
		Password: "SecurePassword123!",
	}

	mockAuthService.On("Login", credentials).Return("", errors.New("invalid credentials"))

	body, _ := json.Marshal(credentials)
	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.Login)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	assert.Equal(t, "invalid credentials\n", rr.Body.String())

	mockAuthService.AssertExpectations(t)
}

func TestPasswordRecovery_Success(t *testing.T) {
	mockAuthService := new(MockAuthService)
	controller := auth.NewAuthController(mockAuthService)

	request := auth.PasswordRecoveryRequest{
		Email: "john@example.com",
	}

	mockAuthService.On("RecoverPassword", request.Email).Return(nil)

	body, _ := json.Marshal(request)
	req, err := http.NewRequest(http.MethodPost, "/password-recovery", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.PasswordRecovery)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	mockAuthService.AssertExpectations(t)
}

func TestPasswordRecovery_InvalidInput(t *testing.T) {
	mockAuthService := new(MockAuthService)
	controller := auth.NewAuthController(mockAuthService)

	req, err := http.NewRequest(http.MethodPost, "/password-recovery", bytes.NewBuffer([]byte("invalid")))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.PasswordRecovery)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "Invalid input\n", rr.Body.String())
}

func TestPasswordRecovery_ErrorWhileRecovering(t *testing.T) {
	mockAuthService := new(MockAuthService)
	controller := auth.NewAuthController(mockAuthService)

	request := auth.PasswordRecoveryRequest{
		Email: "john@example.com",
	}

	mockAuthService.On("RecoverPassword", request.Email).Return(errors.New("error generating recovery token"))

	body, _ := json.Marshal(request)
	req, err := http.NewRequest(http.MethodPost, "/password-recovery", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.PasswordRecovery)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "Error generating recovery token\n", rr.Body.String())

	mockAuthService.AssertExpectations(t)
}
