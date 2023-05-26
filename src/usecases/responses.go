package usecases

import "time"

type AccountResponse struct {
	Role      string
	Email     string
	CreatedAt time.Time
}
