package entities

import "github.com/google/uuid"

type Supplier struct {
	AbstractAccount
}

func NewSupplier(email, passwordHash string, creatorId uuid.UUID) *Planner {
	return &Planner{
		AbstractAccount: NewAbstractAccount(email, passwordHash, creatorId),
	}
}
