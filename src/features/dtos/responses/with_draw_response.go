package responses

import (
	"github.com/google/uuid"
	"time"
)

type WithDrawResponse struct {
	Id                 uuid.UUID
	UniqueName         string
	SupplierId         uuid.UUID
	SupplierName       string
	SupplierEmail      string
	ContractorId       uuid.UUID
	ContractorName     string
	ContractorEmail    string
	ContractId         uuid.UUID
	ContractUniqueName string
	CableAmount        uint
	Status             string
	CreatedAt          time.Time
	Histories          []*HistoryResponse
}
