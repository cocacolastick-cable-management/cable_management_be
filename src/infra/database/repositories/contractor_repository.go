package repositories

import (
	"github.com/cable_management/cable_management_be/src/entities"
	"gorm.io/gorm"
)

type IContractorRepository interface {
	Insert(entity *entities.Contractor) error
	FindByEmail(email string) (*entities.Contractor, error)
}

type ContractorRepository struct {
	db *gorm.DB
}

func NewContractorRepository(db *gorm.DB) *ContractorRepository {
	return &ContractorRepository{db: db}
}

func (ar ContractorRepository) Insert(entity *entities.Contractor) error {
	result := ar.db.Create(entity)
	return result.Error
}

func (ar ContractorRepository) FindByEmail(email string) (*entities.Contractor, error) {

	matchingContractor := &entities.Contractor{}
	result := ar.db.First(matchingContractor, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}

	return matchingContractor, nil
}
