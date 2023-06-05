package repositories

import (
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/google/uuid"
)

type INotificationRepository interface {
	Insert(user *entities.Notification) error
	InsertMany(users []*entities.Notification) error
	FindManyByReceiverId(userId uuid.UUID, withs []string) ([]*entities.Notification, error)
	UpdateIsReadById(id uuid.UUID, isRead bool) error
}
