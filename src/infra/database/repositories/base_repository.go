package repositories

import (
	"github.com/cable_management/cable_management_be/src/entities"
)

type IBaseRepository[T entities.IAbstractEntity] interface {
	Insert(entity T) error
}

//type BaseRepository[T entities.IAbstractEntity] struct {
//	db *gorm.DB
//}
//
//func NewBaseRepository[T entities.IAbstractEntity](db *gorm.DB) BaseRepository[T] {
//	return BaseRepository[T]{db: db}
//}
//
//func (br BaseRepository[T]) Insert(entity T) error {
//	result := br.db.Create(entity)
//	return result.Error
//}
