package services_test

import (
	clients "btmho/app/clients/endereco"
	"btmho/app/models"
	"btmho/app/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock para UserRepository
type MockEnderecoClient struct {
	mock.Mock
}

// FetchCEPData simulates fetching address data based on the provided CEP
func (m *MockEnderecoClient) FetchCEPData(cep string) (*clients.EnderecoDTO, error) {
	args := m.Called(cep)
	return args.Get(0).(*clients.EnderecoDTO), args.Error(1)
}

func (m *MockUserRepository) CreateUser(user *models.Usuario) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByEmail(email string) (*models.Usuario, error) {
	args := m.Called(email)

	if user := args.Get(0); user != nil {
		return user.(*models.Usuario), args.Error(1)
	}
	return nil, args.Error(1)
}

// Mock para PasswordService
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

// Mock para TokenService
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

// Mock para EmailService
type MockEmailService struct {
	mock.Mock
}

func (m *MockEmailService) SendRecoveryEmail(email, token string) error {
	args := m.Called(email, token)
	return args.Error(0)
}

// Teste para RegisterUser
func TestRegisterUser_Success(t *testing.T) {
	userRepo := new(MockUserRepository)
	passwordService := new(MockPasswordService)
	tokenService := new(MockTokenService)
	emailService := new(MockEmailService)
	enderecoClient := new(MockEnderecoClient)

	authService := services.NewAuthService(userRepo, passwordService, tokenService, emailService, enderecoClient)

	user := &models.Usuario{
		NomeCompleto:   "John Doe",
		Email:          "john.doe@example.com",
		Senha:          "SecureP@ssw0rd!",
		ConfirmarSenha: "SecureP@ssw0rd!",
		Endereco: models.Endereco{
			Rua:     "123 Main St",
			Cidade:  "Sample City",
			Estado:  "SP",
			CEP:     "12345678",
			Numero:  "123",
		},
	}

	// Expectativas
  // Mock the expected behavior of FetchCEPData
	validAddress := &clients.EnderecoDTO{
		Rua:     "123 Main St",
		Cidade:  "Sample City",
		Estado:  "SP",
		Cep:     "12345-678",
	}
  enderecoClient.On("FetchCEPData", user.Endereco.CEP).Return(validAddress, nil) // Mock the response for CEP
	userRepo.On("GetUserByEmail", user.Email).Return(nil, nil)                   // Usuário não existe
	passwordService.On("HashPassword", user.Senha).Return("hashedPassword", nil) // Senha com sucesso
	userRepo.On("CreateUser", user).Return(nil)                                  // Criação de usuário bem-sucedida

	// Executa o método
	err := authService.RegisterUser(user)

	// Verifica as asserções
	assert.NoError(t, err)
	userRepo.AssertExpectations(t)
	passwordService.AssertExpectations(t)
}

// Teste para RegisterUser quando o usuário já existe
func TestRegisterUser_UserExists(t *testing.T) {
	userRepo := new(MockUserRepository)
	passwordService := new(MockPasswordService)
	tokenService := new(MockTokenService)
	emailService := new(MockEmailService)
	enderecoClient := new(MockEnderecoClient)

	authService := services.NewAuthService(userRepo, passwordService, tokenService, emailService, enderecoClient)

	user := &models.Usuario{
		NomeCompleto:   "Jane Doe",
		Email:          "jane.doe@example.com",
		Senha:          "SecureP@ssw0rd!",
		ConfirmarSenha: "SecureP@ssw0rd!",
    Endereco: models.Endereco{
			Rua:     "123 Main St",
			Cidade:  "Sample City",
			Estado:  "SP",
			CEP:     "12345678",
			Numero:  "123",
		},
	}

	// Expectativas
	userRepo.On("GetUserByEmail", user.Email).Return(user, nil) // Usuário já existe
  // Mock the expected behavior of FetchCEPData
	validAddress := &clients.EnderecoDTO{
		Rua:     "123 Main St",
		Cidade:  "Sample City",
		Estado:  "SP",
		Cep:     "12345-678",
	}
  enderecoClient.On("FetchCEPData", user.Endereco.CEP).Return(validAddress, nil) // Mock the response for CEP

	// Executa o método
	err := authService.RegisterUser(user)

	// Verifica as asserções
	assert.Error(t, err)
	assert.Equal(t, "user already exists", err.Error())
	userRepo.AssertExpectations(t)
}

// Teste para Login
func TestLogin_Success(t *testing.T) {
	userRepo := new(MockUserRepository)
	passwordService := new(MockPasswordService)
	tokenService := new(MockTokenService)
	emailService := new(MockEmailService)
	enderecoClient := new(MockEnderecoClient)

	authService := services.NewAuthService(userRepo, passwordService, tokenService, emailService, enderecoClient)

	credentials := models.Credenciais{
		Email: "john.doe@example.com",
		Senha: "SecureP@ssw0rd!",
	}

	user := &models.Usuario{
		Id:    "123",
		Email: credentials.Email,
		Senha: "hashedPassword",
	}

	// Expectativas
	userRepo.On("GetUserByEmail", credentials.Email).Return(user, nil)                  // Usuário encontrado
	passwordService.On("CheckPasswordHash", credentials.Senha, user.Senha).Return(true) // Senha correta
	tokenService.On("GenerateJWT", user.Id).Return("token", nil)                        // Gera token com sucesso

	// Executa o método
	token, err := authService.Login(credentials)

	// Verifica as asserções
	assert.NoError(t, err)
	assert.Equal(t, "token", token)
	userRepo.AssertExpectations(t)
	passwordService.AssertExpectations(t)
	tokenService.AssertExpectations(t)
}

// Teste para Login com credenciais inválidas
func TestLogin_InvalidCredentials(t *testing.T) {
	userRepo := new(MockUserRepository)
	passwordService := new(MockPasswordService)
	tokenService := new(MockTokenService)
	emailService := new(MockEmailService)
	enderecoClient := new(MockEnderecoClient)

	authService := services.NewAuthService(userRepo, passwordService, tokenService, emailService, enderecoClient)

	credentials := models.Credenciais{
		Email: "john.doe@example.com",
		Senha: "WrongPassword",
	}

	// Expectativas
	userRepo.On("GetUserByEmail", credentials.Email).Return(nil, nil) // Usuário não encontrado
	// Executa o método
	token, err := authService.Login(credentials)

	// Verifica as asserções
	assert.Error(t, err)
	assert.Equal(t, "invalid credentials", err.Error())
	assert.Empty(t, token)
	userRepo.AssertExpectations(t)
}

// Teste para RecoverPassword
func TestRecoverPassword_Success(t *testing.T) {
	userRepo := new(MockUserRepository)
	passwordService := new(MockPasswordService)
	tokenService := new(MockTokenService)
	emailService := new(MockEmailService)
	enderecoClient := new(MockEnderecoClient)

	authService := services.NewAuthService(userRepo, passwordService, tokenService, emailService, enderecoClient)

	email := "john.doe@example.com"

	user := &models.Usuario{
		Email: email,
	}

	// Expectativas
	userRepo.On("GetUserByEmail", email).Return(user, nil)                               // Usuário encontrado
	tokenService.On("GeneratePasswordRecoveryToken", email).Return("recoveryToken", nil) // Gera token de recuperação
	emailService.On("SendRecoveryEmail", email, "recoveryToken").Return(nil)             // Envia email com sucesso

	// Executa o método
	err := authService.RecoverPassword(email)

	// Verifica as asserções
	assert.NoError(t, err)
	userRepo.AssertExpectations(t)
	tokenService.AssertExpectations(t)
	emailService.AssertExpectations(t)
}

// Teste para RecoverPassword com usuário não encontrado
func TestRecoverPassword_UserNotFound(t *testing.T) {
	userRepo := new(MockUserRepository)
	passwordService := new(MockPasswordService)
	tokenService := new(MockTokenService)
	emailService := new(MockEmailService)
	enderecoClient := new(MockEnderecoClient)

	authService := services.NewAuthService(userRepo, passwordService, tokenService, emailService, enderecoClient)

	email := "non.existent@example.com"

	// Expectativas
	userRepo.On("GetUserByEmail", email).Return(nil, nil) // Usuário não encontrado

	// Executa o método
	err := authService.RecoverPassword(email)

	// Verifica as asserções
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	userRepo.AssertExpectations(t)
}
