package repositories

import (
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/cable_management/cable_management_be/src/infra/db/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
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

func (cr WithDrawRequestRepository) GetList(page uint, size uint, orderBy *string, lastTimestamp *time.Time, withs []string) ([]*entities.WithDrawRequest, error) {

	withDrawReqList := make([]*entities.WithDrawRequest, size)
	query := utils.Pagination(cr.db, page, size, orderBy, lastTimestamp)

	for _, with := range withs {
		query = query.Joins(with)
	}

	query.Find(&withDrawReqList)

	return withDrawReqList, nil
}

func (cr WithDrawRequestRepository) Insert(withDrawReq *entities.WithDrawRequest) error {
	result := cr.db.Create(withDrawReq)
	return result.Error
}
