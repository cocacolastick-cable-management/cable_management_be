package repositories

import (
	"github.com/cable_management/cable_management_be/src/domain/entities"
	"github.com/cable_management/cable_management_be/src/infra/db/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type WithDrawRequestHistoryRepository struct {
	db *gorm.DB
}

func NewWithDrawRequestHistoryRepository(db *gorm.DB) *WithDrawRequestHistoryRepository {
	return &WithDrawRequestHistoryRepository{db: db}
}

func (cr WithDrawRequestHistoryRepository) FindById(id uuid.UUID, withs []string) (*entities.WithDrawRequestHistory, error) {

	matchingWithDrawRequestHistory := &entities.WithDrawRequestHistory{}

	query := cr.db
	for _, with := range withs {
		query = query.Preload(with)
	}

	result := query.First(matchingWithDrawRequestHistory, "with_draw_request_histories.id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return matchingWithDrawRequestHistory, nil
}

func (cr WithDrawRequestHistoryRepository) GetList(page uint, size uint, orderBy *string, lastTimestamp *time.Time, withs []string) ([]*entities.WithDrawRequestHistory, error) {

	historyList := make([]*entities.WithDrawRequestHistory, size)
	query := utils.Pagination(cr.db, page, size, orderBy, lastTimestamp)

	for _, with := range withs {
		query = query.Joins(with)
	}

	query.Find(&historyList)

	return historyList, nil
}

func (cr WithDrawRequestHistoryRepository) Insert(history *entities.WithDrawRequestHistory) error {
	result := cr.db.Create(history)
	return result.Error
}
