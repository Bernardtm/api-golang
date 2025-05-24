package shareds

import (
	"time"
)

type MessagingResponse struct {
	MessageID        int        `json:"messageID"`
	SenderID         int        `json:"senderID"`
	ReceiverID       int        `json:"receiverID"`
	Message          string     `json:"message"`
	SentTime         time.Time  `json:"sent_time"`
	CreationDate     time.Time  `json:"creation_date"`
	ModificationDate *time.Time `json:"ModificationDate"`
	StatusUUID         string     `json:"StatusUUID"`
}
