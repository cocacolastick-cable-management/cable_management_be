package entities

type Admin struct {
	AbstractEntity

	Email        string `gorm:"unique;type:varchar"`
	PasswordHash string `gorm:"unique;type:varchar"`
}

func NewAdmin(email, passwordHash string) *Admin {
	return &Admin{
		AbstractEntity: NewAbstractEntity(),
		Email:          email,
		PasswordHash:   passwordHash,
	}
}
