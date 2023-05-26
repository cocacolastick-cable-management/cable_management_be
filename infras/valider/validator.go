package valider

import (
	"github.com/cable_management/cable_management_be/src/dtos/requests"
	"github.com/cable_management/cable_management_be/src/validations"
	"github.com/go-playground/validator/v10"
)

var (
	Validator *validator.Validate
)

func init() {
	Validator = validator.New()

	Validator.RegisterStructValidation(validations.ValidateCreateUserRequest, requests.CreateUserRequest{})
}
