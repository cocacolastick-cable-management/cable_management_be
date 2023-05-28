package repositories

import (
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/google/uuid"
	"time"
)

type IContractRepository interface {
	Insert(contract *entities.Contract) error
	FindById(id uuid.UUID, withs []string) (*entities.Contract, error)
	GetList(page uint, size uint, orderBy *string, lastTimestamp *time.Time, withs []string) ([]*entities.Contract, error)
}
