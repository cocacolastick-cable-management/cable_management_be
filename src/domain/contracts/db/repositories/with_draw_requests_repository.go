package repositories

import (
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/google/uuid"
)

type IWithDrawRequestRepository interface {
	Insert(withDrawRequest *entities.WithDrawRequest) error
	FindById(id uuid.UUID, withs []string) (*entities.WithDrawRequest, error)
	GetActiveList(withs []string) ([]*entities.WithDrawRequest, error)
}
