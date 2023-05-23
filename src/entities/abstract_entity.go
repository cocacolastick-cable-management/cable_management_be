package entities

import (
	"github.com/google/uuid"
	"time"
)

type AbstractEntity struct {
	Id         uuid.UUID `gorm:"primaryKey;type:varchar"`
	InsertedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
