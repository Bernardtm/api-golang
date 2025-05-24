package status

import (
	"time"
)

type Status struct {
	StatusUUID       string     `json:"status_uuid" db:"status_uuid"`                       // UUID do status (chave primária)
	Name             string     `json:"name" db:"name"`                                     // Nome do status (único)
	CreationDate     *time.Time `json:"creation_date,omitempty" db:"creation_date"`         // Data de criação (padrão: data atual)
	ModificationDate *time.Time `json:"modification_date,omitempty" db:"modification_date"` // Data de modificação (opcional)
}
