package responses

import (
	"github.com/google/uuid"
	"time"
)

type HistoryResponse struct {
	Id           uuid.UUID
	RequestId    uuid.UUID
	CreatorId    uuid.UUID
	CreatorName  string
	CreatorEmail string
	CreatorRole  string
	CreatedAt    time.Time
	Status       string
	Action       string
}
