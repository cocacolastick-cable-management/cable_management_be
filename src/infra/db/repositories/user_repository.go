package repositories

import (
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/cable_management/cable_management_be/src/infra/db/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur UserRepository) FindById(id uuid.UUID) (*entities.User, error) {

	matchingUser := &entities.User{}
	result := ur.db.First(matchingUser, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return matchingUser, nil
}

func (ur UserRepository) FindByEmail(email string) (*entities.User, error) {

	matchingUser := &entities.User{}
	result := ur.db.First(matchingUser, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}

	return matchingUser, nil
}

func (ur UserRepository) GetList(page uint, size uint, orderBy *string, lastTimestamp *time.Time) ([]*entities.User, error) {

	userList := make([]*entities.User, size)
	utils.Pagination(ur.db, page, size, orderBy, lastTimestamp).Find(&userList)
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
