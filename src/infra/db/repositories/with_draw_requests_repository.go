package repositories

import (
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/cable_management/cable_management_be/src/infra/db"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WithDrawRequestRepository struct {
	db *gorm.DB
}

func NewWithDrawRequestRepository(db *gorm.DB) *WithDrawRequestRepository {
	return &WithDrawRequestRepository{db: db}
}

func (cr WithDrawRequestRepository) FindById(id uuid.UUID, withs []string) (*entities.WithDrawRequest, error) {

	matchingWithDrawRequest := &entities.WithDrawRequest{}

	query := cr.db
	for _, with := range withs {
		query = query.Preload(with)
	}

	result := query.First(matchingWithDrawRequest, "with_draw_requests.id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return matchingWithDrawRequest, nil
}

func (cr WithDrawRequestRepository) GetList(withs []string) ([]*entities.WithDrawRequest, error) {

	var withDrawReqList []*entities.WithDrawRequest
	query := db.DB

	for _, with := range withs {
		query = query.Preload(with)
	}

	query.Find(&withDrawReqList)

	return withDrawReqList, nil
}

func (cr WithDrawRequestRepository) Insert(withDrawReq *entities.WithDrawRequest) error {
	result := cr.db.Create(withDrawReq)
	return result.Error
}
