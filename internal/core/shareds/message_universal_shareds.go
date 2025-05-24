package shareds

// ErrorResponse represents a struct of error response
type ErrorResponse struct {
	Message string `json:"message"`
}

// MessageResponse represents a generic success response
type MessageResponse struct {
	Message string `json:"message"`
}

type ApiResponse struct {
	Status  string      `json:"status"` // success, error
	Data    interface{} `json:"data"`
	Message string      `json:"message,omitempty"`
}

type ApiError struct {
	Status  string      `json:"status"` // error
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"` // error details, example: {"name": ["Name is required"]}
	Code    int         `json:"code,omitempty"`
}

type ApiResponsePaginated struct {
	Status     string      `json:"status"` // success, error
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type Pagination struct {
	CurrentPage int `json:"currentPage"`
	PerPage     int `json:"perPage"`
	TotalItems  int `json:"totalItems"`
	TotalPages  int `json:"totalPages"`
}

type PaginatedResponse struct {
	Columns          []TableColumn      `json:"columns,omitempty"`
	SelectableColumn []SelectableColumn `json:"selectableColumns,omitempty"`
	Data             interface{}        `json:"data"`
	Meta             Meta               `json:"meta"`
}

type Meta struct {
	Page       int `json:"currentPage"`
	Count      int `json:"perPage"`
	TotalCount int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
}

type TableColumn struct {
	Key         string `json:"key"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

type SelectableColumn struct {
	Key         string       `json:"key"`
	Placeholder string       `json:"placeholder"`
	Options     []ValueLabel `json:"options"`
}

type ValueLabel struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type ValuesWithOption struct {
	Value  []string `json:"value"`
	Option string   `json:"option"`
}

type OptionsObject struct {
	Options []ValueLabel `json:"options"`
}

type KeyOption struct {
	Key    string `json:"key"`
	Option string `json:"option"`
	Value  string `json:"value"`
}

type ValueWithOption struct {
	Value  string `json:"value"`
	Option string `json:"option"`
}

type ValueWithOptionColorAndFeedback struct {
	Value   string `json:"value"`
	Option  string `json:"option"`
	Color   string `json:"color"`
	Like    bool   `json:"like"`
	Dislike bool   `json:"dislike"`
}
