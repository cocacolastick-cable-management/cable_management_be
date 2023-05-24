package valider

import (
	"github.com/cable_management/cable_management_be/src/services"
	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slices"
)

var (
	Validator *validator.Validate
)

func init() {
	Validator = validator.New()

	err := Validator.RegisterValidation("validateRole", validateRole)
	if err != nil {
		panic(err)
	}
}

func validateRole(fl validator.FieldLevel) bool {

	roles := []string{
		services.AdminRole,
		services.PlannerRole,
		services.SupplierRole,
		services.ContractorRole}

	value := fl.Field().Interface().(string)

	return slices.Contains[string](roles, value)
}
