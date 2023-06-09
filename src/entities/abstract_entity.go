package entities

import (
	"github.com/google/uuid"
	"time"
)

type IAbstractEntity interface{}

type AbstractEntity struct {
	Id         uuid.UUID `gorm:"primaryKey;type:varchar"`
	InsertedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

func NewAbstractEntity() AbstractEntity {
	return AbstractEntity{Id: uuid.New()}
}
