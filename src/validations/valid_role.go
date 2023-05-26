package validations

import (
	"github.com/cable_management/cable_management_be/src/services/_commons"
	"golang.org/x/exp/slices"
)

func ValidateRole(role string) bool {

	roles := []string{
		_commons.AdminRole,
		_commons.PlannerRole,
		_commons.SupplierRole,
		_commons.ContractorRole}

	return slices.Contains[string](roles, role)
}
