package repositories

import (
	"github.com/cable_management/cable_management_be/src/entities"
	"gorm.io/gorm"
)

type IUserRepository interface {
	FindByEmail(email string) (*entities.User, error)
	FindByEmailAndRole(email, role string) (*entities.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur UserRepository) FindByEmail(email string) (*entities.User, error) {

	matchingUser := &entities.User{}
	result := ur.db.First(matchingUser, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}

	return matchingUser, nil
}

func (ur UserRepository) FindByEmailAndRole(email, role string) (*entities.User, error) {

	matchingUser := &entities.User{}
	result := ur.db.First(matchingUser, "email = ? AND role = ?", email, role)
	if result.Error != nil {
		return nil, result.Error
	}

	return matchingUser, nil
}
