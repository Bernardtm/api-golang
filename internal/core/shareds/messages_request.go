package shareds

type MessagesRequest struct {
	SenderUUID  string `json:"sender_uuid" db:"sender_uuid"`
	ContactUUID string `json:"contact_uuid" db:"contact_uuid"`
	ReceiveUUID string `json:"receive_uuid" db:"receive_uuid"`
	Message     string `json:"message" db:"message"`
	Channel     string `json:"channel" db:"channel"`
	StatusUUID  string `json:"status_uuid" db:"status_uuid"`
}
