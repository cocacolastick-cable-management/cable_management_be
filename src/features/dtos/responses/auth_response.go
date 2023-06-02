package responses

import (
	"github.com/cable_management/cable_management_be/src/domain/services"
)

type AuthResponse struct {
	services.AuthData

	Role  string
	Name  string
	Email string
}
