package requests

import "time"

type PaginationRequest struct {
	Page          int16      `validate:"required;min=0"`
	Size          int8       `validate:"required;min=0;max=225"`
	LastTimestamp *time.Time `validate:"omitempty"`
}
