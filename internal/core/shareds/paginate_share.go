package shareds

type PaginateShared struct {
	Page int `json:"page" db:"1"`
	Size int `json:"size" db:"1"`
}
