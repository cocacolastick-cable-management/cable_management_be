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

func (cr ContractRepository) FindById(id uuid.UUID) (*entities.Contract, error) {

	matchingContract := &entities.Contract{}
	result := cr.db.First(matchingContract, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return matchingContract, nil
}

func (cr ContractRepository) GetList(page uint, size uint, orderBy *string, lastTimestamp *time.Time) ([]*entities.Contract, error) {

	contractList := make([]*entities.Contract, size)
	utils.Pagination(cr.db, page, size, orderBy, lastTimestamp).Find(&contractList)
	return contractList, nil
}

func (cr ContractRepository) Insert(user *entities.Contract) error {
	result := cr.db.Create(user)
	return result.Error
}
