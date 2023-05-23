package entities

import "github.com/google/uuid"

type Contractor struct {
	AbstractAccount
}

func NewContractor(email, passwordHash string, creatorId uuid.UUID) *Planner {
	return &Planner{
		AbstractAccount: NewAbstractAccount(email, passwordHash, creatorId),
	}
}
