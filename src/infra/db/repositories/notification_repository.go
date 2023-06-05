package repositories

import (
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (n NotificationRepository) Insert(user *entities.Notification) error {
	result := n.db.Create(user)
	return result.Error
}

func (n NotificationRepository) InsertMany(users []*entities.Notification) error {
	result := n.db.Create(users)
	return result.Error
}

func (n NotificationRepository) FindManyByReceiverId(userId uuid.UUID, withs []string) ([]*entities.Notification, error) {

	var notifList []*entities.Notification
	query := n.db

	for _, with := range withs {
		query = query.Preload(with)
	}

	query.Find(&notifList, "notifications.user_id = ?", userId)
	return notifList, nil
}

func (n NotificationRepository) UpdateIsReadById(id uuid.UUID, isRead bool) error {

	var matchNoti *entities.Notification
	result := n.db.Find(&matchNoti, "notifications.receiver_id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	matchNoti.IsRead = isRead
	result = n.db.Save(matchNoti)

	return result.Error
}
