package repositories

import (
	"github.com/cable_management/cable_management_be/src/entities"
	"gorm.io/gorm"
)

type IPlannerRepository interface {
	Insert(entity *entities.Planner) error
	FindByEmail(email string) (*entities.Planner, error)
}

type PlannerRepository struct {
	db *gorm.DB
}

func NewPlannerRepository(db *gorm.DB) *PlannerRepository {
	return &PlannerRepository{db: db}
}

func (ar PlannerRepository) Insert(entity *entities.Planner) error {
	result := ar.db.Create(entity)
	return result.Error
}

func (ar PlannerRepository) FindByEmail(email string) (*entities.Planner, error) {

	matchingPlanner := &entities.Planner{}
	result := ar.db.First(matchingPlanner, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}

	return matchingPlanner, nil
}
