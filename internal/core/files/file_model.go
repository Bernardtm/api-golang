package files

import "time"

type File struct {
	FileUUID         string     `json:"file_uuid" db:"file_uuid"`                           // UUID do arquivo  (chave primaria)
	Name             string     `json:"file_name" db:"file_name"`                           // Nome do arquivo
	Link             string     `json:"file_link" db:"file_link"`                           // LInk do arquivo
	Folder           string     `json:"file_folder" db:"file_folder"`                       // Pasta do arquivo
	Type             string     `json:"file_type" db:"file_type"`                           // typo do arquivo
	StatusUUID       string     `json:"status_uuid" db:"status_uuid"`                       // Pasta do arquivo
	CreationDate     time.Time  `json:"creation_date,omitempty" db:"creation_date"`         // Data de criação (valor padrão: data atual)
	ModificationDate *time.Time `json:"modification_date,omitempty" db:"modification_date"` // Data de modificação (opcional)
}
