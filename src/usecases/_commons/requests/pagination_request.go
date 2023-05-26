package requests

import "time"

type PaginationRequest struct {
	Page          uint
	Size          uint       `validate:"max=225"`
	LastTimestamp *time.Time `validate:"omitempty"`
}
