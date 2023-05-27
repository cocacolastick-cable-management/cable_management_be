package requests

import "time"

type PaginationRequest struct {
	Page          uint       `validate:"required"`
	Size          uint       `validate:"required,max=225"`
	OrderBy       *string    `validate:"omitempty"`
	LastTimestamp *time.Time `validate:"omitempty"`
}
