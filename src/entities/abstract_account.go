package entities

import (
	"github.com/google/uuid"
	"time"
)

type IAbstractAccount interface{}

type AbstractAccount struct {
	AbstractEntity

	Email        string `gorm:"unique;type:varchar"`
	PasswordHash string `gorm:"unique;type:varchar"`
	CreatedAt    time.Time

	CreatorId uuid.UUID `gorm:"type:varchar"`

	Creator *Admin `gorm:"foreignKey:CreatorId"`
}

func NewAbstractAccount(email, passwordHash string, creatorId uuid.UUID) AbstractAccount {
	return AbstractAccount{
		AbstractEntity: NewAbstractEntity(),
		Email:          email,
		PasswordHash:   passwordHash,
		CreatorId:      creatorId,
	}
}
