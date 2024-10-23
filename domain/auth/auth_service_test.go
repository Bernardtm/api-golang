package auth_test

import (
	clients "btmho/app/clients/address"
	"btmho/app/domain/address"
	"btmho/app/domain/auth"
	"btmho/app/domain/users"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAddressClient struct {
	mock.Mock
}

// FetchCEPData simulates fetching address data based on the provided CEP
func (m *MockAddressClient) FetchCEPData(cep string) (*clients.AddressDTO, error) {
	args := m.Called(cep)
	return args.Get(0).(*clients.AddressDTO), args.Error(1)
}

type MockUserRepository struct {
	mock.Mock
}

// Implements the GetAllUsers method in the mock repository
func (m *MockUserRepository) GetAllUsers() ([]users.UserDTO, error) {
	args := m.Called()

	if user := args.Get(0); user != nil {
		return user.([]users.UserDTO), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) CreateUser(user *users.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByEmail(email string) (*users.User, error) {
	args := m.Called(email)

	if user := args.Get(0); user != nil {
		return user.(*users.User), args.Error(1)
	}
	return nil, args.Error(1)
}

type MockPasswordService struct {
	mock.Mock
}

func (m *MockPasswordService) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *MockPasswordService) CheckPasswordHash(password, hash string) bool {
	args := m.Called(password, hash)
	return args.Bool(0)
}

type MockTokenService struct {
	mock.Mock
}

func (m *MockTokenService) GenerateJWT(id string) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

func (m *MockTokenService) GeneratePasswordRecoveryToken(email string) (string, error) {
	args := m.Called(email)
	return args.String(0), args.Error(1)
}

type MockEmailService struct {
	mock.Mock
}

func (m *MockEmailService) SendRecoveryEmail(email, token string) error {
	args := m.Called(email, token)
	return args.Error(0)
}

func TestRegisterUser_Success(t *testing.T) {
	userRepo := new(MockUserRepository)
	passwordService := new(MockPasswordService)
	tokenService := new(MockTokenService)
	emailService := new(MockEmailService)
	addressClient := new(MockAddressClient)

	authService := auth.NewAuthService(userRepo, passwordService, tokenService, emailService, addressClient)

	user := &users.User{
		FullName:        "John Doe",
		Email:           "john.doe@example.com",
		Password:        "SecureP@ssw0rd!",
		ConfirmPassword: "SecureP@ssw0rd!",
		Address: address.Address{
			Street: "123 Main St",
			City:   "Sample City",
			State:  "SP",
			CEP:    "12345678",
			Number: "123",
		},
	}

	// Mock the expected behavior of FetchCEPData
	validAddress := &clients.AddressDTO{
		Street: "123 Main St",
		City:   "Sample City",
		State:  "SP",
		Cep:    "12345-678",
	}
	addressClient.On("FetchCEPData", user.Address.CEP).Return(validAddress, nil)    // Mock the response for CEP
	userRepo.On("GetUserByEmail", user.Email).Return(nil, nil)                      // Usuário não existe
	passwordService.On("HashPassword", user.Password).Return("hashedPassword", nil) // Password com sucesso
	userRepo.On("CreateUser", user).Return(nil)                                     // Criação de usuário bem-sucedida

	err := authService.RegisterUser(user)

	assert.NoError(t, err)
	userRepo.AssertExpectations(t)
	passwordService.AssertExpectations(t)
}

func TestRegisterUser_UserExists(t *testing.T) {
	userRepo := new(MockUserRepository)
	passwordService := new(MockPasswordService)
	tokenService := new(MockTokenService)
	emailService := new(MockEmailService)
	addressClient := new(MockAddressClient)

	authService := auth.NewAuthService(userRepo, passwordService, tokenService, emailService, addressClient)

	user := &users.User{
		FullName:        "Jane Doe",
		Email:           "jane.doe@example.com",
		Password:        "SecureP@ssw0rd!",
		ConfirmPassword: "SecureP@ssw0rd!",
		Address: address.Address{
			Street: "123 Main St",
			City:   "Sample City",
			State:  "SP",
			CEP:    "12345678",
			Number: "123",
		},
	}

	userRepo.On("GetUserByEmail", user.Email).Return(user, nil) // User already exists
	// Mock the expected behavior of FetchCEPData
	validAddress := &clients.AddressDTO{
		Street: "123 Main St",
		City:   "Sample City",
		State:  "SP",
		Cep:    "12345-678",
	}
	addressClient.On("FetchCEPData", user.Address.CEP).Return(validAddress, nil) // Mock the response for CEP

	err := authService.RegisterUser(user)

	assert.Error(t, err)
	assert.Equal(t, "user already exists", err.Error())
	userRepo.AssertExpectations(t)
}

func TestLogin_Success(t *testing.T) {
	userRepo := new(MockUserRepository)
	passwordService := new(MockPasswordService)
	tokenService := new(MockTokenService)
	emailService := new(MockEmailService)
	addressClient := new(MockAddressClient)

	authService := auth.NewAuthService(userRepo, passwordService, tokenService, emailService, addressClient)

	credentials := auth.Credentials{
		Email:    "john.doe@example.com",
		Password: "SecureP@ssw0rd!",
	}

	user := &users.User{
		Id:       "123",
		Email:    credentials.Email,
		Password: "hashedPassword",
	}

	userRepo.On("GetUserByEmail", credentials.Email).Return(user, nil)                        // Usuer found
	passwordService.On("CheckPasswordHash", credentials.Password, user.Password).Return(true) // Password correct
	tokenService.On("GenerateJWT", user.Id).Return("token", nil)                              // Generate token

	token, err := authService.Login(credentials)

	assert.NoError(t, err)
	assert.Equal(t, "token", token)
	userRepo.AssertExpectations(t)
	passwordService.AssertExpectations(t)
	tokenService.AssertExpectations(t)
}

func TestLogin_InvalidCredentials(t *testing.T) {
	userRepo := new(MockUserRepository)
	passwordService := new(MockPasswordService)
	tokenService := new(MockTokenService)
	emailService := new(MockEmailService)
	addressClient := new(MockAddressClient)

	authService := auth.NewAuthService(userRepo, passwordService, tokenService, emailService, addressClient)

	credentials := auth.Credentials{
		Email:    "john.doe@example.com",
		Password: "WrongPassword",
	}

	userRepo.On("GetUserByEmail", credentials.Email).Return(nil, nil) // User not found
	token, err := authService.Login(credentials)

	assert.Error(t, err)
	assert.Equal(t, "invalid credentials", err.Error())
	assert.Empty(t, token)
	userRepo.AssertExpectations(t)
}

func TestRecoverPassword_Success(t *testing.T) {
	userRepo := new(MockUserRepository)
	passwordService := new(MockPasswordService)
	tokenService := new(MockTokenService)
	emailService := new(MockEmailService)
	addressClient := new(MockAddressClient)

	authService := auth.NewAuthService(userRepo, passwordService, tokenService, emailService, addressClient)

	email := "john.doe@example.com"

	user := &users.User{
		Email: email,
	}

	userRepo.On("GetUserByEmail", email).Return(user, nil)                               // User found
	tokenService.On("GeneratePasswordRecoveryToken", email).Return("recoveryToken", nil) // Generate recovery token
	emailService.On("SendRecoveryEmail", email, "recoveryToken").Return(nil)             // Sends email

	err := authService.RecoverPassword(email)

	assert.NoError(t, err)
	userRepo.AssertExpectations(t)
	tokenService.AssertExpectations(t)
	emailService.AssertExpectations(t)
}

func TestRecoverPassword_UserNotFound(t *testing.T) {
	userRepo := new(MockUserRepository)
	passwordService := new(MockPasswordService)
	tokenService := new(MockTokenService)
	emailService := new(MockEmailService)
	addressClient := new(MockAddressClient)

	authService := auth.NewAuthService(userRepo, passwordService, tokenService, emailService, addressClient)

	email := "non.existent@example.com"

	userRepo.On("GetUserByEmail", email).Return(nil, nil) // User not found

	err := authService.RecoverPassword(email)

	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	userRepo.AssertExpectations(t)
}
