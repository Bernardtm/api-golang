package users

import (
	"bernardtm/backend/internal/core/status"
	"bernardtm/backend/internal/core/storage"

	"golang.org/x/crypto/bcrypt"
)

type UsersService interface {
	GetAll() ([]UserResponse, error)
	GetByID(id string) (UserResponse, error)
	Create(entity UserRequest) (string, error)
	Update(id string, entity UserRequest) error
	Delete(id string) error
	Paginate(page int, size int) ([]UserResponse, error)
}

type usersService struct {
	repo             UserRepository
	statusRepository status.StatusRepository
	storageService   storage.StorageService
}

func NewUsersService(repo UserRepository,
	statusRepository status.StatusRepository,
	storageService storage.StorageService) *usersService {
	return &usersService{
		repo:             repo,
		statusRepository: statusRepository,
		storageService:   storageService,
	}
}

func (s *usersService) GetAll() ([]UserResponse, error) {
	return s.repo.GetAll()
}

func (s *usersService) GetByID(id string) (UserResponse, error) {
	return s.repo.GetByID(id)
}

func (s *usersService) Create(entity UserRequest) (string, error) {
	if entity.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(entity.Password), bcrypt.DefaultCost)
		if err != nil {
			return "", err
		}
		entity.Password = string(hashedPassword)
	}
	return s.repo.Create(entity)
}

func (s *usersService) Update(id string, entity UserRequest) error {
	if entity.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(entity.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		entity.Password = string(hashedPassword)
	}
	return s.repo.Update(id, entity)
}

func (s *usersService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *usersService) Paginate(page int, size int) ([]UserResponse, error) {
	return s.repo.Paginate(page, size)
}
