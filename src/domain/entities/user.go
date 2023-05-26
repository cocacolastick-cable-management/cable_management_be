package entities

import (
	"time"
)

type User struct {
	AbstractEntity

	Email        string `gorm:"uniqueIndex:uni_user;type:varchar"`
	PasswordHash string `gorm:"type:varchar"`
	Role         string `gorm:"uniqueIndex:uni_user;type:varchar"`
	CreatedAt    time.Time
}

func NewUser(role, email, passwordHash string, createdAt time.Time) *User {
	return &User{AbstractEntity: NewAbstractEntity(), Role: role, Email: email, PasswordHash: passwordHash, CreatedAt: createdAt}
}
