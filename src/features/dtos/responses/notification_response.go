package responses

import "github.com/google/uuid"

type NotifResponse struct {
	Id               uuid.UUID
	Action           string
	Message          string
	IsRead           string
	ObjectType       string
	ObjectUniqueName string
	ObjectId         uuid.UUID
	SenderId         uuid.UUID
	ReceiverId       uuid.UUID
}
