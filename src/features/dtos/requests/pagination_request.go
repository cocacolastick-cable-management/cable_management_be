package requests

import "time"

type PaginationRequest struct {
	Page          uint       `json:"page,omitempty" validate:"required"`
	Size          uint       `json:"size,omitempty" validate:"required,max=225"`
	OrderBy       *string    `json:"orderBy,omitempty" validate:"omitempty"`
	LastTimestamp *time.Time `json:"lastTimestamp,omitempty" validate:"omitempty"`
}
