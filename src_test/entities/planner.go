package entities

type Planner struct {
	AbstractAccount
}

func NewPlanner(email, passwordHash string) *Planner {
	return &Planner{
		AbstractAccount: NewAbstractAccount(email, passwordHash),
	}
}
