package message

const (
	Entity          string = "Message"
	Table           string = "messages"
	FieldID         string = "id"
	FieldRoomID     string = "room_id"
	FieldSenderID   string = "sender_id"
	FieldSentAt     string = "sent_at"
	FieldStatus     string = "status"
	FieldContent    string = "content"
	FieldCreatedAt  string = "created_at"
	FieldModifiedAt string = "modified_at"
	FieldDeletedAt  string = "deleted_at"
)

type Status string

const (
	StatusDraft Status = "DRAFT"
	StatusSent  Status = "SENT"
)
