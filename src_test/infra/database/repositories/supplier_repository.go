package repositories

import (
	"github.com/cable_management/cable_management_be/src_test/entities"
	"gorm.io/gorm"
)

type ISupplierRepository interface {
	Insert(entity *entities.Supplier) error
	FindByEmail(email string) (*entities.Supplier, error)
}

type SupplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *SupplierRepository {
	return &SupplierRepository{db: db}
}

func (ar SupplierRepository) Insert(entity *entities.Supplier) error {
	result := ar.db.Create(entity)
	return result.Error
}

func (ar SupplierRepository) FindByEmail(email string) (*entities.Supplier, error) {

	matchingSupplier := &entities.Supplier{}
	result := ar.db.First(matchingSupplier, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}

	return matchingSupplier, nil
}
