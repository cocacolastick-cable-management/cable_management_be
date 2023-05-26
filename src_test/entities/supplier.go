package entities

type Supplier struct {
	AbstractAccount
}

func NewSupplier(email, passwordHash string) *Supplier {
	return &Supplier{
		AbstractAccount: NewAbstractAccount(email, passwordHash),
	}
}
