package auth

import (
	"time"
)

type TwoFactorCodes struct {
	TwoFactorCodeUUID string     `json:"twoFactorCode_uuid"`
	Id                string     `json:"user_uuid"`
	Code              string     `json:"code"`
	CreationDate      time.Time  `json:"creation_date"`
	ModificationDate  *time.Time `json:"modificationDate"`
	ExpirationDate    time.Time  `json:"expiration_date"`
	CurrentTimestamp  time.Time  `json:"current_timestamp"`
	StatusUUID        string     `json:"status_uuid"`
}
