package validations

import (
	"github.com/cable_management/cable_management_be/src/domain/constants"
	"golang.org/x/exp/slices"
)

func ValidateRole(role string) bool {

	roles := []string{
		constants.AdminRole,
		constants.PlannerRole,
		constants.SupplierRole,
		constants.ContractorRole}

	return slices.Contains[string](roles, role)
}
