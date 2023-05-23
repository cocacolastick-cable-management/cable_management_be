package entities

type Contractor struct {
	AbstractAccount
}

func NewContractor(email, passwordHash string) *Planner {
	return &Planner{
		AbstractAccount: NewAbstractAccount(email, passwordHash),
	}
}
