package shareds

type MessagingRequest struct {
	SenderID   int    `json:"senderID" db:"senderID"`
	ReceiverID int    `json:"receiverID" db:"receiverID"`
	Message    string `json:"message" db:"message"`
	StatusUUID int    `json:"StatusUUID" db:"StatusUUID"`
}
