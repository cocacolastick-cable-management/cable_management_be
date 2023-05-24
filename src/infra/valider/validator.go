package valider

import (
	"github.com/cable_management/cable_management_be/src/usecases/adminUsecases"
	"github.com/cable_management/cable_management_be/src/usecases/commUsecases"
	"github.com/go-playground/validator/v10"
)

var (
	Validator *validator.Validate
)

func init() {
	Validator = validator.New()

	err := Validator.RegisterValidation("validateRole", commUsecases.ValidateRole)
	if err != nil {
		panic(err)
	}

	err = Validator.RegisterValidation("validateRoleExceptAdmin", adminUsecases.ValidateRoleExceptAdmin)
	if err != nil {
		panic(err)
	}

	err = Validator.RegisterValidation("validatePassword", adminUsecases.ValidatePassword)
	if err != nil {
		panic(err)
	}

	Validator.RegisterStructValidation(adminUsecases.ValidateCreateAccountRequest, adminUsecases.CreateAccountRequestDto{})
	if err != nil {
		panic(err)
	}
}
