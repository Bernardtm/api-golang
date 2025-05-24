package users

import "bernardtm/backend/internal/core/shareds"

type UserResponse struct {
	Id               string  `json:"user_uuid"`
	Username         string  `json:"username"`
	Email            string  `json:"email"`
	Password         string  `json:"password"`
	TaxNumber        *string `json:"tax_number"`
	Position         *string `json:"position"`
	CreationDate     string  `json:"creation_date"`
	ModificationDate *string `json:"ModificationDate,omitempty"`
	StatusUUID       string  `json:"status_uuid"`
	Phone            *string `json:"phone"`
	ProfileImageLink *string `json:"profile_image_link"`
}

type UserTableFrontEndResponse struct {
	PlayerUserUUID string             `json:"player_user_uuid"`
	Name           shareds.Value      `json:"name"`
	Position       *shareds.Value     `json:"position"`
	StatusColor    shareds.ValueColor `json:"status"`
}

type UserResponseDetails struct {
	Id            string             `json:"user_uuid"`
	Username      string             `json:"username"`
	Email         *string            `json:"email"`
	Phone         *string            `json:"phone"`
	UserStatus    shareds.ValueColor `json:"status"`
	NetWorth      *string            `json:"net_worth"`
	FundsQuantity string             `json:"funds_quantity"`
	Funds         UserFundsResponse  `json:"funds"`
	Position      *string            `json:"position"`
}

type UserFundsResponse struct {
	FundName      *string            `json:"fund_name"`
	FundType      *string            `json:"fund_type"`
	FundUUID      *string            `json:"fund_uuid"`
	FundsQuantity string             `json:"funds_quantity"`
	FundsStatus   shareds.ValueColor `json:"funds_status"`
}

type UserResponseDetailsWithPermissions struct {
	User        UserResponseDetails `json:"user"`
	Permissions UserPermissions     `json:"permissions"`
}

type UserPermissions struct {
	Fields         *[]string `json:"fields"`
	ControlPanels  *[]string `json:"control_panels"`
	Activities     *[]string `json:"actions"`
	Visualizations *[]string `json:"visualizations"`
}
