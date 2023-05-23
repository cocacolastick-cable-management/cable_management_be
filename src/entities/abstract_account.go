package entities

type IAbstractAccount interface{}

type AbstractAccount struct {
	AbstractEntity

	Email        string `gorm:"unique;type:varchar"`
	PasswordHash string `gorm:"unique;type:varchar"`
}

func NewAbstractAccount(email, passwordHash string) AbstractAccount {
	return AbstractAccount{Email: email, PasswordHash: passwordHash}
}
