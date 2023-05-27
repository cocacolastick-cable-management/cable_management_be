package responses

import (
	"github.com/google/uuid"
	"time"
)

type UserResponse struct {
	Id          uuid.UUID
	Role        string
	DisplayName string
	Email       string
	CreatedAt   time.Time
}
