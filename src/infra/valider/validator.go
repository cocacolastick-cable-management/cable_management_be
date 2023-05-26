package valider

import (
	"github.com/go-playground/validator/v10"
)

var (
	Validator *validator.Validate
)

//func init() {
//	Validator = validator.New()
//
//	Validator.RegisterStructValidation(validations.ValidateCreateUserRequest, requests.CreateUserRequest{})
//}

func Init() {
	Validator = validator.New()
}

func SetStructValidations(validations []struct {
	Fn   validator.StructLevelFunc
	Type any
}) {

	if validations == nil {
		return
	}

	for _, validation := range validations {
		Validator.RegisterStructValidation(validation.Fn, validation.Type)
	}
}
