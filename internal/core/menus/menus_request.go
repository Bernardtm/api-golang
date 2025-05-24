package menus

type MenusRequest struct {
	Name       string `json:"name" db:"name"`
	Icon       string `json:"icon" db:"icon"`
	Url        string `json:"url" db:"url"`
	OrderIndex int    `json:"orderIndex" db:"orderIndex"`
	StatusUUID string `json:"status_uuid" db:"status_uuid"`
}
