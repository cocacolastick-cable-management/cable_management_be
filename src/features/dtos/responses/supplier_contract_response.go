package responses

import (
	"github.com/google/uuid"
	"time"
)

type SupplierContractResponse struct {
	Id          uuid.UUID
	UniqueName  string
	CableAmount uint
	Stock       int
	StartDay    time.Time
	EndDay      time.Time
	CreatedAt   time.Time
}
