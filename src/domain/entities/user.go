package entities

import (
	"time"
)

type User struct {
	AbstractEntity

	DisplayName  string `gorm:"type:varchar"`
	Email        string `gorm:"uniqueIndex:uni_user;type:varchar"`
	PasswordHash string `gorm:"type:varchar"`
	Role         string `gorm:"uniqueIndex:uni_user;type:varchar"`
	IsActive     bool   `gorm:"type:boolean"`
	CreatedAt    time.Time
}

func NewUser(role, displayName, email, passwordHash string, isActive bool, createdAt time.Time) *User {
	return &User{AbstractEntity: NewAbstractEntity(), Role: role, DisplayName: displayName, Email: email, IsActive: isActive, PasswordHash: passwordHash, CreatedAt: createdAt}
}
