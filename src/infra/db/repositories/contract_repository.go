package repositories

import (
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ContractRepository struct {
	db *gorm.DB
}

func NewContractRepository(db *gorm.DB) *ContractRepository {
	return &ContractRepository{db: db}
}

func (cr ContractRepository) FindById(id uuid.UUID, withs []string) (*entities.Contract, error) {

	matchingContract := &entities.Contract{}

	query := cr.db
	for _, with := range withs {
		query = query.Preload(with)
	}

	result := query.First(matchingContract, "contracts.id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return matchingContract, nil
}

func (cr ContractRepository) FindByUniqueName(uniqueName string, withs []string) (*entities.Contract, error) {

	matchingContract := &entities.Contract{}

	query := cr.db
	for _, with := range withs {
		query = query.Preload(with)
	}

	result := query.First(matchingContract, "contracts.unique_name = ?", uniqueName)
	if result.Error != nil {
		return nil, result.Error
	}

	return matchingContract, nil
}

func (cr ContractRepository) GetActiveList(withs []string) ([]*entities.Contract, error) {

	var contractList []*entities.Contract
	query := cr.db

	for _, with := range withs {
		if with == "Supplier" {
			continue
		}
		query = query.Preload(with)
	}

	query.
		InnerJoins("Supplier", cr.db.Where(&entities.User{IsActive: true})).
		Find(&contractList)
	return contractList, nil
}

func (cr ContractRepository) Insert(contract *entities.Contract) error {
	result := cr.db.Create(contract)
	return result.Error
}

func (cr ContractRepository) FindManyBySupplierId(supplierId uuid.UUID, withs []string) ([]*entities.Contract, error) {

	var contractList []*entities.Contract
	query := cr.db

	for _, with := range withs {
		query = query.Preload(with)
	}

	query.Find(&contractList, "contracts.supplier_id = ?", supplierId)
	return contractList, nil
}
