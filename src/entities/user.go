package entities

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	AbstractEntity

	Email        string `gorm:"unique;type:varchar"`
	PasswordHash string `gorm:"unique;type:varchar"`
	CreatedAt    time.Time

	CreatorId uuid.UUID `gorm:"type:varchar"`

	Creator *Admin `gorm:"foreignKey:CreatorId"`
}

func NewUser(email string, passwordHash string, createdAt time.Time) *User {
	return &User{AbstractEntity: NewAbstractEntity(), Email: email, PasswordHash: passwordHash, CreatedAt: createdAt}
}
