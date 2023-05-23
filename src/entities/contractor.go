package entities

type Contractor struct {
	AbstractAccount
}

func NewContractor(email, passwordHash string) *Contractor {
	return &Contractor{
		AbstractAccount: NewAbstractAccount(email, passwordHash),
	}
}
