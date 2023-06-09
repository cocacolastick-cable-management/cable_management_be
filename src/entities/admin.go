package entities

type Admin struct {
	AbstractEntity

	Email        string `gorm:"unique;type:varchar"`
	PasswordHash string `gorm:"unique;type:varchar"`

	Planners    []Planner    `gorm:"foreignKey:CreatorId"`
	Suppliers   []Supplier   `gorm:"foreignKey:CreatorId"`
	Contractors []Contractor `gorm:"foreignKey:CreatorId"`
}

func NewAdmin(email, passwordHash string) *Admin {
	return &Admin{
		AbstractEntity: NewAbstractEntity(),
		Email:          email,
		PasswordHash:   passwordHash,
	}
}
