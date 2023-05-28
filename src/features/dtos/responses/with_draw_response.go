package responses

import (
	"github.com/google/uuid"
	"time"
)

type WithDrawResponse struct {
	Id             uuid.UUID
	SupplierId     uuid.UUID
	SupplierName   string
	ContractorId   uuid.UUID
	ContractorName string
	ContractId     uuid.UUID
	CableAmount    uint
	Status         string
	CreatedAt      time.Time
	Histories      []*HistoryResponse
}
