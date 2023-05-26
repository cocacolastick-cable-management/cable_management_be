package _commons

import "time"

type UserResponse struct {
	Role      string
	Email     string
	CreatedAt time.Time
}
