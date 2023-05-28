package repositories

import (
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/cable_management/cable_management_be/src/infra/db/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
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

func (cr ContractRepository) GetList(page uint, size uint, orderBy *string, lastTimestamp *time.Time, withs []string) ([]*entities.Contract, error) {

	contractList := make([]*entities.Contract, size)
	query := utils.Pagination(cr.db, page, size, orderBy, lastTimestamp)

	for _, with := range withs {
		query = query.Joins(with)
	}

	query.Find(&contractList)

	return contractList, nil
}

func (cr ContractRepository) Insert(contract *entities.Contract) error {
	result := cr.db.Create(contract)
	return result.Error
}
