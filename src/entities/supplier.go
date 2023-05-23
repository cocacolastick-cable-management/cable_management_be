package entities

type Supplier struct {
	AbstractAccount
}

func NewSupplier(email, passwordHash string) *Planner {
	return &Planner{
		AbstractAccount: NewAbstractAccount(email, passwordHash),
	}
}
