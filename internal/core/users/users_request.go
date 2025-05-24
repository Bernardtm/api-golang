package users

// UserRequest representes the data required to create/update a user
// @Summary UserRequest
// @Description Struct containing the data required to create/update a user
type UserRequest struct {
	Id           string  `json:"user_uuid" db:"user_uuid" example:"user_uuid"`
	Username     string  `json:"username" db:"username" binding:"required" example:"username"`
	Email        string  `json:"email" db:"email" binding:"required,email" example:"test@test.com"`
	Password     string  `json:"password" db:"password" example:"Password123#"`
	PositionUUID *string `json:"position_uuid" db:"position_uuid" example:""`
	TaxNumber    *string `json:"tax_number" db:"tax_number" example:"12345678901"`
	StatusUUID   string  `json:"status_uuid" db:"status_uuid" example:""`
	Position     *string `json:"position" db:"position" example:""`
	Phone        *string `json:"phone" db:"phone" example:""`
}

type UserVisualizations struct {
	UserUUID          string `json:"user_uuid"`
	VisualizationUUID string `json:"visualization_uuid"`
}

type UserFields struct {
	UserUUID   string `json:"user_uuid"`
	FieldsUUID string `json:"field_uuid"`
}

type UserActivity struct {
	UserUUID     string `json:"user_uuid"`
	ActivityUUID string `json:"activity_uuid"`
}

type UserControlPanel struct {
	UserUUID         string `json:"user_uuid"`
	ControlPanelUUID string `json:"control_panel_uuid"`
}

type UpdateUserPermissionsRequest struct {
	Fields         []string `json:"fields,omitempty"`
	ControlPanels  []string `json:"control_panels,omitempty"`
	Activities     []string `json:"actions,omitempty"`
	Visualizations []string `json:"visualizations,omitempty"`
}

type UserUpdateRequest struct {
	Name             *string `json:"username,omitempty"`
	TaxNumber        *string `json:"taxNumber,omitempty"`
	Position         *string `json:"position,omitempty"`
	Email            *string `json:"email,omitempty"`
	Phone            *string `json:"phone,omitempty"`
	ProfileImageLink *string `json:"profile_image_link,omitempty"`
}
