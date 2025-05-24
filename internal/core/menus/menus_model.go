package menus

type Menu struct {
	MenuUUID         string  `json:"menu_uuid" db:"menu_uuid"`                           // UUID do menu (chave primária)
	Name             string  `json:"name" db:"name"`                                     // Nome do menu
	Icon             *string `json:"icon,omitempty" db:"icon"`                           // Ícone do menu (opcional)
	URL              string  `json:"url" db:"url"`                                       // URL associada ao menu
	OrderIndex       int     `json:"order_index" db:"order_index"`                       // Índice de ordenação
	CreationDate     *string `json:"creation_date,omitempty" db:"creation_date"`         // Data de criação (padrão: data atual)
	ModificationDate *string `json:"modification_date,omitempty" db:"modification_date"` // Data de modificação (opcional)
	StatusUUID       string  `json:"status_uuid" db:"status_uuid"`                       // UUID do status (chave estrangeira)
	Permission       string  `json:"permission,omitempty" db:"permission"`               // Permissão associada ao menu
}
