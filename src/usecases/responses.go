package usecases

import "time"

type UserResponse struct {
	Role      string
	Email     string
	CreatedAt time.Time
}
