package entities

type IAbstractAccount interface{}

type AbstractAccount struct {
	AbstractEntity

	Email        string `gorm:"unique;type:varchar"`
	PasswordHash string `gorm:"unique;type:varchar"`
}
