package entities

type Admin struct {
	AbstractEntity

	Email        string `gorm:"unique;type:varchar"`
	PasswordHash string `gorm:"unique;type:varchar"`

	Users []User `gorm:"foreignKey:CreatorId"`
}

func NewAdmin(email string, passwordHash string) *Admin {
	return &Admin{AbstractEntity: NewAbstractEntity(), Email: email, PasswordHash: passwordHash}
}
