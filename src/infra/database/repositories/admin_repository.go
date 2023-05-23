package repositories

import (
	"github.com/cable_management/cable_management_be/src/entities"
	"gorm.io/gorm"
)

type IAdminRepository interface {
	Insert(entity *entities.Admin) error
	FindByEmail(email string) (*entities.Admin, error)
}

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

func (ar AdminRepository) Insert(entity *entities.Admin) error {
	result := ar.db.Create(entity)
	return result.Error
}

func (ar AdminRepository) FindByEmail(email string) (*entities.Admin, error) {

	matchingAdmin := &entities.Admin{}
	result := ar.db.Find(matchingAdmin, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}

	return matchingAdmin, nil
}
