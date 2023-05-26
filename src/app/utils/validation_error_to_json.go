package utils

import "github.com/go-playground/validator/v10"

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidationErrorToStruct(err validator.ValidationErrors) []*ErrorResponse {
	var errors []*ErrorResponse
	if err != nil {
		for _, err := range err {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
