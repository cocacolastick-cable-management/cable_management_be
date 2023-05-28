package entities

import (
	"github.com/cable_management/cable_management_be/src/domain/constants"
	"github.com/cable_management/cable_management_be/src/domain/errs"
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

func (c Contract) CalCableStock() (int, error) {

	if c.WithDrawRequests == nil {
		return -1, errs.ErrNotIncludeRelationship
	}

	stock := c.CableAmount
	for _, request := range c.WithDrawRequests {
		if request.Status != constants.WD_CanceledStatus {
			stock -= request.CableAmount
		}
	}

	return int(stock), nil
}
