package responses

import (
	"github.com/google/uuid"
	"time"
)

type ContractResponse struct {
	Id           uuid.UUID
	SupplierId   uuid.UUID
	SupplierName string
	StartDay     time.Time
	EndDay       time.Time
	CreatedAt    time.Time
}
