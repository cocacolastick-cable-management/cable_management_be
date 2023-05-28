package entities

import (
	"github.com/google/uuid"
	"time"
)

type WithDrawRequestHistory struct {
	AbstractEntity

	Action    string `gorm:"type:varchar"`
	Status    string `gorm:"type:varchar"`
	CreatedAt time.Time

	CreatorId uuid.UUID `gorm:"type:varchar"`
	RequestId uuid.UUID `gorm:"type:varchar"`

	Creator *User            `gorm:"foreignKey:CreatorId"`
	Request *WithDrawRequest `gorm:"foreignKey:RequestId"`
}

func NewWithDrawRequestHistory(action string, createdAt time.Time, status string, creatorId uuid.UUID, requestId uuid.UUID) *WithDrawRequestHistory {
	return &WithDrawRequestHistory{AbstractEntity: NewAbstractEntity(), Action: action, Status: status, CreatedAt: createdAt, CreatorId: creatorId, RequestId: requestId}
}
