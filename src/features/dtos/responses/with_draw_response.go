package responses

import (
	"github.com/google/uuid"
	"time"
)

type WithDrawResponse struct {
	Id           uuid.UUID
	SupplierId   uuid.UUID
	ContractorId uuid.UUID
	ContractId   uuid.UUID
	CableAmount  uint
	Status       string
	CreatedAt    time.Time
}
