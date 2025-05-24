package storages

import "mime/multipart"

type UploadDto struct {
	FileName         string         `json:"file_name"`          // File name of the attachment
	OriginalFileName string         `json:"original_file_name"` // Original file name of the attachment
	FileStream       multipart.File // File stream of the attachment
}

type StorageProvider interface {
	Upload(email UploadDto) (interface{}, error)
}
