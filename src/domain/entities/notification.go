package entities

import (
	"github.com/google/uuid"
	"time"
)

type Notification struct {
	AbstractEntity

	Action     string `gorm:"type:varchar"`
	IsRead     bool
	ObjectType string    `gorm:"type:varchar"`
	ObjectId   uuid.UUID `gorm:"type:varchar"`
	CreatedAt  time.Time

	SenderId   uuid.UUID
	ReceiverId uuid.UUID

	Sender   *User `gorm:"foreignKey:SenderId"`
	Receiver *User `gorm:"foreignKey:ReceiverId"`
}

func NewNotification(action string, isRead bool, objectType string, objectId uuid.UUID, createdAt time.Time, senderId uuid.UUID, receiverId uuid.UUID) *Notification {
	return &Notification{AbstractEntity: NewAbstractEntity(), Action: action, IsRead: isRead, ObjectType: objectType, ObjectId: objectId, CreatedAt: createdAt, SenderId: senderId, ReceiverId: receiverId}
}
