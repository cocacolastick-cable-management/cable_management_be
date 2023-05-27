package utils

import (
	"gorm.io/gorm"
	"time"
)

func Pagination(db *gorm.DB, page uint, size uint, orderBy *string, lastTimestamp *time.Time) *gorm.DB {

	query := db

	if orderBy != nil {
		query = query.Order(orderBy)
	} else {
		query = query.Order("inserted_at desc")
	}

	if lastTimestamp != nil {
		query = query.Where("created_at < ?", lastTimestamp)
	}

	query = query.
		Offset(int((page - 1) * size)).
		Limit(int(size))

	return query
}
