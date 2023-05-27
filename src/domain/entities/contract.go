package entities

import (
	"github.com/google/uuid"
	"time"
)

type Contract struct {
	AbstractEntity

	StartDay    time.Time
	EndDay      time.Time
	CableAmount uint
	CreatedAt   time.Time
	SupplierId  uuid.UUID

	Supplier         *User              `gorm:"foreignKey:SupplierId"`
	WithDrawRequests []*WithDrawRequest `gorm:"foreignKey:ContractId"`
}

func NewContract(startDay time.Time, endDay time.Time, supplierId uuid.UUID) *Contract {
	return &Contract{AbstractEntity: NewAbstractEntity(), StartDay: startDay, EndDay: endDay, SupplierId: supplierId}
}
