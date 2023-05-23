package repositories

import (
	"github.com/cable_management/cable_management_be/src/entities"
	"gorm.io/gorm"
)

type IAdminRepository interface {
	IBaseRepository[entities.Admin]
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
