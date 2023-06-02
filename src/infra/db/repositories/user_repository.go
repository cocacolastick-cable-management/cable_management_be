package repositories

import (
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (ur UserRepository) GetList(with *string) ([]*entities.User, error) {

	var userList []*entities.User
	ur.db.Find(&userList)
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

func (ur UserRepository) FindManyByRoles(roles []string, withs []string) ([]*entities.User, error) {

	var userList []*entities.User
	query := ur.db
	for _, with := range withs {
		query = query.Preload(with)
	}
	query.Find(&userList, "users.role IN ?", roles)
	return userList, nil
}
