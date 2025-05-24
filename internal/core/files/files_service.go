package files

import (
	"bernardtm/backend/internal/core/status"
	"bernardtm/backend/internal/core/storage"
	"bernardtm/backend/pkg/providers/storages"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"mime/multipart"
	"strconv"
	"time"
)

type FilesService interface {
	Create(data FileRequest, fileStream multipart.File) (string, string, error)
}

type fileService struct {
	filesRepo      FilesRepository
	statusRepo     status.StatusRepository
	storageService storage.StorageService
}

func NewFilesService(filesRepo FilesRepository, statusRepo status.StatusRepository, storageService storage.StorageService) *fileService {
	return &fileService{
		filesRepo:      filesRepo,
		statusRepo:     statusRepo,
		storageService: storageService,
	}
}

func (s *fileService) Create(data FileRequest, fileStream multipart.File) (string, string, error) {

	// Create a SHA-256 hash from the file name
	hash := sha256.New()
	timeString := strconv.Itoa(int(time.Now().Unix()))
	hash.Write([]byte(timeString + data.File.Name)) // Hash the file name (or any part of the file)
	hashString := hex.EncodeToString(hash.Sum(nil))

	// Use the hash as the file name
	fmt.Printf("Generated hash for file name: %s\n", hashString)

	uploadDto := storages.UploadDto{
		FileName:         hashString,
		OriginalFileName: data.File.Name,
		FileStream:       fileStream,
	}

	var result storages.UploadFileOutput
	resultInterface, err := s.storageService.Upload(uploadDto)
	if err != nil {
		return "", "", errors.New("could not upload file")
	}
	result, ok := resultInterface.(storages.UploadFileOutput)
	if !ok {
		return "", "", errors.New("unexpected type from storage service upload")
	}

	status, err := s.getStatusByName("Pending")
	if err != nil {
		return "", "", err
	}

	model := &FileRequest{
		File: File{
			Name:       result.Name,
			Type:       result.Type,
			Folder:     result.Folder,
			Link:       result.Link,
			StatusUUID: status.StatusUUID,
		},
	}

	fileUUID, err := s.filesRepo.Create(*model)
	if err != nil {
		return "", "", err
	}
	return fileUUID, result.Link, nil
}

func (s *fileService) getStatusByName(name string) (status.StatusResponse, error) {
	statusResponse, err := s.statusRepo.GetByName(name)
	if err != nil {
		return status.StatusResponse{}, err
	}
	return statusResponse, nil
}
