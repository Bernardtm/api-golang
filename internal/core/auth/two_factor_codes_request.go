package auth

type TwoFactorCodesRequest struct {
	Id              string `json:"userUUID"`
	Size            int    `json:"size"`
	MinutesToExpiry int    `json:"minutes_to_expiry"`
	IsAlphanumeric  bool   `json:"is_alphanumeric"`
}
