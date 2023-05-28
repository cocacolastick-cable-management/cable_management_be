package repositories

import (
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/google/uuid"
	"time"
)

type IWithDrawRequestRepository interface {
	Insert(withDrawRequest *entities.WithDrawRequest) error
	FindById(id uuid.UUID, withs []string) (*entities.WithDrawRequest, error)
	GetList(page uint, size uint, orderBy *string, lastTimestamp *time.Time, withs []string) ([]*entities.WithDrawRequest, error)
}
