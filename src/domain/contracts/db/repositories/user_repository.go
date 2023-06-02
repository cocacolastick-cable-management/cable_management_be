package repositories

import (
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/google/uuid"
)

type IUserRepository interface {
	Insert(user *entities.User) error
	FindById(id uuid.UUID) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	FindByEmailAndRole(email, role string) (*entities.User, error)
	GetList(with *string) ([]*entities.User, error)
	FindManyByRoles(roles []string, withs []string) ([]*entities.User, error)
}
