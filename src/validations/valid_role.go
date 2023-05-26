package validations

import (
	"github.com/cable_management/cable_management_be/src/services"
	"golang.org/x/exp/slices"
)

func ValidateRole(role string) bool {

	roles := []string{
		services.AdminRole,
		services.PlannerRole,
		services.SupplierRole,
		services.ContractorRole}

	return slices.Contains[string](roles, role)
}
