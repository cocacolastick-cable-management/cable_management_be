package entities

type IAbstractAccount interface{}

type AbstractAccount struct {
	AbstractEntity

	Email        string `gorm:"unique;type:varchar"`
	PasswordHash string `gorm:"unique;type:varchar"`
}

func NewAbstractAccount(email string, passwordHash string) *AbstractAccount {
	return &AbstractAccount{Email: email, PasswordHash: passwordHash}
}
