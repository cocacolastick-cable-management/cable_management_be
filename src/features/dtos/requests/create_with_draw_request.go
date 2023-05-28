package requests

import "github.com/google/uuid"

type CreateWithDrawRequest struct {
	CableAmount  uint
	ContractId   uuid.UUID
	ContractorId uuid.UUID
}
