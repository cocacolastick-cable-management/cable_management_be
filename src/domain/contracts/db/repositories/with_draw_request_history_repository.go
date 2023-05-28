package repositories

import (
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/google/uuid"
	"time"
)

type IWithDrawRequestHistoryRepository interface {
	Insert(history *entities.WithDrawRequestHistory) error
	FindById(id uuid.UUID, withs []string) (*entities.WithDrawRequestHistory, error)
	GetList(page uint, size uint, orderBy *string, lastTimestamp *time.Time, withs []string) ([]*entities.WithDrawRequestHistory, error)
}
