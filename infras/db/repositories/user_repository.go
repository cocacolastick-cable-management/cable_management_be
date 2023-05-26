package repositories

import (
	"github.com/cable_management/cable_management_be/src/entities"
	"gorm.io/gorm"
	"time"
)

type IUserRepository interface {
	Insert(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
	FindByEmailAndRole(email, role string) (*entities.User, error)
	GetList(page uint, size uint, lastTimestamp *time.Time) ([]*entities.User, error)
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

func (ur UserRepository) GetList(page uint, size uint, lastTimestamp *time.Time) ([]*entities.User, error) {

	userList := make([]*entities.User, size)
	ur.db.Offset(int((page - 1) * size)).Limit(int(size)).Find(&userList)
	return userList, nil
}

func (ur UserRepository) FindByEmailAndRole(email, role string) (*entities.User, error) {

	matchingUser := &entities.User{}
	result := ur.db.First(matchingUser, "email = ? AND role = ?", email, role)
	if result.Error != nil {
		return nil, result.Error
	}

	return matchingUser, nil
}

func (ur UserRepository) Insert(user *entities.User) error {
	result := ur.db.Create(user)
	return result.Error
}
