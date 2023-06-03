package repositories

import (
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/google/uuid"
)

type IContractRepository interface {
	Insert(contract *entities.Contract) error
	FindById(id uuid.UUID, withs []string) (*entities.Contract, error)
	FindByUniqueName(uniqueName string, withs []string) (*entities.Contract, error)
	GetActiveList(withs []string) ([]*entities.Contract, error)
}
