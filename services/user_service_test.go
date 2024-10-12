package services_test

import (
	"btmho/app/models"
	"btmho/app/services"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository é um mock do repositório de usuários
type MockUserRepository struct {
	mock.Mock
}

// Implementa o método GetAllUsers no mock do repositório
func (m *MockUserRepository) GetAllUsers() ([]models.UserDTO, error) {
	args := m.Called()
	return args.Get(0).([]models.UserDTO), args.Error(1)
}

func TestListUsers_Success(t *testing.T) {
	// Cria uma instância do mock do repositório de usuários
	mockRepo := new(MockUserRepository)

	// Define o comportamento esperado: uma lista de usuários será retornada
	mockUsers := []models.UserDTO{
		{
			ID:    "123",
			Email: "user1@example.com",
		},
		{
			ID:    "456",
			Email: "user2@example.com",
		},
	}
	mockRepo.On("GetAllUsers").Return(mockUsers, nil)

	// Cria uma instância do serviço de usuários com o mock
	userService := services.NewUserService(mockRepo)

	// Executa o método ListUsers
	users, err := userService.ListUsers()

	// Verifica se não houve erro
	assert.NoError(t, err)
	// Verifica se os usuários retornados são os esperados
	assert.Equal(t, mockUsers, users)

	// Verifica se o método GetAllUsers foi chamado uma vez
	mockRepo.AssertNumberOfCalls(t, "GetAllUsers", 1)
}

func TestListUsers_Error(t *testing.T) {
	// Cria uma instância do mock do repositório de usuários
	mockRepo := new(MockUserRepository)

	// Define o comportamento esperado: um erro será retornado
	mockRepo.On("GetAllUsers").Return(nil, errors.New("database error"))

	// Cria uma instância do serviço de usuários com o mock
	userService := services.NewUserService(mockRepo)

	// Executa o método ListUsers
	users, err := userService.ListUsers()

	// Verifica se o erro foi retornado corretamente
	assert.Error(t, err)
	assert.Nil(t, users)
	assert.EqualError(t, err, "error fetching users")

	// Verifica se o método GetAllUsers foi chamado uma vez
	mockRepo.AssertNumberOfCalls(t, "GetAllUsers", 1)
}
