package users_test

import (
	"btmho/app/domain/users"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock of the user repository
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

func TestListUsers_Success(t *testing.T) {
	// Creates a new userRepository mock instance
	mockRepo := new(MockUserRepository)

	// Defines a expected behaviour: an user list is returned
	mockUsers := []users.UserDTO{
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

	userService := users.NewUserService(mockRepo)

	users, err := userService.ListUsers()

	assert.NoError(t, err)
	assert.Equal(t, mockUsers, users)
	mockRepo.AssertNumberOfCalls(t, "GetAllUsers", 1)
}

func TestListUsers_Error(t *testing.T) {
	mockRepo := new(MockUserRepository)

	mockRepo.On("GetAllUsers").Return(nil, errors.New("database error"))

	userService := users.NewUserService(mockRepo)

	users, err := userService.ListUsers()

	assert.Error(t, err)
	assert.Nil(t, users)
	assert.EqualError(t, err, "error fetching users")

	mockRepo.AssertNumberOfCalls(t, "GetAllUsers", 1)
}
