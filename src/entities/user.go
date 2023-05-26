package entities

import (
	"time"
)

const (
	AdminRole      = "admin"
	PlanerRole     = "planer"
	SupplierRole   = "supplier"
	ContractorRole = "contractor"
)

type User struct {
	AbstractEntity

	Email        string `gorm:"unique;type:varchar"`
	PasswordHash string `gorm:"unique;type:varchar"`
	Role         string `gorm:"type:varchar"`
	CreatedAt    time.Time
}

func NewUser(role, email, passwordHash string, createdAt time.Time) *User {
	return &User{AbstractEntity: NewAbstractEntity(), Role: role, Email: email, PasswordHash: passwordHash, CreatedAt: createdAt}
}
