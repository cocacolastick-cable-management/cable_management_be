package entities

import (
	"github.com/google/uuid"
	"time"
)

type WithDrawRequest struct {
	AbstractEntity

	Status      string `gorm:"type:varchar"`
	Counter     uint   `gorm:"autoIncrement"`
	CableAmount uint
	CreatedAt   time.Time

	ContractId   uuid.UUID
	ContractorId uuid.UUID

	Contract   *Contract                 `gorm:"foreignKey:ContractId"`
	Contractor *User                     `gorm:"foreignKey:ContractorId"`
	Histories  []*WithDrawRequestHistory `gorm:"foreignKey:RequestId"`
}

func NewWithDrawRequest(status string, cableAmount uint, createdAt time.Time, contractId uuid.UUID, contractorId uuid.UUID) *WithDrawRequest {
	return &WithDrawRequest{AbstractEntity: NewAbstractEntity(), Status: status, CableAmount: cableAmount, CreatedAt: createdAt, ContractId: contractId, ContractorId: contractorId}
}
