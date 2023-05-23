package entities

type Admin struct {
	AbstractAccount
}

func NewAdmin(email, passwordHash string) *Admin {
	return &Admin{
		AbstractAccount: NewAbstractAccount(email, passwordHash),
	}
}
