package entities

import "github.com/google/uuid"

type Planner struct {
	AbstractAccount
}

func NewPlanner(email, passwordHash string, creatorId uuid.UUID) *Planner {
	return &Planner{
		AbstractAccount: NewAbstractAccount(email, passwordHash, creatorId),
	}
}
