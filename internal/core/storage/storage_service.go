package storage

import (
	"bernardtm/backend/pkg/providers/storages"
)

type StorageService interface {
	Upload(email storages.UploadDto) (interface{}, error)
}

// StorageService provides methods to interact with a storage provider
type storageService struct {
	provider storages.StorageProvider
}

// NewStorageService creates a new StorageService instance
func NewStorageService(provider storages.StorageProvider) *storageService {
	return &storageService{provider: provider}
}

// Upload uploads a file using the configured provider
func (s *storageService) Upload(dto storages.UploadDto) (interface{}, error) {
	return s.provider.Upload(dto)
}
