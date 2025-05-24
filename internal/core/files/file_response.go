package files

import (
	"time"
)

type FileResponse struct {
	FileUUID         string     `json:"fileUUID"`
	Name             string     `json:"file_name"`
	Link             string     `json:"file_link"`
	Folder           string     `json:"file_folder"`
	Type             string     `json:"file_type"`
	StatusUUID       string     `json:"status_uuid" db:"status_uuid"`
	CreationDate     time.Time  `json:"creation_date,omitempty"`
	ModificationDate *time.Time `json:"modification_date,omitempty"`
}
