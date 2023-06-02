package responses

import (
	"github.com/google/uuid"
	"time"
)

type ContractResponse struct {
	Id            uuid.UUID
	UniqueName    string
	SupplierId    uuid.UUID
	SupplierName  string
	SupplierEmail string
	CableAmount   uint
	Stock         int
	StartDay      time.Time
	EndDay        time.Time
	CreatedAt     time.Time
}
